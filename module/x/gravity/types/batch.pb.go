// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gravity/v1/batch.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// OutgoingTxBatch represents a batch of transactions going from gravity to ETH
type OutgoingTxBatch struct {
	BatchNonce    uint64                `protobuf:"varint,1,opt,name=batch_nonce,json=batchNonce,proto3" json:"batch_nonce,omitempty"`
	BatchTimeout  uint64                `protobuf:"varint,2,opt,name=batch_timeout,json=batchTimeout,proto3" json:"batch_timeout,omitempty"`
	Transactions  []*OutgoingTransferTx `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
	TokenContract string                `protobuf:"bytes,4,opt,name=token_contract,json=tokenContract,proto3" json:"token_contract,omitempty"`
	Block         uint64                `protobuf:"varint,5,opt,name=block,proto3" json:"block,omitempty"`
}

func (m *OutgoingTxBatch) Reset()         { *m = OutgoingTxBatch{} }
func (m *OutgoingTxBatch) String() string { return proto.CompactTextString(m) }
func (*OutgoingTxBatch) ProtoMessage()    {}
func (*OutgoingTxBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_4453b445b0660cab, []int{0}
}
func (m *OutgoingTxBatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutgoingTxBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutgoingTxBatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutgoingTxBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingTxBatch.Merge(m, src)
}
func (m *OutgoingTxBatch) XXX_Size() int {
	return m.Size()
}
func (m *OutgoingTxBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingTxBatch.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingTxBatch proto.InternalMessageInfo

func (m *OutgoingTxBatch) GetBatchNonce() uint64 {
	if m != nil {
		return m.BatchNonce
	}
	return 0
}

func (m *OutgoingTxBatch) GetBatchTimeout() uint64 {
	if m != nil {
		return m.BatchTimeout
	}
	return 0
}

func (m *OutgoingTxBatch) GetTransactions() []*OutgoingTransferTx {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *OutgoingTxBatch) GetTokenContract() string {
	if m != nil {
		return m.TokenContract
	}
	return ""
}

func (m *OutgoingTxBatch) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

// OutgoingTransferTx represents an individual send from gravity to ETH
type OutgoingTransferTx struct {
	Id          uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender      string      `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	DestAddress string      `protobuf:"bytes,3,opt,name=dest_address,json=destAddress,proto3" json:"dest_address,omitempty"`
	Erc20Token  *ERC20Token `protobuf:"bytes,4,opt,name=erc20_token,json=erc20Token,proto3" json:"erc20_token,omitempty"`
	Erc20Fee    *ERC20Token `protobuf:"bytes,5,opt,name=erc20_fee,json=erc20Fee,proto3" json:"erc20_fee,omitempty"`
}

func (m *OutgoingTransferTx) Reset()         { *m = OutgoingTransferTx{} }
func (m *OutgoingTransferTx) String() string { return proto.CompactTextString(m) }
func (*OutgoingTransferTx) ProtoMessage()    {}
func (*OutgoingTransferTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_4453b445b0660cab, []int{1}
}
func (m *OutgoingTransferTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutgoingTransferTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutgoingTransferTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutgoingTransferTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingTransferTx.Merge(m, src)
}
func (m *OutgoingTransferTx) XXX_Size() int {
	return m.Size()
}
func (m *OutgoingTransferTx) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingTransferTx.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingTransferTx proto.InternalMessageInfo

func (m *OutgoingTransferTx) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OutgoingTransferTx) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *OutgoingTransferTx) GetDestAddress() string {
	if m != nil {
		return m.DestAddress
	}
	return ""
}

func (m *OutgoingTransferTx) GetErc20Token() *ERC20Token {
	if m != nil {
		return m.Erc20Token
	}
	return nil
}

func (m *OutgoingTransferTx) GetErc20Fee() *ERC20Token {
	if m != nil {
		return m.Erc20Fee
	}
	return nil
}

// OutgoingLogicCall represents an individual logic call from gravity to ETH
type OutgoingLogicCall struct {
	Transfers            []*ERC20Token `protobuf:"bytes,1,rep,name=transfers,proto3" json:"transfers,omitempty"`
	Fees                 []*ERC20Token `protobuf:"bytes,2,rep,name=fees,proto3" json:"fees,omitempty"`
	LogicContractAddress string        `protobuf:"bytes,3,opt,name=logic_contract_address,json=logicContractAddress,proto3" json:"logic_contract_address,omitempty"`
	Payload              []byte        `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	Timeout              uint64        `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
	InvalidationId       []byte        `protobuf:"bytes,6,opt,name=invalidation_id,json=invalidationId,proto3" json:"invalidation_id,omitempty"`
	InvalidationNonce    uint64        `protobuf:"varint,7,opt,name=invalidation_nonce,json=invalidationNonce,proto3" json:"invalidation_nonce,omitempty"`
	Block                uint64        `protobuf:"varint,8,opt,name=block,proto3" json:"block,omitempty"`
}

