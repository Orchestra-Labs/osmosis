// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: symphony/valsetpref/v1beta1/state.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// ValidatorPreference defines the message structure for
// CreateValidatorSetPreference. It allows a user to set {val_addr, weight} in
// state. If a user does not have a validator set preference list set, and has
// staked, make their preference list default to their current staking
// distribution.
type ValidatorPreference struct {
	// val_oper_address holds the validator address the user wants to delegate
	// funds to.
	ValOperAddress string `protobuf:"bytes,1,opt,name=val_oper_address,json=valOperAddress,proto3" json:"val_oper_address,omitempty" yaml:"val_oper_address"`
	// weight is decimal between 0 and 1, and they all sum to 1.
	Weight cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=weight,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"weight"`
}

func (m *ValidatorPreference) Reset()         { *m = ValidatorPreference{} }
func (m *ValidatorPreference) String() string { return proto.CompactTextString(m) }
func (*ValidatorPreference) ProtoMessage()    {}
func (*ValidatorPreference) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd872c4c526620f8, []int{0}
}
func (m *ValidatorPreference) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorPreference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorPreference.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorPreference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorPreference.Merge(m, src)
}
func (m *ValidatorPreference) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorPreference) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorPreference.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorPreference proto.InternalMessageInfo

// ValidatorSetPreferences defines a delegator's validator set preference.
// It contains a list of (validator, percent_allocation) pairs.
// The percent allocation are arranged in decimal notation from 0 to 1 and must
// add up to 1.
type ValidatorSetPreferences struct {
	// preference holds {valAddr, weight} for the user who created it.
	Preferences []ValidatorPreference `protobuf:"bytes,2,rep,name=preferences,proto3" json:"preferences" yaml:"preferences"`
}

func (m *ValidatorSetPreferences) Reset()         { *m = ValidatorSetPreferences{} }
func (m *ValidatorSetPreferences) String() string { return proto.CompactTextString(m) }
func (*ValidatorSetPreferences) ProtoMessage()    {}
func (*ValidatorSetPreferences) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd872c4c526620f8, []int{1}
}
func (m *ValidatorSetPreferences) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSetPreferences) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSetPreferences.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSetPreferences) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSetPreferences.Merge(m, src)
}
func (m *ValidatorSetPreferences) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSetPreferences) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSetPreferences.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSetPreferences proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ValidatorPreference)(nil), "symphony.valsetpref.v1beta1.ValidatorPreference")
	proto.RegisterType((*ValidatorSetPreferences)(nil), "symphony.valsetpref.v1beta1.ValidatorSetPreferences")
}

func init() {
	proto.RegisterFile("symphony/valsetpref/v1beta1/state.proto", fileDescriptor_dd872c4c526620f8)
}

var fileDescriptor_dd872c4c526620f8 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0xdb, 0xbd, 0x30, 0x78, 0x3b, 0x10, 0xa9, 0xc2, 0xc6, 0x26, 0xed, 0xa8, 0x07, 0x77,
	0x59, 0xe2, 0x26, 0x28, 0xe8, 0xc9, 0xa1, 0x37, 0x41, 0x99, 0xe8, 0xc1, 0xcb, 0x48, 0xdb, 0xff,
	0xda, 0x62, 0xdb, 0x84, 0x24, 0x56, 0xfb, 0x11, 0xbc, 0x79, 0xf5, 0x1b, 0xed, 0xb8, 0xa3, 0x78,
	0x28, 0xba, 0x7d, 0x83, 0x7d, 0x02, 0x59, 0xbb, 0xb9, 0x21, 0xe2, 0x2d, 0xff, 0xe4, 0xf7, 0x24,
	0xcf, 0x93, 0x47, 0xdb, 0x13, 0x69, 0xc4, 0x7c, 0x1a, 0xa7, 0x38, 0x21, 0xa1, 0x00, 0xc9, 0x38,
	0x0c, 0x71, 0xd2, 0xb1, 0x41, 0x92, 0x0e, 0x16, 0x92, 0x48, 0x40, 0x8c, 0x53, 0x49, 0xf5, 0xc6,
	0x12, 0x44, 0x2b, 0x10, 0x2d, 0xc0, 0xfa, 0xb6, 0x47, 0x3d, 0x9a, 0x73, 0x78, 0xbe, 0x2a, 0x24,
	0xf5, 0x1d, 0x8f, 0x52, 0x2f, 0x04, 0x4c, 0x58, 0x80, 0x49, 0x1c, 0x53, 0x49, 0x64, 0x40, 0x63,
	0x51, 0x9c, 0x5a, 0xaf, 0xaa, 0xb6, 0x75, 0x4b, 0xc2, 0xc0, 0x25, 0x92, 0xf2, 0x2b, 0x0e, 0x43,
	0xe0, 0x10, 0x3b, 0xa0, 0x9f, 0x6b, 0x9b, 0x09, 0x09, 0x07, 0x94, 0x01, 0x1f, 0x10, 0xd7, 0xe5,
	0x20, 0x44, 0x4d, 0x6d, 0xaa, 0xad, 0xff, 0xbd, 0xc6, 0x2c, 0x33, 0xab, 0x29, 0x89, 0xc2, 0x63,
	0xeb, 0x27, 0x61, 0xf5, 0x37, 0x12, 0x12, 0x5e, 0x32, 0xe0, 0xa7, 0xc5, 0x86, 0x7e, 0xa2, 0x95,
	0x1f, 0x21, 0xf0, 0x7c, 0x59, 0x2b, 0xe5, 0xe2, 0xdd, 0x51, 0x66, 0x2a, 0xef, 0x99, 0xd9, 0x70,
	0xa8, 0x88, 0xa8, 0x10, 0xee, 0x3d, 0x0a, 0x28, 0x8e, 0x88, 0xf4, 0xd1, 0x05, 0x78, 0xc4, 0x49,
	0xcf, 0xc0, 0xe9, 0x2f, 0x24, 0xd6, 0xb3, 0xaa, 0x55, 0xbf, 0xbd, 0x5d, 0x83, 0x5c, 0xd9, 0x13,
	0x7a, 0xac, 0x55, 0xd8, 0x6a, 0xac, 0x95, 0x9a, 0xff, 0x5a, 0x95, 0xee, 0x3e, 0xfa, 0xe3, 0x7b,
	0xd0, 0x2f, 0x31, 0x7b, 0xf5, 0xb9, 0x9f, 0x59, 0x66, 0xea, 0x45, 0xa0, 0xb5, 0x2b, 0xad, 0xfe,
	0xfa, 0x03, 0xbd, 0x9b, 0xd1, 0xa7, 0xa1, 0x8c, 0x26, 0x86, 0x3a, 0x9e, 0x18, 0xea, 0xc7, 0xc4,
	0x50, 0x5f, 0xa6, 0x86, 0x32, 0x9e, 0x1a, 0xca, 0xdb, 0xd4, 0x50, 0xee, 0x8e, 0xbc, 0x40, 0xfa,
	0x0f, 0x36, 0x72, 0x68, 0x84, 0xf3, 0x60, 0x81, 0x68, 0x87, 0xc4, 0x16, 0xcb, 0x01, 0x27, 0xdd,
	0x43, 0xfc, 0xb4, 0x28, 0xb7, 0x9d, 0xb7, 0x2b, 0x53, 0x06, 0xc2, 0x2e, 0xe7, 0x2d, 0x1c, 0x7c,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x26, 0x7a, 0x13, 0x04, 0x01, 0x02, 0x00, 0x00,
}

func (m *ValidatorPreference) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorPreference) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorPreference) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ValOperAddress) > 0 {
		i -= len(m.ValOperAddress)
		copy(dAtA[i:], m.ValOperAddress)
		i = encodeVarintState(dAtA, i, uint64(len(m.ValOperAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorSetPreferences) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSetPreferences) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSetPreferences) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Preferences) > 0 {
		for iNdEx := len(m.Preferences) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Preferences[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintState(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorPreference) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValOperAddress)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovState(uint64(l))
	return n
}

func (m *ValidatorSetPreferences) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Preferences) > 0 {
		for _, e := range m.Preferences {
			l = e.Size()
			n += 1 + l + sovState(uint64(l))
		}
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorPreference) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: ValidatorPreference: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorPreference: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValOperAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValOperAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *ValidatorSetPreferences) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: ValidatorSetPreferences: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSetPreferences: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Preferences", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Preferences = append(m.Preferences, ValidatorPreference{})
			if err := m.Preferences[len(m.Preferences)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)
