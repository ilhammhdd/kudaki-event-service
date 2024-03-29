// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/user.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	user "github.com/ilhammhdd/kudaki-event-service/entities/aggregates/user"
	events "github.com/ilhammhdd/kudaki-event-service/usecases/events"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AuthenticateUser struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Jwt                  string   `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateUser) Reset()         { *m = AuthenticateUser{} }
func (m *AuthenticateUser) String() string { return proto.CompactTextString(m) }
func (*AuthenticateUser) ProtoMessage()    {}
func (*AuthenticateUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f3a799ad09b2336, []int{0}
}

func (m *AuthenticateUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateUser.Unmarshal(m, b)
}
func (m *AuthenticateUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateUser.Marshal(b, m, deterministic)
}
func (m *AuthenticateUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateUser.Merge(m, src)
}
func (m *AuthenticateUser) XXX_Size() int {
	return xxx_messageInfo_AuthenticateUser.Size(m)
}
func (m *AuthenticateUser) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateUser.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateUser proto.InternalMessageInfo

func (m *AuthenticateUser) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AuthenticateUser) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

type UserAuthenticated struct {
	Uid                  string         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	EventStatus          *events.Status `protobuf:"bytes,2,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	User                 *user.User     `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserAuthenticated) Reset()         { *m = UserAuthenticated{} }
func (m *UserAuthenticated) String() string { return proto.CompactTextString(m) }
func (*UserAuthenticated) ProtoMessage()    {}
func (*UserAuthenticated) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f3a799ad09b2336, []int{1}
}

func (m *UserAuthenticated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAuthenticated.Unmarshal(m, b)
}
func (m *UserAuthenticated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAuthenticated.Marshal(b, m, deterministic)
}
func (m *UserAuthenticated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAuthenticated.Merge(m, src)
}
func (m *UserAuthenticated) XXX_Size() int {
	return xxx_messageInfo_UserAuthenticated.Size(m)
}
func (m *UserAuthenticated) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAuthenticated.DiscardUnknown(m)
}

var xxx_messageInfo_UserAuthenticated proto.InternalMessageInfo

func (m *UserAuthenticated) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserAuthenticated) GetEventStatus() *events.Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *UserAuthenticated) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

type AuthorizeUser struct {
	Uid                  string          `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	UserRole             []user.UserRole `protobuf:"varint,2,rep,packed,name=user_role,json=userRole,proto3,enum=aggregates.user.UserRole" json:"user_role,omitempty"`
	Jwt                  string          `protobuf:"bytes,3,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthorizeUser) Reset()         { *m = AuthorizeUser{} }
func (m *AuthorizeUser) String() string { return proto.CompactTextString(m) }
func (*AuthorizeUser) ProtoMessage()    {}
func (*AuthorizeUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f3a799ad09b2336, []int{2}
}

func (m *AuthorizeUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeUser.Unmarshal(m, b)
}
func (m *AuthorizeUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeUser.Marshal(b, m, deterministic)
}
func (m *AuthorizeUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeUser.Merge(m, src)
}
func (m *AuthorizeUser) XXX_Size() int {
	return xxx_messageInfo_AuthorizeUser.Size(m)
}
func (m *AuthorizeUser) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeUser.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeUser proto.InternalMessageInfo

func (m *AuthorizeUser) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AuthorizeUser) GetUserRole() []user.UserRole {
	if m != nil {
		return m.UserRole
	}
	return nil
}

func (m *AuthorizeUser) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

type UserAuthorized struct {
	Uid                  string         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	EventStatus          *events.Status `protobuf:"bytes,2,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	User                 *user.User     `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserAuthorized) Reset()         { *m = UserAuthorized{} }
func (m *UserAuthorized) String() string { return proto.CompactTextString(m) }
func (*UserAuthorized) ProtoMessage()    {}
func (*UserAuthorized) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f3a799ad09b2336, []int{3}
}

func (m *UserAuthorized) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAuthorized.Unmarshal(m, b)
}
func (m *UserAuthorized) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAuthorized.Marshal(b, m, deterministic)
}
func (m *UserAuthorized) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAuthorized.Merge(m, src)
}
func (m *UserAuthorized) XXX_Size() int {
	return xxx_messageInfo_UserAuthorized.Size(m)
}
func (m *UserAuthorized) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAuthorized.DiscardUnknown(m)
}

var xxx_messageInfo_UserAuthorized proto.InternalMessageInfo

func (m *UserAuthorized) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserAuthorized) GetEventStatus() *events.Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *UserAuthorized) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthenticateUser)(nil), "rpc.AuthenticateUser")
	proto.RegisterType((*UserAuthenticated)(nil), "rpc.UserAuthenticated")
	proto.RegisterType((*AuthorizeUser)(nil), "rpc.AuthorizeUser")
	proto.RegisterType((*UserAuthorized)(nil), "rpc.UserAuthorized")
}

func init() { proto.RegisterFile("grpc/user.proto", fileDescriptor_3f3a799ad09b2336) }

