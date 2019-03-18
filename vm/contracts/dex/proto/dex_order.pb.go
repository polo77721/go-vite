// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dex_order.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Order struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Address              []byte   `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	Side                 bool     `protobuf:"varint,7,opt,name=Side,proto3" json:"Side,omitempty"`
	Type                 int32    `protobuf:"varint,8,opt,name=Type,proto3" json:"Type,omitempty"`
	Price                string   `protobuf:"bytes,9,opt,name=Price,proto3" json:"Price,omitempty"`
	Quantity             []byte   `protobuf:"bytes,10,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Amount               []byte   `protobuf:"bytes,11,opt,name=Amount,proto3" json:"Amount,omitempty"`
	LockedBuyFee         []byte   `protobuf:"bytes,12,opt,name=LockedBuyFee,proto3" json:"LockedBuyFee,omitempty"`
	Status               int32    `protobuf:"varint,13,opt,name=Status,proto3" json:"Status,omitempty"`
	CancelReason         int32    `protobuf:"varint,14,opt,name=CancelReason,proto3" json:"CancelReason,omitempty"`
	ExecutedQuantity     []byte   `protobuf:"bytes,16,opt,name=ExecutedQuantity,proto3" json:"ExecutedQuantity,omitempty"`
	ExecutedAmount       []byte   `protobuf:"bytes,17,opt,name=ExecutedAmount,proto3" json:"ExecutedAmount,omitempty"`
	ExecutedFee          []byte   `protobuf:"bytes,18,opt,name=ExecutedFee,proto3" json:"ExecutedFee,omitempty"`
	RefundToken          []byte   `protobuf:"bytes,19,opt,name=RefundToken,proto3" json:"RefundToken,omitempty"`
	RefundQuantity       []byte   `protobuf:"bytes,20,opt,name=RefundQuantity,proto3" json:"RefundQuantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_5df90bf1b4327f51, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Order) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Order) GetSide() bool {
	if m != nil {
		return m.Side
	}
	return false
}

func (m *Order) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Order) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Order) GetQuantity() []byte {
	if m != nil {
		return m.Quantity
	}
	return nil
}

func (m *Order) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Order) GetLockedBuyFee() []byte {
	if m != nil {
		return m.LockedBuyFee
	}
	return nil
}

func (m *Order) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Order) GetCancelReason() int32 {
	if m != nil {
		return m.CancelReason
	}
	return 0
}

func (m *Order) GetExecutedQuantity() []byte {
	if m != nil {
		return m.ExecutedQuantity
	}
	return nil
}

func (m *Order) GetExecutedAmount() []byte {
	if m != nil {
		return m.ExecutedAmount
	}
	return nil
}

func (m *Order) GetExecutedFee() []byte {
	if m != nil {
		return m.ExecutedFee
	}
	return nil
}

func (m *Order) GetRefundToken() []byte {
	if m != nil {
		return m.RefundToken
	}
	return nil
}

func (m *Order) GetRefundQuantity() []byte {
	if m != nil {
		return m.RefundQuantity
	}
	return nil
}

