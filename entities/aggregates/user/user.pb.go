// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/user/user.proto

package user

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

type UserRole int32

const (
	UserRole_ADMIN       UserRole = 0
	UserRole_USER        UserRole = 1
	UserRole_KUDAKI_TEAM UserRole = 2
	UserRole_ORGANIZER   UserRole = 3
)

var UserRole_name = map[int32]string{
	0: "ADMIN",
	1: "USER",
	2: "KUDAKI_TEAM",
	3: "ORGANIZER",
}

var UserRole_value = map[string]int32{
	"ADMIN":       0,
	"USER":        1,
	"KUDAKI_TEAM": 2,
	"ORGANIZER":   3,
}

func (x UserRole) String() string {
	return proto.EnumName(UserRole_name, int32(x))
}

func (UserRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7100ea213fb19611, []int{0}
}

type AccountType int32

const (
	AccountType_NATIVE   AccountType = 0
	AccountType_GOOGLE   AccountType = 1
	AccountType_FACEBOOK AccountType = 2
)

var AccountType_name = map[int32]string{
	0: "NATIVE",
	1: "GOOGLE",
	2: "FACEBOOK",
}

var AccountType_value = map[string]int32{
	"NATIVE":   0,
	"GOOGLE":   1,
	"FACEBOOK": 2,
}

func (x AccountType) String() string {
	return proto.EnumName(AccountType_name, int32(x))
}

func (AccountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7100ea213fb19611, []int{1}
}

type User struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string               `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Token                string               `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	Role                 UserRole             `protobuf:"varint,6,opt,name=role,proto3,enum=aggregates.user.UserRole" json:"role,omitempty"`
	PhoneNumber          string               `protobuf:"bytes,7,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	AccountType          AccountType          `protobuf:"varint,8,opt,name=account_type,json=accountType,proto3,enum=aggregates.user.AccountType" json:"account_type,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_7100ea213fb19611, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *User) GetRole() UserRole {
	if m != nil {
		return m.Role
	}
	return UserRole_ADMIN
}

func (m *User) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *User) GetAccountType() AccountType {
	if m != nil {
		return m.AccountType
	}
	return AccountType_NATIVE
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("aggregates.user.UserRole", UserRole_name, UserRole_value)
	proto.RegisterEnum("aggregates.user.AccountType", AccountType_name, AccountType_value)
	proto.RegisterType((*User)(nil), "aggregates.user.User")
}

func init() { proto.RegisterFile("aggregates/user/user.proto", fileDescriptor_7100ea213fb19611) }

