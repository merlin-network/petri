// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: contractmanager/genesis.proto

package types

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

// Failure message contains information about ACK failures and can be used to
// replay ACK in case of requirement.
type Failure struct {
	// ChannelId
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// Address of the failed contract
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// id of the failure under specific address
	Id uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// ACK id to restore
	AckId uint64 `protobuf:"varint,4,opt,name=ack_id,json=ackId,proto3" json:"ack_id,omitempty"`
	// Acknowledgement type
	AckType string `protobuf:"bytes,5,opt,name=ack_type,json=ackType,proto3" json:"ack_type,omitempty"`
}

func (m *Failure) Reset()         { *m = Failure{} }
func (m *Failure) String() string { return proto.CompactTextString(m) }
func (*Failure) ProtoMessage()    {}
func (*Failure) Descriptor() ([]byte, []int) {
	return fileDescriptor_c23af9b9805fb076, []int{0}
}
func (m *Failure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Failure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Failure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Failure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Failure.Merge(m, src)
}
func (m *Failure) XXX_Size() int {
	return m.Size()
}
func (m *Failure) XXX_DiscardUnknown() {
	xxx_messageInfo_Failure.DiscardUnknown(m)
}

var xxx_messageInfo_Failure proto.InternalMessageInfo

func (m *Failure) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *Failure) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Failure) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Failure) GetAckId() uint64 {
	if m != nil {
		return m.AckId
	}
	return 0
}

func (m *Failure) GetAckType() string {
	if m != nil {
		return m.AckType
	}
	return ""
}

