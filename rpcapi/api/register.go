package api

import (
	"sort"
	"time"

	"github.com/vitelabs/go-vite/chain"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/consensus"
	"github.com/vitelabs/go-vite/log15"
	"github.com/vitelabs/go-vite/vite"
	"github.com/vitelabs/go-vite/vm/contracts/abi"
)

type RegisterApi struct {
	chain chain.Chain
	cs    consensus.Consensus
	log   log15.Logger
}

func NewRegisterApi(vite *vite.Vite) *RegisterApi {
	return &RegisterApi{
		chain: vite.Chain(),
		cs:    vite.Consensus(),
		log:   log15.New("module", "rpc_api/register_api"),
	}
}

func (r RegisterApi) String() string {
	return "RegisterApi"
}

// Private
func (r *RegisterApi) GetRegisterData(gid types.Gid, name string, nodeAddr types.Address) ([]byte, error) {
	return abi.ABIGovernance.PackMethod(abi.MethodNameRegister, gid, name, nodeAddr)
}

// Private
func (r *RegisterApi) GetCancelRegisterData(gid types.Gid, name string) ([]byte, error) {
	return abi.ABIGovernance.PackMethod(abi.MethodNameRevoke, gid, name)
}


// Private
func (r *RegisterApi) GetUpdateRegistrationData(gid types.Gid, name string, nodeAddr types.Address) ([]byte, error) {
	return abi.ABIGovernance.PackMethod(abi.MethodNameUpdateBlockProducingAddress, gid, name, nodeAddr)
}

type RegistrationInfo struct {
	Name                  string        `json:"name"`
	NodeAddr              types.Address `json:"nodeAddr"`
	PledgeAddr            types.Address `json:"pledgeAddr"`
	RewardWithdrawAddress types.Address `json:"rewardWithdrawAddress"`
	PledgeAmount          string        `json:"pledgeAmount"`
	WithdrawHeight        string        `json:"withdrawHeight"`
	WithdrawTime          int64         `json:"withdrawTime"`
	CancelTime            int64         `json:"cancelTime"`
}

type byRegistrationExpirationHeight []*types.Registration

func (a byRegistrationExpirationHeight) Len() int      { return len(a) }
func (a byRegistrationExpirationHeight) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byRegistrationExpirationHeight) Less(i, j int) bool {
	if a[i].ExpirationHeight == a[j].ExpirationHeight {
		if a[i].RevokeTime == a[j].RevokeTime {
			return a[i].Name > a[j].Name
		} else {
			return a[i].RevokeTime > a[j].RevokeTime
		}
	}
	return a[i].ExpirationHeight > a[j].ExpirationHeight
}

// Deprecated: use contract_getSBPList
func (r *RegisterApi) GetRegistrationList(gid types.Gid, pledgeAddr types.Address) ([]*RegistrationInfo, error) {
	db, err := getVmDb(r.chain, types.AddressGovernance)
	if err != nil {
		return nil, err
	}
	snapshotBlock, err := db.LatestSnapshotBlock()
	if err != nil {
		return nil, err
	}
	list, err := abi.GetRegistrationList(db, gid, pledgeAddr)
	if err != nil {
		return nil, err
	}
	targetList := make([]*RegistrationInfo, len(list))
	if len(list) > 0 {
		sort.Sort(byRegistrationExpirationHeight(list))
		for i, info := range list {
			targetList[i] = &RegistrationInfo{
				Name:                  info.Name,
				NodeAddr:              info.BlockProducingAddress,
				PledgeAddr:            info.StakeAddress,
				RewardWithdrawAddress: info.RewardWithdrawAddress,
				PledgeAmount:          *bigIntToString(info.Amount),
				WithdrawHeight:        Uint64ToString(info.ExpirationHeight),
				WithdrawTime:          getWithdrawTime(snapshotBlock.Timestamp, snapshotBlock.Height, info.ExpirationHeight),
				CancelTime:            info.RevokeTime,
			}
		}
	}
	return targetList, nil
}

type Reward struct {
	BlockReward      string `json:"blockReward"`
	VoteReward       string `json:"voteReward"`
	TotalReward      string `json:"totalReward"`
	BlockNum         string `json:"blockNum"`
	ExpectedBlockNum string `json:"expectedBlockNum"`
	Drained          bool   `json:"drained"`
}

type RewardInfo struct {
	RewardMap map[string]*Reward `json:"rewardMap"`
	StartTime int64              `json:"startTime"`
	EndTime   int64              `json:"endTime"`
}

type RegistParam struct {
	Name string     `json:"name"`
	Gid  *types.Gid `json:"gid"`
}

// Deprecated
func (r *RegisterApi) GetRegisterPledgeAddrList(paramList []*RegistParam) ([]*types.Address, error) {
	if len(paramList) == 0 {
		return nil, nil
	}
	db, err := getVmDb(r.chain, types.AddressGovernance)
	if err != nil {
		return nil, err
	}
	addrList := make([]*types.Address, len(paramList))
	for k, v := range paramList {
		var r *types.Registration
		if v.Gid == nil || *v.Gid == types.DELEGATE_GID {
			if r, err = abi.GetRegistration(db, types.SNAPSHOT_GID, v.Name); err != nil {
				return nil, err
			}
		} else {
			if r, err = abi.GetRegistration(db, *v.Gid, v.Name); err != nil {
				return nil, err
			}
		}
		if r != nil {
			addrList[k] = &r.StakeAddress
		}
	}
	return addrList, nil
}

type CandidateInfo struct {
	Name     string        `json:"name"`
	NodeAddr types.Address `json:"nodeAddr"`
	VoteNum  string        `json:"voteNum"`
}

// Deprecated: usecontract_getSBPVoteList instead
func (r *RegisterApi) GetCandidateList() ([]*CandidateInfo, error) {
	head := r.chain.GetLatestSnapshotBlock()
	details, _, err := r.cs.API().ReadVoteMap((*head.Timestamp).Add(time.Second))
	if err != nil {
		return nil, err
	}
	var result []*CandidateInfo
	for _, v := range details {
		result = append(result, &CandidateInfo{v.Name, v.CurrentAddr, *bigIntToString(v.Balance)})
	}
	return result, nil
}
