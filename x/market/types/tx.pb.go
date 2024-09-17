// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/market/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgSwap represents a message to swap coin to another denom.
type MsgSwap struct {
	Trader    string     `protobuf:"bytes,1,opt,name=trader,proto3" json:"trader,omitempty" yaml:"trader"`
	OfferCoin types.Coin `protobuf:"bytes,2,opt,name=offer_coin,json=offerCoin,proto3" json:"offer_coin" yaml:"offer_coin"`
	AskDenom  string     `protobuf:"bytes,3,opt,name=ask_denom,json=askDenom,proto3" json:"ask_denom,omitempty" yaml:"ask_denom"`
}

func (m *MsgSwap) Reset()         { *m = MsgSwap{} }
func (m *MsgSwap) String() string { return proto.CompactTextString(m) }
func (*MsgSwap) ProtoMessage()    {}
func (*MsgSwap) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b04bdc246eaa07, []int{0}
}
func (m *MsgSwap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwap.Merge(m, src)
}
func (m *MsgSwap) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwap) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwap.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwap proto.InternalMessageInfo

// MsgSwapResponse defines the Msg/Swap response type.
type MsgSwapResponse struct {
	SwapCoin types.Coin `protobuf:"bytes,1,opt,name=swap_coin,json=swapCoin,proto3" json:"swap_coin" yaml:"swap_coin"`
	SwapFee  types.Coin `protobuf:"bytes,2,opt,name=swap_fee,json=swapFee,proto3" json:"swap_fee" yaml:"swap_fee"`
}

