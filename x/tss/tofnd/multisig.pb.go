// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/tss/tofnd/v1beta1/multisig.proto

package tofnd

import (
	fmt "fmt"
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

type KeygenRequest struct {
	KeyUid   string `protobuf:"bytes,1,opt,name=key_uid,json=keyUid,proto3" json:"key_uid,omitempty"`
	PartyUid string `protobuf:"bytes,2,opt,name=party_uid,json=partyUid,proto3" json:"party_uid,omitempty"`
}

func (m *KeygenRequest) Reset()         { *m = KeygenRequest{} }
func (m *KeygenRequest) String() string { return proto.CompactTextString(m) }
func (*KeygenRequest) ProtoMessage()    {}
func (*KeygenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4478b3adebfdcdf0, []int{0}
}
func (m *KeygenRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeygenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeygenRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeygenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeygenRequest.Merge(m, src)
}
func (m *KeygenRequest) XXX_Size() int {
	return m.Size()
}
func (m *KeygenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeygenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeygenRequest proto.InternalMessageInfo

func (m *KeygenRequest) GetKeyUid() string {
	if m != nil {
		return m.KeyUid
	}
	return ""
}

func (m *KeygenRequest) GetPartyUid() string {
	if m != nil {
		return m.PartyUid
	}
	return ""
}

type KeygenResponse struct {
	// Types that are valid to be assigned to KeygenResponse:
	//	*KeygenResponse_PubKey
	//	*KeygenResponse_Error
	KeygenResponse isKeygenResponse_KeygenResponse `protobuf_oneof:"keygen_response"`
}

func (m *KeygenResponse) Reset()         { *m = KeygenResponse{} }
func (m *KeygenResponse) String() string { return proto.CompactTextString(m) }
func (*KeygenResponse) ProtoMessage()    {}
func (*KeygenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4478b3adebfdcdf0, []int{1}
}
func (m *KeygenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeygenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeygenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeygenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeygenResponse.Merge(m, src)
}
func (m *KeygenResponse) XXX_Size() int {
	return m.Size()
}
func (m *KeygenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KeygenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KeygenResponse proto.InternalMessageInfo

type isKeygenResponse_KeygenResponse interface {
	isKeygenResponse_KeygenResponse()
	MarshalTo([]byte) (int, error)
	Size() int
}

type KeygenResponse_PubKey struct {
	PubKey []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3,oneof" json:"pub_key,omitempty"`
}
type KeygenResponse_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (*KeygenResponse_PubKey) isKeygenResponse_KeygenResponse() {}
func (*KeygenResponse_Error) isKeygenResponse_KeygenResponse()  {}

func (m *KeygenResponse) GetKeygenResponse() isKeygenResponse_KeygenResponse {
	if m != nil {
		return m.KeygenResponse
	}
	return nil
}

func (m *KeygenResponse) GetPubKey() []byte {
	if x, ok := m.GetKeygenResponse().(*KeygenResponse_PubKey); ok {
		return x.PubKey
	}
	return nil
}

func (m *KeygenResponse) GetError() string {
	if x, ok := m.GetKeygenResponse().(*KeygenResponse_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*KeygenResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*KeygenResponse_PubKey)(nil),
		(*KeygenResponse_Error)(nil),
	}
}

type SignRequest struct {
	KeyUid    string `protobuf:"bytes,1,opt,name=key_uid,json=keyUid,proto3" json:"key_uid,omitempty"`
	MsgToSign []byte `protobuf:"bytes,2,opt,name=msg_to_sign,json=msgToSign,proto3" json:"msg_to_sign,omitempty"`
	PartyUid  string `protobuf:"bytes,3,opt,name=party_uid,json=partyUid,proto3" json:"party_uid,omitempty"`
	PubKey    []byte `protobuf:"bytes,4,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
}

func (m *SignRequest) Reset()         { *m = SignRequest{} }
func (m *SignRequest) String() string { return proto.CompactTextString(m) }
func (*SignRequest) ProtoMessage()    {}
func (*SignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4478b3adebfdcdf0, []int{2}
}
func (m *SignRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRequest.Merge(m, src)
}
func (m *SignRequest) XXX_Size() int {
	return m.Size()
}
func (m *SignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignRequest proto.InternalMessageInfo

func (m *SignRequest) GetKeyUid() string {
	if m != nil {
		return m.KeyUid
	}
	return ""
}

func (m *SignRequest) GetMsgToSign() []byte {
	if m != nil {
		return m.MsgToSign
	}
	return nil
}

func (m *SignRequest) GetPartyUid() string {
	if m != nil {
		return m.PartyUid
	}
	return ""
}

func (m *SignRequest) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

type SignResponse struct {
	// Types that are valid to be assigned to SignResponse:
	//	*SignResponse_Signature
	//	*SignResponse_Error
	SignResponse isSignResponse_SignResponse `protobuf_oneof:"sign_response"`
}

func (m *SignResponse) Reset()         { *m = SignResponse{} }
func (m *SignResponse) String() string { return proto.CompactTextString(m) }
func (*SignResponse) ProtoMessage()    {}
func (*SignResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4478b3adebfdcdf0, []int{3}
}
func (m *SignResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignResponse.Merge(m, src)
}
func (m *SignResponse) XXX_Size() int {
	return m.Size()
}
func (m *SignResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignResponse proto.InternalMessageInfo

type isSignResponse_SignResponse interface {
	isSignResponse_SignResponse()
	MarshalTo([]byte) (int, error)
	Size() int
}

type SignResponse_Signature struct {
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3,oneof" json:"signature,omitempty"`
}
type SignResponse_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (*SignResponse_Signature) isSignResponse_SignResponse() {}
func (*SignResponse_Error) isSignResponse_SignResponse()     {}

func (m *SignResponse) GetSignResponse() isSignResponse_SignResponse {
	if m != nil {
		return m.SignResponse
	}
	return nil
}

func (m *SignResponse) GetSignature() []byte {
	if x, ok := m.GetSignResponse().(*SignResponse_Signature); ok {
		return x.Signature
	}
	return nil
}

func (m *SignResponse) GetError() string {
	if x, ok := m.GetSignResponse().(*SignResponse_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SignResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SignResponse_Signature)(nil),
		(*SignResponse_Error)(nil),
	}
}

func init() {
	proto.RegisterType((*KeygenRequest)(nil), "axelar.tss.tofnd.v1beta1.KeygenRequest")
	proto.RegisterType((*KeygenResponse)(nil), "axelar.tss.tofnd.v1beta1.KeygenResponse")
	proto.RegisterType((*SignRequest)(nil), "axelar.tss.tofnd.v1beta1.SignRequest")
	proto.RegisterType((*SignResponse)(nil), "axelar.tss.tofnd.v1beta1.SignResponse")
}

func init() {
	proto.RegisterFile("axelar/tss/tofnd/v1beta1/multisig.proto", fileDescriptor_4478b3adebfdcdf0)
}

var fileDescriptor_4478b3adebfdcdf0 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x13, 0x7f, 0x5a, 0x33, 0x6d, 0x2d, 0x66, 0xa1, 0x55, 0x61, 0x90, 0x80, 0xe8, 0xc6,
	0xc4, 0xe2, 0x1b, 0x14, 0x84, 0x62, 0x77, 0xd1, 0x2a, 0xb8, 0x09, 0x49, 0x73, 0x8d, 0x43, 0x9a,
	0x4c, 0x9c, 0x99, 0x68, 0x03, 0x3e, 0x84, 0x8f, 0xe5, 0xb2, 0x4b, 0x97, 0xd2, 0xbe, 0x88, 0x64,
	0x26, 0x6d, 0xe9, 0xa2, 0xb8, 0x9c, 0x7b, 0xcf, 0xfd, 0xe6, 0x1c, 0x0e, 0xba, 0xf0, 0x27, 0x30,
	0xf6, 0x99, 0x23, 0x38, 0x77, 0x04, 0x7d, 0x49, 0x43, 0xe7, 0xbd, 0x1b, 0x80, 0xf0, 0xbb, 0x4e,
	0x92, 0x8f, 0x05, 0xe1, 0x24, 0xb2, 0x33, 0x46, 0x05, 0x35, 0x3b, 0x4a, 0x68, 0x0b, 0xce, 0x6d,
	0x29, 0xb4, 0x2b, 0xe1, 0xc9, 0xf9, 0x46, 0xc4, 0x88, 0x26, 0x09, 0x4d, 0x15, 0xc0, 0xba, 0x45,
	0xad, 0x01, 0x14, 0x11, 0xa4, 0x2e, 0xbc, 0xe5, 0xc0, 0x85, 0x79, 0x84, 0xea, 0x31, 0x14, 0x5e,
	0x4e, 0xc2, 0x8e, 0x7e, 0xa6, 0x5f, 0x1a, 0x6e, 0x2d, 0x86, 0x62, 0x48, 0x42, 0xf3, 0x14, 0x19,
	0x99, 0xcf, 0x84, 0x5a, 0x6d, 0xc9, 0xd5, 0x9e, 0x1c, 0x0c, 0x49, 0x68, 0x3d, 0xa2, 0xfd, 0x05,
	0x86, 0x67, 0x34, 0xe5, 0x60, 0x1e, 0xa3, 0x7a, 0x96, 0x07, 0x5e, 0x0c, 0x85, 0xe4, 0x34, 0xfb,
	0x9a, 0x5b, 0xcb, 0xf2, 0x60, 0x00, 0x85, 0x79, 0x88, 0x76, 0x81, 0x31, 0xca, 0x14, 0xa5, 0xaf,
	0xb9, 0xea, 0xd9, 0x3b, 0x40, 0xed, 0x58, 0x42, 0x3c, 0x56, 0x51, 0xac, 0x4f, 0xd4, 0xb8, 0x27,
	0xd1, 0xff, 0xe6, 0x30, 0x6a, 0x24, 0x3c, 0xf2, 0x04, 0xf5, 0x38, 0x89, 0x52, 0x09, 0x6e, 0xba,
	0x46, 0xc2, 0xa3, 0x07, 0x5a, 0xde, 0xaf, 0x9b, 0xdf, 0x5e, 0x37, 0x5f, 0x52, 0x17, 0x56, 0x77,
	0xe4, 0x61, 0x65, 0xd4, 0x7a, 0x42, 0x4d, 0xf5, 0x7b, 0x95, 0x09, 0x23, 0xa3, 0xc4, 0xfb, 0x22,
	0x67, 0xb0, 0x4c, 0xb5, 0x1a, 0x6d, 0x0c, 0xd6, 0x46, 0xad, 0x52, 0xb4, 0x8c, 0xd5, 0xbb, 0xfb,
	0x9e, 0x61, 0x7d, 0x3a, 0xc3, 0xfa, 0xef, 0x0c, 0xeb, 0x5f, 0x73, 0xac, 0x4d, 0xe7, 0x58, 0xfb,
	0x99, 0x63, 0xed, 0xf9, 0x3a, 0x22, 0xe2, 0x35, 0x0f, 0xec, 0x11, 0x4d, 0x1c, 0xd5, 0x60, 0x0a,
	0xe2, 0x83, 0xb2, 0xb8, 0x7a, 0x5d, 0x8d, 0x28, 0x03, 0x67, 0xb2, 0xaa, 0x35, 0xa8, 0xc9, 0x22,
	0x6f, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xd3, 0xaa, 0x4a, 0x34, 0x02, 0x00, 0x00,
}

