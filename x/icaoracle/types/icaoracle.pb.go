// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/icaoracle/icaoracle.proto

package types

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

type Oracle struct {
	ChainId         string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ConnectionId    string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	ChannelId       string `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	PortId          string `protobuf:"bytes,4,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	IcaAddress      string `protobuf:"bytes,5,opt,name=ica_address,json=icaAddress,proto3" json:"ica_address,omitempty"`
	ContractAddress string `protobuf:"bytes,6,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	ContractCodeId  uint64 `protobuf:"varint,7,opt,name=contract_code_id,json=contractCodeId,proto3" json:"contract_code_id,omitempty"`
	Active          bool   `protobuf:"varint,8,opt,name=active,proto3" json:"active,omitempty"`
}

func (m *Oracle) Reset()         { *m = Oracle{} }
func (m *Oracle) String() string { return proto.CompactTextString(m) }
func (*Oracle) ProtoMessage()    {}
func (*Oracle) Descriptor() ([]byte, []int) {
	return fileDescriptor_842e38c1f0da9e66, []int{0}
}
func (m *Oracle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Oracle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Oracle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Oracle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Oracle.Merge(m, src)
}
func (m *Oracle) XXX_Size() int {
	return m.Size()
}
func (m *Oracle) XXX_DiscardUnknown() {
	xxx_messageInfo_Oracle.DiscardUnknown(m)
}

var xxx_messageInfo_Oracle proto.InternalMessageInfo

func (m *Oracle) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *Oracle) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *Oracle) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *Oracle) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *Oracle) GetIcaAddress() string {
	if m != nil {
		return m.IcaAddress
	}
	return ""
}

func (m *Oracle) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *Oracle) GetContractCodeId() uint64 {
	if m != nil {
		return m.ContractCodeId
	}
	return 0
}

func (m *Oracle) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type Metadata struct {
	UpdateTime   string `protobuf:"bytes,1,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	UpdateHeight string `protobuf:"bytes,2,opt,name=update_height,json=updateHeight,proto3" json:"update_height,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_842e38c1f0da9e66, []int{1}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *Metadata) GetUpdateHeight() string {
	if m != nil {
		return m.UpdateHeight
	}
	return ""
}

