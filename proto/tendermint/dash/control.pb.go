// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/dash/control.proto

package dash

import (
	fmt "fmt"
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

type ControlMessage struct {
	// Types that are valid to be assigned to Sum:
	//	*ControlMessage_ValidatorChallenge
	//	*ControlMessage_ValidatorChallengeResponse
	Sum isControlMessage_Sum `protobuf_oneof:"sum"`
}

func (m *ControlMessage) Reset()         { *m = ControlMessage{} }
func (m *ControlMessage) String() string { return proto.CompactTextString(m) }
func (*ControlMessage) ProtoMessage()    {}
func (*ControlMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b5d1e764dda6dda, []int{0}
}
func (m *ControlMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ControlMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ControlMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ControlMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlMessage.Merge(m, src)
}
func (m *ControlMessage) XXX_Size() int {
	return m.Size()
}
func (m *ControlMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ControlMessage proto.InternalMessageInfo

type isControlMessage_Sum interface {
	isControlMessage_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ControlMessage_ValidatorChallenge struct {
	ValidatorChallenge *ValidatorChallenge `protobuf:"bytes,1,opt,name=validator_challenge,json=validatorChallenge,proto3,oneof" json:"validator_challenge,omitempty"`
}
type ControlMessage_ValidatorChallengeResponse struct {
	ValidatorChallengeResponse *ValidatorChallengeResponse `protobuf:"bytes,2,opt,name=validator_challenge_response,json=validatorChallengeResponse,proto3,oneof" json:"validator_challenge_response,omitempty"`
}

func (*ControlMessage_ValidatorChallenge) isControlMessage_Sum()         {}
func (*ControlMessage_ValidatorChallengeResponse) isControlMessage_Sum() {}

func (m *ControlMessage) GetSum() isControlMessage_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *ControlMessage) GetValidatorChallenge() *ValidatorChallenge {
	if x, ok := m.GetSum().(*ControlMessage_ValidatorChallenge); ok {
		return x.ValidatorChallenge
	}
	return nil
}

func (m *ControlMessage) GetValidatorChallengeResponse() *ValidatorChallengeResponse {
	if x, ok := m.GetSum().(*ControlMessage_ValidatorChallengeResponse); ok {
		return x.ValidatorChallengeResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ControlMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ControlMessage_ValidatorChallenge)(nil),
		(*ControlMessage_ValidatorChallengeResponse)(nil),
	}
}

type ValidatorChallenge struct {
	SenderProTxHash    []byte `protobuf:"bytes,1,opt,name=sender_protxhash,json=senderProtxhash,proto3" json:"sender_protxhash,omitempty"`
	RecipientProTxHash []byte `protobuf:"bytes,2,opt,name=recipient_protxhash,json=recipientProtxhash,proto3" json:"recipient_protxhash,omitempty"`
	SenderNodeID       string `protobuf:"bytes,3,opt,name=sender_node_id,json=senderNodeId,proto3" json:"sender_node_id,omitempty"`
	RecipientNodeID    string `protobuf:"bytes,4,opt,name=recipient_node_id,json=recipientNodeId,proto3" json:"recipient_node_id,omitempty"`
	SenderHeight       int64  `protobuf:"varint,5,opt,name=sender_height,json=senderHeight,proto3" json:"sender_height,omitempty"`
	QuorumHash         []byte `protobuf:"bytes,6,opt,name=quorum_hash,json=quorumHash,proto3" json:"quorum_hash,omitempty"`
	Token              []byte `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	Signature          []byte `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ValidatorChallenge) Reset()         { *m = ValidatorChallenge{} }
func (m *ValidatorChallenge) String() string { return proto.CompactTextString(m) }
func (*ValidatorChallenge) ProtoMessage()    {}
func (*ValidatorChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b5d1e764dda6dda, []int{1}
}
func (m *ValidatorChallenge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorChallenge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorChallenge.Merge(m, src)
}
func (m *ValidatorChallenge) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorChallenge proto.InternalMessageInfo

func (m *ValidatorChallenge) GetSenderProTxHash() []byte {
	if m != nil {
		return m.SenderProTxHash
	}
	return nil
}

func (m *ValidatorChallenge) GetRecipientProTxHash() []byte {
	if m != nil {
		return m.RecipientProTxHash
	}
	return nil
}

func (m *ValidatorChallenge) GetSenderNodeID() string {
	if m != nil {
		return m.SenderNodeID
	}
	return ""
}

func (m *ValidatorChallenge) GetRecipientNodeID() string {
	if m != nil {
		return m.RecipientNodeID
	}
	return ""
}

func (m *ValidatorChallenge) GetSenderHeight() int64 {
	if m != nil {
		return m.SenderHeight
	}
	return 0
}

func (m *ValidatorChallenge) GetQuorumHash() []byte {
	if m != nil {
		return m.QuorumHash
	}
	return nil
}

func (m *ValidatorChallenge) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *ValidatorChallenge) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type ValidatorChallengeResponse struct {
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ValidatorChallengeResponse) Reset()         { *m = ValidatorChallengeResponse{} }
func (m *ValidatorChallengeResponse) String() string { return proto.CompactTextString(m) }
func (*ValidatorChallengeResponse) ProtoMessage()    {}
func (*ValidatorChallengeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b5d1e764dda6dda, []int{2}
}
func (m *ValidatorChallengeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorChallengeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorChallengeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorChallengeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorChallengeResponse.Merge(m, src)
}
func (m *ValidatorChallengeResponse) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorChallengeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorChallengeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorChallengeResponse proto.InternalMessageInfo