func (m *OutgoingLogicCall) Reset()         { *m = OutgoingLogicCall{} }
func (m *OutgoingLogicCall) String() string { return proto.CompactTextString(m) }
func (*OutgoingLogicCall) ProtoMessage()    {}
func (*OutgoingLogicCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_4453b445b0660cab, []int{2}
}
func (m *OutgoingLogicCall) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutgoingLogicCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutgoingLogicCall.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutgoingLogicCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingLogicCall.Merge(m, src)
}
func (m *OutgoingLogicCall) XXX_Size() int {
	return m.Size()
}
func (m *OutgoingLogicCall) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingLogicCall.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingLogicCall proto.InternalMessageInfo

func (m *OutgoingLogicCall) GetTransfers() []*ERC20Token {
	if m != nil {
		return m.Transfers
	}
	return nil
}

func (m *OutgoingLogicCall) GetFees() []*ERC20Token {
	if m != nil {
		return m.Fees
	}
	return nil
}

func (m *OutgoingLogicCall) GetLogicContractAddress() string {
	if m != nil {
		return m.LogicContractAddress
	}
	return ""
}

func (m *OutgoingLogicCall) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *OutgoingLogicCall) GetTimeout() uint64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *OutgoingLogicCall) GetInvalidationId() []byte {
	if m != nil {
		return m.InvalidationId
	}
	return nil
}

func (m *OutgoingLogicCall) GetInvalidationNonce() uint64 {
	if m != nil {
		return m.InvalidationNonce
	}
	return 0
}

func (m *OutgoingLogicCall) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func init() {
	proto.RegisterType((*OutgoingTxBatch)(nil), "gravity.v1.OutgoingTxBatch")
	proto.RegisterType((*OutgoingTransferTx)(nil), "gravity.v1.OutgoingTransferTx")
	proto.RegisterType((*OutgoingLogicCall)(nil), "gravity.v1.OutgoingLogicCall")
}

func init() { proto.RegisterFile("gravity/v1/batch.proto", fileDescriptor_4453b445b0660cab) }

