// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wormhole/events.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

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

type EventGuardianSetUpdate struct {
	OldIndex uint32 `protobuf:"varint,1,opt,name=old_index,json=oldIndex,proto3" json:"old_index,omitempty"`
	NewIndex uint32 `protobuf:"varint,2,opt,name=new_index,json=newIndex,proto3" json:"new_index,omitempty"`
}

func (m *EventGuardianSetUpdate) Reset()         { *m = EventGuardianSetUpdate{} }
func (m *EventGuardianSetUpdate) String() string { return proto.CompactTextString(m) }
func (*EventGuardianSetUpdate) ProtoMessage()    {}
func (*EventGuardianSetUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_486bfc4df1202b88, []int{0}
}
func (m *EventGuardianSetUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventGuardianSetUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventGuardianSetUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventGuardianSetUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventGuardianSetUpdate.Merge(m, src)
}
func (m *EventGuardianSetUpdate) XXX_Size() int {
	return m.Size()
}
func (m *EventGuardianSetUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventGuardianSetUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_EventGuardianSetUpdate proto.InternalMessageInfo

func (m *EventGuardianSetUpdate) GetOldIndex() uint32 {
	if m != nil {
		return m.OldIndex
	}
	return 0
}

func (m *EventGuardianSetUpdate) GetNewIndex() uint32 {
	if m != nil {
		return m.NewIndex
	}
	return 0
}

type EventPostedMessage struct {
	Emitter  []byte `protobuf:"bytes,1,opt,name=emitter,proto3" json:"emitter,omitempty"`
	Sequence uint64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Nonce    uint32 `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Time     uint64 `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	Payload  []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *EventPostedMessage) Reset()         { *m = EventPostedMessage{} }
func (m *EventPostedMessage) String() string { return proto.CompactTextString(m) }
func (*EventPostedMessage) ProtoMessage()    {}
func (*EventPostedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_486bfc4df1202b88, []int{1}
}
func (m *EventPostedMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventPostedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventPostedMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventPostedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventPostedMessage.Merge(m, src)
}
func (m *EventPostedMessage) XXX_Size() int {
	return m.Size()
}
func (m *EventPostedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EventPostedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EventPostedMessage proto.InternalMessageInfo

func (m *EventPostedMessage) GetEmitter() []byte {
	if m != nil {
		return m.Emitter
	}
	return nil
}

func (m *EventPostedMessage) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *EventPostedMessage) GetNonce() uint32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *EventPostedMessage) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *EventPostedMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type EventGuardianRegistered struct {
	GuardianKey  []byte `protobuf:"bytes,1,opt,name=guardian_key,json=guardianKey,proto3" json:"guardian_key,omitempty"`
	ValidatorKey []byte `protobuf:"bytes,2,opt,name=validator_key,json=validatorKey,proto3" json:"validator_key,omitempty"`
}

func (m *EventGuardianRegistered) Reset()         { *m = EventGuardianRegistered{} }
func (m *EventGuardianRegistered) String() string { return proto.CompactTextString(m) }
func (*EventGuardianRegistered) ProtoMessage()    {}
func (*EventGuardianRegistered) Descriptor() ([]byte, []int) {
	return fileDescriptor_486bfc4df1202b88, []int{2}
}
func (m *EventGuardianRegistered) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventGuardianRegistered) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventGuardianRegistered.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventGuardianRegistered) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventGuardianRegistered.Merge(m, src)
}
func (m *EventGuardianRegistered) XXX_Size() int {
	return m.Size()
}
func (m *EventGuardianRegistered) XXX_DiscardUnknown() {
	xxx_messageInfo_EventGuardianRegistered.DiscardUnknown(m)
}

var xxx_messageInfo_EventGuardianRegistered proto.InternalMessageInfo

func (m *EventGuardianRegistered) GetGuardianKey() []byte {
	if m != nil {
		return m.GuardianKey
	}
	return nil
}

func (m *EventGuardianRegistered) GetValidatorKey() []byte {
	if m != nil {
		return m.ValidatorKey
	}
	return nil
}

type EventConsensusSetUpdate struct {
	OldIndex uint32 `protobuf:"varint,1,opt,name=old_index,json=oldIndex,proto3" json:"old_index,omitempty"`
	NewIndex uint32 `protobuf:"varint,2,opt,name=new_index,json=newIndex,proto3" json:"new_index,omitempty"`
}

