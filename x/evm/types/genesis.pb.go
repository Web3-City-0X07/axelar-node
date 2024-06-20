// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/evm/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	utils "github.com/axelarnetwork/axelar-core/utils"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState represents the genesis state
type GenesisState struct {
	Chains []GenesisState_Chain `protobuf:"bytes,3,rep,name=chains,proto3" json:"chains"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3ecf743c731821, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

type GenesisState_Chain struct {
	Params                  Params                 `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	BurnerInfos             []BurnerInfo           `protobuf:"bytes,2,rep,name=burner_infos,json=burnerInfos,proto3" json:"burner_infos"`
	CommandQueue            utils.QueueState       `protobuf:"bytes,3,opt,name=command_queue,json=commandQueue,proto3" json:"command_queue"`
	ConfirmedDeposits       []ERC20Deposit         `protobuf:"bytes,4,rep,name=confirmed_deposits,json=confirmedDeposits,proto3" json:"confirmed_deposits"`
	BurnedDeposits          []ERC20Deposit         `protobuf:"bytes,5,rep,name=burned_deposits,json=burnedDeposits,proto3" json:"burned_deposits"`
	CommandBatches          []CommandBatchMetadata `protobuf:"bytes,8,rep,name=command_batches,json=commandBatches,proto3" json:"command_batches"`
	Gateway                 Gateway                `protobuf:"bytes,9,opt,name=gateway,proto3" json:"gateway"`
	Tokens                  []ERC20TokenMetadata   `protobuf:"bytes,10,rep,name=tokens,proto3" json:"tokens"`
	Events                  []Event                `protobuf:"bytes,11,rep,name=events,proto3" json:"events"`
	ConfirmedEventQueue     utils.QueueState       `protobuf:"bytes,12,opt,name=confirmed_event_queue,json=confirmedEventQueue,proto3" json:"confirmed_event_queue"`
	LegacyConfirmedDeposits []ERC20Deposit         `protobuf:"bytes,13,rep,name=legacy_confirmed_deposits,json=legacyConfirmedDeposits,proto3" json:"legacy_confirmed_deposits"`
	LegacyBurnedDeposits    []ERC20Deposit         `protobuf:"bytes,14,rep,name=legacy_burned_deposits,json=legacyBurnedDeposits,proto3" json:"legacy_burned_deposits"`
}

func (m *GenesisState_Chain) Reset()         { *m = GenesisState_Chain{} }
func (m *GenesisState_Chain) String() string { return proto.CompactTextString(m) }
func (*GenesisState_Chain) ProtoMessage()    {}
func (*GenesisState_Chain) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3ecf743c731821, []int{0, 0}
}
func (m *GenesisState_Chain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState_Chain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState_Chain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState_Chain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState_Chain.Merge(m, src)
}
func (m *GenesisState_Chain) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState_Chain) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState_Chain.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState_Chain proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "axelar.evm.v1beta1.GenesisState")
	proto.RegisterType((*GenesisState_Chain)(nil), "axelar.evm.v1beta1.GenesisState.Chain")
}

func init() { proto.RegisterFile("axelar/evm/v1beta1/genesis.proto", fileDescriptor_dd3ecf743c731821) }

