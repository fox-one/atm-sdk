package atm

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Side int32

const (
	Side_Ask Side = 0
	Side_Bid Side = 1
)

var Side_name = map[int32]string{
	0: "Ask",
	1: "Bid",
}

var Side_value = map[string]int32{
	"Ask": 0,
	"Bid": 1,
}

func (x Side) String() string {
	return proto.EnumName(Side_name, int32(x))
}

func (Side) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

type Strategy int32

const (
	Strategy_Market Strategy = 0
	Strategy_Limit  Strategy = 1
	Strategy_Follow Strategy = 2
	Strategy_Flex   Strategy = 3
)

var Strategy_name = map[int32]string{
	0: "Market",
	1: "Limit",
	2: "Follow",
	3: "Flex",
}

var Strategy_value = map[string]int32{
	"Market": 0,
	"Limit":  1,
	"Follow": 2,
	"Flex":   3,
}

func (x Strategy) String() string {
	return proto.EnumName(Strategy_name, int32(x))
}

func (Strategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

type SortOrder int32

const (
	SortOrder_DESC SortOrder = 0
	SortOrder_ASC  SortOrder = 1
)

var SortOrder_name = map[int32]string{
	0: "DESC",
	1: "ASC",
}

var SortOrder_value = map[string]int32{
	"DESC": 0,
	"ASC":  1,
}

func (x SortOrder) String() string {
	return proto.EnumName(SortOrder_name, int32(x))
}

func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

type Book_State int32

const (
	Book_Pending Book_State = 0
	Book_Paid    Book_State = 1
	Book_Done    Book_State = 2
)

var Book_State_name = map[int32]string{
	0: "Pending",
	1: "Paid",
	2: "Done",
}

var Book_State_value = map[string]int32{
	"Pending": 0,
	"Paid":    1,
	"Done":    2,
}

func (x Book_State) String() string {
	return proto.EnumName(Book_State_name, int32(x))
}

func (Book_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1, 0}
}

type Order_State int32

const (
	Order_Trading   Order_State = 0
	Order_Filled    Order_State = 1
	Order_Cancelled Order_State = 2
	Order_Rejected  Order_State = 3
	Order_Timeout   Order_State = 4
)

var Order_State_name = map[int32]string{
	0: "Trading",
	1: "Filled",
	2: "Cancelled",
	3: "Rejected",
	4: "Timeout",
}

var Order_State_value = map[string]int32{
	"Trading":   0,
	"Filled":    1,
	"Cancelled": 2,
	"Rejected":  3,
	"Timeout":   4,
}

func (x Order_State) String() string {
	return proto.EnumName(Order_State_name, int32(x))
}

func (Order_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2, 0}
}

type Pagination struct {
	NextCursor           string   `protobuf:"bytes,1,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
	HasNext              bool     `protobuf:"varint,2,opt,name=has_next,json=hasNext,proto3" json:"has_next,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pagination.Unmarshal(m, b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
}
func (m *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(m, src)
}
func (m *Pagination) XXX_Size() int {
	return xxx_messageInfo_Pagination.Size(m)
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetNextCursor() string {
	if m != nil {
		return m.NextCursor
	}
	return ""
}

func (m *Pagination) GetHasNext() bool {
	if m != nil {
		return m.HasNext
	}
	return false
}

type Book struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	State                Book_State           `protobuf:"varint,3,opt,name=state,proto3,enum=fox.atm.service.Book_State" json:"state,omitempty"`
	MerchantId           string               `protobuf:"bytes,4,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	BrokerId             string               `protobuf:"bytes,5,opt,name=broker_id,json=brokerId,proto3" json:"broker_id,omitempty"`
	TraceId              string               `protobuf:"bytes,6,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	UserId               string               `protobuf:"bytes,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Payer                string               `protobuf:"bytes,8,opt,name=payer,proto3" json:"payer,omitempty"`
	SnapshotId           string               `protobuf:"bytes,9,opt,name=snapshot_id,json=snapshotId,proto3" json:"snapshot_id,omitempty"`
	ReceiptId            string               `protobuf:"bytes,10,opt,name=receipt_id,json=receiptId,proto3" json:"receipt_id,omitempty"`
	Memo                 string               `protobuf:"bytes,11,opt,name=memo,proto3" json:"memo,omitempty"`
	Funds                string               `protobuf:"bytes,12,opt,name=funds,proto3" json:"funds,omitempty"`
	PaySymbol            string               `protobuf:"bytes,13,opt,name=pay_symbol,json=paySymbol,proto3" json:"pay_symbol,omitempty"`
	FillSymbol           string               `protobuf:"bytes,14,opt,name=fill_symbol,json=fillSymbol,proto3" json:"fill_symbol,omitempty"`
	Strategy             Strategy             `protobuf:"varint,15,opt,name=strategy,proto3,enum=fox.atm.service.Strategy" json:"strategy,omitempty"`
	Price                string               `protobuf:"bytes,16,opt,name=price,proto3" json:"price,omitempty"`
	Discount             string               `protobuf:"bytes,17,opt,name=discount,proto3" json:"discount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Book) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Book) GetState() Book_State {
	if m != nil {
		return m.State
	}
	return Book_Pending
}

func (m *Book) GetMerchantId() string {
	if m != nil {
		return m.MerchantId
	}
	return ""
}

func (m *Book) GetBrokerId() string {
	if m != nil {
		return m.BrokerId
	}
	return ""
}

func (m *Book) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Book) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Book) GetPayer() string {
	if m != nil {
		return m.Payer
	}
	return ""
}

func (m *Book) GetSnapshotId() string {
	if m != nil {
		return m.SnapshotId
	}
	return ""
}

func (m *Book) GetReceiptId() string {
	if m != nil {
		return m.ReceiptId
	}
	return ""
}

func (m *Book) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Book) GetFunds() string {
	if m != nil {
		return m.Funds
	}
	return ""
}

func (m *Book) GetPaySymbol() string {
	if m != nil {
		return m.PaySymbol
	}
	return ""
}

func (m *Book) GetFillSymbol() string {
	if m != nil {
		return m.FillSymbol
	}
	return ""
}

func (m *Book) GetStrategy() Strategy {
	if m != nil {
		return m.Strategy
	}
	return Strategy_Market
}

func (m *Book) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Book) GetDiscount() string {
	if m != nil {
		return m.Discount
	}
	return ""
}

type Order struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CancelledAt          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=cancelled_at,json=cancelledAt,proto3" json:"cancelled_at,omitempty"`
	UserId               string               `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MerchantId           string               `protobuf:"bytes,6,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	State                Order_State          `protobuf:"varint,7,opt,name=state,proto3,enum=fox.atm.service.Order_State" json:"state,omitempty"`
	PaySymbol            string               `protobuf:"bytes,8,opt,name=pay_symbol,json=paySymbol,proto3" json:"pay_symbol,omitempty"`
	FillSymbol           string               `protobuf:"bytes,9,opt,name=fill_symbol,json=fillSymbol,proto3" json:"fill_symbol,omitempty"`
	Symbol               string               `protobuf:"bytes,10,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Side                 Side                 `protobuf:"varint,11,opt,name=side,proto3,enum=fox.atm.service.Side" json:"side,omitempty"`
	Strategy             Strategy             `protobuf:"varint,12,opt,name=strategy,proto3,enum=fox.atm.service.Strategy" json:"strategy,omitempty"`
	Price                string               `protobuf:"bytes,13,opt,name=price,proto3" json:"price,omitempty"`
	Discount             string               `protobuf:"bytes,14,opt,name=discount,proto3" json:"discount,omitempty"`
	Funds                string               `protobuf:"bytes,15,opt,name=funds,proto3" json:"funds,omitempty"`
	FilledFunds          string               `protobuf:"bytes,16,opt,name=filled_funds,json=filledFunds,proto3" json:"filled_funds,omitempty"`
	FilledAmount         string               `protobuf:"bytes,17,opt,name=filled_amount,json=filledAmount,proto3" json:"filled_amount,omitempty"`
	ExtraFilledAmount    string               `protobuf:"bytes,18,opt,name=extra_filled_amount,json=extraFilledAmount,proto3" json:"extra_filled_amount,omitempty"`
	FeeAmount            string               `protobuf:"bytes,19,opt,name=fee_amount,json=feeAmount,proto3" json:"fee_amount,omitempty"`
	AveragePrice         string               `protobuf:"bytes,20,opt,name=average_price,json=averagePrice,proto3" json:"average_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
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

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Order) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Order) GetCancelledAt() *timestamp.Timestamp {
	if m != nil {
		return m.CancelledAt
	}
	return nil
}

