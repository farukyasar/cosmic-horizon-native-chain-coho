// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: game/game.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Deposit struct {
	Address         string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Amount          github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	Staking         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=staking,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"staking"`
	Unbonding       github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=unbonding,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"unbonding"`
	RewardClaimTime time.Time                              `protobuf:"bytes,5,opt,name=reward_claim_time,json=rewardClaimTime,proto3,stdtime" json:"reward_claim_time"`
}

func (m *Deposit) Reset()         { *m = Deposit{} }
func (m *Deposit) String() string { return proto.CompactTextString(m) }
func (*Deposit) ProtoMessage()    {}
func (*Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a9278d664c0c01e, []int{0}
}
func (m *Deposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Deposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Deposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Deposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deposit.Merge(m, src)
}
func (m *Deposit) XXX_Size() int {
	return m.Size()
}
func (m *Deposit) XXX_DiscardUnknown() {
	xxx_messageInfo_Deposit.DiscardUnknown(m)
}

var xxx_messageInfo_Deposit proto.InternalMessageInfo

func (m *Deposit) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Deposit) GetRewardClaimTime() time.Time {
	if m != nil {
		return m.RewardClaimTime
	}
	return time.Time{}
}

type Unbonding struct {
	Id             uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StakerAddress  string                                 `protobuf:"bytes,2,opt,name=staker_address,json=stakerAddress,proto3" json:"staker_address,omitempty"`
	CreationHeight int64                                  `protobuf:"varint,3,opt,name=creation_height,json=creationHeight,proto3" json:"creation_height,omitempty"`
	CompletionTime time.Time                              `protobuf:"bytes,4,opt,name=completion_time,json=completionTime,proto3,stdtime" json:"completion_time"`
	Amount         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *Unbonding) Reset()         { *m = Unbonding{} }
func (m *Unbonding) String() string { return proto.CompactTextString(m) }
func (*Unbonding) ProtoMessage()    {}
func (*Unbonding) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a9278d664c0c01e, []int{1}
}
func (m *Unbonding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Unbonding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Unbonding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Unbonding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unbonding.Merge(m, src)
}
func (m *Unbonding) XXX_Size() int {
	return m.Size()
}
func (m *Unbonding) XXX_DiscardUnknown() {
	xxx_messageInfo_Unbonding.DiscardUnknown(m)
}

var xxx_messageInfo_Unbonding proto.InternalMessageInfo

func (m *Unbonding) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Unbonding) GetStakerAddress() string {
	if m != nil {
		return m.StakerAddress
	}
	return ""
}

func (m *Unbonding) GetCreationHeight() int64 {
	if m != nil {
		return m.CreationHeight
	}
	return 0
}

func (m *Unbonding) GetCompletionTime() time.Time {
	if m != nil {
		return m.CompletionTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Deposit)(nil), "cosmichorizon.coho.game.Deposit")
	proto.RegisterType((*Unbonding)(nil), "cosmichorizon.coho.game.Unbonding")
}

func init() { proto.RegisterFile("game/game.proto", fileDescriptor_2a9278d664c0c01e) }