var fileDescriptor_3f3a799ad09b2336 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0x4d, 0x4f, 0xea, 0x40,
	0x14, 0x4d, 0x29, 0x79, 0x79, 0x0c, 0x0f, 0x1e, 0x96, 0x60, 0x6a, 0x57, 0x84, 0x15, 0x2e, 0x98,
	0x1a, 0x4c, 0x88, 0x31, 0x6e, 0xd4, 0x7f, 0x50, 0xe3, 0xc6, 0x0d, 0x19, 0xda, 0x9b, 0x76, 0xa0,
	0x74, 0x9a, 0xf9, 0x00, 0xe3, 0xc6, 0xc4, 0x9d, 0xff, 0xda, 0xcc, 0xad, 0x7c, 0x88, 0xb8, 0x75,
	0x37, 0x73, 0xe6, 0x9e, 0x73, 0xcf, 0x99, 0x7b, 0xc9, 0xff, 0x54, 0x96, 0x71, 0x68, 0x14, 0x48,
	0x5a, 0x4a, 0xa1, 0x85, 0xe7, 0xca, 0x32, 0x0e, 0xba, 0xb0, 0x82, 0x42, 0xab, 0x50, 0x69, 0xa6,
	0x8d, 0xaa, 0x5e, 0x82, 0x80, 0xa5, 0xa9, 0x84, 0x94, 0x69, 0x50, 0x48, 0xd8, 0x63, 0x0d, 0x26,
	0xa4, 0x73, 0x6b, 0x74, 0x06, 0x85, 0xe6, 0x31, 0xd3, 0xf0, 0xa8, 0x40, 0x7a, 0x1d, 0xe2, 0x1a,
	0x9e, 0xf8, 0x4e, 0xdf, 0x19, 0x36, 0x22, 0x7b, 0xb4, 0xc8, 0x7c, 0xad, 0xfd, 0x5a, 0x85, 0xcc,
	0xd7, 0x7a, 0xf0, 0xe6, 0x90, 0x13, 0x5b, 0xbc, 0x4f, 0x4e, 0x8e, 0x30, 0x2f, 0xc8, 0x3f, 0xb4,
	0x34, 0xad, 0x1c, 0xa1, 0x44, 0x73, 0xdc, 0xa2, 0x08, 0xd2, 0x07, 0x04, 0xa3, 0x26, 0xde, 0xaa,
	0x8b, 0x77, 0x4e, 0xea, 0xd6, 0x9f, 0xef, 0x62, 0x65, 0x8f, 0xee, 0xcc, 0x53, 0xf4, 0x6d, 0xbb,
	0x46, 0x58, 0x32, 0x58, 0x90, 0x96, 0xed, 0x2f, 0x24, 0x7f, 0xf9, 0xc9, 0xf9, 0x84, 0x34, 0x6c,
	0xe9, 0x54, 0x8a, 0x1c, 0xfc, 0x5a, 0xdf, 0x1d, 0xb6, 0xc7, 0x67, 0xc7, 0x25, 0x45, 0x0e, 0xd1,
	0x5f, 0xf3, 0x79, 0xda, 0x24, 0x76, 0x77, 0x89, 0x5f, 0x49, 0x7b, 0x13, 0x18, 0x1b, 0xfe, 0x76,
	0xda, 0xf1, 0xbb, 0x43, 0xea, 0x98, 0xf2, 0x9e, 0x78, 0x07, 0x5f, 0xcf, 0x45, 0xe1, 0xf5, 0xa8,
	0x2c, 0x63, 0x7a, 0x38, 0xcc, 0xe0, 0x14, 0xe1, 0xef, 0xa3, 0xba, 0xd9, 0xcd, 0xcf, 0xc6, 0xa9,
	0x34, 0xbc, 0xad, 0xc6, 0xf6, 0x4f, 0x83, 0xee, 0x17, 0x81, 0x2a, 0xfa, 0xdd, 0xf5, 0xd3, 0x55,
	0xca, 0x75, 0x66, 0x66, 0x34, 0x16, 0xcb, 0x90, 0xe7, 0x19, 0x5b, 0x2e, 0xb3, 0x24, 0x09, 0x17,
	0x26, 0x61, 0x0b, 0x3e, 0xc2, 0x84, 0x23, 0x05, 0x72, 0xc5, 0x63, 0x08, 0xe1, 0x59, 0x83, 0x2c,
	0x58, 0xae, 0x42, 0xbb, 0xb3, 0xb3, 0x3f, 0xb8, 0x79, 0x97, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x25, 0xbc, 0xfd, 0x51, 0xc2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	UserAuthentication(ctx context.Context, in *AuthenticateUser, opts ...grpc.CallOption) (*UserAuthenticated, error)
	UserAuthorization(ctx context.Context, in *AuthorizeUser, opts ...grpc.CallOption) (*UserAuthorized, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserAuthentication(ctx context.Context, in *AuthenticateUser, opts ...grpc.CallOption) (*UserAuthenticated, error) {
	out := new(UserAuthenticated)
	err := c.cc.Invoke(ctx, "/rpc.User/UserAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserAuthorization(ctx context.Context, in *AuthorizeUser, opts ...grpc.CallOption) (*UserAuthorized, error) {
	out := new(UserAuthorized)
	err := c.cc.Invoke(ctx, "/rpc.User/UserAuthorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	UserAuthentication(context.Context, *AuthenticateUser) (*UserAuthenticated, error)
	UserAuthorization(context.Context, *AuthorizeUser) (*UserAuthorized, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) UserAuthentication(ctx context.Context, req *AuthenticateUser) (*UserAuthenticated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAuthentication not implemented")
}
func (*UnimplementedUserServer) UserAuthorization(ctx context.Context, req *AuthorizeUser) (*UserAuthorized, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAuthorization not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_UserAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/UserAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserAuthentication(ctx, req.(*AuthenticateUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/UserAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserAuthorization(ctx, req.(*AuthorizeUser))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserAuthentication",
			Handler:    _User_UserAuthentication_Handler,
		},
		{
			MethodName: "UserAuthorization",
			Handler:    _User_UserAuthorization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/user.proto",
}