func (m *Order) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Order) GetMerchantId() string {
	if m != nil {
		return m.MerchantId
	}
	return ""
}

func (m *Order) GetState() Order_State {
	if m != nil {
		return m.State
	}
	return Order_Trading
}

func (m *Order) GetPaySymbol() string {
	if m != nil {
		return m.PaySymbol
	}
	return ""
}

func (m *Order) GetFillSymbol() string {
	if m != nil {
		return m.FillSymbol
	}
	return ""
}

func (m *Order) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Order) GetSide() Side {
	if m != nil {
		return m.Side
	}
	return Side_Ask
}

func (m *Order) GetStrategy() Strategy {
	if m != nil {
		return m.Strategy
	}
	return Strategy_Market
}

func (m *Order) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Order) GetDiscount() string {
	if m != nil {
		return m.Discount
	}
	return ""
}

func (m *Order) GetFunds() string {
	if m != nil {
		return m.Funds
	}
	return ""
}

func (m *Order) GetFilledFunds() string {
	if m != nil {
		return m.FilledFunds
	}
	return ""
}

func (m *Order) GetFilledAmount() string {
	if m != nil {
		return m.FilledAmount
	}
	return ""
}

func (m *Order) GetExtraFilledAmount() string {
	if m != nil {
		return m.ExtraFilledAmount
	}
	return ""
}

func (m *Order) GetFeeAmount() string {
	if m != nil {
		return m.FeeAmount
	}
	return ""
}

func (m *Order) GetAveragePrice() string {
	if m != nil {
		return m.AveragePrice
	}
	return ""
}