func (m *ValidatorChallengeResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*ControlMessage)(nil), "tendermint.dash.ControlMessage")
	proto.RegisterType((*ValidatorChallenge)(nil), "tendermint.dash.ValidatorChallenge")
	proto.RegisterType((*ValidatorChallengeResponse)(nil), "tendermint.dash.ValidatorChallengeResponse")
}

func init() { proto.RegisterFile("tendermint/dash/control.proto", fileDescriptor_6b5d1e764dda6dda) }

var fileDescriptor_6b5d1e764dda6dda = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6b, 0xdb, 0x30,
	0x18, 0xb5, 0x9a, 0xa5, 0x5b, 0xbf, 0x66, 0x4d, 0xa7, 0x94, 0x61, 0x42, 0xe7, 0x84, 0xf4, 0x12,
	0x18, 0x38, 0xb0, 0x41, 0x0f, 0x3b, 0x6c, 0x90, 0x0e, 0x96, 0x1d, 0x36, 0x8a, 0x37, 0x7a, 0xd8,
	0xc5, 0xa8, 0x91, 0xb0, 0xc4, 0x12, 0x29, 0x93, 0xe4, 0xd0, 0x9f, 0xd1, 0x9f, 0xb5, 0x63, 0x8f,
	0x3b, 0x85, 0xe2, 0xfc, 0x91, 0x61, 0xd9, 0x9e, 0xd3, 0x74, 0x85, 0xde, 0xa4, 0xf7, 0xbe, 0xf7,
	0xde, 0xf7, 0x12, 0x0b, 0x5e, 0x59, 0x26, 0x29, 0xd3, 0x73, 0x21, 0xed, 0x88, 0x12, 0xc3, 0x47,
	0x53, 0x25, 0xad, 0x56, 0xb3, 0x70, 0xa1, 0x95, 0x55, 0xb8, 0x5d, 0xd3, 0x61, 0x4e, 0x77, 0x8f,
	0x12, 0x95, 0x28, 0xc7, 0x8d, 0xf2, 0x53, 0x31, 0x36, 0xb8, 0x45, 0x70, 0x70, 0x56, 0x08, 0xbf,
	0x30, 0x63, 0x48, 0xc2, 0xf0, 0x05, 0x74, 0x96, 0x64, 0x26, 0x28, 0xb1, 0x4a, 0xc7, 0x53, 0x4e,
	0x66, 0x33, 0x26, 0x13, 0xe6, 0xa3, 0x3e, 0x1a, 0xee, 0xbf, 0x39, 0x09, 0xb7, 0x7c, 0xc3, 0x8b,
	0x6a, 0xf6, 0xac, 0x1a, 0x9d, 0x78, 0x11, 0x5e, 0xde, 0x43, 0xb1, 0x82, 0xe3, 0xff, 0xf8, 0xc6,
	0x9a, 0x99, 0x85, 0x92, 0x86, 0xf9, 0x3b, 0x2e, 0xe0, 0xf5, 0x23, 0x02, 0xa2, 0x52, 0x32, 0xf1,
	0xa2, 0xee, 0xf2, 0x41, 0x76, 0xdc, 0x84, 0x86, 0x49, 0xe7, 0x83, 0xeb, 0x06, 0xe0, 0xfb, 0x1e,
	0xf8, 0x3d, 0x1c, 0x1a, 0x97, 0x14, 0xe7, 0xbf, 0xc4, 0x15, 0x27, 0x86, 0xbb, 0x8e, 0xad, 0x71,
	0x27, 0x5b, 0xf5, 0xda, 0xdf, 0x1c, 0x77, 0xae, 0xd5, 0xf7, 0xab, 0x09, 0x31, 0x3c, 0x6a, 0x9b,
	0x0a, 0x28, 0x66, 0xf1, 0x27, 0xe8, 0x68, 0x36, 0x15, 0x0b, 0xc1, 0xa4, 0xdd, 0xb0, 0xd8, 0x71,
	0x16, 0x2f, 0xb3, 0x55, 0x0f, 0x47, 0x15, 0x5d, 0xbb, 0x60, 0xbd, 0x81, 0x95, 0x46, 0xa7, 0x70,
	0x50, 0x2e, 0x22, 0x15, 0x65, 0xb1, 0xa0, 0x7e, 0xa3, 0x8f, 0x86, 0x7b, 0xe3, 0xc3, 0x6c, 0xd5,
	0x6b, 0x15, 0x6b, 0x7c, 0x55, 0x94, 0x7d, 0xfe, 0x18, 0xb5, 0x4c, 0x7d, 0xa3, 0xf8, 0x03, 0xbc,
	0xa8, 0x17, 0xa8, 0xa4, 0x4f, 0x9c, 0xd4, 0x35, 0xf8, 0x17, 0x5f, 0xaa, 0xdb, 0xfa, 0x0e, 0x40,
	0xf1, 0x09, 0x3c, 0x2f, 0x83, 0x39, 0x13, 0x09, 0xb7, 0x7e, 0xb3, 0x8f, 0x86, 0x8d, 0x2a, 0x65,
	0xe2, 0x30, 0xdc, 0x83, 0xfd, 0x5f, 0xa9, 0xd2, 0xe9, 0x3c, 0x76, 0xf5, 0x76, 0xf3, 0x7a, 0x11,
	0x14, 0x50, 0x5e, 0x09, 0x1f, 0x41, 0xd3, 0xaa, 0x9f, 0x4c, 0xfa, 0x4f, 0x1d, 0x55, 0x5c, 0xf0,
	0x31, 0xec, 0x19, 0x91, 0x48, 0x62, 0x53, 0xcd, 0xfc, 0x67, 0x8e, 0xa9, 0x81, 0xc1, 0x3b, 0xe8,
	0x3e, 0xfc, 0xaf, 0xde, 0xd5, 0xa2, 0x2d, 0xed, 0xf8, 0xfc, 0x77, 0x16, 0xa0, 0x9b, 0x2c, 0x40,
	0xb7, 0x59, 0x80, 0xae, 0xd7, 0x81, 0x77, 0xb3, 0x0e, 0xbc, 0x3f, 0xeb, 0xc0, 0xfb, 0x71, 0x9a,
	0x08, 0xcb, 0xd3, 0xcb, 0x70, 0xaa, 0xe6, 0xa3, 0x8d, 0xc7, 0xb1, 0x71, 0x2c, 0x3e, 0xff, 0xad,
	0x87, 0x73, 0xb9, 0xeb, 0xe0, 0xb7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xd1, 0x2e, 0x90,
	0x52, 0x03, 0x00, 0x00,
}

