// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/axelarnet/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	utils "github.com/axelarnetwork/axelar-core/utils"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
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

type GenesisState struct {
	Params           Params                                        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	CollectorAddress github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=collector_address,json=collectorAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"collector_address,omitempty"`
	Chains           []CosmosChain                                 `protobuf:"bytes,3,rep,name=chains,proto3" json:"chains"`
	TransferQueue    utils.QueueState                              `protobuf:"bytes,5,opt,name=transfer_queue,json=transferQueue,proto3" json:"transfer_queue"`
	IBCTransfers     []IBCTransfer                                 `protobuf:"bytes,7,rep,name=ibc_transfers,json=ibcTransfers,proto3" json:"ibc_transfers"`
	SeqIDMapping     map[string]uint64                             `protobuf:"bytes,8,rep,name=seq_id_mapping,json=seqIdMapping,proto3" json:"seq_id_mapping" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_66d88d9f35b1824e, []int{0}
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

func init() {
	proto.RegisterType((*GenesisState)(nil), "axelar.axelarnet.v1beta1.GenesisState")
	proto.RegisterMapType((map[string]uint64)(nil), "axelar.axelarnet.v1beta1.GenesisState.SeqIdMappingEntry")
}

func init() {
	proto.RegisterFile("axelar/axelarnet/v1beta1/genesis.proto", fileDescriptor_66d88d9f35b1824e)
}

var fileDescriptor_66d88d9f35b1824e = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x3d, 0x6f, 0xd3, 0x40,
	0x18, 0xc7, 0xe3, 0xc6, 0x49, 0xc3, 0x35, 0xad, 0x52, 0x2b, 0x83, 0x95, 0xc1, 0x09, 0x88, 0xa2,
	0x2c, 0xb1, 0x95, 0x22, 0xa1, 0xaa, 0x03, 0xa8, 0x09, 0x08, 0xb5, 0x52, 0x25, 0x70, 0x3b, 0x31,
	0x10, 0xce, 0xe7, 0xc3, 0xb5, 0xe2, 0xf8, 0x9c, 0xbb, 0x73, 0x69, 0xbe, 0x05, 0x23, 0x2b, 0xdf,
	0x26, 0x63, 0x47, 0xa6, 0x08, 0x92, 0x99, 0x2f, 0xc0, 0x84, 0x7c, 0x77, 0x31, 0x46, 0xc8, 0x62,
	0xba, 0xb7, 0xdf, 0xff, 0xff, 0xbc, 0xdc, 0x03, 0x9e, 0xc0, 0x3b, 0x1c, 0x41, 0xea, 0xc8, 0x25,
	0xc6, 0xdc, 0xb9, 0x1d, 0x7a, 0x98, 0xc3, 0xa1, 0x13, 0xe0, 0x18, 0xb3, 0x90, 0xd9, 0x09, 0x25,
	0x9c, 0x18, 0xa6, 0x04, 0xec, 0x9c, 0xb3, 0x15, 0xd7, 0x69, 0x07, 0x24, 0x20, 0x02, 0x72, 0xb2,
	0x9d, 0xe4, 0x3b, 0x47, 0xa5, 0xbe, 0x09, 0xa4, 0x70, 0xa6, 0x6c, 0x3b, 0x8f, 0x4b, 0x31, 0xbe,
	0x48, 0xf0, 0x96, 0x7a, 0xa8, 0xa8, 0x94, 0x87, 0x11, 0xcb, 0x89, 0x79, 0x8a, 0x53, 0x4c, 0x25,
	0xf2, 0xe8, 0xa7, 0x0e, 0x9a, 0xaf, 0x65, 0xc6, 0x57, 0x1c, 0x72, 0x6c, 0x3c, 0x07, 0x75, 0x19,
	0xc9, 0xd4, 0x7a, 0x5a, 0x7f, 0xef, 0xb8, 0x67, 0x97, 0x55, 0x60, 0xbf, 0x11, 0xdc, 0x48, 0x5f,
	0xae, 0xba, 0x15, 0x57, 0xa9, 0x8c, 0xf7, 0xe0, 0x10, 0x91, 0x28, 0xc2, 0x88, 0x13, 0x3a, 0x81,
	0xbe, 0x4f, 0x31, 0x63, 0xe6, 0x4e, 0x4f, 0xeb, 0x37, 0x47, 0xc3, 0x5f, 0xab, 0xee, 0x20, 0x08,
	0xf9, 0x4d, 0xea, 0xd9, 0x88, 0xcc, 0x1c, 0x44, 0xd8, 0x8c, 0x30, 0xb5, 0x0c, 0x98, 0x3f, 0x55,
	0xc9, 0x9f, 0x21, 0x74, 0x26, 0x85, 0x6e, 0x2b, 0xf7, 0x52, 0x37, 0xc6, 0x18, 0xd4, 0xd1, 0x0d,
	0x0c, 0x63, 0x66, 0x56, 0x7b, 0xd5, 0xfe, 0xde, 0xf1, 0x51, 0x79, 0x7e, 0x63, 0x61, 0x3c, 0xce,
	0xe8, 0x6d, 0x92, 0x52, 0x6a, 0x5c, 0x82, 0x03, 0x4e, 0x61, 0xcc, 0x3e, 0x62, 0x3a, 0x11, 0xed,
	0x30, 0x6b, 0x7f, 0x17, 0x2b, 0x3a, 0x96, 0x1b, 0xbd, 0xcd, 0x10, 0xd1, 0x1e, 0xe5, 0xb3, 0xbf,
	0x55, 0x8b, 0x17, 0xe3, 0x03, 0xd8, 0x0f, 0x3d, 0x34, 0xd9, 0x5e, 0x32, 0x73, 0xf7, 0x7f, 0xa9,
	0x9d, 0x8f, 0xc6, 0xd7, 0x8a, 0x1e, 0xb5, 0x33, 0xcb, 0xf5, 0xaa, 0xdb, 0x2c, 0x5c, 0x32, 0xb7,
	0x19, 0x7a, 0x28, 0x3f, 0x19, 0x1c, 0x1c, 0x30, 0x3c, 0x9f, 0x84, 0xfe, 0x64, 0x06, 0x93, 0x24,
	0x8c, 0x03, 0xb3, 0x21, 0x42, 0x9c, 0x94, 0x87, 0x28, 0xfe, 0xaa, 0x7d, 0x85, 0xe7, 0xe7, 0xfe,
	0xa5, 0x94, 0xbe, 0x8a, 0x39, 0x5d, 0xfc, 0x89, 0x9a, 0x3d, 0xbd, 0x54, 0x4f, 0x6e, 0x93, 0x15,
	0xc0, 0xce, 0x0b, 0x70, 0xf8, 0x8f, 0xd0, 0x68, 0x81, 0xea, 0x14, 0x2f, 0xc4, 0x74, 0x3c, 0x70,
	0xb3, 0xad, 0xd1, 0x06, 0xb5, 0x5b, 0x18, 0xa5, 0x58, 0x7c, 0xb3, 0xee, 0xca, 0xc3, 0xe9, 0xce,
	0x89, 0x76, 0xaa, 0x7f, 0xf9, 0xda, 0xd5, 0x2e, 0xf4, 0x86, 0xde, 0xaa, 0x5d, 0xe8, 0x8d, 0x7a,
	0x6b, 0x77, 0x74, 0xbd, 0xfc, 0x61, 0x55, 0x96, 0x6b, 0x4b, 0xbb, 0x5f, 0x5b, 0xda, 0xf7, 0xb5,
	0xa5, 0x7d, 0xde, 0x58, 0x95, 0xfb, 0x8d, 0x55, 0xf9, 0xb6, 0xb1, 0x2a, 0xef, 0x9e, 0x15, 0xa6,
	0x23, 0xaf, 0xe8, 0x13, 0xa1, 0x53, 0x75, 0x1a, 0x20, 0x42, 0xb1, 0x73, 0x57, 0x18, 0x7b, 0x31,
	0x31, 0x5e, 0x5d, 0x0c, 0xf3, 0xd3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x91, 0xe2, 0x82,
	0x96, 0x03, 0x00, 0x00,
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
	if len(m.SeqIDMapping) > 0 {
		keysForSeqIDMapping := make([]string, 0, len(m.SeqIDMapping))
		for k := range m.SeqIDMapping {
			keysForSeqIDMapping = append(keysForSeqIDMapping, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForSeqIDMapping)
		for iNdEx := len(keysForSeqIDMapping) - 1; iNdEx >= 0; iNdEx-- {
			v := m.SeqIDMapping[string(keysForSeqIDMapping[iNdEx])]
			baseI := i
			i = encodeVarintGenesis(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i -= len(keysForSeqIDMapping[iNdEx])
			copy(dAtA[i:], keysForSeqIDMapping[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(keysForSeqIDMapping[iNdEx])))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.IBCTransfers) > 0 {
		for iNdEx := len(m.IBCTransfers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IBCTransfers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	{
		size, err := m.TransferQueue.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
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
	if len(m.CollectorAddress) > 0 {
		i -= len(m.CollectorAddress)
		copy(dAtA[i:], m.CollectorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.CollectorAddress)))
		i--
		dAtA[i] = 0x12
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.CollectorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Chains) > 0 {
		for _, e := range m.Chains {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.TransferQueue.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.IBCTransfers) > 0 {
		for _, e := range m.IBCTransfers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SeqIDMapping) > 0 {
		for k, v := range m.SeqIDMapping {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenesis(uint64(len(k))) + 1 + sovGenesis(uint64(v))
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
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
				return fmt.Errorf("proto: wrong wireType = %d for field CollectorAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectorAddress = append(m.CollectorAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.CollectorAddress == nil {
				m.CollectorAddress = []byte{}
			}
			iNdEx = postIndex
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
			m.Chains = append(m.Chains, CosmosChain{})
			if err := m.Chains[len(m.Chains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferQueue", wireType)
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
			if err := m.TransferQueue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IBCTransfers", wireType)
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
			m.IBCTransfers = append(m.IBCTransfers, IBCTransfer{})
			if err := m.IBCTransfers[len(m.IBCTransfers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeqIDMapping", wireType)
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
			if m.SeqIDMapping == nil {
				m.SeqIDMapping = make(map[string]uint64)
			}
			var mapkey string
			var mapvalue uint64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.SeqIDMapping[mapkey] = mapvalue
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