type OrderReport struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Date                 string               `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	UserId               string               `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MerchantId           string               `protobuf:"bytes,5,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	Symbol               string               `protobuf:"bytes,6,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Side                 string               `protobuf:"bytes,7,opt,name=side,proto3" json:"side,omitempty"`
	FilledAmount         string               `protobuf:"bytes,8,opt,name=filled_amount,json=filledAmount,proto3" json:"filled_amount,omitempty"`
	ObtainedAmount       string               `protobuf:"bytes,9,opt,name=obtained_amount,json=obtainedAmount,proto3" json:"obtained_amount,omitempty"`
	FeeAmount            string               `protobuf:"bytes,10,opt,name=fee_amount,json=feeAmount,proto3" json:"fee_amount,omitempty"`
	FeeAsset             string               `protobuf:"bytes,11,opt,name=fee_asset,json=feeAsset,proto3" json:"fee_asset,omitempty"`
	Count                int32                `protobuf:"varint,12,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OrderReport) Reset()         { *m = OrderReport{} }
func (m *OrderReport) String() string { return proto.CompactTextString(m) }
func (*OrderReport) ProtoMessage()    {}
func (*OrderReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *OrderReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderReport.Unmarshal(m, b)
}
func (m *OrderReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderReport.Marshal(b, m, deterministic)
}
func (m *OrderReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderReport.Merge(m, src)
}
func (m *OrderReport) XXX_Size() int {
	return xxx_messageInfo_OrderReport.Size(m)
}
func (m *OrderReport) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderReport.DiscardUnknown(m)
}

var xxx_messageInfo_OrderReport proto.InternalMessageInfo

func (m *OrderReport) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrderReport) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *OrderReport) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *OrderReport) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *OrderReport) GetMerchantId() string {
	if m != nil {
		return m.MerchantId
	}
	return ""
}

func (m *OrderReport) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *OrderReport) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *OrderReport) GetFilledAmount() string {
	if m != nil {
		return m.FilledAmount
	}
	return ""
}

func (m *OrderReport) GetObtainedAmount() string {
	if m != nil {
		return m.ObtainedAmount
	}
	return ""
}

func (m *OrderReport) GetFeeAmount() string {
	if m != nil {
		return m.FeeAmount
	}
	return ""
}

func (m *OrderReport) GetFeeAsset() string {
	if m != nil {
		return m.FeeAsset
	}
	return ""
}

func (m *OrderReport) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Role                 string   `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	BrokerId             string   `protobuf:"bytes,4,opt,name=broker_id,json=brokerId,proto3" json:"broker_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *User) GetBrokerId() string {
	if m != nil {
		return m.BrokerId
	}
	return ""
}

type UserServiceReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceReq) Reset()         { *m = UserServiceReq{} }
func (m *UserServiceReq) String() string { return proto.CompactTextString(m) }
func (*UserServiceReq) ProtoMessage()    {}
func (*UserServiceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *UserServiceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceReq.Unmarshal(m, b)
}
func (m *UserServiceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceReq.Marshal(b, m, deterministic)
}
func (m *UserServiceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceReq.Merge(m, src)
}
func (m *UserServiceReq) XXX_Size() int {
	return xxx_messageInfo_UserServiceReq.Size(m)
}
func (m *UserServiceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceReq proto.InternalMessageInfo

type UserServiceReq_Me struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceReq_Me) Reset()         { *m = UserServiceReq_Me{} }
func (m *UserServiceReq_Me) String() string { return proto.CompactTextString(m) }
func (*UserServiceReq_Me) ProtoMessage()    {}
func (*UserServiceReq_Me) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5, 0}
}

func (m *UserServiceReq_Me) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceReq_Me.Unmarshal(m, b)
}
func (m *UserServiceReq_Me) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceReq_Me.Marshal(b, m, deterministic)
}
func (m *UserServiceReq_Me) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceReq_Me.Merge(m, src)
}
func (m *UserServiceReq_Me) XXX_Size() int {
	return xxx_messageInfo_UserServiceReq_Me.Size(m)
}
func (m *UserServiceReq_Me) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceReq_Me.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceReq_Me proto.InternalMessageInfo

type MerchantServiceReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantServiceReq) Reset()         { *m = MerchantServiceReq{} }
func (m *MerchantServiceReq) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceReq) ProtoMessage()    {}
func (*MerchantServiceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *MerchantServiceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceReq.Unmarshal(m, b)
}
func (m *MerchantServiceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceReq.Marshal(b, m, deterministic)
}
func (m *MerchantServiceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceReq.Merge(m, src)
}
func (m *MerchantServiceReq) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceReq.Size(m)
}
func (m *MerchantServiceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceReq.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceReq proto.InternalMessageInfo

type MerchantServiceReq_CreateBook struct {
	TraceId              string   `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ReceiptId            string   `protobuf:"bytes,3,opt,name=receipt_id,json=receiptId,proto3" json:"receipt_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantServiceReq_CreateBook) Reset()         { *m = MerchantServiceReq_CreateBook{} }
func (m *MerchantServiceReq_CreateBook) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceReq_CreateBook) ProtoMessage()    {}
func (*MerchantServiceReq_CreateBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6, 0}
}

func (m *MerchantServiceReq_CreateBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceReq_CreateBook.Unmarshal(m, b)
}
func (m *MerchantServiceReq_CreateBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceReq_CreateBook.Marshal(b, m, deterministic)
}
func (m *MerchantServiceReq_CreateBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceReq_CreateBook.Merge(m, src)
}
func (m *MerchantServiceReq_CreateBook) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceReq_CreateBook.Size(m)
}
func (m *MerchantServiceReq_CreateBook) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceReq_CreateBook.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceReq_CreateBook proto.InternalMessageInfo

func (m *MerchantServiceReq_CreateBook) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *MerchantServiceReq_CreateBook) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *MerchantServiceReq_CreateBook) GetReceiptId() string {
	if m != nil {
		return m.ReceiptId
	}
	return ""
}

type MerchantServiceReq_ReadOrder struct {
	TraceId              string   `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantServiceReq_ReadOrder) Reset()         { *m = MerchantServiceReq_ReadOrder{} }
func (m *MerchantServiceReq_ReadOrder) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceReq_ReadOrder) ProtoMessage()    {}
func (*MerchantServiceReq_ReadOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6, 1}
}