var fileDescriptor_4453b445b0660cab = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x51, 0x6b, 0x1a, 0x41,
	0x10, 0xf6, 0x8c, 0x31, 0x71, 0x34, 0x86, 0x2c, 0x41, 0x8e, 0x52, 0xae, 0xd6, 0x52, 0x2a, 0x05,
	0xbd, 0xc4, 0x04, 0xfa, 0x5c, 0xa5, 0x85, 0x42, 0x69, 0xe1, 0xf0, 0xa9, 0x14, 0x64, 0xbd, 0x1d,
	0xcf, 0x25, 0xe7, 0xad, 0xdc, 0xae, 0xa2, 0xff, 0xa2, 0x3f, 0xab, 0x2f, 0x81, 0x3c, 0xe6, 0xb1,
	0xe8, 0x1f, 0x29, 0x37, 0x77, 0x67, 0x4c, 0x4b, 0xf2, 0x76, 0xf3, 0xcd, 0x37, 0x33, 0xdf, 0xcc,
	0x7e, 0x07, 0x8d, 0x20, 0xe6, 0x4b, 0x69, 0xd6, 0xee, 0xf2, 0xd2, 0x1d, 0x73, 0xe3, 0x4f, 0xbb,
	0xf3, 0x58, 0x19, 0xc5, 0x20, 0xc3, 0xbb, 0xcb, 0xcb, 0x17, 0x2f, 0xf7, 0x38, 0xdc, 0x18, 0xd4,
	0x86, 0x1b, 0xa9, 0xa2, 0x94, 0xd9, 0xba, 0xb7, 0xe0, 0xf4, 0xfb, 0xc2, 0x04, 0x4a, 0x46, 0xc1,
	0x70, 0xd5, 0x4f, 0x7a, 0xb0, 0x57, 0x50, 0xa5, 0x66, 0xa3, 0x48, 0x45, 0x3e, 0xda, 0x56, 0xd3,
	0x6a, 0x97, 0x3c, 0x20, 0xe8, 0x5b, 0x82, 0xb0, 0x37, 0x70, 0x92, 0x12, 0x8c, 0x9c, 0xa1, 0x5a,
	0x18, 0xbb, 0x48, 0x94, 0x1a, 0x81, 0xc3, 0x14, 0x63, 0x7d, 0xa8, 0x99, 0x98, 0x47, 0x9a, 0xfb,
	0xc9, 0x38, 0x6d, 0x1f, 0x34, 0x0f, 0xda, 0xd5, 0x9e, 0xd3, 0x7d, 0x90, 0xd6, 0xdd, 0x0d, 0x4e,
	0x78, 0x13, 0x8c, 0x87, 0x2b, 0xef, 0x51, 0x0d, 0x7b, 0x0b, 0x75, 0xa3, 0x6e, 0x30, 0x1a, 0xf9,
	0x2a, 0x32, 0x31, 0xf7, 0x8d, 0x5d, 0x6a, 0x5a, 0xed, 0x8a, 0x77, 0x42, 0xe8, 0x20, 0x03, 0xd9,
	0x39, 0x1c, 0x8e, 0x43, 0xe5, 0xdf, 0xd8, 0x87, 0xa4, 0x23, 0x0d, 0x5a, 0xb7, 0x16, 0xb0, 0xff,
	0x27, 0xb0, 0x3a, 0x14, 0xa5, 0xc8, 0x96, 0x2a, 0x4a, 0xc1, 0x1a, 0x50, 0xd6, 0x18, 0x09, 0x8c,
	0x69, 0x8b, 0x8a, 0x97, 0x45, 0xec, 0x35, 0xd4, 0x04, 0x6a, 0x33, 0xe2, 0x42, 0xc4, 0xa8, 0x13,
	0xfd, 0x49, 0xb6, 0x9a, 0x60, 0x1f, 0x53, 0x88, 0x7d, 0x80, 0x2a, 0xc6, 0x7e, 0xef, 0x62, 0x44,
	0x72, 0x48, 0x5b, 0xb5, 0xd7, 0xd8, 0xdf, 0xf0, 0x93, 0x37, 0xe8, 0x5d, 0x0c, 0x93, 0xac, 0x07,
	0x44, 0xa5, 0x6f, 0x76, 0x05, 0x95, 0xb4, 0x70, 0x82, 0x48, 0xa2, 0x9f, 0x2e, 0x3b, 0x26, 0xe2,
	0x67, 0xc4, 0xd6, 0x6d, 0x11, 0xce, 0xf2, 0x7d, 0xbe, 0xaa, 0x40, 0xfa, 0x03, 0x1e, 0x86, 0xec,
	0x1a, 0x2a, 0x26, 0x5b, 0x4e, 0xdb, 0x16, 0xdd, 0xf8, 0xa9, 0x56, 0x0f, 0x44, 0xf6, 0x1e, 0x4a,
	0x13, 0x44, 0x6d, 0x17, 0x9f, 0x2d, 0x20, 0x0e, 0xbb, 0x86, 0x46, 0x98, 0x8c, 0xdb, 0x3d, 0xc2,
	0x3f, 0x27, 0x39, 0xa7, 0x6c, 0xfe, 0x18, 0xf9, 0x6d, 0x6c, 0x38, 0x9a, 0xf3, 0x75, 0xa8, 0xb8,
	0xa0, 0xbb, 0xd4, 0xbc, 0x3c, 0x4c, 0x32, 0xb9, 0x6f, 0xd2, 0xf7, 0xca, 0x43, 0xf6, 0x0e, 0x4e,
	0x65, 0xb4, 0xe4, 0xa1, 0x14, 0x64, 0xd1, 0x91, 0x14, 0x76, 0x99, 0x6a, 0xeb, 0xfb, 0xf0, 0x17,
	0xc1, 0x3a, 0xc0, 0x1e, 0x11, 0x53, 0xa3, 0x1e, 0x51, 0xb7, 0xb3, 0xfd, 0x4c, 0xea, 0xd7, 0x9d,
	0x3f, 0x8e, 0xf7, 0xfc, 0xd1, 0xff, 0xf9, 0x7b, 0xe3, 0x58, 0x77, 0x1b, 0xc7, 0xfa, 0xb3, 0x71,
	0xac, 0x5f, 0x5b, 0xa7, 0x70, 0xb7, 0x75, 0x0a, 0xf7, 0x5b, 0xa7, 0xf0, 0xa3, 0x1f, 0x48, 0x33,
	0x5d, 0x8c, 0xbb, 0xbe, 0x9a, 0xb9, 0x3c, 0x34, 0x53, 0xe4, 0x9d, 0x08, 0x8d, 0xeb, 0x2b, 0x3d,
	0x53, 0xba, 0x93, 0xdd, 0xaa, 0x33, 0x8e, 0xa5, 0x08, 0xd0, 0x9d, 0x29, 0xb1, 0x08, 0xd1, 0x5d,
	0xb9, 0xf9, 0x7f, 0x66, 0xd6, 0x73, 0xd4, 0xe3, 0x32, 0xfd, 0x5f, 0x57, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xb6, 0x22, 0x58, 0x81, 0xa3, 0x03, 0x00, 0x00,
}

