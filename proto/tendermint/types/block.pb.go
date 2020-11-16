// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/types/block.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Block struct {
	Header        Header         `protobuf:"bytes,1,opt,name=header,proto3" json:"header"`
	Data          Data           `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
	CoreChainLock *CoreChainLock `protobuf:"bytes,100,opt,name=core_chain_lock,json=coreChainLock,proto3" json:"core_chain_lock,omitempty"`
	Evidence      EvidenceList   `protobuf:"bytes,3,opt,name=evidence,proto3" json:"evidence"`
	LastCommit    *Commit        `protobuf:"bytes,4,opt,name=last_commit,json=lastCommit,proto3" json:"last_commit,omitempty"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_70840e82f4357ab1, []int{0}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Block.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return m.Size()
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() Header {
	if m != nil {
		return m.Header
	}
	return Header{}
}

func (m *Block) GetData() Data {
	if m != nil {
		return m.Data
	}
	return Data{}
}

func (m *Block) GetCoreChainLock() *CoreChainLock {
	if m != nil {
		return m.CoreChainLock
	}
	return nil
}

func (m *Block) GetEvidence() EvidenceList {
	if m != nil {
		return m.Evidence
	}
	return EvidenceList{}
}

func (m *Block) GetLastCommit() *Commit {
	if m != nil {
		return m.LastCommit
	}
	return nil
}

// TODO: Remove later
type BlockTestOne struct {
	FirstField  string `protobuf:"bytes,1,opt,name=firstField,proto3" json:"firstField,omitempty"`
	SecondField string `protobuf:"bytes,2,opt,name=secondField,proto3" json:"secondField,omitempty"`
}

func (m *BlockTestOne) Reset()         { *m = BlockTestOne{} }
func (m *BlockTestOne) String() string { return proto.CompactTextString(m) }
func (*BlockTestOne) ProtoMessage()    {}
func (*BlockTestOne) Descriptor() ([]byte, []int) {
	return fileDescriptor_70840e82f4357ab1, []int{1}
}
func (m *BlockTestOne) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockTestOne) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockTestOne.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockTestOne) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockTestOne.Merge(m, src)
}
func (m *BlockTestOne) XXX_Size() int {
	return m.Size()
}
func (m *BlockTestOne) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockTestOne.DiscardUnknown(m)
}

var xxx_messageInfo_BlockTestOne proto.InternalMessageInfo

func (m *BlockTestOne) GetFirstField() string {
	if m != nil {
		return m.FirstField
	}
	return ""
}

func (m *BlockTestOne) GetSecondField() string {
	if m != nil {
		return m.SecondField
	}
	return ""
}

// TODO: Remove later
type BlockTestTwo struct {
	FirstField string `protobuf:"bytes,1,opt,name=firstField,proto3" json:"firstField,omitempty"`
	ThirdField string `protobuf:"bytes,100,opt,name=thirdField,proto3" json:"thirdField,omitempty"`
}

func (m *BlockTestTwo) Reset()         { *m = BlockTestTwo{} }
func (m *BlockTestTwo) String() string { return proto.CompactTextString(m) }
func (*BlockTestTwo) ProtoMessage()    {}
func (*BlockTestTwo) Descriptor() ([]byte, []int) {
	return fileDescriptor_70840e82f4357ab1, []int{2}
}
func (m *BlockTestTwo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockTestTwo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockTestTwo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockTestTwo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockTestTwo.Merge(m, src)
}
func (m *BlockTestTwo) XXX_Size() int {
	return m.Size()
}
func (m *BlockTestTwo) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockTestTwo.DiscardUnknown(m)
}

var xxx_messageInfo_BlockTestTwo proto.InternalMessageInfo

func (m *BlockTestTwo) GetFirstField() string {
	if m != nil {
		return m.FirstField
	}
	return ""
}

func (m *BlockTestTwo) GetThirdField() string {
	if m != nil {
		return m.ThirdField
	}
	return ""
}