func (m *MerchantServiceReq_ReadOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceReq_ReadOrder.Unmarshal(m, b)
}
func (m *MerchantServiceReq_ReadOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceReq_ReadOrder.Marshal(b, m, deterministic)
}
func (m *MerchantServiceReq_ReadOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceReq_ReadOrder.Merge(m, src)
}
func (m *MerchantServiceReq_ReadOrder) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceReq_ReadOrder.Size(m)
}
func (m *MerchantServiceReq_ReadOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceReq_ReadOrder.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceReq_ReadOrder proto.InternalMessageInfo

func (m *MerchantServiceReq_ReadOrder) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

type MerchantServiceReq_ListOrders struct {
	Symbol               string    `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Side                 string    `protobuf:"bytes,2,opt,name=side,proto3" json:"side,omitempty"`
	Strategy             string    `protobuf:"bytes,3,opt,name=strategy,proto3" json:"strategy,omitempty"`
	State                string    `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	UserId               string    `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Order                SortOrder `protobuf:"varint,6,opt,name=order,proto3,enum=fox.atm.service.SortOrder" json:"order,omitempty"`
	Cursor               string    `protobuf:"bytes,7,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit                int64     `protobuf:"varint,8,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MerchantServiceReq_ListOrders) Reset()         { *m = MerchantServiceReq_ListOrders{} }
func (m *MerchantServiceReq_ListOrders) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceReq_ListOrders) ProtoMessage()    {}
func (*MerchantServiceReq_ListOrders) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6, 2}
}

func (m *MerchantServiceReq_ListOrders) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceReq_ListOrders.Unmarshal(m, b)
}
func (m *MerchantServiceReq_ListOrders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceReq_ListOrders.Marshal(b, m, deterministic)
}
func (m *MerchantServiceReq_ListOrders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceReq_ListOrders.Merge(m, src)
}
func (m *MerchantServiceReq_ListOrders) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceReq_ListOrders.Size(m)
}
func (m *MerchantServiceReq_ListOrders) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceReq_ListOrders.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceReq_ListOrders proto.InternalMessageInfo

func (m *MerchantServiceReq_ListOrders) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *MerchantServiceReq_ListOrders) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *MerchantServiceReq_ListOrders) GetStrategy() string {
	if m != nil {
		return m.Strategy
	}
	return ""
}

func (m *MerchantServiceReq_ListOrders) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *MerchantServiceReq_ListOrders) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *MerchantServiceReq_ListOrders) GetOrder() SortOrder {
	if m != nil {
		return m.Order
	}
	return SortOrder_DESC
}

func (m *MerchantServiceReq_ListOrders) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func (m *MerchantServiceReq_ListOrders) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type MerchantServiceReq_CancelOrder struct {
	TraceId              string   `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantServiceReq_CancelOrder) Reset()         { *m = MerchantServiceReq_CancelOrder{} }
func (m *MerchantServiceReq_CancelOrder) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceReq_CancelOrder) ProtoMessage()    {}
func (*MerchantServiceReq_CancelOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6, 3}
}

func (m *MerchantServiceReq_CancelOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceReq_CancelOrder.Unmarshal(m, b)
}
func (m *MerchantServiceReq_CancelOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceReq_CancelOrder.Marshal(b, m, deterministic)
}
func (m *MerchantServiceReq_CancelOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceReq_CancelOrder.Merge(m, src)
}
func (m *MerchantServiceReq_CancelOrder) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceReq_CancelOrder.Size(m)
}
func (m *MerchantServiceReq_CancelOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceReq_CancelOrder.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceReq_CancelOrder proto.InternalMessageInfo

func (m *MerchantServiceReq_CancelOrder) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

type MerchantServiceReq_ListOrderReports struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Cursor               string   `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantServiceReq_ListOrderReports) Reset()         { *m = MerchantServiceReq_ListOrderReports{} }
func (m *MerchantServiceReq_ListOrderReports) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceReq_ListOrderReports) ProtoMessage()    {}
func (*MerchantServiceReq_ListOrderReports) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6, 4}
}

func (m *MerchantServiceReq_ListOrderReports) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceReq_ListOrderReports.Unmarshal(m, b)
}
func (m *MerchantServiceReq_ListOrderReports) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceReq_ListOrderReports.Marshal(b, m, deterministic)
}
func (m *MerchantServiceReq_ListOrderReports) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceReq_ListOrderReports.Merge(m, src)
}
func (m *MerchantServiceReq_ListOrderReports) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceReq_ListOrderReports.Size(m)
}
func (m *MerchantServiceReq_ListOrderReports) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceReq_ListOrderReports.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceReq_ListOrderReports proto.InternalMessageInfo

func (m *MerchantServiceReq_ListOrderReports) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *MerchantServiceReq_ListOrderReports) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func (m *MerchantServiceReq_ListOrderReports) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type MerchantServiceResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantServiceResp) Reset()         { *m = MerchantServiceResp{} }
func (m *MerchantServiceResp) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceResp) ProtoMessage()    {}
func (*MerchantServiceResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *MerchantServiceResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceResp.Unmarshal(m, b)
}
func (m *MerchantServiceResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceResp.Marshal(b, m, deterministic)
}
func (m *MerchantServiceResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceResp.Merge(m, src)
}
func (m *MerchantServiceResp) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceResp.Size(m)
}
func (m *MerchantServiceResp) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceResp.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceResp proto.InternalMessageInfo

type MerchantServiceResp_CancelOrder struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantServiceResp_CancelOrder) Reset()         { *m = MerchantServiceResp_CancelOrder{} }
func (m *MerchantServiceResp_CancelOrder) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceResp_CancelOrder) ProtoMessage()    {}
func (*MerchantServiceResp_CancelOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7, 0}
}

