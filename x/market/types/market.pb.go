// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: symphony/market/v1beta1/market.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params defines the parameters for the market module.
type Params struct {
	ExchangePool cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=exchange_pool,json=exchangePool,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"exchange_pool" yaml:"exchange_pool"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bc7086493ef9038, []int{0}
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
	proto.RegisterType((*Params)(nil), "symphony.market.v1beta1.Params")
}

func init() {
	proto.RegisterFile("symphony/market/v1beta1/market.proto", fileDescriptor_8bc7086493ef9038)
}

var fileDescriptor_8bc7086493ef9038 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0xae, 0xcc, 0x2d,
	0xc8, 0xc8, 0xcf, 0xab, 0xd4, 0xcf, 0x4d, 0x2c, 0xca, 0x4e, 0x2d, 0xd1, 0x2f, 0x33, 0x4c, 0x4a,
	0x2d, 0x49, 0x34, 0x84, 0x72, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0x61, 0xaa, 0xf4,
	0xa0, 0xc2, 0x50, 0x55, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x35, 0xfa, 0x20, 0x16, 0x44,
	0xb9, 0x52, 0x09, 0x17, 0x5b, 0x40, 0x62, 0x51, 0x62, 0x6e, 0xb1, 0x50, 0x02, 0x17, 0x6f, 0x6a,
	0x45, 0x72, 0x46, 0x62, 0x5e, 0x7a, 0x6a, 0x7c, 0x41, 0x7e, 0x7e, 0x8e, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x8f, 0x93, 0xf5, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9, 0x4b, 0x27, 0xe7, 0x17, 0xe7,
	0xe6, 0x17, 0x17, 0xa7, 0x64, 0xeb, 0x65, 0xe6, 0xeb, 0xe7, 0x26, 0x96, 0x64, 0xe8, 0xf9, 0xa4,
	0xa6, 0x27, 0x26, 0x57, 0xba, 0xa4, 0x26, 0x7f, 0xba, 0x27, 0x2f, 0x52, 0x99, 0x98, 0x9b, 0x63,
	0xa5, 0x84, 0x62, 0x82, 0x52, 0x10, 0x0f, 0x8c, 0x1f, 0x90, 0x9f, 0x9f, 0x63, 0xc5, 0x31, 0x63,
	0x81, 0x3c, 0xc3, 0x8b, 0x05, 0xf2, 0x8c, 0x4e, 0x3e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0x65, 0x94, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f,
	0xb6, 0x30, 0xb3, 0x58, 0x37, 0x27, 0x31, 0xa9, 0x18, 0xc6, 0xd1, 0x2f, 0x33, 0x32, 0xd3, 0xaf,
	0x80, 0x85, 0x40, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x2b, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x53, 0x2d, 0x6e, 0x5b, 0x21, 0x01, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if !this.ExchangePool.Equal(that1.ExchangePool) {
		return false
	}
	return true
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
		size := m.ExchangePool.Size()
		i -= size
		if _, err := m.ExchangePool.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMarket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMarket(dAtA []byte, offset int, v uint64) int {
	offset -= sovMarket(v)
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
	l = m.ExchangePool.Size()
	n += 1 + l + sovMarket(uint64(l))
	return n
}

func sovMarket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMarket(x uint64) (n int) {
	return sovMarket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarket
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
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangePool", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExchangePool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarket
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
func skipMarket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMarket
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
					return 0, ErrIntOverflowMarket
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
					return 0, ErrIntOverflowMarket
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
				return 0, ErrInvalidLengthMarket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMarket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMarket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMarket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMarket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMarket = fmt.Errorf("proto: unexpected end of group")
)