func (m *MsgSwapResponse) Reset()         { *m = MsgSwapResponse{} }
func (m *MsgSwapResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSwapResponse) ProtoMessage()    {}
func (*MsgSwapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b04bdc246eaa07, []int{1}
}
func (m *MsgSwapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapResponse.Merge(m, src)
}
func (m *MsgSwapResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapResponse proto.InternalMessageInfo

func (m *MsgSwapResponse) GetSwapCoin() types.Coin {
	if m != nil {
		return m.SwapCoin
	}
	return types.Coin{}
}

func (m *MsgSwapResponse) GetSwapFee() types.Coin {
	if m != nil {
		return m.SwapFee
	}
	return types.Coin{}
}

// MsgSwapSend represents a message to swap coin and send all result coin to
// recipient
type MsgSwapSend struct {
	FromAddress string     `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	ToAddress   string     `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty" yaml:"to_address"`
	OfferCoin   types.Coin `protobuf:"bytes,3,opt,name=offer_coin,json=offerCoin,proto3" json:"offer_coin" yaml:"offer_coin"`
	AskDenom    string     `protobuf:"bytes,4,opt,name=ask_denom,json=askDenom,proto3" json:"ask_denom,omitempty" yaml:"ask_denom"`
}

func (m *MsgSwapSend) Reset()         { *m = MsgSwapSend{} }
func (m *MsgSwapSend) String() string { return proto.CompactTextString(m) }
func (*MsgSwapSend) ProtoMessage()    {}
func (*MsgSwapSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b04bdc246eaa07, []int{2}
}
func (m *MsgSwapSend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapSend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapSend.Merge(m, src)
}
func (m *MsgSwapSend) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapSend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapSend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapSend proto.InternalMessageInfo

// MsgSwapSendResponse defines the Msg/SwapSend response type.
type MsgSwapSendResponse struct {
	SwapCoin types.Coin `protobuf:"bytes,1,opt,name=swap_coin,json=swapCoin,proto3" json:"swap_coin" yaml:"swap_coin"`
	SwapFee  types.Coin `protobuf:"bytes,2,opt,name=swap_fee,json=swapFee,proto3" json:"swap_fee" yaml:"swap_fee"`
}

func (m *MsgSwapSendResponse) Reset()         { *m = MsgSwapSendResponse{} }
func (m *MsgSwapSendResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSwapSendResponse) ProtoMessage()    {}
func (*MsgSwapSendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b04bdc246eaa07, []int{3}
}
func (m *MsgSwapSendResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapSendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapSendResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapSendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapSendResponse.Merge(m, src)
}
func (m *MsgSwapSendResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapSendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapSendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapSendResponse proto.InternalMessageInfo

func (m *MsgSwapSendResponse) GetSwapCoin() types.Coin {
	if m != nil {
		return m.SwapCoin
	}
	return types.Coin{}
}

func (m *MsgSwapSendResponse) GetSwapFee() types.Coin {
	if m != nil {
		return m.SwapFee
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*MsgSwap)(nil), "osmosis.market.v1beta1.MsgSwap")
	proto.RegisterType((*MsgSwapResponse)(nil), "osmosis.market.v1beta1.MsgSwapResponse")
	proto.RegisterType((*MsgSwapSend)(nil), "osmosis.market.v1beta1.MsgSwapSend")
	proto.RegisterType((*MsgSwapSendResponse)(nil), "osmosis.market.v1beta1.MsgSwapSendResponse")
}

func init() { proto.RegisterFile("osmosis/market/v1beta1/tx.proto", fileDescriptor_91b04bdc246eaa07) }

var fileDescriptor_91b04bdc246eaa07 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xf6, 0x26, 0x55, 0x1b, 0x6f, 0x40, 0xa5, 0x6e, 0xa1, 0x69, 0x0e, 0x76, 0xb5, 0x1c, 0x28,
	0x42, 0xd8, 0x4a, 0x40, 0x1c, 0x72, 0x23, 0x20, 0x4e, 0x44, 0xaa, 0x9c, 0x1b, 0x42, 0x8a, 0xd6,
	0xf1, 0xda, 0x44, 0xa9, 0xbd, 0x96, 0x77, 0xe9, 0xcf, 0x1b, 0x70, 0x84, 0x37, 0xe8, 0x33, 0x20,
	0xc1, 0x23, 0xa0, 0x1e, 0x7b, 0xe4, 0x64, 0xa1, 0xe4, 0xc2, 0x39, 0x4f, 0x80, 0xf6, 0xc7, 0x4e,
	0x90, 0x90, 0x02, 0x07, 0x0e, 0xbd, 0xcd, 0x78, 0xbe, 0xef, 0x9b, 0x6f, 0x76, 0xbc, 0x0b, 0x1d,
	0xca, 0x12, 0xca, 0x26, 0xcc, 0x4b, 0x70, 0x3e, 0x25, 0xdc, 0x3b, 0xed, 0x04, 0x84, 0xe3, 0x8e,
	0xc7, 0xcf, 0xdd, 0x2c, 0xa7, 0x9c, 0x5a, 0xf7, 0x34, 0xc0, 0x55, 0x00, 0x57, 0x03, 0xda, 0x7b,
	0x31, 0x8d, 0xa9, 0x84, 0x78, 0x22, 0x52, 0xe8, 0xb6, 0x3d, 0x96, 0x70, 0x2f, 0xc0, 0x8c, 0x54,
	0x5a, 0x63, 0x3a, 0x49, 0x55, 0x1d, 0x7d, 0x03, 0x70, 0x6b, 0xc0, 0xe2, 0xe1, 0x19, 0xce, 0xac,
	0x87, 0x70, 0x93, 0xe7, 0x38, 0x24, 0x79, 0x0b, 0x1c, 0x82, 0x23, 0xb3, 0xbf, 0xb3, 0x28, 0x9c,
	0xdb, 0x17, 0x38, 0x39, 0xe9, 0x21, 0xf5, 0x1d, 0xf9, 0x1a, 0x60, 0x0d, 0x21, 0xa4, 0x51, 0x44,
	0xf2, 0x91, 0x90, 0x6a, 0xd5, 0x0e, 0xc1, 0x51, 0xb3, 0x7b, 0xe0, 0xaa, 0x5e, 0xae, 0xe8, 0x55,
	0xda, 0x72, 0x5f, 0xd0, 0x49, 0xda, 0x3f, 0xb8, 0x2a, 0x1c, 0x63, 0x51, 0x38, 0x3b, 0x4a, 0x6d,
	0x49, 0x45, 0xbe, 0x29, 0x13, 0x81, 0xb2, 0x3a, 0xd0, 0xc4, 0x6c, 0x3a, 0x0a, 0x49, 0x4a, 0x93,
	0x56, 0x5d, 0x5a, 0xd8, 0x5b, 0x14, 0xce, 0x1d, 0x45, 0xaa, 0x4a, 0xc8, 0x6f, 0x60, 0x36, 0x7d,
	0x29, 0xc2, 0x5e, 0xe3, 0xc3, 0xa5, 0x63, 0xfc, 0xbc, 0x74, 0x0c, 0xf4, 0x19, 0xc0, 0x6d, 0x3d,
	0x88, 0x4f, 0x58, 0x46, 0x53, 0x46, 0xac, 0x63, 0x68, 0xb2, 0x33, 0x9c, 0x29, 0x93, 0x60, 0x9d,
	0xc9, 0x96, 0x36, 0xa9, 0xfb, 0x55, 0x4c, 0xe4, 0x37, 0x44, 0x2c, 0x2d, 0x0e, 0xa0, 0x8c, 0x47,
	0x11, 0x21, 0xeb, 0xa7, 0xde, 0xd7, 0x82, 0xdb, 0x2b, 0x82, 0x11, 0x21, 0xc8, 0xdf, 0x12, 0xe1,
	0x2b, 0x42, 0xd0, 0xa7, 0x1a, 0x6c, 0x6a, 0xd3, 0x43, 0x92, 0x86, 0x56, 0x0f, 0xde, 0x8a, 0x72,
	0x9a, 0x8c, 0x70, 0x18, 0xe6, 0x84, 0x31, 0xbd, 0x87, 0xfd, 0x45, 0xe1, 0xec, 0x2a, 0x8d, 0xd5,
	0x2a, 0xf2, 0x9b, 0x22, 0x7d, 0xae, 0x32, 0xeb, 0x29, 0x84, 0x9c, 0x56, 0xcc, 0x9a, 0x64, 0xde,
	0x5d, 0x9e, 0xf9, 0xb2, 0x86, 0x7c, 0x93, 0xd3, 0x92, 0xf5, 0xfb, 0x22, 0xeb, 0xff, 0x61, 0x91,
	0x1b, 0xff, 0xb8, 0xc8, 0xaf, 0x00, 0xee, 0xae, 0x9c, 0xc9, 0x8d, 0x59, 0x66, 0xf7, 0x0b, 0x80,
	0xf5, 0x01, 0x8b, 0xad, 0x63, 0xb8, 0x21, 0xaf, 0x93, 0xe3, 0xfe, 0xf9, 0xa6, 0xba, 0x7a, 0xba,
	0xf6, 0x83, 0x35, 0x80, 0x6a, 0xf4, 0xb7, 0xb0, 0x51, 0xfd, 0x22, 0xf7, 0xd7, 0x90, 0x04, 0xa8,
	0xfd, 0xe8, 0x2f, 0x40, 0xa5, 0x7a, 0xff, 0xf5, 0xd5, 0xcc, 0x06, 0xd7, 0x33, 0x1b, 0xfc, 0x98,
	0xd9, 0xe0, 0xe3, 0xdc, 0x36, 0xae, 0xe7, 0xb6, 0xf1, 0x7d, 0x6e, 0x1b, 0x6f, 0xba, 0xf1, 0x84,
	0xbf, 0x7b, 0x1f, 0xb8, 0x63, 0x9a, 0x78, 0x5a, 0xf0, 0xf1, 0x09, 0x0e, 0x58, 0x99, 0x78, 0xa7,
	0xdd, 0x67, 0xde, 0x79, 0xf9, 0x52, 0xf1, 0x8b, 0x8c, 0xb0, 0x60, 0x53, 0xbe, 0x2b, 0x4f, 0x7e,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x0f, 0xc2, 0xc0, 0xc8, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Swap defines a method for swapping coin from one denom to another
	// denom.
	Swap(ctx context.Context, in *MsgSwap, opts ...grpc.CallOption) (*MsgSwapResponse, error)
	// SwapSend defines a method for swapping and sending coin from a account to
	// other account.
	SwapSend(ctx context.Context, in *MsgSwapSend, opts ...grpc.CallOption) (*MsgSwapSendResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Swap(ctx context.Context, in *MsgSwap, opts ...grpc.CallOption) (*MsgSwapResponse, error) {
	out := new(MsgSwapResponse)
	err := c.cc.Invoke(ctx, "/osmosis.market.v1beta1.Msg/Swap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SwapSend(ctx context.Context, in *MsgSwapSend, opts ...grpc.CallOption) (*MsgSwapSendResponse, error) {
	out := new(MsgSwapSendResponse)
	err := c.cc.Invoke(ctx, "/osmosis.market.v1beta1.Msg/SwapSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Swap defines a method for swapping coin from one denom to another
	// denom.
	Swap(context.Context, *MsgSwap) (*MsgSwapResponse, error)
	// SwapSend defines a method for swapping and sending coin from a account to
	// other account.
	SwapSend(context.Context, *MsgSwapSend) (*MsgSwapSendResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Swap(ctx context.Context, req *MsgSwap) (*MsgSwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Swap not implemented")
}
func (*UnimplementedMsgServer) SwapSend(ctx context.Context, req *MsgSwapSend) (*MsgSwapSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapSend not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Swap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Swap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.market.v1beta1.Msg/Swap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Swap(ctx, req.(*MsgSwap))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SwapSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapSend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.market.v1beta1.Msg/SwapSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapSend(ctx, req.(*MsgSwapSend))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.market.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Swap",
			Handler:    _Msg_Swap_Handler,
		},
		{
			MethodName: "SwapSend",
			Handler:    _Msg_SwapSend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/market/v1beta1/tx.proto",
}

func (m *MsgSwap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AskDenom) > 0 {
		i -= len(m.AskDenom)
		copy(dAtA[i:], m.AskDenom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AskDenom)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.OfferCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Trader) > 0 {
		i -= len(m.Trader)
		copy(dAtA[i:], m.Trader)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Trader)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSwapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SwapFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.SwapCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgSwapSend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapSend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapSend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AskDenom) > 0 {
		i -= len(m.AskDenom)
		copy(dAtA[i:], m.AskDenom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AskDenom)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.OfferCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSwapSendResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapSendResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapSendResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SwapFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.SwapCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSwap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Trader)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.OfferCoin.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.AskDenom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSwapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SwapCoin.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.SwapFee.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSwapSend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.OfferCoin.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.AskDenom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSwapSendResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SwapCoin.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.SwapFee.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSwap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSwap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trader", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Trader = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OfferCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AskDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AskDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSwapResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSwapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSwapSend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSwapSend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapSend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OfferCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AskDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AskDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSwapSendResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSwapSendResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapSendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