func (m *MerchantServiceResp_CancelOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceResp_CancelOrder.Unmarshal(m, b)
}
func (m *MerchantServiceResp_CancelOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceResp_CancelOrder.Marshal(b, m, deterministic)
}
func (m *MerchantServiceResp_CancelOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceResp_CancelOrder.Merge(m, src)
}
func (m *MerchantServiceResp_CancelOrder) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceResp_CancelOrder.Size(m)
}
func (m *MerchantServiceResp_CancelOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceResp_CancelOrder.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceResp_CancelOrder proto.InternalMessageInfo

type MerchantServiceResp_ListOrders struct {
	Orders               []*Order    `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	Pagination           *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MerchantServiceResp_ListOrders) Reset()         { *m = MerchantServiceResp_ListOrders{} }
func (m *MerchantServiceResp_ListOrders) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceResp_ListOrders) ProtoMessage()    {}
func (*MerchantServiceResp_ListOrders) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7, 1}
}

func (m *MerchantServiceResp_ListOrders) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceResp_ListOrders.Unmarshal(m, b)
}
func (m *MerchantServiceResp_ListOrders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceResp_ListOrders.Marshal(b, m, deterministic)
}
func (m *MerchantServiceResp_ListOrders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceResp_ListOrders.Merge(m, src)
}
func (m *MerchantServiceResp_ListOrders) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceResp_ListOrders.Size(m)
}
func (m *MerchantServiceResp_ListOrders) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceResp_ListOrders.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceResp_ListOrders proto.InternalMessageInfo

func (m *MerchantServiceResp_ListOrders) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func (m *MerchantServiceResp_ListOrders) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type MerchantServiceResp_ListOrderReports struct {
	Reports              []*OrderReport `protobuf:"bytes,1,rep,name=reports,proto3" json:"reports,omitempty"`
	Pagination           *Pagination    `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MerchantServiceResp_ListOrderReports) Reset()         { *m = MerchantServiceResp_ListOrderReports{} }
func (m *MerchantServiceResp_ListOrderReports) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceResp_ListOrderReports) ProtoMessage()    {}
func (*MerchantServiceResp_ListOrderReports) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7, 2}
}

func (m *MerchantServiceResp_ListOrderReports) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceResp_ListOrderReports.Unmarshal(m, b)
}
func (m *MerchantServiceResp_ListOrderReports) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceResp_ListOrderReports.Marshal(b, m, deterministic)
}
func (m *MerchantServiceResp_ListOrderReports) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceResp_ListOrderReports.Merge(m, src)
}
func (m *MerchantServiceResp_ListOrderReports) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceResp_ListOrderReports.Size(m)
}
func (m *MerchantServiceResp_ListOrderReports) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceResp_ListOrderReports.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceResp_ListOrderReports proto.InternalMessageInfo

func (m *MerchantServiceResp_ListOrderReports) GetReports() []*OrderReport {
	if m != nil {
		return m.Reports
	}
	return nil
}