var fileDescriptor_dd3ecf743c731821 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x63, 0x92, 0xa6, 0x65, 0x93, 0xb6, 0x62, 0x29, 0xe0, 0x06, 0xc9, 0x09, 0x1c, 0x50,
	0x2e, 0xd8, 0x6d, 0x38, 0x80, 0xe0, 0x96, 0x14, 0x55, 0x08, 0xf1, 0xaf, 0x80, 0x90, 0x2a, 0xa4,
	0x68, 0xed, 0x4c, 0x1c, 0xab, 0xb1, 0x37, 0x78, 0x37, 0x69, 0xf3, 0x14, 0x70, 0xe4, 0xca, 0x2b,
	0xf0, 0x14, 0x39, 0xf6, 0xc8, 0x09, 0x41, 0xf2, 0x22, 0xc8, 0xb3, 0x6b, 0x13, 0x5a, 0x73, 0xc8,
	0xcd, 0x9e, 0xfd, 0xbe, 0xdf, 0xcc, 0xee, 0x7e, 0x5a, 0xd2, 0x60, 0x67, 0x30, 0x64, 0xb1, 0x03,
	0x93, 0xd0, 0x99, 0xec, 0xbb, 0x20, 0xd9, 0xbe, 0xe3, 0x43, 0x04, 0x22, 0x10, 0xf6, 0x28, 0xe6,
	0x92, 0x53, 0xaa, 0x14, 0x36, 0x4c, 0x42, 0x5b, 0x2b, 0x6a, 0x77, 0xb4, 0x6b, 0x2c, 0x83, 0xa1,
	0xc8, 0x7c, 0x9f, 0xc6, 0x30, 0x86, 0x58, 0xd9, 0x6a, 0x3b, 0x3e, 0xf7, 0x39, 0x7e, 0x3a, 0xc9,
	0x97, 0xae, 0xd6, 0x73, 0xda, 0x8d, 0x58, 0xcc, 0x42, 0xdd, 0xad, 0x66, 0xe5, 0x08, 0xe4, 0x74,
	0x04, 0x7a, 0xfd, 0xee, 0xe7, 0x0d, 0x52, 0x3d, 0x54, 0xf3, 0xbd, 0x95, 0x4c, 0x02, 0x3d, 0x20,
	0x65, 0x6f, 0xc0, 0x82, 0x48, 0x98, 0xc5, 0x46, 0xb1, 0x59, 0x69, 0xdd, 0xb3, 0x2f, 0xcf, 0x6b,
	0x2f, 0x3b, 0xec, 0x4e, 0x22, 0x6f, 0x97, 0x66, 0x3f, 0xeb, 0x85, 0x23, 0xed, 0xad, 0x7d, 0x5f,
	0x27, 0x6b, 0x58, 0xa7, 0x8f, 0x48, 0x59, 0x0d, 0x64, 0x1a, 0x0d, 0xa3, 0x59, 0x69, 0xd5, 0xf2,
	0x78, 0xaf, 0x51, 0x91, 0x32, 0x94, 0x9e, 0x1e, 0x92, 0xaa, 0x3b, 0x8e, 0x23, 0x88, 0xbb, 0x41,
	0xd4, 0xe7, 0xc2, 0xbc, 0x82, 0xf3, 0x58, 0x79, 0xfe, 0x36, 0xea, 0x9e, 0x45, 0x7d, 0xae, 0x19,
	0x15, 0x37, 0xab, 0x08, 0xfa, 0x9c, 0x6c, 0x7a, 0x3c, 0x0c, 0x59, 0xd4, 0xeb, 0xe2, 0x91, 0x9a,
	0x45, 0x9c, 0xa4, 0x91, 0x92, 0xf0, 0xd4, 0x33, 0xd6, 0x9b, 0x44, 0x82, 0x3b, 0xd3, 0xac, 0xaa,
	0x36, 0xe3, 0x02, 0x7d, 0x4f, 0xa8, 0xc7, 0xa3, 0x7e, 0x10, 0x87, 0xd0, 0xeb, 0xf6, 0x60, 0xc4,
	0x45, 0x20, 0x85, 0x59, 0xc2, 0xd9, 0x1a, 0x79, 0xb3, 0x3d, 0x3d, 0xea, 0xb4, 0xf6, 0x0e, 0x94,
	0x50, 0x13, 0xaf, 0x65, 0x04, 0x5d, 0x17, 0xf4, 0x15, 0xd9, 0xc6, 0x91, 0x97, 0x98, 0x6b, 0x2b,
	0x31, 0xb7, 0x94, 0x3d, 0x03, 0x7e, 0x20, 0xdb, 0xe9, 0xa6, 0x5d, 0x26, 0xbd, 0x01, 0x08, 0x73,
	0x03, 0x81, 0xcd, 0x3c, 0x60, 0x47, 0x49, 0xdb, 0x89, 0xf2, 0x05, 0x48, 0xd6, 0x63, 0x92, 0xa5,
	0x60, 0x6f, 0x69, 0x0d, 0x04, 0x7d, 0x42, 0xd6, 0x7d, 0x26, 0xe1, 0x94, 0x4d, 0xcd, 0xab, 0x78,
	0x8e, 0xb7, 0x73, 0x13, 0xa2, 0x24, 0x9a, 0x91, 0x3a, 0x92, 0x74, 0x49, 0x7e, 0x02, 0x91, 0x30,
	0xc9, 0xff, 0xd3, 0x85, 0xbb, 0x7b, 0x97, 0xc8, 0x2e, 0x8c, 0xa2, 0xbd, 0xf4, 0x21, 0x29, 0xc3,
	0x04, 0x22, 0x29, 0xcc, 0x0a, 0x52, 0x76, 0x73, 0x29, 0x89, 0x22, 0x35, 0x2a, 0x39, 0x3d, 0x26,
	0x37, 0xfe, 0x5e, 0x1e, 0xd6, 0x74, 0x22, 0xaa, 0x2b, 0x25, 0xe2, 0x7a, 0x06, 0xc1, 0x26, 0x2a,
	0x18, 0x2e, 0xd9, 0x1d, 0x82, 0xcf, 0xbc, 0x69, 0x37, 0x27, 0x1f, 0x9b, 0x2b, 0xdd, 0xe5, 0x2d,
	0x05, 0xea, 0x5c, 0x4a, 0xc9, 0x47, 0x72, 0x53, 0xf7, 0xb8, 0x18, 0x96, 0xad, 0x95, 0x1a, 0xec,
	0x28, 0x4a, 0xfb, 0x9f, 0xc8, 0x3c, 0x2e, 0x7d, 0xfd, 0x56, 0x37, 0xda, 0x2f, 0x67, 0xbf, 0xad,
	0xc2, 0x6c, 0x6e, 0x19, 0xe7, 0x73, 0xcb, 0xf8, 0x35, 0xb7, 0x8c, 0x2f, 0x0b, 0xab, 0x70, 0xbe,
	0xb0, 0x0a, 0x3f, 0x16, 0x56, 0xe1, 0x78, 0xcf, 0x0f, 0xe4, 0x60, 0xec, 0xda, 0x1e, 0x0f, 0x1d,
	0xd5, 0x2b, 0x02, 0x79, 0xca, 0xe3, 0x13, 0xfd, 0x77, 0xdf, 0xe3, 0x31, 0x38, 0x67, 0xf8, 0xde,
	0xe0, 0x3b, 0xe3, 0x96, 0xf1, 0xa1, 0x79, 0xf0, 0x27, 0x00, 0x00, 0xff, 0xff, 0xba, 0xe6, 0x48,
	0xba, 0x1a, 0x05, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chains) > 0 {
		for iNdEx := len(m.Chains) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Chains[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState_Chain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState_Chain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState_Chain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LegacyBurnedDeposits) > 0 {
		for iNdEx := len(m.LegacyBurnedDeposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LegacyBurnedDeposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if len(m.LegacyConfirmedDeposits) > 0 {
		for iNdEx := len(m.LegacyConfirmedDeposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LegacyConfirmedDeposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	{
		size, err := m.ConfirmedEventQueue.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	{
		size, err := m.Gateway.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.CommandBatches) > 0 {
		for iNdEx := len(m.CommandBatches) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CommandBatches[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.BurnedDeposits) > 0 {
		for iNdEx := len(m.BurnedDeposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BurnedDeposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ConfirmedDeposits) > 0 {
		for iNdEx := len(m.ConfirmedDeposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ConfirmedDeposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size, err := m.CommandQueue.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.BurnerInfos) > 0 {
		for iNdEx := len(m.BurnerInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BurnerInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Chains) > 0 {
		for _, e := range m.Chains {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState_Chain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.BurnerInfos) > 0 {
		for _, e := range m.BurnerInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.CommandQueue.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ConfirmedDeposits) > 0 {
		for _, e := range m.ConfirmedDeposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BurnedDeposits) > 0 {
		for _, e := range m.BurnedDeposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CommandBatches) > 0 {
		for _, e := range m.CommandBatches {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Gateway.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Tokens) > 0 {
		for _, e := range m.Tokens {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.ConfirmedEventQueue.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.LegacyConfirmedDeposits) > 0 {
		for _, e := range m.LegacyConfirmedDeposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LegacyBurnedDeposits) > 0 {
		for _, e := range m.LegacyBurnedDeposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chains = append(m.Chains, GenesisState_Chain{})
			if err := m.Chains[len(m.Chains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState_Chain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Chain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Chain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnerInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurnerInfos = append(m.BurnerInfos, BurnerInfo{})
			if err := m.BurnerInfos[len(m.BurnerInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommandQueue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommandQueue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmedDeposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConfirmedDeposits = append(m.ConfirmedDeposits, ERC20Deposit{})
			if err := m.ConfirmedDeposits[len(m.ConfirmedDeposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnedDeposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurnedDeposits = append(m.BurnedDeposits, ERC20Deposit{})
			if err := m.BurnedDeposits[len(m.BurnedDeposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommandBatches", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommandBatches = append(m.CommandBatches, CommandBatchMetadata{})
			if err := m.CommandBatches[len(m.CommandBatches)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gateway", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Gateway.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, ERC20TokenMetadata{})
			if err := m.Tokens[len(m.Tokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmedEventQueue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ConfirmedEventQueue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LegacyConfirmedDeposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LegacyConfirmedDeposits = append(m.LegacyConfirmedDeposits, ERC20Deposit{})
			if err := m.LegacyConfirmedDeposits[len(m.LegacyConfirmedDeposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LegacyBurnedDeposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LegacyBurnedDeposits = append(m.LegacyBurnedDeposits, ERC20Deposit{})
			if err := m.LegacyBurnedDeposits[len(m.LegacyBurnedDeposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)