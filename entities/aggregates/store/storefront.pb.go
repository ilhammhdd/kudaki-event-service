// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/store/storefront.proto

package store

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

type Storefront struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	TotalItem            int32                `protobuf:"varint,4,opt,name=total_item,json=totalItem,proto3" json:"total_item,omitempty"`
	Rating               float64              `protobuf:"fixed64,5,opt,name=rating,proto3" json:"rating,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Storefront) Reset()         { *m = Storefront{} }
func (m *Storefront) String() string { return proto.CompactTextString(m) }
func (*Storefront) ProtoMessage()    {}
func (*Storefront) Descriptor() ([]byte, []int) {
	return fileDescriptor_11e209133ba8f19a, []int{0}
}

func (m *Storefront) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storefront.Unmarshal(m, b)
}
func (m *Storefront) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storefront.Marshal(b, m, deterministic)
}
func (m *Storefront) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storefront.Merge(m, src)
}
func (m *Storefront) XXX_Size() int {
	return xxx_messageInfo_Storefront.Size(m)
}
func (m *Storefront) XXX_DiscardUnknown() {
	xxx_messageInfo_Storefront.DiscardUnknown(m)
}

var xxx_messageInfo_Storefront proto.InternalMessageInfo

func (m *Storefront) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Storefront) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Storefront) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *Storefront) GetTotalItem() int32 {
	if m != nil {
		return m.TotalItem
	}
	return 0
}

func (m *Storefront) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Storefront) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Storefront)(nil), "aggregates.store.Storefront")
}

func init() { proto.RegisterFile("aggregates/store/storefront.proto", fileDescriptor_11e209133ba8f19a) }

var fileDescriptor_11e209133ba8f19a = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0xc6, 0x49, 0xff, 0xf1, 0x6e, 0x5e, 0x10, 0xc9, 0x41, 0x96, 0x8a, 0xb8, 0x7a, 0xda, 0x4b,
	0x13, 0xd0, 0x93, 0x47, 0x15, 0x0f, 0x5e, 0x57, 0xbd, 0x78, 0x29, 0x69, 0x33, 0xcd, 0x0e, 0xdd,
	0x6c, 0x4a, 0x32, 0xe9, 0xd7, 0xf3, 0xab, 0x49, 0xb3, 0xad, 0x82, 0x97, 0x90, 0x79, 0x7e, 0xcf,
	0xcc, 0x3c, 0x0c, 0xbf, 0xd1, 0xd6, 0x06, 0xb0, 0x9a, 0x20, 0xaa, 0x48, 0x3e, 0xc0, 0xf0, 0x6e,
	0x82, 0xef, 0x49, 0xee, 0x82, 0x27, 0x2f, 0xce, 0x7f, 0x2d, 0x32, 0xc3, 0xf9, 0xb5, 0xf5, 0xde,
	0x76, 0xa0, 0x32, 0x5f, 0xa5, 0x8d, 0x22, 0x74, 0x10, 0x49, 0xbb, 0xdd, 0xd0, 0x72, 0xfb, 0xc5,
	0x38, 0x7f, 0xfb, 0x99, 0x23, 0xce, 0xf8, 0x08, 0x4d, 0xc9, 0x2a, 0x56, 0x8f, 0x9b, 0x11, 0x1a,
	0x21, 0xf8, 0x24, 0x25, 0x34, 0xe5, 0xa8, 0x62, 0x75, 0xd1, 0xe4, 0xbf, 0xb8, 0xe4, 0x45, 0x8a,
	0x10, 0x96, 0x19, 0x8c, 0x33, 0xf8, 0x77, 0x10, 0x3e, 0x0e, 0xf0, 0x8a, 0x73, 0xf2, 0xa4, 0xbb,
	0x25, 0x12, 0xb8, 0x72, 0x52, 0xb1, 0x7a, 0xda, 0x14, 0x59, 0x79, 0x25, 0x70, 0xe2, 0x82, 0xcf,
	0x82, 0x26, 0xec, 0x6d, 0x39, 0xad, 0x58, 0xcd, 0x9a, 0x63, 0x25, 0x1e, 0x38, 0x5f, 0x07, 0xd0,
	0x04, 0x66, 0xa9, 0xa9, 0x9c, 0x55, 0xac, 0xfe, 0x7f, 0x37, 0x97, 0x43, 0x78, 0x79, 0x0a, 0x2f,
	0xdf, 0x4f, 0xe1, 0x9b, 0xe2, 0xe8, 0x7e, 0xa4, 0xa7, 0x97, 0xcf, 0x67, 0x8b, 0xd4, 0xa6, 0x95,
	0x5c, 0x7b, 0xa7, 0xb0, 0x6b, 0xb5, 0x73, 0xad, 0x31, 0x6a, 0x9b, 0x8c, 0xde, 0xe2, 0x02, 0xf6,
	0xd0, 0xd3, 0x22, 0x42, 0xd8, 0xe3, 0x1a, 0x14, 0xf4, 0x84, 0x84, 0x10, 0xd5, 0xdf, 0x63, 0xae,
	0x66, 0x79, 0xcb, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x28, 0xaf, 0xff, 0x67, 0x01,
	0x00, 0x00,
}