type Metric struct {
	Key      string    `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value    string    `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Metadata *Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_842e38c1f0da9e66, []int{2}
}
func (m *Metric) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return m.Size()
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Metric) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Metric) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type PendingMetricUpdate struct {
	Metric        *Metric `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	OracleChainId string  `protobuf:"bytes,2,opt,name=oracle_chain_id,json=oracleChainId,proto3" json:"oracle_chain_id,omitempty"`
}

func (m *PendingMetricUpdate) Reset()         { *m = PendingMetricUpdate{} }
func (m *PendingMetricUpdate) String() string { return proto.CompactTextString(m) }
func (*PendingMetricUpdate) ProtoMessage()    {}
func (*PendingMetricUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_842e38c1f0da9e66, []int{3}
}
func (m *PendingMetricUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingMetricUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingMetricUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingMetricUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingMetricUpdate.Merge(m, src)
}
func (m *PendingMetricUpdate) XXX_Size() int {
	return m.Size()
}
func (m *PendingMetricUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingMetricUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_PendingMetricUpdate proto.InternalMessageInfo

func (m *PendingMetricUpdate) GetMetric() *Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (m *PendingMetricUpdate) GetOracleChainId() string {
	if m != nil {
		return m.OracleChainId
	}
	return ""
}

func init() {
	proto.RegisterType((*Oracle)(nil), "stride.icaoracle.Oracle")
	proto.RegisterType((*Metadata)(nil), "stride.icaoracle.Metadata")
	proto.RegisterType((*Metric)(nil), "stride.icaoracle.Metric")
	proto.RegisterType((*PendingMetricUpdate)(nil), "stride.icaoracle.PendingMetricUpdate")
}

func init() { proto.RegisterFile("stride/icaoracle/icaoracle.proto", fileDescriptor_842e38c1f0da9e66) }

var fileDescriptor_842e38c1f0da9e66 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x9b, 0xfd, 0x93, 0x66, 0xa7, 0x2c, 0x5b, 0x19, 0x04, 0x05, 0x89, 0x50, 0x15, 0x09,
	0x95, 0x03, 0x29, 0x2a, 0x82, 0x3b, 0xec, 0x85, 0x48, 0xbb, 0x62, 0x15, 0xe0, 0xc2, 0x25, 0x72,
	0xed, 0x51, 0x63, 0xd1, 0xd8, 0x91, 0xe3, 0x56, 0xec, 0x4b, 0x20, 0x1e, 0x8b, 0xe3, 0x1e, 0x39,
	0xa2, 0xf6, 0x45, 0x50, 0x6c, 0x27, 0x54, 0x68, 0x6f, 0x33, 0xdf, 0xf7, 0x8b, 0x67, 0x34, 0xf9,
	0x60, 0x5c, 0x1b, 0x2d, 0x38, 0xce, 0x04, 0xa3, 0x4a, 0x53, 0xb6, 0xda, 0xab, 0x92, 0x4a, 0x2b,
	0xa3, 0xc8, 0xd0, 0x11, 0x49, 0xa7, 0x4f, 0x7e, 0x1c, 0x40, 0xf8, 0xd1, 0x96, 0xe4, 0x11, 0x44,
	0xac, 0xa0, 0x42, 0xe6, 0x82, 0x8f, 0x82, 0x71, 0x30, 0x3d, 0xc9, 0xfa, 0xb6, 0x4f, 0x39, 0x79,
	0x06, 0xa7, 0x4c, 0x49, 0x89, 0xcc, 0x08, 0x65, 0xfd, 0x03, 0xeb, 0xdf, 0xf9, 0x27, 0xa6, 0x9c,
	0x3c, 0x01, 0x60, 0x05, 0x95, 0x12, 0x57, 0x0d, 0x71, 0x68, 0x89, 0x13, 0xaf, 0xa4, 0x9c, 0x3c,
	0x84, 0x7e, 0xa5, 0xb4, 0x69, 0xbc, 0x23, 0xeb, 0x85, 0x4d, 0x9b, 0x72, 0xf2, 0x14, 0x06, 0x82,
	0xd1, 0x9c, 0x72, 0xae, 0xb1, 0xae, 0x47, 0xc7, 0xd6, 0x04, 0xc1, 0xe8, 0x3b, 0xa7, 0x90, 0x17,
	0x30, 0x64, 0x4a, 0x1a, 0x4d, 0x99, 0xe9, 0xa8, 0xd0, 0x52, 0x67, 0xad, 0xde, 0xa2, 0xd3, 0x3d,
	0x94, 0x29, 0x8e, 0xcd, 0xb4, 0xfe, 0x38, 0x98, 0x1e, 0x65, 0x77, 0x5b, 0xfd, 0x5c, 0x71, 0x4c,
	0x39, 0x79, 0x00, 0x21, 0x65, 0x46, 0x6c, 0x70, 0x14, 0x8d, 0x83, 0x69, 0x94, 0xf9, 0x6e, 0x72,
	0x05, 0xd1, 0x25, 0x1a, 0xca, 0xa9, 0xa1, 0xcd, 0x66, 0xeb, 0x8a, 0x53, 0x83, 0xb9, 0x11, 0x25,
	0xfa, 0xa3, 0x80, 0x93, 0x3e, 0x8b, 0x12, 0x9b, 0xbb, 0x78, 0xa0, 0x40, 0xb1, 0x2c, 0x4c, 0x7b,
	0x17, 0x27, 0x7e, 0xb0, 0xda, 0xa4, 0x80, 0xf0, 0x12, 0x8d, 0x16, 0x8c, 0x0c, 0xe1, 0xf0, 0x1b,
	0x5e, 0xfb, 0x77, 0x9a, 0x92, 0xdc, 0x87, 0xe3, 0x0d, 0x5d, 0xad, 0xd1, 0x7f, 0xe8, 0x1a, 0xf2,
	0x16, 0xa2, 0xd2, 0xef, 0x60, 0xef, 0x38, 0x98, 0x3f, 0x4e, 0xfe, 0xff, 0x73, 0x49, 0xbb, 0x65,
	0xd6, 0xb1, 0x13, 0x05, 0xf7, 0xae, 0x50, 0x72, 0x21, 0x97, 0x6e, 0xe0, 0x17, 0xbb, 0x06, 0x79,
	0x05, 0x61, 0x69, 0x7b, 0x3b, 0x79, 0x30, 0x1f, 0xdd, 0xfa, 0x98, 0x16, 0x2c, 0xf3, 0x1c, 0x79,
	0x0e, 0x67, 0xce, 0xc8, 0xbb, 0x44, 0xb8, 0x05, 0x4f, 0x9d, 0x7c, 0xee, 0x72, 0xf1, 0xfe, 0xe2,
	0xd7, 0x36, 0x0e, 0x6e, 0xb6, 0x71, 0xf0, 0x67, 0x1b, 0x07, 0x3f, 0x77, 0x71, 0xef, 0x66, 0x17,
	0xf7, 0x7e, 0xef, 0xe2, 0xde, 0xd7, 0xf9, 0x52, 0x98, 0x62, 0xbd, 0x48, 0x98, 0x2a, 0x67, 0x9f,
	0xec, 0xb4, 0x97, 0x17, 0x74, 0x51, 0xcf, 0x7c, 0x44, 0x37, 0x6f, 0x66, 0xdf, 0xf7, 0x72, 0x6a,
	0xae, 0x2b, 0xac, 0x17, 0xa1, 0x0d, 0xe9, 0xeb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0xdc,
	0xc3, 0x11, 0xc8, 0x02, 0x00, 0x00,
}

func (m *Oracle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Oracle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Oracle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.ContractCodeId != 0 {
		i = encodeVarintIcaoracle(dAtA, i, uint64(m.ContractCodeId))
		i--
		dAtA[i] = 0x38
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.IcaAddress) > 0 {
		i -= len(m.IcaAddress)
		copy(dAtA[i:], m.IcaAddress)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.IcaAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UpdateHeight) > 0 {
		i -= len(m.UpdateHeight)
		copy(dAtA[i:], m.UpdateHeight)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.UpdateHeight)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UpdateTime) > 0 {
		i -= len(m.UpdateTime)
		copy(dAtA[i:], m.UpdateTime)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.UpdateTime)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Metric) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metric) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metric) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIcaoracle(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PendingMetricUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingMetricUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingMetricUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OracleChainId) > 0 {
		i -= len(m.OracleChainId)
		copy(dAtA[i:], m.OracleChainId)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.OracleChainId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Metric != nil {
		{
			size, err := m.Metric.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIcaoracle(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIcaoracle(dAtA []byte, offset int, v uint64) int {
	offset -= sovIcaoracle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Oracle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.IcaAddress)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	if m.ContractCodeId != 0 {
		n += 1 + sovIcaoracle(uint64(m.ContractCodeId))
	}
	if m.Active {
		n += 2
	}
	return n
}

func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UpdateTime)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.UpdateHeight)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	return n
}

func (m *Metric) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	return n
}

func (m *PendingMetricUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metric != nil {
		l = m.Metric.Size()
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.OracleChainId)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	return n
}

func sovIcaoracle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIcaoracle(x uint64) (n int) {
	return sovIcaoracle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Oracle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIcaoracle
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
			return fmt.Errorf("proto: Oracle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Oracle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
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
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
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
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
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
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
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
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IcaAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
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
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IcaAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
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
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractCodeId", wireType)
			}
			m.ContractCodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ContractCodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Active = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipIcaoracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIcaoracle
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
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIcaoracle
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
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
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdateTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
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
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdateHeight = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIcaoracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIcaoracle
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
func (m *Metric) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIcaoracle
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
			return fmt.Errorf("proto: Metric: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metric: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
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
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
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
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
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
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIcaoracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIcaoracle
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
func (m *PendingMetricUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIcaoracle
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
			return fmt.Errorf("proto: PendingMetricUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingMetricUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metric", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
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
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metric == nil {
				m.Metric = &Metric{}
			}
			if err := m.Metric.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
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
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIcaoracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIcaoracle
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
func skipIcaoracle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIcaoracle
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
					return 0, ErrIntOverflowIcaoracle
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
					return 0, ErrIntOverflowIcaoracle
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
				return 0, ErrInvalidLengthIcaoracle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIcaoracle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIcaoracle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIcaoracle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIcaoracle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIcaoracle = fmt.Errorf("proto: unexpected end of group")
)