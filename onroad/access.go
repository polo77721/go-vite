package onroad

import (
	"fmt"

	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/vm_db"
)

func (manager *Manager) GetOnRoadBlockByAddr(addr *types.Address, pageNum, pageCount uint64) ([]*ledger.AccountBlock, error) {
	log := manager.log.New("method", "newSignalToWorker", "addr", addr)
	hasOnroads, err := manager.chain.HasOnRoadBlocks(*addr)
	if err != nil {
		return nil, err
	}
	if !hasOnroads {
		return nil, nil
	}
	blocks := make([]*ledger.AccountBlock, 0)
	hashList, err := manager.chain.GetOnRoadBlocksHashList(*addr, int(pageNum), int(pageCount))
	if err != nil {
		log.Error(fmt.Sprintf("GetOnRoadBlocksHashList failed, err:%v", err))
		return nil, err
	}
	if hashList == nil {
		return nil, nil
	}
	for _, v := range hashList {
		onroad, err := manager.chain.GetAccountBlockByHash(v)
		if err != nil {
			log.Error(fmt.Sprintf("GetAccountBlockByHash failed, err:%v", err))
		}
		if onroad != nil {
			blocks = append(blocks, onroad)
		}
	}
	return blocks, nil
}

func (manager *Manager) insertBlockToPool(block *vm_db.VmAccountBlock) error {
	return manager.pool.AddDirectAccountBlock(block.AccountBlock.AccountAddress, block)
}

func (manager *Manager) hasOnRoadBlocks(addr *types.Address) (bool, error) {
	return manager.chain.HasOnRoadBlocks(*addr)
}

func (manager *Manager) deleteDirect(sendBlock *ledger.AccountBlock) error {
	manager.chain.DeleteOnRoad(sendBlock.ToAddress, sendBlock.Hash)
	return nil
}

type reactFunc func(address types.Address)

func (manager *Manager) addContractLis(gid types.Gid, f reactFunc) {
	manager.newContractListener.Store(gid, f)
}

func (manager *Manager) removeContractLis(gid types.Gid) {
	manager.newContractListener.Delete(gid)
}

func (manager *Manager) newSignalToWorker(block *ledger.AccountBlock) {
	newLog := manager.log.New("method", "newSignalToWorker", "Hash", block.Hash)
	toAddr := block.ToAddress
	if types.IsContractAddr(toAddr) {
		meta, err := manager.chain.GetContractMeta(toAddr)
		if err != nil {
			newLog.Error(fmt.Sprintf("GetContractMeta, err:%v", err))
			return
		}
		if meta == nil {
			newLog.Error("GetContractMeta fail.", "contract", toAddr)
			return
		}
		if f, ok := manager.newContractListener.Load(meta.Gid); ok {
			f.(reactFunc)(toAddr)
		}
	}
}