type OrderTokenInfo struct {
	TradeToken           []byte   `protobuf:"bytes,1,opt,name=TradeToken,proto3" json:"TradeToken,omitempty"`
	QuoteToken           []byte   `protobuf:"bytes,2,opt,name=QuoteToken,proto3" json:"QuoteToken,omitempty"`
	TradeTokenDecimals   int32    `protobuf:"varint,3,opt,name=TradeTokenDecimals,proto3" json:"TradeTokenDecimals,omitempty"`
	QuoteTokenDecimals   int32    `protobuf:"varint,4,opt,name=QuoteTokenDecimals,proto3" json:"QuoteTokenDecimals,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderTokenInfo) Reset()         { *m = OrderTokenInfo{} }
func (m *OrderTokenInfo) String() string { return proto.CompactTextString(m) }
func (*OrderTokenInfo) ProtoMessage()    {}
func (*OrderTokenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5df90bf1b4327f51, []int{1}
}

func (m *OrderTokenInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderTokenInfo.Unmarshal(m, b)
}
func (m *OrderTokenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderTokenInfo.Marshal(b, m, deterministic)
}
func (m *OrderTokenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderTokenInfo.Merge(m, src)
}
func (m *OrderTokenInfo) XXX_Size() int {
	return xxx_messageInfo_OrderTokenInfo.Size(m)
}
func (m *OrderTokenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderTokenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OrderTokenInfo proto.InternalMessageInfo

func (m *OrderTokenInfo) GetTradeToken() []byte {
	if m != nil {
		return m.TradeToken
	}
	return nil
}

func (m *OrderTokenInfo) GetQuoteToken() []byte {
	if m != nil {
		return m.QuoteToken
	}
	return nil
}

func (m *OrderTokenInfo) GetTradeTokenDecimals() int32 {
	if m != nil {
		return m.TradeTokenDecimals
	}
	return 0
}

func (m *OrderTokenInfo) GetQuoteTokenDecimals() int32 {
	if m != nil {
		return m.QuoteTokenDecimals
	}
	return 0
}

type OrderInfo struct {
	Order                *Order          `protobuf:"bytes,1,opt,name=Order,proto3" json:"Order,omitempty"`
	OrderTokenInfo       *OrderTokenInfo `protobuf:"bytes,2,opt,name=OrderTokenInfo,proto3" json:"OrderTokenInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OrderInfo) Reset()         { *m = OrderInfo{} }
func (m *OrderInfo) String() string { return proto.CompactTextString(m) }
func (*OrderInfo) ProtoMessage()    {}
func (*OrderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5df90bf1b4327f51, []int{2}
}

func (m *OrderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderInfo.Unmarshal(m, b)
}
func (m *OrderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderInfo.Marshal(b, m, deterministic)
}
func (m *OrderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderInfo.Merge(m, src)
}
func (m *OrderInfo) XXX_Size() int {
	return xxx_messageInfo_OrderInfo.Size(m)
}
func (m *OrderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OrderInfo proto.InternalMessageInfo

func (m *OrderInfo) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *OrderInfo) GetOrderTokenInfo() *OrderTokenInfo {
	if m != nil {
		return m.OrderTokenInfo
	}
	return nil
}

type OrderUpdateInfo struct {
	Id                   []byte          `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Status               int32           `protobuf:"varint,13,opt,name=Status,proto3" json:"Status,omitempty"`
	CancelReason         int32           `protobuf:"varint,14,opt,name=CancelReason,proto3" json:"CancelReason,omitempty"`
	ExecutedQuantity     []byte          `protobuf:"bytes,16,opt,name=ExecutedQuantity,proto3" json:"ExecutedQuantity,omitempty"`
	ExecutedAmount       []byte          `protobuf:"bytes,17,opt,name=ExecutedAmount,proto3" json:"ExecutedAmount,omitempty"`
	ExecutedFee          []byte          `protobuf:"bytes,18,opt,name=ExecutedFee,proto3" json:"ExecutedFee,omitempty"`
	RefundToken          []byte          `protobuf:"bytes,19,opt,name=RefundToken,proto3" json:"RefundToken,omitempty"`
	RefundQuantity       []byte          `protobuf:"bytes,20,opt,name=RefundQuantity,proto3" json:"RefundQuantity,omitempty"`
	OrderTokenInfo       *OrderTokenInfo `protobuf:"bytes,2,opt,name=OrderTokenInfo,proto3" json:"OrderTokenInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OrderUpdateInfo) Reset()         { *m = OrderUpdateInfo{} }
func (m *OrderUpdateInfo) String() string { return proto.CompactTextString(m) }
func (*OrderUpdateInfo) ProtoMessage()    {}
func (*OrderUpdateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5df90bf1b4327f51, []int{3}
}

func (m *OrderUpdateInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderUpdateInfo.Unmarshal(m, b)
}
func (m *OrderUpdateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderUpdateInfo.Marshal(b, m, deterministic)
}
func (m *OrderUpdateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderUpdateInfo.Merge(m, src)
}
func (m *OrderUpdateInfo) XXX_Size() int {
	return xxx_messageInfo_OrderUpdateInfo.Size(m)
}
func (m *OrderUpdateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderUpdateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OrderUpdateInfo proto.InternalMessageInfo

func (m *OrderUpdateInfo) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *OrderUpdateInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *OrderUpdateInfo) GetCancelReason() int32 {
	if m != nil {
		return m.CancelReason
	}
	return 0
}

func (m *OrderUpdateInfo) GetExecutedQuantity() []byte {
	if m != nil {
		return m.ExecutedQuantity
	}
	return nil
}

func (m *OrderUpdateInfo) GetExecutedAmount() []byte {
	if m != nil {
		return m.ExecutedAmount
	}
	return nil
}

func (m *OrderUpdateInfo) GetExecutedFee() []byte {
	if m != nil {
		return m.ExecutedFee
	}
	return nil
}

func (m *OrderUpdateInfo) GetRefundToken() []byte {
	if m != nil {
		return m.RefundToken
	}
	return nil
}

func (m *OrderUpdateInfo) GetRefundQuantity() []byte {
	if m != nil {
		return m.RefundQuantity
	}
	return nil
}

func (m *OrderUpdateInfo) GetOrderTokenInfo() *OrderTokenInfo {
	if m != nil {
		return m.OrderTokenInfo
	}
	return nil
}

type OrderNode struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=Order,proto3" json:"Order,omitempty"`
	ForwardOnLevel       [][]byte `protobuf:"bytes,2,rep,name=ForwardOnLevel,proto3" json:"ForwardOnLevel,omitempty"`
	BackwardOnLevel      [][]byte `protobuf:"bytes,3,rep,name=BackwardOnLevel,proto3" json:"BackwardOnLevel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderNode) Reset()         { *m = OrderNode{} }
func (m *OrderNode) String() string { return proto.CompactTextString(m) }
func (*OrderNode) ProtoMessage()    {}
func (*OrderNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_5df90bf1b4327f51, []int{4}
}

func (m *OrderNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderNode.Unmarshal(m, b)
}
func (m *OrderNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderNode.Marshal(b, m, deterministic)
}
func (m *OrderNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderNode.Merge(m, src)
}
func (m *OrderNode) XXX_Size() int {
	return xxx_messageInfo_OrderNode.Size(m)
}
func (m *OrderNode) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderNode.DiscardUnknown(m)
}

var xxx_messageInfo_OrderNode proto.InternalMessageInfo

func (m *OrderNode) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *OrderNode) GetForwardOnLevel() [][]byte {
	if m != nil {
		return m.ForwardOnLevel
	}
	return nil
}

func (m *OrderNode) GetBackwardOnLevel() [][]byte {
	if m != nil {
		return m.BackwardOnLevel
	}
	return nil
}

type OrderListMeta struct {
	Header               []byte   `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Tail                 []byte   `protobuf:"bytes,2,opt,name=Tail,proto3" json:"Tail,omitempty"`
	Length               int32    `protobuf:"varint,3,opt,name=Length,proto3" json:"Length,omitempty"`
	Level                int32    `protobuf:"varint,4,opt,name=Level,proto3" json:"Level,omitempty"`
	ForwardOnLevel       [][]byte `protobuf:"bytes,5,rep,name=ForwardOnLevel,proto3" json:"ForwardOnLevel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderListMeta) Reset()         { *m = OrderListMeta{} }
func (m *OrderListMeta) String() string { return proto.CompactTextString(m) }
func (*OrderListMeta) ProtoMessage()    {}
func (*OrderListMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_5df90bf1b4327f51, []int{5}
}

func (m *OrderListMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderListMeta.Unmarshal(m, b)
}
func (m *OrderListMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderListMeta.Marshal(b, m, deterministic)
}
func (m *OrderListMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderListMeta.Merge(m, src)
}
func (m *OrderListMeta) XXX_Size() int {
	return xxx_messageInfo_OrderListMeta.Size(m)
}
func (m *OrderListMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderListMeta.DiscardUnknown(m)
}

var xxx_messageInfo_OrderListMeta proto.InternalMessageInfo

func (m *OrderListMeta) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *OrderListMeta) GetTail() []byte {
	if m != nil {
		return m.Tail
	}
	return nil
}

func (m *OrderListMeta) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *OrderListMeta) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *OrderListMeta) GetForwardOnLevel() [][]byte {
	if m != nil {
		return m.ForwardOnLevel
	}
	return nil
}

type OrderBook struct {
	MarketSide           []string `protobuf:"bytes,1,rep,name=MarketSide,proto3" json:"MarketSide,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderBook) Reset()         { *m = OrderBook{} }
func (m *OrderBook) String() string { return proto.CompactTextString(m) }
func (*OrderBook) ProtoMessage()    {}
func (*OrderBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_5df90bf1b4327f51, []int{6}
}

func (m *OrderBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderBook.Unmarshal(m, b)
}
func (m *OrderBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderBook.Marshal(b, m, deterministic)
}
func (m *OrderBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBook.Merge(m, src)
}
func (m *OrderBook) XXX_Size() int {
	return xxx_messageInfo_OrderBook.Size(m)
}
func (m *OrderBook) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBook.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBook proto.InternalMessageInfo

func (m *OrderBook) GetMarketSide() []string {
	if m != nil {
		return m.MarketSide
	}
	return nil
}

type Transaction struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	TakerSide            bool     `protobuf:"varint,2,opt,name=TakerSide,proto3" json:"TakerSide,omitempty"`
	TakerId              []byte   `protobuf:"bytes,3,opt,name=TakerId,proto3" json:"TakerId,omitempty"`
	MakerId              []byte   `protobuf:"bytes,4,opt,name=MakerId,proto3" json:"MakerId,omitempty"`
	Price                string   `protobuf:"bytes,5,opt,name=Price,proto3" json:"Price,omitempty"`
	Quantity             []byte   `protobuf:"bytes,6,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Amount               []byte   `protobuf:"bytes,7,opt,name=Amount,proto3" json:"Amount,omitempty"`
	TakerFee             []byte   `protobuf:"bytes,8,opt,name=TakerFee,proto3" json:"TakerFee,omitempty"`
	MakerFee             []byte   `protobuf:"bytes,9,opt,name=MakerFee,proto3" json:"MakerFee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_5df90bf1b4327f51, []int{7}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Transaction) GetTakerSide() bool {
	if m != nil {
		return m.TakerSide
	}
	return false
}

func (m *Transaction) GetTakerId() []byte {
	if m != nil {
		return m.TakerId
	}
	return nil
}

func (m *Transaction) GetMakerId() []byte {
	if m != nil {
		return m.MakerId
	}
	return nil
}

func (m *Transaction) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Transaction) GetQuantity() []byte {
	if m != nil {
		return m.Quantity
	}
	return nil
}