func (m *EventConsensusSetUpdate) Reset()         { *m = EventConsensusSetUpdate{} }
func (m *EventConsensusSetUpdate) String() string { return proto.CompactTextString(m) }
func (*EventConsensusSetUpdate) ProtoMessage()    {}
func (*EventConsensusSetUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_486bfc4df1202b88, []int{3}
}
func (m *EventConsensusSetUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventConsensusSetUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventConsensusSetUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventConsensusSetUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventConsensusSetUpdate.Merge(m, src)
}
func (m *EventConsensusSetUpdate) XXX_Size() int {
	return m.Size()
}
func (m *EventConsensusSetUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventConsensusSetUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_EventConsensusSetUpdate proto.InternalMessageInfo

func (m *EventConsensusSetUpdate) GetOldIndex() uint32 {
	if m != nil {
		return m.OldIndex
	}
	return 0
}

func (m *EventConsensusSetUpdate) GetNewIndex() uint32 {
	if m != nil {
		return m.NewIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*EventGuardianSetUpdate)(nil), "certusone.wormholechain.wormhole.EventGuardianSetUpdate")
	proto.RegisterType((*EventPostedMessage)(nil), "certusone.wormholechain.wormhole.EventPostedMessage")
	proto.RegisterType((*EventGuardianRegistered)(nil), "certusone.wormholechain.wormhole.EventGuardianRegistered")
	proto.RegisterType((*EventConsensusSetUpdate)(nil), "certusone.wormholechain.wormhole.EventConsensusSetUpdate")
}

func init() { proto.RegisterFile("wormhole/events.proto", fileDescriptor_486bfc4df1202b88) }

var fileDescriptor_486bfc4df1202b88 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xbd, 0x4e, 0xeb, 0x30,
	0x14, 0x6e, 0x7a, 0xdb, 0x7b, 0x7b, 0x7d, 0xdb, 0xc5, 0xba, 0x40, 0x04, 0x52, 0x54, 0xc2, 0xc2,
	0x42, 0x33, 0xc0, 0x13, 0x80, 0x10, 0x42, 0x15, 0x08, 0xa5, 0x62, 0x61, 0xa9, 0xdc, 0xf8, 0x28,
	0xb5, 0x48, 0xec, 0x60, 0x9f, 0xb4, 0xcd, 0x4b, 0x20, 0x1e, 0x8b, 0xb1, 0x23, 0x23, 0x6a, 0x5f,
	0x04, 0xc5, 0x4d, 0x22, 0xb1, 0xb3, 0xf9, 0xfb, 0xf1, 0x77, 0x6c, 0x7f, 0x26, 0x7b, 0x4b, 0xa5,
	0xd3, 0xb9, 0x4a, 0x20, 0x80, 0x05, 0x48, 0x34, 0xa3, 0x4c, 0x2b, 0x54, 0x74, 0x18, 0x81, 0xc6,
	0xdc, 0x28, 0x09, 0xa3, 0xda, 0x10, 0xcd, 0x99, 0x90, 0x0d, 0xf2, 0x43, 0xb2, 0x7f, 0x5d, 0xee,
	0xb8, 0xc9, 0x99, 0xe6, 0x82, 0xc9, 0x09, 0xe0, 0x63, 0xc6, 0x19, 0x02, 0x3d, 0x22, 0x7f, 0x55,
	0xc2, 0xa7, 0x42, 0x72, 0x58, 0xb9, 0xce, 0xd0, 0x39, 0x1d, 0x84, 0x3d, 0x95, 0xf0, 0xdb, 0x12,
	0x97, 0xa2, 0x84, 0x65, 0x25, 0xb6, 0x77, 0xa2, 0x84, 0xa5, 0x15, 0xfd, 0x57, 0x87, 0x50, 0x1b,
	0xfa, 0xa0, 0x0c, 0x02, 0xbf, 0x03, 0x63, 0x58, 0x0c, 0xd4, 0x25, 0x7f, 0x20, 0x15, 0x88, 0xa0,
	0x6d, 0x5c, 0x3f, 0xac, 0x21, 0x3d, 0x24, 0x3d, 0x03, 0x2f, 0x39, 0xc8, 0x08, 0x6c, 0x58, 0x27,
	0x6c, 0x30, 0xfd, 0x4f, 0xba, 0x52, 0x95, 0xc2, 0x2f, 0x3b, 0x65, 0x07, 0x28, 0x25, 0x1d, 0x14,
	0x29, 0xb8, 0x1d, 0xeb, 0xb6, 0xeb, 0x32, 0x3f, 0x63, 0x45, 0xa2, 0x18, 0x77, 0xbb, 0xbb, 0xfc,
	0x0a, 0xfa, 0x8c, 0x1c, 0x7c, 0xbb, 0x64, 0x08, 0xb1, 0x30, 0x08, 0x1a, 0x38, 0x3d, 0x26, 0xfd,
	0xb8, 0x62, 0xa7, 0xcf, 0x50, 0x54, 0x27, 0xfb, 0x57, 0x73, 0x63, 0x28, 0xe8, 0x09, 0x19, 0x2c,
	0x58, 0x22, 0x38, 0x43, 0xa5, 0xad, 0xa7, 0x6d, 0x3d, 0xfd, 0x86, 0x1c, 0x43, 0xe1, 0x4f, 0xaa,
	0x11, 0x57, 0x4a, 0x1a, 0x90, 0x26, 0x37, 0x3f, 0xf0, 0x90, 0x97, 0xf7, 0xef, 0x1b, 0xcf, 0x59,
	0x6f, 0x3c, 0xe7, 0x73, 0xe3, 0x39, 0x6f, 0x5b, 0xaf, 0xb5, 0xde, 0x7a, 0xad, 0x8f, 0xad, 0xd7,
	0x7a, 0xba, 0x88, 0x05, 0xce, 0xf3, 0xd9, 0x28, 0x52, 0x69, 0xd0, 0x74, 0x1c, 0xd4, 0xad, 0x9e,
	0xd9, 0x92, 0x83, 0x55, 0x43, 0x04, 0x58, 0x64, 0x60, 0x66, 0xbf, 0xed, 0xaf, 0x38, 0xff, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x1e, 0x16, 0xf1, 0x8c, 0x2e, 0x02, 0x00, 0x00,
}