func (m *MerchantServiceResp_ListOrderReports) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterEnum("fox.atm.service.Side", Side_name, Side_value)
	proto.RegisterEnum("fox.atm.service.Strategy", Strategy_name, Strategy_value)
	proto.RegisterEnum("fox.atm.service.SortOrder", SortOrder_name, SortOrder_value)
	proto.RegisterEnum("fox.atm.service.Book_State", Book_State_name, Book_State_value)
	proto.RegisterEnum("fox.atm.service.Order_State", Order_State_name, Order_State_value)
	proto.RegisterType((*Pagination)(nil), "fox.atm.service.Pagination")
	proto.RegisterType((*Book)(nil), "fox.atm.service.Book")
	proto.RegisterType((*Order)(nil), "fox.atm.service.Order")
	proto.RegisterType((*OrderReport)(nil), "fox.atm.service.OrderReport")
	proto.RegisterType((*User)(nil), "fox.atm.service.User")
	proto.RegisterType((*UserServiceReq)(nil), "fox.atm.service.UserServiceReq")
	proto.RegisterType((*UserServiceReq_Me)(nil), "fox.atm.service.UserServiceReq.Me")
	proto.RegisterType((*MerchantServiceReq)(nil), "fox.atm.service.MerchantServiceReq")
	proto.RegisterType((*MerchantServiceReq_CreateBook)(nil), "fox.atm.service.MerchantServiceReq.CreateBook")
	proto.RegisterType((*MerchantServiceReq_ReadOrder)(nil), "fox.atm.service.MerchantServiceReq.ReadOrder")
	proto.RegisterType((*MerchantServiceReq_ListOrders)(nil), "fox.atm.service.MerchantServiceReq.ListOrders")
	proto.RegisterType((*MerchantServiceReq_CancelOrder)(nil), "fox.atm.service.MerchantServiceReq.CancelOrder")
	proto.RegisterType((*MerchantServiceReq_ListOrderReports)(nil), "fox.atm.service.MerchantServiceReq.ListOrderReports")
	proto.RegisterType((*MerchantServiceResp)(nil), "fox.atm.service.MerchantServiceResp")
	proto.RegisterType((*MerchantServiceResp_CancelOrder)(nil), "fox.atm.service.MerchantServiceResp.CancelOrder")
	proto.RegisterType((*MerchantServiceResp_ListOrders)(nil), "fox.atm.service.MerchantServiceResp.ListOrders")
	proto.RegisterType((*MerchantServiceResp_ListOrderReports)(nil), "fox.atm.service.MerchantServiceResp.ListOrderReports")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 1261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcd, 0x92, 0xdb, 0x44,
	0x10, 0x8e, 0x64, 0xd9, 0x96, 0xda, 0x6b, 0xaf, 0x32, 0xf9, 0x41, 0x51, 0x08, 0x59, 0x9c, 0x2a,
	0x30, 0xa9, 0x42, 0x1b, 0x0c, 0x81, 0x4a, 0x51, 0x14, 0xe5, 0x2c, 0xa4, 0x30, 0x95, 0x0d, 0x5b,
	0xf2, 0x72, 0x81, 0x83, 0x6b, 0xd6, 0x1a, 0x7b, 0xc5, 0x5a, 0x3f, 0x68, 0xc6, 0x61, 0x7d, 0xe2,
	0xc0, 0x81, 0x97, 0xe0, 0x05, 0x78, 0x0f, 0x2e, 0x3c, 0x07, 0x6f, 0xc0, 0x13, 0x50, 0xf3, 0x23,
	0x5b, 0x5a, 0xd9, 0xec, 0x06, 0x72, 0x92, 0xa6, 0xfb, 0xeb, 0xe9, 0x9e, 0xe9, 0xaf, 0xbb, 0x07,
	0xda, 0x94, 0x64, 0x2f, 0xc3, 0x09, 0xf1, 0xd2, 0x2c, 0x61, 0x09, 0xda, 0x9d, 0x26, 0xe7, 0x1e,
	0x66, 0x91, 0xa7, 0xc4, 0xee, 0xfd, 0x59, 0x92, 0xcc, 0xe6, 0x64, 0x5f, 0xa8, 0x4f, 0x16, 0xd3,
	0x7d, 0x16, 0x46, 0x84, 0x32, 0x1c, 0xa5, 0xd2, 0xa2, 0xfb, 0x15, 0xc0, 0x11, 0x9e, 0x85, 0x31,
	0x66, 0x61, 0x12, 0xa3, 0xfb, 0xd0, 0x8a, 0xc9, 0x39, 0x1b, 0x4f, 0x16, 0x19, 0x4d, 0x32, 0x47,
	0xdb, 0xd3, 0x7a, 0x96, 0x0f, 0x5c, 0x74, 0x20, 0x24, 0xe8, 0x0e, 0x98, 0xa7, 0x98, 0x8e, 0xb9,
	0xc4, 0xd1, 0xf7, 0xb4, 0x9e, 0xe9, 0x37, 0x4f, 0x31, 0x7d, 0x41, 0xce, 0x59, 0xf7, 0x4f, 0x03,
	0x8c, 0xa7, 0x49, 0x72, 0x86, 0x3a, 0xa0, 0x87, 0x81, 0xb2, 0xd5, 0xc3, 0x00, 0x3d, 0x01, 0x98,
	0x64, 0x04, 0x33, 0x12, 0x8c, 0xb1, 0xb4, 0x6a, 0xf5, 0x5d, 0x4f, 0x06, 0xe6, 0xe5, 0x81, 0x79,
	0xc7, 0x79, 0x60, 0xbe, 0xa5, 0xd0, 0x03, 0x86, 0x3e, 0x80, 0x3a, 0x65, 0x98, 0x11, 0xa7, 0xb6,
	0xa7, 0xf5, 0x3a, 0xfd, 0xbb, 0xde, 0x85, 0xf3, 0x79, 0xdc, 0xa1, 0x37, 0xe2, 0x10, 0x5f, 0x22,
	0xf9, 0x11, 0x22, 0x92, 0x4d, 0x4e, 0x71, 0xcc, 0xc6, 0x61, 0xe0, 0x18, 0xf2, 0x08, 0xb9, 0x68,
	0x18, 0xa0, 0xbb, 0x60, 0x9d, 0x64, 0xc9, 0x19, 0xc9, 0xb8, 0xba, 0x2e, 0xd4, 0xa6, 0x14, 0x0c,
	0x03, 0x7e, 0x3e, 0x96, 0xe1, 0x09, 0xe1, 0xba, 0x86, 0xd0, 0x35, 0xc5, 0x7a, 0x18, 0xa0, 0x37,
	0xa0, 0xb9, 0xa0, 0xd2, 0xaa, 0x29, 0x34, 0x0d, 0xbe, 0x1c, 0x06, 0xe8, 0x26, 0xd4, 0x53, 0xbc,
	0x24, 0x99, 0x63, 0x0a, 0xb1, 0x5c, 0xf0, 0x38, 0x68, 0x8c, 0x53, 0x7a, 0x9a, 0x88, 0x38, 0x2c,
	0x19, 0x47, 0x2e, 0x1a, 0x06, 0xe8, 0x1e, 0x40, 0x46, 0x26, 0x24, 0x4c, 0x85, 0x1e, 0x84, 0xde,
	0x52, 0x92, 0x61, 0x80, 0x10, 0x18, 0x11, 0x89, 0x12, 0xa7, 0x25, 0x14, 0xe2, 0x9f, 0x7b, 0x9a,
	0x2e, 0xe2, 0x80, 0x3a, 0x3b, 0xd2, 0x93, 0x58, 0xf0, 0x8d, 0x52, 0xbc, 0x1c, 0xd3, 0x65, 0x74,
	0x92, 0xcc, 0x9d, 0xb6, 0xdc, 0x28, 0xc5, 0xcb, 0x91, 0x10, 0xf0, 0x40, 0xa6, 0xe1, 0x7c, 0x9e,
	0xeb, 0x3b, 0x32, 0x10, 0x2e, 0x52, 0x80, 0xc7, 0x60, 0x52, 0x96, 0x61, 0x46, 0x66, 0x4b, 0x67,
	0x57, 0xdc, 0xf3, 0x9d, 0xca, 0x3d, 0x8f, 0x14, 0xc0, 0x5f, 0x41, 0xc5, 0xb1, 0xb3, 0x70, 0x42,
	0x1c, 0x5b, 0x1d, 0x9b, 0x2f, 0x90, 0x0b, 0x66, 0x10, 0xd2, 0x49, 0xb2, 0x88, 0x99, 0x73, 0x5d,
	0x5e, 0x6e, 0xbe, 0xee, 0xf6, 0xa0, 0x2e, 0x52, 0x85, 0x5a, 0xd0, 0x3c, 0x22, 0x71, 0x10, 0xc6,
	0x33, 0xfb, 0x1a, 0x32, 0xc1, 0x38, 0xc2, 0x61, 0x60, 0x6b, 0xfc, 0xef, 0x8b, 0x24, 0x26, 0xb6,
	0xde, 0xfd, 0xa3, 0x01, 0xf5, 0x6f, 0xb2, 0x80, 0x64, 0xaf, 0x93, 0x4c, 0x4f, 0x00, 0x16, 0x69,
	0x90, 0x9b, 0xd6, 0x2e, 0x37, 0x55, 0xe8, 0x01, 0x43, 0x9f, 0xc1, 0xce, 0x04, 0xc7, 0x13, 0x32,
	0x9f, 0x4b, 0x63, 0xe3, 0x52, 0xe3, 0xd6, 0x0a, 0x3f, 0x60, 0x45, 0xea, 0xd4, 0x4b, 0xd4, 0xb9,
	0x40, 0xd6, 0x46, 0x85, 0xac, 0xfd, 0xbc, 0x00, 0x9a, 0x22, 0x31, 0x6f, 0x56, 0x12, 0x23, 0x6e,
	0xa9, 0x5c, 0x01, 0x65, 0x3e, 0x98, 0x97, 0xf0, 0xc1, 0xaa, 0xf0, 0xe1, 0x36, 0x34, 0x94, 0x4e,
	0x92, 0x52, 0xad, 0xd0, 0x7b, 0x60, 0xd0, 0x30, 0x20, 0x82, 0x91, 0x9d, 0xfe, 0xad, 0x2a, 0x47,
	0xc2, 0x80, 0xf8, 0x02, 0x52, 0xa2, 0xd4, 0xce, 0x7f, 0xa0, 0x54, 0x7b, 0x1b, 0xa5, 0x3a, 0x65,
	0x4a, 0xad, 0x2b, 0x62, 0xb7, 0x58, 0x11, 0x6f, 0xc3, 0x0e, 0x3f, 0x0f, 0x09, 0xc6, 0x52, 0x29,
	0x19, 0xda, 0x92, 0xb2, 0x67, 0x02, 0xf2, 0x00, 0xda, 0x0a, 0x82, 0xa3, 0x02, 0x59, 0x95, 0xdd,
	0x40, 0xc8, 0x90, 0x07, 0x37, 0xc8, 0x39, 0xcb, 0xf0, 0xb8, 0x0c, 0x45, 0x02, 0x7a, 0x5d, 0xa8,
	0x9e, 0x15, 0xf1, 0xf7, 0x00, 0xa6, 0x84, 0xe4, 0xb0, 0x1b, 0xf2, 0xe6, 0xa7, 0x84, 0x28, 0xf5,
	0x03, 0x68, 0xe3, 0x97, 0x24, 0xc3, 0x33, 0x32, 0x96, 0xc7, 0xbc, 0x29, 0x7d, 0x2a, 0xe1, 0x11,
	0x97, 0x75, 0xbf, 0x2e, 0x14, 0xc9, 0x71, 0x86, 0x55, 0x91, 0x00, 0x34, 0xa4, 0x27, 0x5b, 0x43,
	0x6d, 0xb0, 0x0e, 0x72, 0x72, 0xd9, 0x3a, 0xda, 0x01, 0xd3, 0x27, 0x3f, 0x90, 0x09, 0x23, 0x81,
	0x5d, 0x13, 0x56, 0x61, 0x44, 0x92, 0x05, 0xb3, 0x8d, 0xee, 0xdf, 0x3a, 0xb4, 0x04, 0x41, 0x7c,
	0x92, 0x26, 0x19, 0x7b, 0x9d, 0xc5, 0x84, 0xc0, 0x08, 0xf2, 0xc6, 0x6c, 0xf9, 0xe2, 0xbf, 0x48,
	0x73, 0xe3, 0xdf, 0x68, 0x5e, 0xaf, 0xd0, 0x7c, 0x4d, 0xb9, 0x46, 0x89, 0x72, 0x48, 0x51, 0x4e,
	0x36, 0x5c, 0xc9, 0xad, 0x4a, 0xe6, 0xcc, 0x0d, 0x99, 0x7b, 0x17, 0x76, 0x93, 0x13, 0x86, 0xc3,
	0x78, 0x0d, 0x93, 0x44, 0xef, 0xe4, 0xe2, 0x8d, 0x29, 0x83, 0x8b, 0x29, 0xbb, 0x0b, 0x96, 0x50,
	0x53, 0x4a, 0x98, 0x6a, 0xc5, 0x26, 0xd7, 0xf2, 0x35, 0x27, 0x9f, 0x64, 0x25, 0xa7, 0x78, 0xdd,
	0x97, 0x8b, 0xee, 0xf7, 0x60, 0x7c, 0x4b, 0x37, 0x74, 0x2e, 0x04, 0x46, 0x8c, 0x23, 0x22, 0xae,
	0xd9, 0xf2, 0xc5, 0x3f, 0x97, 0x65, 0xc9, 0x7c, 0x75, 0x8b, 0xfc, 0xbf, 0x3c, 0x9f, 0x8c, 0xf2,
	0x7c, 0xea, 0xde, 0x86, 0x0e, 0xdf, 0x7c, 0x24, 0x6b, 0xc8, 0x27, 0x3f, 0xba, 0x06, 0xe8, 0x87,
	0xa4, 0xfb, 0x8b, 0x01, 0xe8, 0x50, 0xdd, 0x67, 0x41, 0x39, 0x06, 0x38, 0x10, 0x29, 0x13, 0x83,
	0xb9, 0x38, 0xdc, 0xb4, 0xad, 0xc3, 0x4d, 0x2f, 0xa5, 0xae, 0x3c, 0xa5, 0x6a, 0x17, 0xa6, 0x94,
	0xfb, 0x0e, 0x58, 0x3e, 0xc1, 0x81, 0xec, 0xd5, 0xdb, 0xf7, 0x77, 0xff, 0xd2, 0x00, 0x9e, 0x87,
	0x94, 0x09, 0x20, 0x2d, 0xe4, 0x5b, 0xdb, 0x98, 0x6f, 0xbd, 0x90, 0x6f, 0xb7, 0xd0, 0x4b, 0xa4,
	0xff, 0x52, 0xc3, 0x90, 0xed, 0x51, 0xde, 0x93, 0x6a, 0x80, 0x5b, 0xdb, 0xed, 0x23, 0xa8, 0x27,
	0x3c, 0x00, 0xc1, 0xb2, 0x4e, 0xdf, 0xad, 0xf6, 0xa4, 0x24, 0x93, 0x21, 0xfa, 0x12, 0xc8, 0x03,
	0x55, 0x6f, 0x21, 0x35, 0xf3, 0xe5, 0x8a, 0x3b, 0x9e, 0x87, 0x51, 0x28, 0xc9, 0x57, 0xf3, 0xe5,
	0xc2, 0xed, 0x41, 0x4b, 0x56, 0xe6, 0xa5, 0xf7, 0x71, 0x0c, 0xf6, 0xea, 0x3a, 0x64, 0x71, 0xd2,
	0x55, 0x49, 0x69, 0x85, 0x92, 0x5a, 0xfb, 0xd7, 0x37, 0xfb, 0xaf, 0x15, 0xfc, 0x77, 0x7f, 0xd7,
	0xe1, 0x46, 0x85, 0x05, 0x34, 0x75, 0xdb, 0xa5, 0xb8, 0xdc, 0x65, 0x29, 0x17, 0x1e, 0x34, 0xc4,
	0x59, 0xa9, 0xa3, 0xed, 0xd5, 0x7a, 0xad, 0xfe, 0xed, 0xcd, 0x33, 0xc6, 0x57, 0x28, 0xf4, 0x29,
	0x1f, 0x2f, 0xf9, 0x8b, 0x51, 0x35, 0x8d, 0xea, 0xc3, 0x6c, 0xfd, 0xa8, 0xf4, 0x0b, 0x70, 0xf7,
	0x57, 0x6d, 0xc3, 0xc1, 0x3f, 0x86, 0x66, 0x26, 0x7f, 0x55, 0x08, 0x5b, 0xc6, 0x9c, 0xc4, 0xfb,
	0x39, 0xf8, 0x7f, 0x45, 0xf2, 0xd0, 0x01, 0x83, 0x0f, 0x2c, 0xd4, 0x84, 0xda, 0x80, 0x9e, 0xd9,
	0xd7, 0xf8, 0xcf, 0x53, 0xfe, 0x0c, 0x79, 0xf8, 0x09, 0x98, 0xf9, 0x6c, 0xe2, 0x7d, 0xf7, 0x10,
	0x67, 0x67, 0x84, 0xd9, 0xd7, 0x90, 0x05, 0xf5, 0xe7, 0xfc, 0x9a, 0x6d, 0x4d, 0xb4, 0xe3, 0x64,
	0x3e, 0x4f, 0x7e, 0xb2, 0x75, 0xfe, 0x6a, 0x79, 0x36, 0x27, 0xe7, 0x76, 0xed, 0xe1, 0x5b, 0x60,
	0xad, 0x08, 0x24, 0x1e, 0x33, 0x5f, 0x8e, 0x0e, 0xe4, 0xc6, 0x83, 0xd1, 0x81, 0xad, 0xf5, 0x5f,
	0x40, 0xab, 0x50, 0xbc, 0xe8, 0x73, 0x5e, 0xb9, 0xa8, 0x5b, 0x09, 0xb8, 0x5c, 0xe0, 0xde, 0x21,
	0x71, 0x6f, 0x6d, 0xc4, 0xf4, 0x7f, 0x33, 0x60, 0xf7, 0x42, 0xba, 0xd1, 0xa8, 0x54, 0xf1, 0x5e,
	0xc5, 0xb0, 0xda, 0x24, 0xbc, 0x35, 0x7e, 0x83, 0x23, 0xb1, 0x8d, 0x5f, 0xac, 0xf2, 0xf7, 0xaf,
	0xb2, 0xe7, 0x0a, 0xee, 0x6e, 0xa1, 0x13, 0x8a, 0xca, 0x24, 0xbc, 0xca, 0xa6, 0x6b, 0xbc, 0xbb,
	0x7f, 0x39, 0x9e, 0xa6, 0x05, 0x03, 0xf4, 0xf3, 0x06, 0xde, 0x7d, 0xf4, 0x4a, 0x4e, 0x95, 0x95,
	0xfb, 0xf8, 0xd5, 0x5c, 0xe7, 0xce, 0xd2, 0x72, 0x6f, 0xd8, 0xbf, 0x52, 0x66, 0x0a, 0x45, 0xfb,
	0xe8, 0x4a, 0x6e, 0x0b, 0x16, 0x4f, 0x9b, 0xdf, 0xd5, 0xe5, 0x0c, 0x6f, 0x88, 0xcf, 0x87, 0xff,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xba, 0x1a, 0xea, 0x2d, 0x0e, 0x00, 0x00,
}