func (m *ControlMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ControlMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ControlMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *ControlMessage_ValidatorChallenge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ControlMessage_ValidatorChallenge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ValidatorChallenge != nil {
		{
			size, err := m.ValidatorChallenge.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintControl(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *ControlMessage_ValidatorChallengeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ControlMessage_ValidatorChallengeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ValidatorChallengeResponse != nil {
		{
			size, err := m.ValidatorChallengeResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintControl(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *ValidatorChallenge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorChallenge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorChallenge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintControl(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintControl(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.QuorumHash) > 0 {
		i -= len(m.QuorumHash)
		copy(dAtA[i:], m.QuorumHash)
		i = encodeVarintControl(dAtA, i, uint64(len(m.QuorumHash)))
		i--
		dAtA[i] = 0x32
	}
	if m.SenderHeight != 0 {
		i = encodeVarintControl(dAtA, i, uint64(m.SenderHeight))
		i--
		dAtA[i] = 0x28
	}
	if len(m.RecipientNodeID) > 0 {
		i -= len(m.RecipientNodeID)
		copy(dAtA[i:], m.RecipientNodeID)
		i = encodeVarintControl(dAtA, i, uint64(len(m.RecipientNodeID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SenderNodeID) > 0 {
		i -= len(m.SenderNodeID)
		copy(dAtA[i:], m.SenderNodeID)
		i = encodeVarintControl(dAtA, i, uint64(len(m.SenderNodeID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RecipientProTxHash) > 0 {
		i -= len(m.RecipientProTxHash)
		copy(dAtA[i:], m.RecipientProTxHash)
		i = encodeVarintControl(dAtA, i, uint64(len(m.RecipientProTxHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SenderProTxHash) > 0 {
		i -= len(m.SenderProTxHash)
		copy(dAtA[i:], m.SenderProTxHash)
		i = encodeVarintControl(dAtA, i, uint64(len(m.SenderProTxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorChallengeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorChallengeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorChallengeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintControl(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintControl(dAtA []byte, offset int, v uint64) int {
	offset -= sovControl(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ControlMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *ControlMessage_ValidatorChallenge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValidatorChallenge != nil {
		l = m.ValidatorChallenge.Size()
		n += 1 + l + sovControl(uint64(l))
	}
	return n
}
func (m *ControlMessage_ValidatorChallengeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValidatorChallengeResponse != nil {
		l = m.ValidatorChallengeResponse.Size()
		n += 1 + l + sovControl(uint64(l))
	}
	return n
}
func (m *ValidatorChallenge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SenderProTxHash)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	l = len(m.RecipientProTxHash)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	l = len(m.SenderNodeID)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	l = len(m.RecipientNodeID)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	if m.SenderHeight != 0 {
		n += 1 + sovControl(uint64(m.SenderHeight))
	}
	l = len(m.QuorumHash)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	return n
}

func (m *ValidatorChallengeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	return n
}

func sovControl(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozControl(x uint64) (n int) {
	return sovControl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ControlMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
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
			return fmt.Errorf("proto: ControlMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ControlMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorChallenge", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
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
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ValidatorChallenge{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &ControlMessage_ValidatorChallenge{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorChallengeResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
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
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ValidatorChallengeResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &ControlMessage_ValidatorChallengeResponse{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthControl
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
func (m *ValidatorChallenge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
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
			return fmt.Errorf("proto: ValidatorChallenge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorChallenge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderProTxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
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
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SenderProTxHash = append(m.SenderProTxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.SenderProTxHash == nil {
				m.SenderProTxHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientProTxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
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
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientProTxHash = append(m.RecipientProTxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.RecipientProTxHash == nil {
				m.RecipientProTxHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderNodeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
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
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SenderNodeID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientNodeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
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
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientNodeID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderHeight", wireType)
			}
			m.SenderHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SenderHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuorumHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
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
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuorumHash = append(m.QuorumHash[:0], dAtA[iNdEx:postIndex]...)
			if m.QuorumHash == nil {
				m.QuorumHash = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
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
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = append(m.Token[:0], dAtA[iNdEx:postIndex]...)
			if m.Token == nil {
				m.Token = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
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
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthControl
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
func (m *ValidatorChallengeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
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
			return fmt.Errorf("proto: ValidatorChallengeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorChallengeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
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
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthControl
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
func skipControl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowControl
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
					return 0, ErrIntOverflowControl
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
					return 0, ErrIntOverflowControl
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
				return 0, ErrInvalidLengthControl
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupControl
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthControl
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthControl        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowControl          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupControl = fmt.Errorf("proto: unexpected end of group")
)
