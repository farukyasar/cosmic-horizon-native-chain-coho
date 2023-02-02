// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: game/params.proto

package types

import (
	fmt "fmt"
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

// Params defines the parameters for the module.
type Params struct {
	Owner            string                                  `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	DepositDenom     string                                  `protobuf:"bytes,2,opt,name=deposit_denom,json=depositDenom,proto3" json:"deposit_denom,omitempty"`
	StakingInflation uint64                                  `protobuf:"varint,3,opt,name=staking_inflation,json=stakingInflation,proto3" json:"staking_inflation,omitempty"`
	UnstakingTime    time.Duration                           `protobuf:"bytes,4,opt,name=unstaking_time,json=unstakingTime,proto3,stdduration" json:"unstaking_time"`
	SwapFeeCollector string                                  `protobuf:"bytes,5,opt,name=swap_fee_collector,json=swapFeeCollector,proto3" json:"swap_fee_collector,omitempty"`
	SwapFee          github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,6,opt,name=swap_fee,json=swapFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"swap_fee"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd39677cdca32c4, []int{0}
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

func (m *Params) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Params) GetDepositDenom() string {
	if m != nil {
		return m.DepositDenom
	}
	return ""
}

func (m *Params) GetStakingInflation() uint64 {
	if m != nil {
		return m.StakingInflation
	}
	return 0
}

func (m *Params) GetUnstakingTime() time.Duration {
	if m != nil {
		return m.UnstakingTime
	}
	return 0
}

func (m *Params) GetSwapFeeCollector() string {
	if m != nil {
		return m.SwapFeeCollector
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "cosmichorizon.qwoyn.game.Params")
}

func init() { proto.RegisterFile("game/params.proto", fileDescriptor_4fd39677cdca32c4) }

var fileDescriptor_4fd39677cdca32c4 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x31, 0x8f, 0xd3, 0x30,
	0x1c, 0xc5, 0xe3, 0xa3, 0x57, 0x0e, 0xc3, 0xa1, 0x3b, 0xeb, 0x86, 0x70, 0x43, 0x52, 0xc1, 0x40,
	0x25, 0xee, 0x6c, 0x09, 0x36, 0xc6, 0x5e, 0x55, 0x89, 0x4e, 0x28, 0x62, 0x62, 0x89, 0xd2, 0xc4,
	0x4d, 0xad, 0xc6, 0xfe, 0x87, 0xd8, 0x51, 0x29, 0x9f, 0x82, 0xb1, 0x23, 0x1f, 0x84, 0x0f, 0xd0,
	0xb1, 0x23, 0x62, 0x28, 0xa8, 0xfd, 0x22, 0xc8, 0x4e, 0xc2, 0xd0, 0x29, 0xf1, 0x7b, 0x3f, 0x3f,
	0x3d, 0xeb, 0xe1, 0xeb, 0x3c, 0x91, 0x9c, 0x95, 0x49, 0x95, 0x48, 0x4d, 0xcb, 0x0a, 0x0c, 0x10,
	0x3f, 0x05, 0x2d, 0x45, 0xba, 0x80, 0x4a, 0x7c, 0x03, 0x45, 0xbf, 0xac, 0x60, 0xad, 0xa8, 0xc5,
	0x6e, 0x6f, 0x72, 0xc8, 0xc1, 0x41, 0xcc, 0xfe, 0x35, 0xfc, 0x6d, 0x90, 0x03, 0xe4, 0x05, 0x67,
	0xee, 0x34, 0xab, 0xe7, 0x2c, 0xab, 0xab, 0xc4, 0x08, 0x50, 0xad, 0x1f, 0x9e, 0xfa, 0x46, 0x48,
	0xae, 0x4d, 0x22, 0xcb, 0x06, 0x78, 0xf9, 0xf3, 0x0c, 0xf7, 0x3f, 0xba, 0x06, 0xe4, 0x06, 0x9f,
	0xc3, 0x4a, 0xf1, 0xca, 0x47, 0x03, 0x34, 0x7c, 0x12, 0x35, 0x07, 0xf2, 0x0a, 0x5f, 0x66, 0xbc,
	0x04, 0x2d, 0x4c, 0x9c, 0x71, 0x05, 0xd2, 0x3f, 0x73, 0xee, 0xb3, 0x56, 0x1c, 0x5b, 0x8d, 0xbc,
	0xc1, 0xd7, 0xda, 0x24, 0x4b, 0xa1, 0xf2, 0x58, 0xa8, 0x79, 0xe1, 0x1a, 0xf8, 0x8f, 0x06, 0x68,
	0xd8, 0x8b, 0xae, 0x5a, 0xe3, 0x43, 0xa7, 0x93, 0x29, 0x7e, 0x5e, 0xab, 0x0e, 0xb7, 0x7d, 0xfc,
	0xde, 0x00, 0x0d, 0x9f, 0xbe, 0x7d, 0x41, 0x9b, 0xb2, 0xb4, 0x2b, 0x4b, 0xc7, 0xed, 0x63, 0x46,
	0x17, 0xdb, 0x7d, 0xe8, 0x6d, 0xfe, 0x84, 0x28, 0xba, 0xfc, 0x7f, 0xf5, 0x93, 0x90, 0x9c, 0xdc,
	0x61, 0xa2, 0x57, 0x49, 0x19, 0xcf, 0x39, 0x8f, 0x53, 0x28, 0x0a, 0x9e, 0x1a, 0xa8, 0xfc, 0x73,
	0x57, 0xf1, 0xca, 0x3a, 0x13, 0xce, 0x1f, 0x3a, 0x9d, 0x4c, 0xf1, 0x45, 0x47, 0xfb, 0x7d, 0xcb,
	0x8c, 0x98, 0x0d, 0xfe, 0xbd, 0x0f, 0x5f, 0xe7, 0xc2, 0x2c, 0xea, 0x19, 0x4d, 0x41, 0x32, 0x3b,
	0x01, 0xe8, 0xf6, 0x73, 0xaf, 0xb3, 0x25, 0x33, 0xeb, 0x92, 0x6b, 0xfa, 0x00, 0x42, 0x45, 0x8f,
	0xdb, 0xd0, 0xf7, 0xbd, 0xcd, 0x8f, 0xd0, 0x1b, 0x4d, 0xb6, 0x87, 0x00, 0xed, 0x0e, 0x01, 0xfa,
	0x7b, 0x08, 0xd0, 0xf7, 0x63, 0xe0, 0xed, 0x8e, 0x81, 0xf7, 0xeb, 0x18, 0x78, 0x9f, 0xef, 0x4e,
	0x12, 0x45, 0x7a, 0xdf, 0xae, 0xca, 0xdc, 0xaa, 0xec, 0x2b, 0x73, 0xf3, 0xbb, 0xec, 0x59, 0xdf,
	0xbd, 0xf9, 0xdd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x09, 0xb5, 0xc4, 0x13, 0x02, 0x00,
	0x00,
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
		size := m.SwapFee.Size()
		i -= size
		if _, err := m.SwapFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.SwapFeeCollector) > 0 {
		i -= len(m.SwapFeeCollector)
		copy(dAtA[i:], m.SwapFeeCollector)
		i = encodeVarintParams(dAtA, i, uint64(len(m.SwapFeeCollector)))
		i--
		dAtA[i] = 0x2a
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.UnstakingTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.UnstakingTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.StakingInflation != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.StakingInflation))
		i--
		dAtA[i] = 0x18
	}
	if len(m.DepositDenom) > 0 {
		i -= len(m.DepositDenom)
		copy(dAtA[i:], m.DepositDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DepositDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
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
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DepositDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.StakingInflation != 0 {
		n += 1 + sovParams(uint64(m.StakingInflation))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.UnstakingTime)
	n += 1 + l + sovParams(uint64(l))
	l = len(m.SwapFeeCollector)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.SwapFee.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingInflation", wireType)
			}
			m.StakingInflation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StakingInflation |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnstakingTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.UnstakingTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFeeCollector", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SwapFeeCollector = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
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
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