func (m *EventGuardianSetUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventGuardianSetUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventGuardianSetUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NewIndex != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.NewIndex))
		i--
		dAtA[i] = 0x10
	}
	if m.OldIndex != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.OldIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventPostedMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventPostedMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventPostedMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Time != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x20
	}
	if m.Nonce != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x18
	}
	if m.Sequence != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Emitter) > 0 {
		i -= len(m.Emitter)
		copy(dAtA[i:], m.Emitter)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Emitter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventGuardianRegistered) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventGuardianRegistered) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventGuardianRegistered) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorKey) > 0 {
		i -= len(m.ValidatorKey)
		copy(dAtA[i:], m.ValidatorKey)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ValidatorKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.GuardianKey) > 0 {
		i -= len(m.GuardianKey)
		copy(dAtA[i:], m.GuardianKey)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.GuardianKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventConsensusSetUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventConsensusSetUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventConsensusSetUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NewIndex != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.NewIndex))
		i--
		dAtA[i] = 0x10
	}
	if m.OldIndex != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.OldIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventGuardianSetUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OldIndex != 0 {
		n += 1 + sovEvents(uint64(m.OldIndex))
	}
	if m.NewIndex != 0 {
		n += 1 + sovEvents(uint64(m.NewIndex))
	}
	return n
}

func (m *EventPostedMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Emitter)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovEvents(uint64(m.Sequence))
	}
	if m.Nonce != 0 {
		n += 1 + sovEvents(uint64(m.Nonce))
	}
	if m.Time != 0 {
		n += 1 + sovEvents(uint64(m.Time))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventGuardianRegistered) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GuardianKey)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ValidatorKey)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventConsensusSetUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OldIndex != 0 {
		n += 1 + sovEvents(uint64(m.OldIndex))
	}
	if m.NewIndex != 0 {
		n += 1 + sovEvents(uint64(m.NewIndex))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventGuardianSetUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventGuardianSetUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventGuardianSetUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldIndex", wireType)
			}
			m.OldIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OldIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewIndex", wireType)
			}
			m.NewIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventPostedMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventPostedMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventPostedMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Emitter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Emitter = append(m.Emitter[:0], dAtA[iNdEx:postIndex]...)
			if m.Emitter == nil {
				m.Emitter = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventGuardianRegistered) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventGuardianRegistered: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventGuardianRegistered: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuardianKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GuardianKey = append(m.GuardianKey[:0], dAtA[iNdEx:postIndex]...)
			if m.GuardianKey == nil {
				m.GuardianKey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorKey = append(m.ValidatorKey[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorKey == nil {
				m.ValidatorKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventConsensusSetUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventConsensusSetUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventConsensusSetUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldIndex", wireType)
			}
			m.OldIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OldIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewIndex", wireType)
			}
			m.NewIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
