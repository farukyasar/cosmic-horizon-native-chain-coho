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
	Owner            string                                 `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	DepositDenom     string                                 `protobuf:"bytes,2,opt,name=deposit_denom,json=depositDenom,proto3" json:"deposit_denom,omitempty"`
	StakingInflation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=staking_inflation,json=stakingInflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"staking_inflation"`
	UnstakingTime    time.Duration                          `protobuf:"bytes,4,opt,name=unstaking_time,json=unstakingTime,proto3,stdduration" json:"unstaking_time"`
	SwapFeeCollector string                                 `protobuf:"bytes,5,opt,name=swap_fee_collector,json=swapFeeCollector,proto3" json:"swap_fee_collector,omitempty"`
	SwapFee          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=swap_fee,json=swapFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_fee"`
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
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0xae, 0xd3, 0x30,
	0x14, 0x86, 0x93, 0xcb, 0xbd, 0xe5, 0x62, 0x28, 0x6a, 0xad, 0x0e, 0xa1, 0x43, 0x52, 0x81, 0x84,
	0x3a, 0xb4, 0xb6, 0x04, 0x1b, 0x63, 0xa9, 0x2a, 0x95, 0x09, 0x55, 0x4c, 0x30, 0x44, 0x69, 0xe2,
	0xa6, 0x56, 0x63, 0x9f, 0x10, 0x3b, 0x2a, 0xe5, 0x29, 0x18, 0x3b, 0xf2, 0x38, 0x1d, 0x3b, 0x22,
	0x86, 0x82, 0xda, 0x91, 0x97, 0x40, 0x76, 0x12, 0x86, 0x8e, 0x4c, 0xf6, 0x39, 0xff, 0xef, 0xdf,
	0xe7, 0x93, 0x0e, 0xea, 0xa6, 0x91, 0x60, 0x34, 0x8f, 0x8a, 0x48, 0x28, 0x92, 0x17, 0xa0, 0x01,
	0x7b, 0x31, 0x28, 0xc1, 0xe3, 0x35, 0x14, 0xfc, 0x2b, 0x48, 0xf2, 0x79, 0x0b, 0x3b, 0x49, 0x8c,
	0xad, 0xdf, 0x4b, 0x21, 0x05, 0x6b, 0xa2, 0xe6, 0x56, 0xf9, 0xfb, 0x7e, 0x0a, 0x90, 0x66, 0x8c,
	0xda, 0x6a, 0x59, 0xae, 0x68, 0x52, 0x16, 0x91, 0xe6, 0x20, 0x6b, 0x3d, 0xb8, 0xd6, 0x35, 0x17,
	0x4c, 0xe9, 0x48, 0xe4, 0x95, 0xe1, 0xf9, 0x9f, 0x1b, 0xd4, 0x7a, 0x6f, 0x27, 0xc0, 0x3d, 0x74,
	0x07, 0x5b, 0xc9, 0x0a, 0xcf, 0x1d, 0xb8, 0xc3, 0x47, 0x8b, 0xaa, 0xc0, 0x2f, 0x50, 0x3b, 0x61,
	0x39, 0x28, 0xae, 0xc3, 0x84, 0x49, 0x10, 0xde, 0x8d, 0x55, 0x9f, 0xd4, 0xcd, 0xa9, 0xe9, 0xe1,
	0x4f, 0xa8, 0xab, 0x74, 0xb4, 0xe1, 0x32, 0x0d, 0xb9, 0x5c, 0x65, 0x76, 0x02, 0xef, 0x81, 0x31,
	0x4e, 0xc8, 0xe1, 0x14, 0x38, 0x3f, 0x4f, 0xc1, 0xcb, 0x94, 0xeb, 0x75, 0xb9, 0x24, 0x31, 0x08,
	0x6a, 0x20, 0x41, 0xd5, 0xc7, 0x58, 0x25, 0x1b, 0xaa, 0x77, 0x39, 0x53, 0x64, 0xca, 0xe2, 0x45,
	0xa7, 0x0e, 0x9a, 0x37, 0x39, 0xf8, 0x1d, 0x7a, 0x5a, 0xca, 0x26, 0xde, 0xcc, 0xef, 0xdd, 0x0e,
	0xdc, 0xe1, 0xe3, 0x57, 0xcf, 0x48, 0x05, 0x47, 0x1a, 0x38, 0x32, 0xad, 0xe1, 0x27, 0xf7, 0xe6,
	0xd3, 0xfd, 0xaf, 0xc0, 0x5d, 0xb4, 0xff, 0x3d, 0xfd, 0xc0, 0x05, 0xc3, 0x23, 0x84, 0xd5, 0x36,
	0xca, 0xc3, 0x15, 0x63, 0x61, 0x0c, 0x59, 0xc6, 0x62, 0x0d, 0x85, 0x77, 0x67, 0x91, 0x3a, 0x46,
	0x99, 0x31, 0xf6, 0xb6, 0xe9, 0xe3, 0x39, 0xba, 0x6f, 0xdc, 0x5e, 0xeb, 0xbf, 0x68, 0x1e, 0xd6,
	0x99, 0x6f, 0x6e, 0xf7, 0xdf, 0x03, 0x67, 0x32, 0x3b, 0x9c, 0x7d, 0xf7, 0x78, 0xf6, 0xdd, 0xdf,
	0x67, 0xdf, 0xfd, 0x76, 0xf1, 0x9d, 0xe3, 0xc5, 0x77, 0x7e, 0x5c, 0x7c, 0xe7, 0xe3, 0xe8, 0x2a,
	0x90, 0xc7, 0xe3, 0x7a, 0x09, 0xa8, 0x5d, 0x02, 0xfa, 0x85, 0xda, 0x6d, 0xb1, 0xd1, 0xcb, 0x96,
	0x45, 0x7e, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0x41, 0xc9, 0x09, 0x0d, 0x42, 0x02, 0x00, 0x00,
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
	{
		size := m.StakingInflation.Size()
		i -= size
		if _, err := m.StakingInflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
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
	l = m.StakingInflation.Size()
	n += 1 + l + sovParams(uint64(l))
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingInflation", wireType)
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
			if err := m.StakingInflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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