func (m *Transaction) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Transaction) GetTakerFee() []byte {
	if m != nil {
		return m.TakerFee
	}
	return nil
}

func (m *Transaction) GetMakerFee() []byte {
	if m != nil {
		return m.MakerFee
	}
	return nil
}

type OrderFail struct {
	OrderInfo            *OrderInfo `protobuf:"bytes,1,opt,name=OrderInfo,proto3" json:"OrderInfo,omitempty"`
	ErrCode              string     `protobuf:"bytes,2,opt,name=ErrCode,proto3" json:"ErrCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OrderFail) Reset()         { *m = OrderFail{} }
func (m *OrderFail) String() string { return proto.CompactTextString(m) }
func (*OrderFail) ProtoMessage()    {}
func (*OrderFail) Descriptor() ([]byte, []int) {
	return fileDescriptor_5df90bf1b4327f51, []int{8}
}

func (m *OrderFail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderFail.Unmarshal(m, b)
}
func (m *OrderFail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderFail.Marshal(b, m, deterministic)
}
func (m *OrderFail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderFail.Merge(m, src)
}
func (m *OrderFail) XXX_Size() int {
	return xxx_messageInfo_OrderFail.Size(m)
}
func (m *OrderFail) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderFail.DiscardUnknown(m)
}

var xxx_messageInfo_OrderFail proto.InternalMessageInfo

func (m *OrderFail) GetOrderInfo() *OrderInfo {
	if m != nil {
		return m.OrderInfo
	}
	return nil
}

func (m *OrderFail) GetErrCode() string {
	if m != nil {
		return m.ErrCode
	}
	return ""
}

type MarketInfo struct {
	TradeTokenDecimals   int32    `protobuf:"varint,1,opt,name=TradeTokenDecimals,proto3" json:"TradeTokenDecimals,omitempty"`
	QuoteTokenDecimals   int32    `protobuf:"varint,2,opt,name=QuoteTokenDecimals,proto3" json:"QuoteTokenDecimals,omitempty"`
	Creator              []byte   `protobuf:"bytes,3,opt,name=Creator,proto3" json:"Creator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarketInfo) Reset()         { *m = MarketInfo{} }
func (m *MarketInfo) String() string { return proto.CompactTextString(m) }
func (*MarketInfo) ProtoMessage()    {}
func (*MarketInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5df90bf1b4327f51, []int{9}
}

func (m *MarketInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketInfo.Unmarshal(m, b)
}
func (m *MarketInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketInfo.Marshal(b, m, deterministic)
}
func (m *MarketInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketInfo.Merge(m, src)
}
func (m *MarketInfo) XXX_Size() int {
	return xxx_messageInfo_MarketInfo.Size(m)
}
func (m *MarketInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MarketInfo proto.InternalMessageInfo

func (m *MarketInfo) GetTradeTokenDecimals() int32 {
	if m != nil {
		return m.TradeTokenDecimals
	}
	return 0
}

func (m *MarketInfo) GetQuoteTokenDecimals() int32 {
	if m != nil {
		return m.QuoteTokenDecimals
	}
	return 0
}

func (m *MarketInfo) GetCreator() []byte {
	if m != nil {
		return m.Creator
	}
	return nil
}

type NewMarket struct {
	TradeToken           []byte   `protobuf:"bytes,1,opt,name=TradeToken,proto3" json:"TradeToken,omitempty"`
	QuoteToken           []byte   `protobuf:"bytes,2,opt,name=QuoteToken,proto3" json:"QuoteToken,omitempty"`
	TradeTokenSymbol     string   `protobuf:"bytes,3,opt,name=TradeTokenSymbol,proto3" json:"TradeTokenSymbol,omitempty"`
	QuoteTokenSymbol     string   `protobuf:"bytes,4,opt,name=QuoteTokenSymbol,proto3" json:"QuoteTokenSymbol,omitempty"`
	TradeTokenDecimals   int32    `protobuf:"varint,5,opt,name=TradeTokenDecimals,proto3" json:"TradeTokenDecimals,omitempty"`
	QuoteTokenDecimals   int32    `protobuf:"varint,6,opt,name=QuoteTokenDecimals,proto3" json:"QuoteTokenDecimals,omitempty"`
	Creator              []byte   `protobuf:"bytes,7,opt,name=Creator,proto3" json:"Creator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewMarket) Reset()         { *m = NewMarket{} }
func (m *NewMarket) String() string { return proto.CompactTextString(m) }
func (*NewMarket) ProtoMessage()    {}
func (*NewMarket) Descriptor() ([]byte, []int) {
	return fileDescriptor_5df90bf1b4327f51, []int{10}
}

func (m *NewMarket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewMarket.Unmarshal(m, b)
}
func (m *NewMarket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewMarket.Marshal(b, m, deterministic)
}
func (m *NewMarket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewMarket.Merge(m, src)
}
func (m *NewMarket) XXX_Size() int {
	return xxx_messageInfo_NewMarket.Size(m)
}
func (m *NewMarket) XXX_DiscardUnknown() {
	xxx_messageInfo_NewMarket.DiscardUnknown(m)
}

var xxx_messageInfo_NewMarket proto.InternalMessageInfo

func (m *NewMarket) GetTradeToken() []byte {
	if m != nil {
		return m.TradeToken
	}
	return nil
}

func (m *NewMarket) GetQuoteToken() []byte {
	if m != nil {
		return m.QuoteToken
	}
	return nil
}

func (m *NewMarket) GetTradeTokenSymbol() string {
	if m != nil {
		return m.TradeTokenSymbol
	}
	return ""
}

func (m *NewMarket) GetQuoteTokenSymbol() string {
	if m != nil {
		return m.QuoteTokenSymbol
	}
	return ""
}

func (m *NewMarket) GetTradeTokenDecimals() int32 {
	if m != nil {
		return m.TradeTokenDecimals
	}
	return 0
}

func (m *NewMarket) GetQuoteTokenDecimals() int32 {
	if m != nil {
		return m.QuoteTokenDecimals
	}
	return 0
}

func (m *NewMarket) GetCreator() []byte {
	if m != nil {
		return m.Creator
	}
	return nil
}

func init() {
	proto.RegisterType((*Order)(nil), "proto.Order")
	proto.RegisterType((*OrderTokenInfo)(nil), "proto.OrderTokenInfo")
	proto.RegisterType((*OrderInfo)(nil), "proto.OrderInfo")
	proto.RegisterType((*OrderUpdateInfo)(nil), "proto.OrderUpdateInfo")
	proto.RegisterType((*OrderNode)(nil), "proto.OrderNode")
	proto.RegisterType((*OrderListMeta)(nil), "proto.OrderListMeta")
	proto.RegisterType((*OrderBook)(nil), "proto.OrderBook")
	proto.RegisterType((*Transaction)(nil), "proto.Transaction")
	proto.RegisterType((*OrderFail)(nil), "proto.OrderFail")
	proto.RegisterType((*MarketInfo)(nil), "proto.MarketInfo")
	proto.RegisterType((*NewMarket)(nil), "proto.NewMarket")
}

func init() { proto.RegisterFile("dex_order.proto", fileDescriptor_5df90bf1b4327f51) }

var fileDescriptor_5df90bf1b4327f51 = []byte{
	// 728 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x54, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0x95, 0x9d, 0x5f, 0x4f, 0x42, 0xc8, 0x37, 0x1f, 0xdf, 0xa7, 0x11, 0xaa, 0xaa, 0xc8, 0x0b,
	0x14, 0x51, 0x29, 0x0b, 0xba, 0xee, 0x02, 0x28, 0xa8, 0x91, 0x12, 0x28, 0x26, 0xac, 0xab, 0x21,
	0x73, 0x69, 0xa3, 0x04, 0x0f, 0x1a, 0x8f, 0x0b, 0x59, 0x76, 0xd3, 0x7d, 0x1f, 0xa0, 0xef, 0xd0,
	0x77, 0xe8, 0x6b, 0xf4, 0x25, 0xfa, 0x06, 0xd5, 0xdc, 0x19, 0x27, 0xce, 0x9f, 0x04, 0x62, 0xd9,
	0x55, 0x7c, 0xce, 0x3d, 0xe3, 0x39, 0xbe, 0xf7, 0xe6, 0x90, 0x6d, 0x01, 0x0f, 0x1f, 0xa4, 0x12,
	0xa0, 0x3a, 0x77, 0x4a, 0x6a, 0x49, 0x4b, 0xf8, 0x13, 0xfe, 0x2c, 0x90, 0xd2, 0xb9, 0xa1, 0x69,
	0x83, 0xf8, 0x5d, 0xc1, 0xbc, 0x96, 0xd7, 0xae, 0x47, 0x7e, 0x57, 0x50, 0x46, 0x2a, 0x87, 0x42,
	0x28, 0x48, 0x12, 0xe6, 0x23, 0x99, 0x41, 0x4a, 0x49, 0xf1, 0x72, 0x24, 0x80, 0x55, 0x5a, 0x5e,
	0xbb, 0x1a, 0xe1, 0xb3, 0xe1, 0x06, 0xd3, 0x3b, 0x60, 0xd5, 0x96, 0xd7, 0x2e, 0x45, 0xf8, 0x4c,
	0x77, 0x48, 0xe9, 0xbd, 0x1a, 0x0d, 0x81, 0x05, 0x2d, 0xaf, 0x1d, 0x44, 0x16, 0xd0, 0x5d, 0x52,
	0xbd, 0x48, 0x79, 0xac, 0x47, 0x7a, 0xca, 0x08, 0xbe, 0x78, 0x86, 0xe9, 0xff, 0xa4, 0x7c, 0x78,
	0x2b, 0xd3, 0x58, 0xb3, 0x1a, 0x56, 0x1c, 0xa2, 0x21, 0xa9, 0xf7, 0xe4, 0x70, 0x0c, 0xe2, 0x28,
	0x9d, 0x9e, 0x02, 0xb0, 0x3a, 0x56, 0x17, 0x38, 0x73, 0xf6, 0x52, 0x73, 0x9d, 0x26, 0x6c, 0x0b,
	0x3d, 0x38, 0x64, 0xce, 0x1e, 0xf3, 0x78, 0x08, 0x93, 0x08, 0x78, 0x22, 0x63, 0xd6, 0xc0, 0xea,
	0x02, 0x47, 0xf7, 0x49, 0xf3, 0xe4, 0x01, 0x86, 0xa9, 0x06, 0x31, 0xf3, 0xd6, 0xc4, 0x3b, 0x56,
	0x78, 0xba, 0x47, 0x1a, 0x19, 0xe7, 0xbc, 0xfe, 0x83, 0xca, 0x25, 0x96, 0xb6, 0x48, 0x2d, 0x63,
	0x8c, 0x65, 0x8a, 0xa2, 0x3c, 0x65, 0x14, 0x11, 0xdc, 0xa4, 0xb1, 0x18, 0xc8, 0x31, 0xc4, 0xec,
	0x5f, 0xab, 0xc8, 0x51, 0xe6, 0x2e, 0x0b, 0x67, 0xae, 0x76, 0xec, 0x5d, 0x8b, 0x6c, 0xf8, 0xc3,
	0x23, 0x0d, 0x9c, 0x22, 0x1e, 0xeb, 0xc6, 0x37, 0x92, 0xbe, 0x24, 0x64, 0xa0, 0xb8, 0x00, 0xfb,
	0x6e, 0x3b, 0xd6, 0x1c, 0x63, 0xea, 0x17, 0xa9, 0xd4, 0xae, 0x6e, 0x27, 0x9c, 0x63, 0x68, 0x87,
	0xd0, 0xb9, 0xfa, 0x2d, 0x0c, 0x47, 0xb7, 0x7c, 0x92, 0xb0, 0x02, 0x36, 0x6f, 0x4d, 0xc5, 0xe8,
	0xe7, 0xa7, 0x67, 0xfa, 0xa2, 0xd5, 0xaf, 0x56, 0xc2, 0x98, 0x04, 0xe8, 0x18, 0xcd, 0x86, 0x6e,
	0x09, 0xd1, 0x67, 0xed, 0xa0, 0x6e, 0x77, 0xb4, 0x83, 0x5c, 0xe4, 0xf6, 0xf3, 0xcd, 0xf2, 0x27,
	0xa2, 0xe9, 0xda, 0xc1, 0x7f, 0x79, 0xf1, 0xac, 0x18, 0x2d, 0x89, 0xc3, 0x5f, 0x3e, 0xd9, 0x46,
	0xea, 0xea, 0x4e, 0x70, 0x0d, 0x78, 0xed, 0xf2, 0xca, 0xff, 0xc5, 0x2b, 0xf4, 0xdc, 0xf6, 0x7e,
	0xf1, 0xdc, 0x3c, 0xcf, 0xa4, 0x80, 0x47, 0xcd, 0x73, 0x8f, 0x34, 0x4e, 0xa5, 0xba, 0xe7, 0x4a,
	0x9c, 0xc7, 0x3d, 0xf8, 0x0c, 0x13, 0xe6, 0xb7, 0x0a, 0xc6, 0xd8, 0x22, 0x4b, 0xdb, 0x64, 0xfb,
	0x88, 0x0f, 0xc7, 0x79, 0x61, 0x01, 0x85, 0xcb, 0x74, 0xf8, 0xcd, 0x23, 0x5b, 0xf8, 0xee, 0xde,
	0x28, 0xd1, 0x7d, 0xd0, 0xdc, 0x0c, 0xf4, 0x1d, 0xf0, 0xcc, 0x48, 0x3d, 0x72, 0x08, 0xd3, 0x8a,
	0x8f, 0x26, 0x6e, 0xed, 0xf1, 0xd9, 0x68, 0x7b, 0x10, 0x7f, 0xd4, 0x9f, 0xdc, 0x92, 0x3b, 0x64,
	0x52, 0xcc, 0xde, 0x6a, 0x77, 0xd9, 0x82, 0x35, 0xee, 0x4b, 0xeb, 0xdc, 0x87, 0xaf, 0x5c, 0x5b,
	0x8e, 0xa4, 0x1c, 0x9b, 0xff, 0x5c, 0x9f, 0xab, 0x31, 0x68, 0x8c, 0x4f, 0xaf, 0x55, 0x68, 0x07,
	0x51, 0x8e, 0x09, 0x7f, 0x7b, 0xa4, 0x36, 0x50, 0x3c, 0x4e, 0xf8, 0x50, 0x8f, 0x64, 0xbc, 0xb2,
	0x9f, 0x2f, 0x48, 0x30, 0xe0, 0x63, 0x50, 0x78, 0xdc, 0xc7, 0xf4, 0x9d, 0x13, 0x26, 0xb0, 0x11,
	0x74, 0x05, 0x7e, 0x41, 0x3d, 0xca, 0xa0, 0xa9, 0xf4, 0x5d, 0xa5, 0x68, 0x2b, 0x0e, 0xce, 0x23,
	0xba, 0xb4, 0x29, 0xa2, 0xcb, 0x1b, 0x23, 0xba, 0xb2, 0x10, 0xd1, 0xbb, 0xa4, 0x8a, 0xd7, 0x99,
	0x45, 0xad, 0xda, 0x33, 0x19, 0x36, 0xb5, 0x7e, 0x56, 0x0b, 0x6c, 0x2d, 0xc3, 0xe1, 0x95, 0x6b,
	0xd0, 0xa9, 0x99, 0x41, 0x27, 0x17, 0x0a, 0x6e, 0x77, 0x9a, 0xf9, 0xdd, 0xc1, 0xd5, 0xcb, 0xe5,
	0x06, 0x23, 0x95, 0x13, 0xa5, 0x8e, 0xa5, 0x6b, 0x47, 0x10, 0x65, 0x30, 0xfc, 0xea, 0x65, 0xbd,
	0x46, 0xe1, 0xfa, 0x34, 0xf3, 0x9e, 0x98, 0x66, 0xfe, 0xa6, 0x34, 0x33, 0x46, 0x8e, 0x15, 0x70,
	0x2d, 0x55, 0xd6, 0x7b, 0x07, 0xc3, 0xef, 0x3e, 0x09, 0xce, 0xe0, 0xde, 0x7a, 0x79, 0x76, 0x2a,
	0xef, 0x93, 0xe6, 0x5c, 0x7d, 0x39, 0xbd, 0xbd, 0x96, 0x13, 0xbc, 0x30, 0x88, 0x56, 0x78, 0xa3,
	0x9d, 0x9f, 0x74, 0xda, 0xa2, 0xd5, 0x2e, 0xf3, 0x1b, 0xfa, 0x53, 0x7a, 0x62, 0x7f, 0xca, 0x8f,
	0xe9, 0x4f, 0x65, 0xa1, 0x3f, 0xd7, 0x65, 0x1c, 0xef, 0xeb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x75, 0x83, 0x4b, 0x64, 0xa1, 0x08, 0x00, 0x00,
}
