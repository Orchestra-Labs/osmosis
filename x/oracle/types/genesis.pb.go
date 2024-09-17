// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/oracle/v1beta1/genesis.proto

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

// GenesisState defines the oracle module's genesis state.
type GenesisState struct {
	Params                        Params                         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	FeederDelegations             []FeederDelegation             `protobuf:"bytes,2,rep,name=feeder_delegations,json=feederDelegations,proto3" json:"feeder_delegations"`
	ExchangeRates                 ExchangeRateTuples             `protobuf:"bytes,3,rep,name=exchange_rates,json=exchangeRates,proto3,castrepeated=ExchangeRateTuples" json:"exchange_rates"`
	MissCounters                  []MissCounter                  `protobuf:"bytes,4,rep,name=miss_counters,json=missCounters,proto3" json:"miss_counters"`
	AggregateExchangeRatePrevotes []AggregateExchangeRatePrevote `protobuf:"bytes,5,rep,name=aggregate_exchange_rate_prevotes,json=aggregateExchangeRatePrevotes,proto3" json:"aggregate_exchange_rate_prevotes"`
	AggregateExchangeRateVotes    []AggregateExchangeRateVote    `protobuf:"bytes,6,rep,name=aggregate_exchange_rate_votes,json=aggregateExchangeRateVotes,proto3" json:"aggregate_exchange_rate_votes"`
	TobinTaxes                    []TobinTax                     `protobuf:"bytes,7,rep,name=tobin_taxes,json=tobinTaxes,proto3" json:"tobin_taxes"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_00d991d274be17e0, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetFeederDelegations() []FeederDelegation {
	if m != nil {
		return m.FeederDelegations
	}
	return nil
}

func (m *GenesisState) GetExchangeRates() ExchangeRateTuples {
	if m != nil {
		return m.ExchangeRates
	}
	return nil
}

func (m *GenesisState) GetMissCounters() []MissCounter {
	if m != nil {
		return m.MissCounters
	}
	return nil
}

func (m *GenesisState) GetAggregateExchangeRatePrevotes() []AggregateExchangeRatePrevote {
	if m != nil {
		return m.AggregateExchangeRatePrevotes
	}
	return nil
}

func (m *GenesisState) GetAggregateExchangeRateVotes() []AggregateExchangeRateVote {
	if m != nil {
		return m.AggregateExchangeRateVotes
	}
	return nil
}

func (m *GenesisState) GetTobinTaxes() []TobinTax {
	if m != nil {
		return m.TobinTaxes
	}
	return nil
}

// FeederDelegation is the address for where oracle feeder authority are
// delegated to. By default this struct is only used at genesis to feed in
// default feeder addresses.
type FeederDelegation struct {
	FeederAddress    string `protobuf:"bytes,1,opt,name=feeder_address,json=feederAddress,proto3" json:"feeder_address,omitempty"`
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

func (m *FeederDelegation) Reset()         { *m = FeederDelegation{} }
func (m *FeederDelegation) String() string { return proto.CompactTextString(m) }
func (*FeederDelegation) ProtoMessage()    {}
func (*FeederDelegation) Descriptor() ([]byte, []int) {
	return fileDescriptor_00d991d274be17e0, []int{1}
}
func (m *FeederDelegation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeederDelegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeederDelegation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeederDelegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeederDelegation.Merge(m, src)
}
func (m *FeederDelegation) XXX_Size() int {
	return m.Size()
}
func (m *FeederDelegation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeederDelegation.DiscardUnknown(m)
}

var xxx_messageInfo_FeederDelegation proto.InternalMessageInfo

func (m *FeederDelegation) GetFeederAddress() string {
	if m != nil {
		return m.FeederAddress
	}
	return ""
}

func (m *FeederDelegation) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

// MissCounter defines an miss counter and validator address pair used in
// oracle module's genesis state
type MissCounter struct {
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	MissCounter      uint64 `protobuf:"varint,2,opt,name=miss_counter,json=missCounter,proto3" json:"miss_counter,omitempty"`
}

func (m *MissCounter) Reset()         { *m = MissCounter{} }
func (m *MissCounter) String() string { return proto.CompactTextString(m) }
func (*MissCounter) ProtoMessage()    {}
func (*MissCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_00d991d274be17e0, []int{2}
}
func (m *MissCounter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MissCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MissCounter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MissCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissCounter.Merge(m, src)
}
func (m *MissCounter) XXX_Size() int {
	return m.Size()
}
func (m *MissCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_MissCounter.DiscardUnknown(m)
}

var xxx_messageInfo_MissCounter proto.InternalMessageInfo

func (m *MissCounter) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *MissCounter) GetMissCounter() uint64 {
	if m != nil {
		return m.MissCounter
	}
	return 0
}

// TobinTax defines an denom and tobin_tax pair used in
// oracle module's genesis state
type TobinTax struct {
	Denom    string                      `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	TobinTax cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=tobin_tax,json=tobinTax,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"tobin_tax"`
}

func (m *TobinTax) Reset()         { *m = TobinTax{} }
func (m *TobinTax) String() string { return proto.CompactTextString(m) }
func (*TobinTax) ProtoMessage()    {}
func (*TobinTax) Descriptor() ([]byte, []int) {
	return fileDescriptor_00d991d274be17e0, []int{3}
}
func (m *TobinTax) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TobinTax) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TobinTax.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TobinTax) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TobinTax.Merge(m, src)
}
func (m *TobinTax) XXX_Size() int {
	return m.Size()
}
func (m *TobinTax) XXX_DiscardUnknown() {
	xxx_messageInfo_TobinTax.DiscardUnknown(m)
}

var xxx_messageInfo_TobinTax proto.InternalMessageInfo

func (m *TobinTax) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "osmosis.oracle.v1beta1.GenesisState")
	proto.RegisterType((*FeederDelegation)(nil), "osmosis.oracle.v1beta1.FeederDelegation")
	proto.RegisterType((*MissCounter)(nil), "osmosis.oracle.v1beta1.MissCounter")
	proto.RegisterType((*TobinTax)(nil), "osmosis.oracle.v1beta1.TobinTax")
}

func init() {
	proto.RegisterFile("osmosis/oracle/v1beta1/genesis.proto", fileDescriptor_00d991d274be17e0)
}

var fileDescriptor_00d991d274be17e0 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x6e, 0xd3, 0x4e,
	0x14, 0xc5, 0xe3, 0x7e, 0xfd, 0xdb, 0x49, 0x53, 0xb5, 0xa3, 0xea, 0xaf, 0x28, 0xa8, 0x4e, 0x48,
	0x40, 0x0a, 0x42, 0xd8, 0x4a, 0x40, 0xac, 0x58, 0xd0, 0x50, 0xe8, 0xa6, 0xa0, 0xca, 0x44, 0x2c,
	0x90, 0x2a, 0x6b, 0x6c, 0xdf, 0x38, 0x16, 0xb6, 0x27, 0xf2, 0x9d, 0x44, 0x29, 0x5b, 0x5e, 0x80,
	0xe7, 0x60, 0xcd, 0x43, 0x74, 0xd9, 0x25, 0x62, 0x51, 0x50, 0xf2, 0x22, 0xc8, 0xe3, 0xc9, 0x07,
	0x25, 0x46, 0x62, 0xe7, 0xb9, 0xfe, 0x9d, 0x73, 0xee, 0x1d, 0x5f, 0x99, 0xdc, 0xe3, 0x18, 0x71,
	0x0c, 0xd0, 0xe4, 0x09, 0x73, 0x43, 0x30, 0x47, 0x2d, 0x07, 0x04, 0x6b, 0x99, 0x3e, 0xc4, 0x80,
	0x01, 0x1a, 0x83, 0x84, 0x0b, 0x4e, 0xff, 0x57, 0x94, 0x91, 0x51, 0x86, 0xa2, 0x2a, 0x87, 0x3e,
	0xf7, 0xb9, 0x44, 0xcc, 0xf4, 0x29, 0xa3, 0x2b, 0x8d, 0x1c, 0x4f, 0x25, 0x96, 0x50, 0xfd, 0xeb,
	0x26, 0xd9, 0x3d, 0xcd, 0x42, 0xde, 0x0a, 0x26, 0x80, 0x3e, 0x23, 0x5b, 0x03, 0x96, 0xb0, 0x08,
	0xcb, 0x5a, 0x4d, 0x6b, 0x16, 0xdb, 0xba, 0xb1, 0x3a, 0xd4, 0x38, 0x97, 0x54, 0x67, 0xe3, 0xea,
	0xa6, 0x5a, 0xb0, 0x94, 0x86, 0x5e, 0x10, 0xda, 0x03, 0xf0, 0x20, 0xb1, 0x3d, 0x08, 0xc1, 0x67,
	0x22, 0xe0, 0x31, 0x96, 0xd7, 0x6a, 0xeb, 0xcd, 0x62, 0xbb, 0x99, 0xe7, 0xf4, 0x4a, 0x2a, 0x4e,
	0xe6, 0x02, 0xe5, 0x79, 0xd0, 0xbb, 0x55, 0x47, 0x1a, 0x92, 0x3d, 0x18, 0xbb, 0x7d, 0x16, 0xfb,
	0x60, 0x27, 0x4c, 0x00, 0x96, 0xd7, 0xa5, 0xf5, 0x83, 0x3c, 0xeb, 0x97, 0x8a, 0xb6, 0x98, 0x80,
	0xee, 0x70, 0x10, 0x42, 0xa7, 0x92, 0x7a, 0x7f, 0xf9, 0x51, 0xa5, 0x7f, 0xbc, 0x42, 0xab, 0x04,
	0x4b, 0x35, 0xa4, 0x6f, 0x48, 0x29, 0x0a, 0x10, 0x6d, 0x97, 0x0f, 0x63, 0x01, 0x09, 0x96, 0x37,
	0x64, 0x58, 0x23, 0x2f, 0xec, 0x75, 0x80, 0xf8, 0x22, 0x63, 0xd5, 0x08, 0xbb, 0xd1, 0xa2, 0x84,
	0xf4, 0x93, 0x46, 0x6a, 0xcc, 0xf7, 0x93, 0x74, 0x1c, 0xb0, 0x7f, 0x1b, 0xc4, 0x1e, 0x24, 0x30,
	0xe2, 0xe9, 0x40, 0x9b, 0x32, 0xe3, 0x49, 0x5e, 0xc6, 0xf1, 0x4c, 0xbf, 0xdc, 0xfe, 0x79, 0x26,
	0x56, 0xa1, 0x47, 0xec, 0x2f, 0x0c, 0xd2, 0x8f, 0xe4, 0x28, 0xaf, 0x89, 0xac, 0x83, 0x2d, 0xd9,
	0x41, 0xeb, 0x9f, 0x3a, 0x78, 0xb7, 0x88, 0xaf, 0xb0, 0x3c, 0x00, 0xe9, 0x29, 0x29, 0x0a, 0xee,
	0x04, 0xb1, 0x2d, 0xd8, 0x18, 0xb0, 0xfc, 0x9f, 0x4c, 0xaa, 0xe5, 0x25, 0x75, 0x53, 0xb4, 0xcb,
	0xc6, 0xca, 0x98, 0x08, 0x75, 0x06, 0xac, 0xf7, 0xc8, 0xfe, 0xed, 0xad, 0xa1, 0xf7, 0xc9, 0x9e,
	0xda, 0x3d, 0xe6, 0x79, 0x09, 0x60, 0xb6, 0xc1, 0x3b, 0x56, 0x29, 0xab, 0x1e, 0x67, 0x45, 0xfa,
	0x90, 0x1c, 0x8c, 0x58, 0x18, 0x78, 0x4c, 0xf0, 0x05, 0xb9, 0x26, 0xc9, 0xfd, 0xf9, 0x0b, 0x05,
	0xd7, 0x2f, 0x48, 0x71, 0xe9, 0xab, 0xae, 0xd6, 0x6a, 0xab, 0xb5, 0xf4, 0x2e, 0xd9, 0x5d, 0x5e,
	0x1f, 0x99, 0xb1, 0x61, 0x15, 0x97, 0x56, 0xa2, 0xee, 0x90, 0xed, 0xd9, 0x90, 0xf4, 0x90, 0x6c,
	0x7a, 0x10, 0xf3, 0x48, 0xf9, 0x65, 0x07, 0xfa, 0x9c, 0xec, 0xcc, 0x6f, 0x2c, 0xeb, 0xb2, 0xd3,
	0x48, 0x6f, 0xe3, 0xfb, 0x4d, 0xf5, 0x8e, 0x2b, 0xef, 0x0d, 0xbd, 0x0f, 0x46, 0xc0, 0xcd, 0x88,
	0x89, 0xbe, 0x71, 0x06, 0x3e, 0x73, 0x2f, 0x4f, 0xc0, 0xb5, 0xb6, 0x67, 0x97, 0xd5, 0x39, 0xbb,
	0x9a, 0xe8, 0xda, 0xf5, 0x44, 0xd7, 0x7e, 0x4e, 0x74, 0xed, 0xf3, 0x54, 0x2f, 0x5c, 0x4f, 0xf5,
	0xc2, 0xb7, 0xa9, 0x5e, 0x78, 0xdf, 0xf6, 0x03, 0xd1, 0x1f, 0x3a, 0x86, 0xcb, 0x23, 0x53, 0x7d,
	0x82, 0x47, 0x21, 0x73, 0x70, 0x76, 0x30, 0x47, 0xed, 0xa7, 0xe6, 0x78, 0xf6, 0xfb, 0x10, 0x97,
	0x03, 0x40, 0x67, 0x4b, 0xfe, 0x36, 0x1e, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x76, 0x50, 0xf2,
	0x51, 0xb1, 0x04, 0x00, 0x00,
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
	if len(m.TobinTaxes) > 0 {
		for iNdEx := len(m.TobinTaxes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TobinTaxes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.AggregateExchangeRateVotes) > 0 {
		for iNdEx := len(m.AggregateExchangeRateVotes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AggregateExchangeRateVotes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.AggregateExchangeRatePrevotes) > 0 {
		for iNdEx := len(m.AggregateExchangeRatePrevotes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AggregateExchangeRatePrevotes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.MissCounters) > 0 {
		for iNdEx := len(m.MissCounters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MissCounters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ExchangeRates) > 0 {
		for iNdEx := len(m.ExchangeRates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExchangeRates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.FeederDelegations) > 0 {
		for iNdEx := len(m.FeederDelegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeederDelegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
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

func (m *FeederDelegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeederDelegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeederDelegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FeederAddress) > 0 {
		i -= len(m.FeederAddress)
		copy(dAtA[i:], m.FeederAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.FeederAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MissCounter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MissCounter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MissCounter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MissCounter != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MissCounter))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TobinTax) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TobinTax) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TobinTax) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TobinTax.Size()
		i -= size
		if _, err := m.TobinTax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
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
	if len(m.FeederDelegations) > 0 {
		for _, e := range m.FeederDelegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ExchangeRates) > 0 {
		for _, e := range m.ExchangeRates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MissCounters) > 0 {
		for _, e := range m.MissCounters {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AggregateExchangeRatePrevotes) > 0 {
		for _, e := range m.AggregateExchangeRatePrevotes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AggregateExchangeRateVotes) > 0 {
		for _, e := range m.AggregateExchangeRateVotes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TobinTaxes) > 0 {
		for _, e := range m.TobinTaxes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *FeederDelegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FeederAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *MissCounter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MissCounter != 0 {
		n += 1 + sovGenesis(uint64(m.MissCounter))
	}
	return n
}

func (m *TobinTax) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.TobinTax.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field FeederDelegations", wireType)
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
			m.FeederDelegations = append(m.FeederDelegations, FeederDelegation{})
			if err := m.FeederDelegations[len(m.FeederDelegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangeRates", wireType)
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
			m.ExchangeRates = append(m.ExchangeRates, ExchangeRateTuple{})
			if err := m.ExchangeRates[len(m.ExchangeRates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissCounters", wireType)
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
			m.MissCounters = append(m.MissCounters, MissCounter{})
			if err := m.MissCounters[len(m.MissCounters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateExchangeRatePrevotes", wireType)
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
			m.AggregateExchangeRatePrevotes = append(m.AggregateExchangeRatePrevotes, AggregateExchangeRatePrevote{})
			if err := m.AggregateExchangeRatePrevotes[len(m.AggregateExchangeRatePrevotes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateExchangeRateVotes", wireType)
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
			m.AggregateExchangeRateVotes = append(m.AggregateExchangeRateVotes, AggregateExchangeRateVote{})
			if err := m.AggregateExchangeRateVotes[len(m.AggregateExchangeRateVotes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TobinTaxes", wireType)
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
			m.TobinTaxes = append(m.TobinTaxes, TobinTax{})
			if err := m.TobinTaxes[len(m.TobinTaxes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *FeederDelegation) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: FeederDelegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeederDelegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeederAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeederAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
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
func (m *MissCounter) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MissCounter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MissCounter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissCounter", wireType)
			}
			m.MissCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissCounter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *TobinTax) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TobinTax: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TobinTax: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TobinTax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TobinTax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