func (m *KeygenRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeygenRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeygenRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PartyUid) > 0 {
		i -= len(m.PartyUid)
		copy(dAtA[i:], m.PartyUid)
		i = encodeVarintMultisig(dAtA, i, uint64(len(m.PartyUid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.KeyUid) > 0 {
		i -= len(m.KeyUid)
		copy(dAtA[i:], m.KeyUid)
		i = encodeVarintMultisig(dAtA, i, uint64(len(m.KeyUid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KeygenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeygenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeygenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.KeygenResponse != nil {
		{
			size := m.KeygenResponse.Size()
			i -= size
			if _, err := m.KeygenResponse.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *KeygenResponse_PubKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeygenResponse_PubKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PubKey != nil {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintMultisig(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *KeygenResponse_Error) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeygenResponse_Error) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Error)
	copy(dAtA[i:], m.Error)
	i = encodeVarintMultisig(dAtA, i, uint64(len(m.Error)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *SignRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintMultisig(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PartyUid) > 0 {
		i -= len(m.PartyUid)
		copy(dAtA[i:], m.PartyUid)
		i = encodeVarintMultisig(dAtA, i, uint64(len(m.PartyUid)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MsgToSign) > 0 {
		i -= len(m.MsgToSign)
		copy(dAtA[i:], m.MsgToSign)
		i = encodeVarintMultisig(dAtA, i, uint64(len(m.MsgToSign)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.KeyUid) > 0 {
		i -= len(m.KeyUid)
		copy(dAtA[i:], m.KeyUid)
		i = encodeVarintMultisig(dAtA, i, uint64(len(m.KeyUid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SignResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SignResponse != nil {
		{
			size := m.SignResponse.Size()
			i -= size
			if _, err := m.SignResponse.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *SignResponse_Signature) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignResponse_Signature) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Signature != nil {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintMultisig(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *SignResponse_Error) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignResponse_Error) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Error)
	copy(dAtA[i:], m.Error)
	i = encodeVarintMultisig(dAtA, i, uint64(len(m.Error)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func encodeVarintMultisig(dAtA []byte, offset int, v uint64) int {
	offset -= sovMultisig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KeygenRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.KeyUid)
	if l > 0 {
		n += 1 + l + sovMultisig(uint64(l))
	}
	l = len(m.PartyUid)
	if l > 0 {
		n += 1 + l + sovMultisig(uint64(l))
	}
	return n
}

func (m *KeygenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.KeygenResponse != nil {
		n += m.KeygenResponse.Size()
	}
	return n
}

func (m *KeygenResponse_PubKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PubKey != nil {
		l = len(m.PubKey)
		n += 1 + l + sovMultisig(uint64(l))
	}
	return n
}
func (m *KeygenResponse_Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Error)
	n += 1 + l + sovMultisig(uint64(l))
	return n
}
func (m *SignRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.KeyUid)
	if l > 0 {
		n += 1 + l + sovMultisig(uint64(l))
	}
	l = len(m.MsgToSign)
	if l > 0 {
		n += 1 + l + sovMultisig(uint64(l))
	}
	l = len(m.PartyUid)
	if l > 0 {
		n += 1 + l + sovMultisig(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovMultisig(uint64(l))
	}
	return n
}

func (m *SignResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignResponse != nil {
		n += m.SignResponse.Size()
	}
	return n
}

func (m *SignResponse_Signature) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Signature != nil {
		l = len(m.Signature)
		n += 1 + l + sovMultisig(uint64(l))
	}
	return n
}
func (m *SignResponse_Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Error)
	n += 1 + l + sovMultisig(uint64(l))
	return n
}

func sovMultisig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMultisig(x uint64) (n int) {
	return sovMultisig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeygenRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMultisig
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
			return fmt.Errorf("proto: KeygenRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeygenRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyUid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyUid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartyUid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PartyUid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMultisig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMultisig
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
func (m *KeygenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMultisig
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
			return fmt.Errorf("proto: KeygenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeygenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.KeygenResponse = &KeygenResponse_PubKey{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeygenResponse = &KeygenResponse_Error{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMultisig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMultisig
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
func (m *SignRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMultisig
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
			return fmt.Errorf("proto: SignRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyUid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyUid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgToSign", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgToSign = append(m.MsgToSign[:0], dAtA[iNdEx:postIndex]...)
			if m.MsgToSign == nil {
				m.MsgToSign = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartyUid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PartyUid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMultisig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMultisig
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
func (m *SignResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMultisig
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
			return fmt.Errorf("proto: SignResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.SignResponse = &SignResponse_Signature{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignResponse = &SignResponse_Error{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMultisig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMultisig
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
func skipMultisig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMultisig
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
					return 0, ErrIntOverflowMultisig
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
					return 0, ErrIntOverflowMultisig
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
				return 0, ErrInvalidLengthMultisig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMultisig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMultisig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMultisig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMultisig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMultisig = fmt.Errorf("proto: unexpected end of group")
)
