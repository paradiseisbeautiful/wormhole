// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wormhole/guardian_set.proto

package types

import (
	bytes "bytes"
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

type GuardianSet struct {
	Index          uint32   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Keys           [][]byte `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	ExpirationTime uint64   `protobuf:"varint,3,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
}

func (m *GuardianSet) Reset()         { *m = GuardianSet{} }
func (m *GuardianSet) String() string { return proto.CompactTextString(m) }
func (*GuardianSet) ProtoMessage()    {}
func (*GuardianSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6a773f49e89397, []int{0}
}
func (m *GuardianSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GuardianSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GuardianSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GuardianSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuardianSet.Merge(m, src)
}
func (m *GuardianSet) XXX_Size() int {
	return m.Size()
}
func (m *GuardianSet) XXX_DiscardUnknown() {
	xxx_messageInfo_GuardianSet.DiscardUnknown(m)
}

var xxx_messageInfo_GuardianSet proto.InternalMessageInfo

func (m *GuardianSet) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *GuardianSet) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *GuardianSet) GetExpirationTime() uint64 {
	if m != nil {
		return m.ExpirationTime
	}
	return 0
}

func init() {
	proto.RegisterType((*GuardianSet)(nil), "certusone.wormholechain.wormhole.GuardianSet")
}

func init() { proto.RegisterFile("wormhole/guardian_set.proto", fileDescriptor_3a6a773f49e89397) }

var fileDescriptor_3a6a773f49e89397 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0xcf, 0x2f, 0xca,
	0xcd, 0xc8, 0xcf, 0x49, 0xd5, 0x4f, 0x2f, 0x4d, 0x2c, 0x4a, 0xc9, 0x4c, 0xcc, 0x8b, 0x2f, 0x4e,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0x4e, 0x2d, 0x2a, 0x29, 0x2d, 0xce,
	0xcf, 0x4b, 0xd5, 0x83, 0x29, 0x4b, 0xce, 0x48, 0xcc, 0xcc, 0x83, 0xf3, 0xa4, 0x44, 0xd2, 0xf3,
	0xd3, 0xf3, 0xc1, 0x8a, 0xf5, 0x41, 0x2c, 0x88, 0x3e, 0xa5, 0x54, 0x2e, 0x6e, 0x77, 0xa8, 0x69,
	0xc1, 0xa9, 0x25, 0x42, 0x22, 0x5c, 0xac, 0x99, 0x79, 0x29, 0xa9, 0x15, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xbc, 0x41, 0x10, 0x8e, 0x90, 0x10, 0x17, 0x4b, 0x76, 0x6a, 0x65, 0xb1, 0x04, 0x93, 0x02,
	0xb3, 0x06, 0x4f, 0x10, 0x98, 0x2d, 0xa4, 0xc6, 0xc5, 0x97, 0x5a, 0x51, 0x90, 0x59, 0x94, 0x58,
	0x92, 0x99, 0x9f, 0x17, 0x92, 0x99, 0x9b, 0x2a, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x12, 0x84, 0x26,
	0x6a, 0xc5, 0xf2, 0x62, 0x81, 0x3c, 0xa3, 0x93, 0xdf, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9,
	0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e,
	0xcb, 0x31, 0x44, 0x99, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xc3,
	0xfd, 0xa0, 0x0f, 0x73, 0xb5, 0x2e, 0xd8, 0x13, 0xfa, 0x15, 0x70, 0x01, 0xfd, 0x92, 0xca, 0x82,
	0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xeb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x44, 0x6e,
	0x52, 0x14, 0x01, 0x00, 0x00,
}

func (this *GuardianSet) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GuardianSet)
	if !ok {
		that2, ok := that.(GuardianSet)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Index != that1.Index {
		return false
	}
	if len(this.Keys) != len(that1.Keys) {
		return false
	}
	for i := range this.Keys {
		if !bytes.Equal(this.Keys[i], that1.Keys[i]) {
			return false
		}
	}
	if this.ExpirationTime != that1.ExpirationTime {
		return false
	}
	return true
}
func (m *GuardianSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GuardianSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GuardianSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpirationTime != 0 {
		i = encodeVarintGuardianSet(dAtA, i, uint64(m.ExpirationTime))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Keys[iNdEx])
			copy(dAtA[i:], m.Keys[iNdEx])
			i = encodeVarintGuardianSet(dAtA, i, uint64(len(m.Keys[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Index != 0 {
		i = encodeVarintGuardianSet(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGuardianSet(dAtA []byte, offset int, v uint64) int {
	offset -= sovGuardianSet(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GuardianSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovGuardianSet(uint64(m.Index))
	}
	if len(m.Keys) > 0 {
		for _, b := range m.Keys {
			l = len(b)
			n += 1 + l + sovGuardianSet(uint64(l))
		}
	}
	if m.ExpirationTime != 0 {
		n += 1 + sovGuardianSet(uint64(m.ExpirationTime))
	}
	return n
}

func sovGuardianSet(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGuardianSet(x uint64) (n int) {
	return sovGuardianSet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GuardianSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuardianSet
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
			return fmt.Errorf("proto: GuardianSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GuardianSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardianSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardianSet
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
				return ErrInvalidLengthGuardianSet
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGuardianSet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, make([]byte, postIndex-iNdEx))
			copy(m.Keys[len(m.Keys)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationTime", wireType)
			}
			m.ExpirationTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardianSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpirationTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGuardianSet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGuardianSet
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
func skipGuardianSet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGuardianSet
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
					return 0, ErrIntOverflowGuardianSet
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
					return 0, ErrIntOverflowGuardianSet
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
				return 0, ErrInvalidLengthGuardianSet
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGuardianSet
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGuardianSet
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGuardianSet        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGuardianSet          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGuardianSet = fmt.Errorf("proto: unexpected end of group")
)