var fileDescriptor_7100ea213fb19611 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xc1, 0x6e, 0x9b, 0x40,
	0x14, 0x0c, 0x18, 0xbb, 0xf0, 0x70, 0x13, 0xb4, 0xea, 0x81, 0x5a, 0x95, 0xea, 0xf6, 0x64, 0x45,
	0x32, 0x48, 0xc9, 0xa9, 0xa7, 0x08, 0xd7, 0xd4, 0xb2, 0xdc, 0x18, 0x69, 0x6b, 0xf7, 0x90, 0x8b,
	0xb5, 0x86, 0x57, 0xbc, 0x32, 0xb0, 0x68, 0x59, 0x52, 0xe5, 0xef, 0xfa, 0x69, 0x95, 0x97, 0xba,
	0xa9, 0xd2, 0x0b, 0x7a, 0x33, 0xef, 0xcd, 0x68, 0x18, 0x2d, 0x8c, 0x58, 0x9e, 0x4b, 0xcc, 0x99,
	0xc2, 0x26, 0x6c, 0x1b, 0x94, 0xfa, 0x13, 0xd4, 0x52, 0x28, 0x41, 0xae, 0x9e, 0x77, 0xc1, 0x89,
	0x1e, 0xbd, 0xcf, 0x85, 0xc8, 0x0b, 0x0c, 0xf5, 0x7a, 0xdf, 0xfe, 0x08, 0x15, 0x2f, 0xb1, 0x51,
	0xac, 0xac, 0x3b, 0xc5, 0xc7, 0x5f, 0x26, 0x58, 0xdb, 0x06, 0x25, 0xb9, 0x04, 0x93, 0x67, 0xbe,
	0x31, 0x36, 0x26, 0x3d, 0x6a, 0xf2, 0x8c, 0x10, 0xb0, 0xda, 0x96, 0x67, 0xbe, 0x39, 0x36, 0x26,
	0x0e, 0xd5, 0x33, 0x79, 0x03, 0x7d, 0x2c, 0x19, 0x2f, 0xfc, 0x9e, 0x26, 0x3b, 0x40, 0x46, 0x60,
	0xd7, 0xac, 0x69, 0x7e, 0x0a, 0x99, 0xf9, 0x96, 0x5e, 0xfc, 0xc5, 0x27, 0x85, 0x12, 0x47, 0xac,
	0xfc, 0x7e, 0xa7, 0xd0, 0x80, 0x4c, 0xc1, 0x92, 0xa2, 0x40, 0x7f, 0x30, 0x36, 0x26, 0x97, 0x37,
	0x6f, 0x83, 0x17, 0xa9, 0x83, 0x53, 0x20, 0x2a, 0x0a, 0xa4, 0xfa, 0x8c, 0x7c, 0x80, 0x61, 0x7d,
	0x10, 0x15, 0xee, 0xaa, 0xb6, 0xdc, 0xa3, 0xf4, 0x5f, 0x69, 0x2f, 0x57, 0x73, 0x6b, 0x4d, 0x91,
	0x3b, 0x18, 0xb2, 0x34, 0x15, 0x6d, 0xa5, 0x76, 0xea, 0xa9, 0x46, 0xdf, 0xd6, 0xce, 0xef, 0xfe,
	0x73, 0x8e, 0xba, 0xa3, 0xcd, 0x53, 0x8d, 0xd4, 0x65, 0xcf, 0x80, 0x7c, 0x02, 0x48, 0x25, 0x32,
	0x85, 0xd9, 0x8e, 0x29, 0xdf, 0x19, 0x1b, 0x13, 0xf7, 0x66, 0x14, 0x74, 0xed, 0x05, 0xe7, 0xf6,
	0x82, 0xcd, 0xb9, 0x3d, 0xea, 0xfc, 0xb9, 0x8e, 0xd4, 0xf5, 0x1d, 0xd8, 0xe7, 0xc0, 0xc4, 0x81,
	0x7e, 0x34, 0xbf, 0x5f, 0xae, 0xbd, 0x0b, 0x62, 0x83, 0xb5, 0xfd, 0x16, 0x53, 0xcf, 0x20, 0x57,
	0xe0, 0xae, 0xb6, 0xf3, 0x68, 0xb5, 0xdc, 0x6d, 0xe2, 0xe8, 0xde, 0x33, 0xc9, 0x6b, 0x70, 0x12,
	0xba, 0x88, 0xd6, 0xcb, 0x87, 0x98, 0x7a, 0xbd, 0xeb, 0x5b, 0x70, 0xff, 0xc9, 0x45, 0x00, 0x06,
	0xeb, 0x68, 0xb3, 0xfc, 0x1e, 0x7b, 0x17, 0xa7, 0x79, 0x91, 0x24, 0x8b, 0xaf, 0xb1, 0x67, 0x90,
	0x21, 0xd8, 0x5f, 0xa2, 0xcf, 0xf1, 0x2c, 0x49, 0x56, 0x9e, 0x39, 0x9b, 0x3f, 0xcc, 0x72, 0xae,
	0x0e, 0xed, 0x3e, 0x48, 0x45, 0x19, 0xf2, 0xe2, 0xc0, 0xca, 0xf2, 0x90, 0x65, 0xe1, 0xb1, 0xcd,
	0xd8, 0x91, 0x4f, 0xf1, 0x11, 0x2b, 0x35, 0x6d, 0x50, 0x3e, 0xf2, 0x14, 0x43, 0xac, 0x14, 0x57,
	0x1c, 0x9b, 0xf0, 0xc5, 0xdb, 0xd9, 0x0f, 0xf4, 0xaf, 0xdd, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0xc6, 0x30, 0xc0, 0x35, 0x55, 0x02, 0x00, 0x00,
}