func init() {
	proto.RegisterType((*Block)(nil), "tendermint.types.Block")
	proto.RegisterType((*BlockTestOne)(nil), "tendermint.types.BlockTestOne")
	proto.RegisterType((*BlockTestTwo)(nil), "tendermint.types.BlockTestTwo")
}

func init() { proto.RegisterFile("tendermint/types/block.proto", fileDescriptor_70840e82f4357ab1) }

var fileDescriptor_70840e82f4357ab1 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x93, 0x5a, 0x8b, 0x4e, 0x15, 0x65, 0x10, 0x09, 0x45, 0xa6, 0xa5, 0x2b, 0x57, 0x89,
	0x28, 0x28, 0xee, 0xa4, 0x55, 0x71, 0x51, 0xff, 0x42, 0x57, 0x6e, 0x4a, 0x3a, 0x73, 0x6d, 0x86,
	0x36, 0x99, 0x92, 0x8c, 0x8a, 0x6f, 0xe1, 0x0b, 0xb9, 0xef, 0xb2, 0x4b, 0x57, 0x22, 0xed, 0x8b,
	0x48, 0x6e, 0xa2, 0x0d, 0x0d, 0xe2, 0x26, 0x5c, 0xce, 0xf9, 0xce, 0xc9, 0xcc, 0x70, 0xc9, 0x9e,
	0x86, 0x50, 0x40, 0x14, 0xc8, 0x50, 0x3b, 0xfa, 0x75, 0x0c, 0xb1, 0xd3, 0x1f, 0x29, 0x3e, 0xb4,
	0xc7, 0x91, 0xd2, 0x8a, 0x6e, 0x2f, 0x5c, 0x1b, 0xdd, 0xda, 0xce, 0x40, 0x0d, 0x14, 0x9a, 0x4e,
	0x32, 0xa5, 0x5c, 0xad, 0xd8, 0x82, 0xdf, 0xcc, 0xad, 0x17, 0x5c, 0x78, 0x96, 0x02, 0x42, 0x0e,
	0x29, 0xd0, 0x7c, 0x2f, 0x91, 0xd5, 0x56, 0xf2, 0x5b, 0x7a, 0x4c, 0x2a, 0x3e, 0x78, 0x02, 0x22,
	0xcb, 0x6c, 0x98, 0xfb, 0xd5, 0x43, 0xcb, 0x5e, 0x3e, 0x81, 0x7d, 0x85, 0x7e, 0xab, 0x3c, 0xf9,
	0xac, 0x1b, 0x6e, 0x46, 0xd3, 0x03, 0x52, 0x16, 0x9e, 0xf6, 0xac, 0x12, 0xa6, 0x76, 0x8b, 0xa9,
	0x73, 0x4f, 0x7b, 0x59, 0x06, 0x49, 0x7a, 0x4d, 0xb6, 0xb8, 0x8a, 0xa0, 0xc7, 0x7d, 0x4f, 0x86,
	0xbd, 0xe4, 0xe7, 0x96, 0xc0, 0x70, 0xbd, 0x18, 0x6e, 0xab, 0x08, 0xda, 0x09, 0xd7, 0x51, 0x7c,
	0x88, 0x2d, 0xa6, 0xbb, 0xc9, 0xf3, 0x22, 0x3d, 0x23, 0x6b, 0x3f, 0x97, 0xb2, 0x56, 0xb0, 0x87,
	0x15, 0x7b, 0x2e, 0x32, 0xa2, 0x23, 0x63, 0x9d, 0x1d, 0xe6, 0x37, 0x45, 0x4f, 0x49, 0x75, 0xe4,
	0xc5, 0xba, 0xc7, 0x55, 0x10, 0x48, 0x6d, 0x95, 0xff, 0xba, 0x7f, 0x1b, 0x7d, 0x97, 0x24, 0x70,
	0x3a, 0x37, 0xef, 0xc8, 0x06, 0x3e, 0x5f, 0x17, 0x62, 0x7d, 0x1b, 0x02, 0x65, 0x84, 0x3c, 0xca,
	0x28, 0xd6, 0x97, 0x12, 0x46, 0x02, 0x5f, 0x72, 0xdd, 0xcd, 0x29, 0xb4, 0x41, 0xaa, 0x31, 0x70,
	0x15, 0x8a, 0x14, 0x28, 0x21, 0x90, 0x97, 0x9a, 0x37, 0xb9, 0xc6, 0xee, 0x8b, 0xfa, 0xb7, 0x91,
	0x11, 0xa2, 0x7d, 0x19, 0x65, 0x85, 0x22, 0xf5, 0x17, 0x4a, 0xeb, 0x7e, 0x32, 0x63, 0xe6, 0x74,
	0xc6, 0xcc, 0xaf, 0x19, 0x33, 0xdf, 0xe6, 0xcc, 0x98, 0xce, 0x99, 0xf1, 0x31, 0x67, 0xc6, 0xc3,
	0xc9, 0x40, 0x6a, 0xff, 0xa9, 0x6f, 0x73, 0x15, 0x38, 0xf9, 0x3d, 0x59, 0x8c, 0xe9, 0xb6, 0x2d,
	0xef, 0x50, 0xbf, 0x82, 0xfa, 0xd1, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0xda, 0x2a, 0x4a,
	0xc2, 0x02, 0x00, 0x00,
}

