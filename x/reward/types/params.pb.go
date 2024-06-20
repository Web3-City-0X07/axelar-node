// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/reward/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params represent the genesis parameters for the module
type Params struct {
	ExternalChainVotingInflationRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=external_chain_voting_inflation_rate,json=externalChainVotingInflationRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"external_chain_voting_inflation_rate"`
	KeyMgmtRelativeInflationRate     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=key_mgmt_relative_inflation_rate,json=keyMgmtRelativeInflationRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"key_mgmt_relative_inflation_rate"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc8c8df034e5ffb0, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "axelar.reward.v1beta1.Params")
}

func init() {
	proto.RegisterFile("axelar/reward/v1beta1/params.proto", fileDescriptor_bc8c8df034e5ffb0)
}

var fileDescriptor_bc8c8df034e5ffb0 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x40, 0x93, 0x0e, 0x1d, 0xa2, 0x6f, 0xaa, 0x3e, 0x24, 0x84, 0x90, 0x5b, 0x55, 0x08, 0xb1,
	0xd4, 0x56, 0xd5, 0x37, 0x28, 0x2c, 0x0c, 0x48, 0xd0, 0x81, 0x81, 0xc5, 0x72, 0xd3, 0x8b, 0x1b,
	0x25, 0xb6, 0x23, 0xe7, 0x92, 0x26, 0x0b, 0xe2, 0x11, 0x78, 0xac, 0x8c, 0x1d, 0x11, 0x43, 0x05,
	0xc9, 0x8b, 0xa0, 0xfc, 0x49, 0xc0, 0xc8, 0x64, 0x5b, 0x3e, 0x3e, 0x47, 0xba, 0xf6, 0xa6, 0x22,
	0x83, 0x48, 0x58, 0x66, 0x61, 0x27, 0xec, 0x86, 0xa5, 0xf3, 0x35, 0xa0, 0x98, 0xb3, 0x58, 0x58,
	0xa1, 0x12, 0x1a, 0x5b, 0x83, 0x66, 0x74, 0xd4, 0x32, 0xb4, 0x65, 0x68, 0xc7, 0x9c, 0xfc, 0x97,
	0x46, 0x9a, 0x86, 0x60, 0xf5, 0xae, 0x85, 0xa7, 0x2f, 0x03, 0x6f, 0x78, 0xdb, 0xbc, 0x1e, 0x3d,
	0x7b, 0x67, 0x90, 0x21, 0x58, 0x2d, 0x22, 0xee, 0x6f, 0x45, 0xa0, 0x79, 0x6a, 0x30, 0xd0, 0x92,
	0x07, 0xfa, 0x31, 0x12, 0x18, 0x18, 0xcd, 0xad, 0x40, 0x38, 0x76, 0x27, 0xee, 0xc5, 0xbf, 0x25,
	0x2d, 0x0e, 0x63, 0xe7, 0xfd, 0x30, 0x3e, 0x97, 0x01, 0x6e, 0x9f, 0xd6, 0xd4, 0x37, 0x8a, 0xf9,
	0x26, 0x51, 0x26, 0xe9, 0x96, 0x59, 0xb2, 0x09, 0x19, 0xe6, 0x31, 0x24, 0xf4, 0x0a, 0xfc, 0xd5,
	0xa4, 0x77, 0x5f, 0xd6, 0xea, 0xfb, 0xc6, 0x7c, 0xdd, 0x8b, 0x57, 0x02, 0x61, 0x94, 0x7a, 0x93,
	0x10, 0x72, 0xae, 0xa4, 0x42, 0x6e, 0xa1, 0xbe, 0x48, 0xe1, 0x77, 0x7b, 0xf0, 0xa7, 0xf6, 0x69,
	0x08, 0xf9, 0x8d, 0x54, 0xb8, 0xea, 0xac, 0x3f, 0xba, 0xcb, 0xbb, 0xe2, 0x93, 0x38, 0x45, 0x49,
	0xdc, 0x7d, 0x49, 0xdc, 0x8f, 0x92, 0xb8, 0xaf, 0x15, 0x71, 0xf6, 0x15, 0x71, 0xde, 0x2a, 0xe2,
	0x3c, 0x2c, 0xbe, 0x35, 0xda, 0xc1, 0x6a, 0xc0, 0x9d, 0xb1, 0x61, 0x77, 0x9a, 0xf9, 0xc6, 0x02,
	0xcb, 0xfa, 0x1f, 0x69, 0xa2, 0xeb, 0x61, 0x33, 0xdc, 0xc5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5b, 0x16, 0x92, 0x11, 0xaf, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.KeyMgmtRelativeInflationRate.Size()
		i -= size
		if _, err := m.KeyMgmtRelativeInflationRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ExternalChainVotingInflationRate.Size()
		i -= size
		if _, err := m.ExternalChainVotingInflationRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ExternalChainVotingInflationRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.KeyMgmtRelativeInflationRate.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalChainVotingInflationRate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExternalChainVotingInflationRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyMgmtRelativeInflationRate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.KeyMgmtRelativeInflationRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