// GenesisState defines the contractmanager module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// List of the contract failures
	FailuresList []Failure `protobuf:"bytes,2,rep,name=failures_list,json=failuresList,proto3" json:"failures_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c23af9b9805fb076, []int{1}
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

func (m *GenesisState) GetFailuresList() []Failure {
	if m != nil {
		return m.FailuresList
	}
	return nil
}

func init() {
	proto.RegisterType((*Failure)(nil), "petri.contractmanager.Failure")
	proto.RegisterType((*GenesisState)(nil), "petri.contractmanager.GenesisState")
}

func init() { proto.RegisterFile("contractmanager/genesis.proto", fileDescriptor_c23af9b9805fb076) }

var fileDescriptor_c23af9b9805fb076 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x3f, 0x4f, 0x02, 0x31,
	0x18, 0xc6, 0xaf, 0xc7, 0x3f, 0x29, 0xe8, 0xd0, 0x68, 0x3c, 0x89, 0x1c, 0x17, 0x26, 0x16, 0xef,
	0x12, 0x4c, 0xdc, 0x5c, 0x18, 0x34, 0x44, 0x07, 0x82, 0x4e, 0x2e, 0xa4, 0xb4, 0xf5, 0x68, 0x80,
	0xf6, 0xd2, 0x96, 0x44, 0x76, 0x3f, 0x80, 0xb3, 0x9f, 0x88, 0x91, 0xd1, 0xc9, 0x18, 0xf8, 0x22,
	0xe6, 0xee, 0xca, 0x72, 0x09, 0xdb, 0xdb, 0xa7, 0xcf, 0xfb, 0xcb, 0xfb, 0x3c, 0xb0, 0x4d, 0xa4,
	0x30, 0x0a, 0x13, 0xb3, 0xc4, 0x02, 0xc7, 0x4c, 0x45, 0x31, 0x13, 0x4c, 0x73, 0x1d, 0x26, 0x4a,
	0x1a, 0x89, 0x2e, 0x05, 0x5b, 0x19, 0x25, 0x45, 0x58, 0xb0, 0xb5, 0xce, 0x63, 0x19, 0xcb, 0xcc,
	0x13, 0xa5, 0x53, 0x6e, 0x6f, 0x5d, 0x17, 0x69, 0x09, 0x56, 0x78, 0x69, 0x61, 0xdd, 0x4f, 0x00,
	0x6b, 0x0f, 0x98, 0x2f, 0x56, 0x8a, 0xa1, 0x36, 0x84, 0x64, 0x86, 0x85, 0x60, 0x8b, 0x09, 0xa7,
	0x1e, 0x08, 0x40, 0xaf, 0x3e, 0xae, 0x5b, 0x65, 0x48, 0x91, 0x07, 0x6b, 0x98, 0x52, 0xc5, 0xb4,
	0xf6, 0xdc, 0xec, 0xef, 0xf0, 0x44, 0x67, 0xd0, 0xe5, 0xd4, 0x2b, 0x05, 0xa0, 0x57, 0x1e, 0xbb,
	0x9c, 0xa2, 0x0b, 0x58, 0xc5, 0x64, 0x9e, 0x42, 0xca, 0x99, 0x56, 0xc1, 0x64, 0x3e, 0xa4, 0xe8,
	0x0a, 0x9e, 0xa4, 0xb2, 0x59, 0x27, 0xcc, 0xab, 0x58, 0x02, 0x99, 0xbf, 0xae, 0x13, 0xd6, 0xfd,
	0x06, 0xb0, 0xf9, 0x98, 0xa7, 0x7c, 0x31, 0xd8, 0x30, 0x74, 0x0f, 0xab, 0xf9, 0x9d, 0xd9, 0x1d,
	0x8d, 0x7e, 0x27, 0x3c, 0x92, 0x3a, 0x1c, 0x65, 0xb6, 0x41, 0x79, 0xf3, 0xdb, 0x71, 0xc6, 0x76,
	0x09, 0x3d, 0xc1, 0xd3, 0xf7, 0x3c, 0x95, 0x9e, 0x2c, 0xb8, 0x36, 0x9e, 0x1b, 0x94, 0x7a, 0x8d,
	0x7e, 0x70, 0x94, 0x62, 0x3b, 0xb0, 0x98, 0xe6, 0x61, 0xf9, 0x99, 0x6b, 0x33, 0x18, 0x6d, 0x76,
	0x3e, 0xd8, 0xee, 0x7c, 0xf0, 0xb7, 0xf3, 0xc1, 0xd7, 0xde, 0x77, 0xb6, 0x7b, 0xdf, 0xf9, 0xd9,
	0xfb, 0xce, 0xdb, 0x5d, 0xcc, 0xcd, 0x6c, 0x35, 0x0d, 0x89, 0x5c, 0x46, 0x96, 0x7c, 0x23, 0x55,
	0x7c, 0x98, 0xa3, 0x8f, 0xa8, 0x58, 0x7e, 0x1a, 0x5e, 0x4f, 0xab, 0x59, 0xf9, 0xb7, 0xff, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x91, 0x99, 0x4b, 0xa1, 0xea, 0x01, 0x00, 0x00,
}

func (m *Failure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Failure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Failure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AckType) > 0 {
		i -= len(m.AckType)
		copy(dAtA[i:], m.AckType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AckType)))
		i--
		dAtA[i] = 0x2a
	}
	if m.AckId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AckId))
		i--
		dAtA[i] = 0x20
	}
	if m.Id != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	if len(m.FailuresList) > 0 {
		for iNdEx := len(m.FailuresList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FailuresList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *Failure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovGenesis(uint64(m.Id))
	}
	if m.AckId != 0 {
		n += 1 + sovGenesis(uint64(m.AckId))
	}
	l = len(m.AckType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FailuresList) > 0 {
		for _, e := range m.FailuresList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Failure) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Failure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Failure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
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
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckId", wireType)
			}
			m.AckId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AckId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckType", wireType)
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
			m.AckType = string(dAtA[iNdEx:postIndex])
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
				return fmt.Errorf("proto: wrong wireType = %d for field FailuresList", wireType)
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
			m.FailuresList = append(m.FailuresList, Failure{})
			if err := m.FailuresList[len(m.FailuresList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