func (m *Block) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Block) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Block) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CoreChainLock != nil {
		{
			size, err := m.CoreChainLock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0xa2
	}
	if m.LastCommit != nil {
		{
			size, err := m.LastCommit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Evidence.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBlock(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBlock(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBlock(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BlockTestOne) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockTestOne) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockTestOne) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SecondField) > 0 {
		i -= len(m.SecondField)
		copy(dAtA[i:], m.SecondField)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.SecondField)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FirstField) > 0 {
		i -= len(m.FirstField)
		copy(dAtA[i:], m.FirstField)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.FirstField)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockTestTwo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockTestTwo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockTestTwo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ThirdField) > 0 {
		i -= len(m.ThirdField)
		copy(dAtA[i:], m.ThirdField)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.ThirdField)))
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0xa2
	}
	if len(m.FirstField) > 0 {
		i -= len(m.FirstField)
		copy(dAtA[i:], m.FirstField)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.FirstField)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Block) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Header.Size()
	n += 1 + l + sovBlock(uint64(l))
	l = m.Data.Size()
	n += 1 + l + sovBlock(uint64(l))
	l = m.Evidence.Size()
	n += 1 + l + sovBlock(uint64(l))
	if m.LastCommit != nil {
		l = m.LastCommit.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.CoreChainLock != nil {
		l = m.CoreChainLock.Size()
		n += 2 + l + sovBlock(uint64(l))
	}
	return n
}

func (m *BlockTestOne) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FirstField)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.SecondField)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	return n
}

func (m *BlockTestTwo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FirstField)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.ThirdField)
	if l > 0 {
		n += 2 + l + sovBlock(uint64(l))
	}
	return n
}

func sovBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Block) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: Block: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Block: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Evidence.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastCommit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastCommit == nil {
				m.LastCommit = &Commit{}
			}
			if err := m.LastCommit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 100:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoreChainLock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CoreChainLock == nil {
				m.CoreChainLock = &CoreChainLock{}
			}
			if err := m.CoreChainLock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBlock
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
func (m *BlockTestOne) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: BlockTestOne: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockTestOne: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstField", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstField = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecondField", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecondField = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBlock
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
func (m *BlockTestTwo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: BlockTestTwo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockTestTwo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstField", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstField = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 100:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThirdField", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ThirdField = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBlock
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
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
				return 0, ErrInvalidLengthBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlock = fmt.Errorf("proto: unexpected end of group")
)