func (m *OutgoingTxBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutgoingTxBatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutgoingTxBatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x28
	}
	if len(m.TokenContract) > 0 {
		i -= len(m.TokenContract)
		copy(dAtA[i:], m.TokenContract)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.TokenContract)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Transactions) > 0 {
		for iNdEx := len(m.Transactions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Transactions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBatch(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.BatchTimeout != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.BatchTimeout))
		i--
		dAtA[i] = 0x10
	}
	if m.BatchNonce != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.BatchNonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OutgoingTransferTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutgoingTransferTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutgoingTransferTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Erc20Fee != nil {
		{
			size, err := m.Erc20Fee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBatch(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Erc20Token != nil {
		{
			size, err := m.Erc20Token.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBatch(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.DestAddress) > 0 {
		i -= len(m.DestAddress)
		copy(dAtA[i:], m.DestAddress)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.DestAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OutgoingLogicCall) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutgoingLogicCall) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutgoingLogicCall) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x40
	}
	if m.InvalidationNonce != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.InvalidationNonce))
		i--
		dAtA[i] = 0x38
	}
	if len(m.InvalidationId) > 0 {
		i -= len(m.InvalidationId)
		copy(dAtA[i:], m.InvalidationId)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.InvalidationId)))
		i--
		dAtA[i] = 0x32
	}
	if m.Timeout != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.LogicContractAddress) > 0 {
		i -= len(m.LogicContractAddress)
		copy(dAtA[i:], m.LogicContractAddress)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.LogicContractAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Fees) > 0 {
		for iNdEx := len(m.Fees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBatch(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Transfers) > 0 {
		for iNdEx := len(m.Transfers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Transfers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBatch(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintBatch(dAtA []byte, offset int, v uint64) int {
	offset -= sovBatch(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OutgoingTxBatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BatchNonce != 0 {
		n += 1 + sovBatch(uint64(m.BatchNonce))
	}
	if m.BatchTimeout != 0 {
		n += 1 + sovBatch(uint64(m.BatchTimeout))
	}
	if len(m.Transactions) > 0 {
		for _, e := range m.Transactions {
			l = e.Size()
			n += 1 + l + sovBatch(uint64(l))
		}
	}
	l = len(m.TokenContract)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovBatch(uint64(m.Block))
	}
	return n
}

func (m *OutgoingTransferTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovBatch(uint64(m.Id))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	l = len(m.DestAddress)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	if m.Erc20Token != nil {
		l = m.Erc20Token.Size()
		n += 1 + l + sovBatch(uint64(l))
	}
	if m.Erc20Fee != nil {
		l = m.Erc20Fee.Size()
		n += 1 + l + sovBatch(uint64(l))
	}
	return n
}

func (m *OutgoingLogicCall) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Transfers) > 0 {
		for _, e := range m.Transfers {
			l = e.Size()
			n += 1 + l + sovBatch(uint64(l))
		}
	}
	if len(m.Fees) > 0 {
		for _, e := range m.Fees {
			l = e.Size()
			n += 1 + l + sovBatch(uint64(l))
		}
	}
	l = len(m.LogicContractAddress)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovBatch(uint64(m.Timeout))
	}
	l = len(m.InvalidationId)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	if m.InvalidationNonce != 0 {
		n += 1 + sovBatch(uint64(m.InvalidationNonce))
	}
	if m.Block != 0 {
		n += 1 + sovBatch(uint64(m.Block))
	}
	return n
}

func sovBatch(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBatch(x uint64) (n int) {
	return sovBatch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OutgoingTxBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OutgoingTxBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutgoingTxBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchNonce", wireType)
			}
			m.BatchNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchTimeout", wireType)
			}
			m.BatchTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transactions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transactions = append(m.Transactions, &OutgoingTransferTx{})
			if err := m.Transactions[len(m.Transactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenContract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenContract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBatch
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OutgoingTransferTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OutgoingTransferTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutgoingTransferTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Erc20Token == nil {
				m.Erc20Token = &ERC20Token{}
			}
			if err := m.Erc20Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Erc20Fee == nil {
				m.Erc20Fee = &ERC20Token{}
			}
			if err := m.Erc20Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBatch
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OutgoingLogicCall) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OutgoingLogicCall: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutgoingLogicCall: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transfers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transfers = append(m.Transfers, &ERC20Token{})
			if err := m.Transfers[len(m.Transfers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fees = append(m.Fees, &ERC20Token{})
			if err := m.Fees[len(m.Fees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogicContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogicContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvalidationId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InvalidationId = append(m.InvalidationId[:0], dAtA[iNdEx:postIndex]...)
			if m.InvalidationId == nil {
				m.InvalidationId = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvalidationNonce", wireType)
			}
			m.InvalidationNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InvalidationNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBatch
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBatch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBatch
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthBatch
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBatch
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBatch
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBatch        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBatch          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBatch = fmt.Errorf("proto: unexpected end of group")
)