var fileDescriptor_2a9278d664c0c01e = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcd, 0x6e, 0xd4, 0x30,
	0x14, 0x85, 0x27, 0xe9, 0xb4, 0xc3, 0x18, 0x31, 0x23, 0x22, 0x24, 0xa2, 0x59, 0x64, 0xaa, 0x4a,
	0x40, 0x25, 0x54, 0x5b, 0x85, 0x27, 0x60, 0xf8, 0x51, 0x91, 0x40, 0x42, 0x11, 0x6c, 0xd8, 0x8c,
	0x1c, 0xdb, 0x38, 0x56, 0xc7, 0xb9, 0x91, 0xed, 0xf0, 0xb7, 0xe4, 0x09, 0xfa, 0x58, 0x5d, 0x76,
	0x83, 0x84, 0x58, 0x14, 0x34, 0xf3, 0x22, 0xc8, 0x4e, 0x42, 0x51, 0x77, 0xb4, 0x9b, 0xc4, 0xbe,
	0xc7, 0xe7, 0xe4, 0xe6, 0xb3, 0x2e, 0x9a, 0x4a, 0xaa, 0x05, 0xf1, 0x0f, 0x5c, 0x1b, 0x70, 0x90,
	0xdc, 0x65, 0x60, 0xb5, 0x62, 0x25, 0x18, 0xf5, 0x15, 0x2a, 0xcc, 0xa0, 0x04, 0xec, 0xe5, 0x59,
	0xe6, 0x05, 0xb0, 0xa4, 0xa0, 0x56, 0x90, 0x8f, 0x87, 0x85, 0x70, 0xf4, 0x90, 0x30, 0x50, 0x55,
	0x6b, 0x9c, 0x65, 0x12, 0x40, 0xae, 0x04, 0x09, 0xbb, 0xa2, 0xf9, 0x40, 0x78, 0x63, 0xa8, 0x53,
	0xd0, 0xeb, 0xf3, 0xcb, 0xba, 0x53, 0x5a, 0x58, 0x47, 0x75, 0xdd, 0x1d, 0xb8, 0x23, 0x41, 0x42,
	0x58, 0x12, 0xbf, 0x6a, 0xab, 0x7b, 0xdf, 0x63, 0x34, 0x7a, 0x26, 0x6a, 0xb0, 0xca, 0x25, 0x29,
	0x1a, 0x51, 0xce, 0x8d, 0xb0, 0x36, 0x8d, 0x76, 0xa3, 0xfd, 0x71, 0xde, 0x6f, 0x93, 0x17, 0x68,
	0x87, 0x6a, 0x68, 0x2a, 0x97, 0xc6, 0x5e, 0x58, 0xe0, 0xd3, 0xf3, 0xf9, 0xe0, 0xe7, 0xf9, 0xfc,
	0xbe, 0x54, 0xae, 0x6c, 0x0a, 0xcc, 0x40, 0x93, 0xae, 0xff, 0xf6, 0x75, 0x60, 0xf9, 0x31, 0x71,
	0x5f, 0x6a, 0x61, 0xf1, 0xcb, 0xca, 0xe5, 0x9d, 0x3b, 0x39, 0x42, 0x23, 0xeb, 0xe8, 0xb1, 0xaa,
	0x64, 0xba, 0x75, 0xa5, 0xa0, 0xde, 0x9e, 0xbc, 0x42, 0xe3, 0xa6, 0x2a, 0xa0, 0xe2, 0x3e, 0x6b,
	0x78, 0xa5, 0xac, 0x8b, 0x80, 0xe4, 0x0d, 0xba, 0x6d, 0xc4, 0x27, 0x6a, 0xf8, 0x92, 0xad, 0xa8,
	0xd2, 0x4b, 0xcf, 0x2e, 0xdd, 0xde, 0x8d, 0xf6, 0x6f, 0x3e, 0x9a, 0xe1, 0x16, 0x2c, 0xee, 0xc1,
	0xe2, 0xb7, 0x3d, 0xd8, 0xc5, 0x0d, 0xff, 0xc5, 0x93, 0x5f, 0xf3, 0x28, 0x9f, 0xb6, 0xf6, 0xa7,
	0xde, 0xed, 0xf5, 0xbd, 0x6f, 0x31, 0x1a, 0xbf, 0xfb, 0x9b, 0x3f, 0x41, 0xb1, 0xe2, 0x01, 0xea,
	0x30, 0x8f, 0x15, 0x4f, 0xee, 0xa1, 0x89, 0xff, 0x11, 0x61, 0x96, 0x3d, 0xf0, 0xc0, 0x35, 0xbf,
	0xd5, 0x56, 0x9f, 0x74, 0xd8, 0x1f, 0xa0, 0x29, 0x33, 0x22, 0xdc, 0xf2, 0xb2, 0x14, 0x4a, 0x96,
	0x2e, 0x60, 0xdb, 0xca, 0x27, 0x7d, 0xf9, 0x28, 0x54, 0x93, 0xd7, 0x68, 0xca, 0x40, 0xd7, 0x2b,
	0x11, 0x8e, 0x86, 0xee, 0x87, 0xff, 0xd1, 0xfd, 0xe4, 0xc2, 0xec, 0xe5, 0x7f, 0xae, 0x7b, 0xfb,
	0x3a, 0xd7, 0xbd, 0x78, 0x7e, 0xba, 0xce, 0xa2, 0xb3, 0x75, 0x16, 0xfd, 0x5e, 0x67, 0xd1, 0xc9,
	0x26, 0x1b, 0x9c, 0x6d, 0xb2, 0xc1, 0x8f, 0x4d, 0x36, 0x78, 0xff, 0xf0, 0x52, 0x92, 0x62, 0x07,
	0xdd, 0x48, 0x10, 0x3f, 0x12, 0xe4, 0x73, 0x98, 0x99, 0x36, 0xb2, 0xd8, 0x09, 0xcd, 0x3f, 0xfe,
	0x13, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x40, 0xf9, 0xef, 0x4d, 0x03, 0x00, 0x00,
}

func (m *Deposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Deposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Deposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.RewardClaimTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.RewardClaimTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGame(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	{
		size := m.Unbonding.Size()
		i -= size
		if _, err := m.Unbonding.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGame(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Staking.Size()
		i -= size
		if _, err := m.Staking.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGame(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGame(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Unbonding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Unbonding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Unbonding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGame(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CompletionTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CompletionTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGame(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if m.CreationHeight != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.CreationHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.StakerAddress) > 0 {
		i -= len(m.StakerAddress)
		copy(dAtA[i:], m.StakerAddress)
		i = encodeVarintGame(dAtA, i, uint64(len(m.StakerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGame(dAtA []byte, offset int, v uint64) int {
	offset -= sovGame(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Deposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovGame(uint64(l))
	l = m.Staking.Size()
	n += 1 + l + sovGame(uint64(l))
	l = m.Unbonding.Size()
	n += 1 + l + sovGame(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.RewardClaimTime)
	n += 1 + l + sovGame(uint64(l))
	return n
}

func (m *Unbonding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovGame(uint64(m.Id))
	}
	l = len(m.StakerAddress)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	if m.CreationHeight != 0 {
		n += 1 + sovGame(uint64(m.CreationHeight))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CompletionTime)
	n += 1 + l + sovGame(uint64(l))
	l = m.Amount.Size()
	n += 1 + l + sovGame(uint64(l))
	return n
}

func sovGame(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGame(x uint64) (n int) {
	return sovGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Deposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: Deposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Deposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staking", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Staking.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unbonding", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Unbonding.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardClaimTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.RewardClaimTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGame
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
func (m *Unbonding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: Unbonding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Unbonding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationHeight", wireType)
			}
			m.CreationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CompletionTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGame
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
func skipGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGame
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
					return 0, ErrIntOverflowGame
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
					return 0, ErrIntOverflowGame
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
				return 0, ErrInvalidLengthGame
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGame
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGame
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGame        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGame          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGame = fmt.Errorf("proto: unexpected end of group")
)
