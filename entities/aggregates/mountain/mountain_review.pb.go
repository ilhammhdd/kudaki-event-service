// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/mountain/mountain_review.proto

package mountain

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MountainReview struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Mountain             *Mountain            `protobuf:"bytes,4,opt,name=mountain,proto3" json:"mountain,omitempty"`
	Difficulty           float64              `protobuf:"fixed64,5,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Comment              string               `protobuf:"bytes,6,opt,name=comment,proto3" json:"comment,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MountainReview) Reset()         { *m = MountainReview{} }
func (m *MountainReview) String() string { return proto.CompactTextString(m) }
func (*MountainReview) ProtoMessage()    {}
func (*MountainReview) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a245687c67fd166, []int{0}
}

func (m *MountainReview) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MountainReview.Unmarshal(m, b)
}
func (m *MountainReview) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MountainReview.Marshal(b, m, deterministic)
}
func (m *MountainReview) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MountainReview.Merge(m, src)
}
func (m *MountainReview) XXX_Size() int {
	return xxx_messageInfo_MountainReview.Size(m)
}
func (m *MountainReview) XXX_DiscardUnknown() {
	xxx_messageInfo_MountainReview.DiscardUnknown(m)
}

var xxx_messageInfo_MountainReview proto.InternalMessageInfo

func (m *MountainReview) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MountainReview) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *MountainReview) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *MountainReview) GetMountain() *Mountain {
	if m != nil {
		return m.Mountain
	}
	return nil
}

func (m *MountainReview) GetDifficulty() float64 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *MountainReview) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *MountainReview) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*MountainReview)(nil), "aggregates.mountain.MountainReview")
}

func init() {
	proto.RegisterFile("aggregates/mountain/mountain_review.proto", fileDescriptor_7a245687c67fd166)
}

var fileDescriptor_7a245687c67fd166 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xbf, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0x95, 0xb4, 0xdf, 0xfe, 0xf0, 0x57, 0xea, 0x60, 0x16, 0xab, 0x08, 0x88, 0x3a, 0x85,
	0xa1, 0xb6, 0x04, 0x53, 0x47, 0x18, 0x10, 0x0b, 0x4b, 0x04, 0x0b, 0x4b, 0xe5, 0xc6, 0xd7, 0xf4,
	0xd4, 0x3a, 0xae, 0x9c, 0x73, 0x11, 0x23, 0xff, 0x39, 0xaa, 0xdb, 0x14, 0x86, 0x8a, 0xed, 0xee,
	0xde, 0x7b, 0xfa, 0xf8, 0x99, 0xdd, 0xea, 0xaa, 0xf2, 0x50, 0x69, 0x82, 0x46, 0x59, 0x17, 0x6a,
	0xd2, 0x58, 0x9f, 0x86, 0xb9, 0x87, 0x1d, 0xc2, 0x87, 0xdc, 0x7a, 0x47, 0x8e, 0x5f, 0xfc, 0x58,
	0x65, 0xeb, 0x18, 0x4f, 0xfe, 0xca, 0x1f, 0x82, 0xe3, 0x9b, 0xca, 0xb9, 0x6a, 0x03, 0x2a, 0x6e,
	0x8b, 0xb0, 0x54, 0x84, 0x16, 0x1a, 0xd2, 0x76, 0x7b, 0x30, 0x4c, 0xbe, 0x52, 0x36, 0x7a, 0x39,
	0x66, 0x8a, 0x88, 0xe4, 0x23, 0x96, 0xa2, 0x11, 0x49, 0x96, 0xe4, 0x9d, 0x22, 0x45, 0xc3, 0x39,
	0xeb, 0x86, 0x80, 0x46, 0xa4, 0x59, 0x92, 0x0f, 0x8b, 0x38, 0xf3, 0x4b, 0x36, 0x0c, 0x0d, 0xf8,
	0x79, 0x14, 0x3a, 0x51, 0x18, 0xec, 0x0f, 0x6f, 0x7b, 0x71, 0xc6, 0x06, 0xed, 0x33, 0x44, 0x37,
	0x4b, 0xf2, 0xff, 0x77, 0x57, 0xf2, 0x4c, 0x01, 0x79, 0xe2, 0x9e, 0xec, 0xfc, 0x9a, 0x31, 0x83,
	0xcb, 0x25, 0x96, 0x61, 0x43, 0x9f, 0xe2, 0x5f, 0x96, 0xe4, 0x49, 0xf1, 0xeb, 0xc2, 0x05, 0xeb,
	0x97, 0xce, 0x5a, 0xa8, 0x49, 0xf4, 0x22, 0xb5, 0x5d, 0xf9, 0x8c, 0xb1, 0xd2, 0x83, 0x26, 0x30,
	0x73, 0x4d, 0xa2, 0x1f, 0xb1, 0x63, 0x79, 0xa8, 0x2f, 0xdb, 0xfa, 0xf2, 0xb5, 0xad, 0x5f, 0x0c,
	0x8f, 0xee, 0x07, 0x7a, 0x7c, 0x7e, 0x7f, 0xaa, 0x90, 0x56, 0x61, 0x21, 0x4b, 0x67, 0x15, 0x6e,
	0x56, 0xda, 0xda, 0x95, 0x31, 0x6a, 0x1d, 0x8c, 0x5e, 0xe3, 0x14, 0x76, 0x50, 0xd3, 0xb4, 0x01,
	0xbf, 0xc3, 0x12, 0x14, 0xd4, 0x84, 0x84, 0xd0, 0xa8, 0x33, 0xbf, 0xbf, 0xe8, 0x45, 0xd0, 0xfd,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0xe1, 0xde, 0x93, 0xdb, 0x01, 0x00, 0x00,
}
