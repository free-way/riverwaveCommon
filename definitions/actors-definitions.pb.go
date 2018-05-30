// Code generated by protoc-gen-go. DO NOT EDIT.
// source: actors-definitions.proto

package definitions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=LastName" json:"LastName,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=Password" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_actors_definitions_b98b19bd138f898a, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
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

type GeneralResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=Status" json:"Status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeneralResponse) Reset()         { *m = GeneralResponse{} }
func (m *GeneralResponse) String() string { return proto.CompactTextString(m) }
func (*GeneralResponse) ProtoMessage()    {}
func (*GeneralResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_actors_definitions_b98b19bd138f898a, []int{1}
}
func (m *GeneralResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralResponse.Unmarshal(m, b)
}
func (m *GeneralResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralResponse.Marshal(b, m, deterministic)
}
func (dst *GeneralResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralResponse.Merge(dst, src)
}
func (m *GeneralResponse) XXX_Size() int {
	return xxx_messageInfo_GeneralResponse.Size(m)
}
func (m *GeneralResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralResponse proto.InternalMessageInfo

func (m *GeneralResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GeneralResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetAllUsersResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=Status" json:"Status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	Users                []*User  `protobuf:"bytes,3,rep,name=users" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllUsersResponse) Reset()         { *m = GetAllUsersResponse{} }
func (m *GetAllUsersResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllUsersResponse) ProtoMessage()    {}
func (*GetAllUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_actors_definitions_b98b19bd138f898a, []int{2}
}
func (m *GetAllUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllUsersResponse.Unmarshal(m, b)
}
func (m *GetAllUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllUsersResponse.Marshal(b, m, deterministic)
}
func (dst *GetAllUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllUsersResponse.Merge(dst, src)
}
func (m *GetAllUsersResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllUsersResponse.Size(m)
}
func (m *GetAllUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllUsersResponse proto.InternalMessageInfo

func (m *GetAllUsersResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetAllUsersResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetAllUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_actors_definitions_b98b19bd138f898a, []int{3}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type CreateUserRequest struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=LastName" json:"LastName,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=Email" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=Password" json:"Password,omitempty"`
	RoleID               int32    `protobuf:"varint,5,opt,name=RoleID" json:"RoleID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_actors_definitions_b98b19bd138f898a, []int{4}
}
func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (dst *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(dst, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *CreateUserRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserRequest) GetRoleID() int32 {
	if m != nil {
		return m.RoleID
	}
	return 0
}

type EditUserRequest struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=LastName" json:"LastName,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=Password" json:"Password,omitempty"`
	RoleID               int32    `protobuf:"varint,6,opt,name=RoleID" json:"RoleID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditUserRequest) Reset()         { *m = EditUserRequest{} }
func (m *EditUserRequest) String() string { return proto.CompactTextString(m) }
func (*EditUserRequest) ProtoMessage()    {}
func (*EditUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_actors_definitions_b98b19bd138f898a, []int{5}
}
func (m *EditUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditUserRequest.Unmarshal(m, b)
}
func (m *EditUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditUserRequest.Marshal(b, m, deterministic)
}
func (dst *EditUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditUserRequest.Merge(dst, src)
}
func (m *EditUserRequest) XXX_Size() int {
	return xxx_messageInfo_EditUserRequest.Size(m)
}
func (m *EditUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EditUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EditUserRequest proto.InternalMessageInfo

func (m *EditUserRequest) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *EditUserRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *EditUserRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *EditUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *EditUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *EditUserRequest) GetRoleID() int32 {
	if m != nil {
		return m.RoleID
	}
	return 0
}

type DeleteUserRequest struct {
	UserID               int32    `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_actors_definitions_b98b19bd138f898a, []int{6}
}
func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(dst, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type AuthenticationParams struct {
	Email                string   `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationParams) Reset()         { *m = AuthenticationParams{} }
func (m *AuthenticationParams) String() string { return proto.CompactTextString(m) }
func (*AuthenticationParams) ProtoMessage()    {}
func (*AuthenticationParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_actors_definitions_b98b19bd138f898a, []int{7}
}
func (m *AuthenticationParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationParams.Unmarshal(m, b)
}
func (m *AuthenticationParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationParams.Marshal(b, m, deterministic)
}
func (dst *AuthenticationParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationParams.Merge(dst, src)
}
func (m *AuthenticationParams) XXX_Size() int {
	return xxx_messageInfo_AuthenticationParams.Size(m)
}
func (m *AuthenticationParams) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationParams.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationParams proto.InternalMessageInfo

func (m *AuthenticationParams) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthenticationParams) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "definitions.User")
	proto.RegisterType((*GeneralResponse)(nil), "definitions.GeneralResponse")
	proto.RegisterType((*GetAllUsersResponse)(nil), "definitions.GetAllUsersResponse")
	proto.RegisterType((*Empty)(nil), "definitions.Empty")
	proto.RegisterType((*CreateUserRequest)(nil), "definitions.CreateUserRequest")
	proto.RegisterType((*EditUserRequest)(nil), "definitions.EditUserRequest")
	proto.RegisterType((*DeleteUserRequest)(nil), "definitions.DeleteUserRequest")
	proto.RegisterType((*AuthenticationParams)(nil), "definitions.AuthenticationParams")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ActorsService service

type ActorsServiceClient interface {
	GetAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllUsersResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
	EditUser(ctx context.Context, in *EditUserRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
	Authenticate(ctx context.Context, in *AuthenticationParams, opts ...grpc.CallOption) (*GeneralResponse, error)
}

type actorsServiceClient struct {
	cc *grpc.ClientConn
}

func NewActorsServiceClient(cc *grpc.ClientConn) ActorsServiceClient {
	return &actorsServiceClient{cc}
}

func (c *actorsServiceClient) GetAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllUsersResponse, error) {
	out := new(GetAllUsersResponse)
	err := grpc.Invoke(ctx, "/definitions.ActorsService/GetAllUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actorsServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := grpc.Invoke(ctx, "/definitions.ActorsService/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actorsServiceClient) EditUser(ctx context.Context, in *EditUserRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := grpc.Invoke(ctx, "/definitions.ActorsService/EditUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actorsServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := grpc.Invoke(ctx, "/definitions.ActorsService/DeleteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actorsServiceClient) Authenticate(ctx context.Context, in *AuthenticationParams, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := grpc.Invoke(ctx, "/definitions.ActorsService/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ActorsService service

type ActorsServiceServer interface {
	GetAllUsers(context.Context, *Empty) (*GetAllUsersResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*GeneralResponse, error)
	EditUser(context.Context, *EditUserRequest) (*GeneralResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*GeneralResponse, error)
	Authenticate(context.Context, *AuthenticationParams) (*GeneralResponse, error)
}

func RegisterActorsServiceServer(s *grpc.Server, srv ActorsServiceServer) {
	s.RegisterService(&_ActorsService_serviceDesc, srv)
}

func _ActorsService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActorsServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.ActorsService/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActorsServiceServer).GetAllUsers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActorsService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActorsServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.ActorsService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActorsServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActorsService_EditUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActorsServiceServer).EditUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.ActorsService/EditUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActorsServiceServer).EditUser(ctx, req.(*EditUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActorsService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActorsServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.ActorsService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActorsServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActorsService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActorsServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.ActorsService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActorsServiceServer).Authenticate(ctx, req.(*AuthenticationParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActorsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "definitions.ActorsService",
	HandlerType: (*ActorsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllUsers",
			Handler:    _ActorsService_GetAllUsers_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _ActorsService_CreateUser_Handler,
		},
		{
			MethodName: "EditUser",
			Handler:    _ActorsService_EditUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _ActorsService_DeleteUser_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _ActorsService_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "actors-definitions.proto",
}

func init() {
	proto.RegisterFile("actors-definitions.proto", fileDescriptor_actors_definitions_b98b19bd138f898a)
}

var fileDescriptor_actors_definitions_b98b19bd138f898a = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0xe3, 0xd8, 0x6d, 0x27, 0x40, 0x95, 0xa1, 0xaa, 0xac, 0xa8, 0x42, 0x66, 0x2f, 0x44,
	0x42, 0xf4, 0x50, 0xbe, 0x20, 0xaa, 0x43, 0x1b, 0x29, 0xa0, 0xe2, 0x8a, 0x0f, 0x58, 0x92, 0x01,
	0x56, 0x72, 0x6c, 0xb3, 0xbb, 0x06, 0x71, 0xe5, 0x07, 0x38, 0xf2, 0x05, 0xfc, 0x27, 0xda, 0xb5,
	0xdd, 0x78, 0xd3, 0x86, 0x48, 0x5c, 0x7a, 0x7c, 0x33, 0x9e, 0xa7, 0x37, 0xef, 0xcd, 0x1a, 0x22,
	0xbe, 0xd0, 0x85, 0x54, 0xaf, 0x96, 0xf4, 0x49, 0xe4, 0x42, 0x8b, 0x22, 0x57, 0x67, 0xa5, 0x2c,
	0x74, 0x81, 0x83, 0x4e, 0x89, 0xfd, 0xf4, 0xa0, 0xff, 0x41, 0x91, 0xc4, 0x27, 0xd0, 0x9b, 0x25,
	0x91, 0x17, 0x7b, 0xe3, 0x20, 0xed, 0xcd, 0x12, 0x3c, 0x85, 0xc3, 0x37, 0x42, 0x2a, 0xfd, 0x8e,
	0xaf, 0x28, 0xea, 0xc5, 0xde, 0xf8, 0x30, 0x5d, 0x17, 0x70, 0x04, 0x07, 0x73, 0xde, 0x34, 0x7d,
	0xdb, 0xbc, 0xc5, 0x78, 0x0c, 0xc1, 0x74, 0xc5, 0x45, 0x16, 0xf5, 0x6d, 0xa3, 0x06, 0x66, 0xe2,
	0x9a, 0x2b, 0xf5, 0xbd, 0x90, 0xcb, 0x28, 0xa8, 0x27, 0x5a, 0xcc, 0x2e, 0xe0, 0xe8, 0x92, 0x72,
	0x92, 0x3c, 0x4b, 0x49, 0x95, 0x45, 0xae, 0x08, 0x4f, 0x20, 0xbc, 0xd1, 0x5c, 0x57, 0xaa, 0x91,
	0xd4, 0x20, 0x8c, 0x60, 0xff, 0x2d, 0x29, 0xc5, 0x3f, 0xb7, 0xa2, 0x5a, 0xc8, 0x4a, 0x78, 0x7a,
	0x49, 0x7a, 0x92, 0x65, 0x66, 0x1d, 0xf5, 0xff, 0x44, 0xf8, 0x02, 0x82, 0xca, 0x50, 0x44, 0x7e,
	0xec, 0x8f, 0x07, 0xe7, 0xc3, 0xb3, 0xae, 0x85, 0x86, 0x3c, 0xad, 0xfb, 0x6c, 0xdf, 0x2c, 0x5a,
	0xea, 0x1f, 0xec, 0xb7, 0x07, 0xc3, 0x0b, 0x49, 0x5c, 0x93, 0x6d, 0xd3, 0xd7, 0x8a, 0x94, 0x76,
	0x1d, 0xf4, 0xfe, 0xe5, 0x60, 0x6f, 0x9b, 0x83, 0xfe, 0x36, 0x07, 0xfb, 0xae, 0x83, 0x66, 0xcb,
	0xb4, 0xc8, 0x68, 0x96, 0x58, 0x6f, 0x83, 0xb4, 0x41, 0xec, 0x8f, 0x07, 0x47, 0xd3, 0xa5, 0xd0,
	0x5d, 0x5d, 0x0f, 0x98, 0x74, 0x47, 0x67, 0xe8, 0xe8, 0x7c, 0x09, 0xc3, 0x84, 0x32, 0x72, 0x0d,
	0x3c, 0x81, 0xd0, 0xc0, 0x5b, 0xb1, 0x0d, 0x62, 0x57, 0x70, 0x3c, 0xa9, 0xf4, 0x17, 0xca, 0xb5,
	0x58, 0x70, 0x13, 0xcb, 0x35, 0x97, 0x7c, 0xa5, 0x8c, 0x1c, 0xb2, 0x72, 0x6a, 0xb3, 0x6b, 0x60,
	0xe4, 0x94, 0xad, 0x9c, 0xc6, 0xe8, 0x16, 0x9f, 0xff, 0xf2, 0xe1, 0xf1, 0xc4, 0xbe, 0x93, 0x1b,
	0x92, 0xdf, 0xc4, 0x82, 0x70, 0x0a, 0x83, 0xce, 0x15, 0x21, 0x3a, 0xe1, 0xdb, 0xb4, 0x47, 0xb1,
	0x53, 0xbb, 0xe7, 0xe6, 0xd8, 0x1e, 0xce, 0x01, 0xd6, 0x07, 0x81, 0xcf, 0x9c, 0x89, 0x3b, 0x97,
	0x32, 0x3a, 0xdd, 0x60, 0x74, 0x9e, 0x02, 0xdb, 0xc3, 0x2b, 0x38, 0x68, 0x43, 0x44, 0xf7, 0xdb,
	0x8d, 0x6c, 0x77, 0x32, 0xcd, 0x01, 0xd6, 0x3e, 0x6f, 0xe8, 0xba, 0x13, 0xc0, 0x4e, 0xb6, 0xf7,
	0xf0, 0xa8, 0x13, 0x04, 0xe1, 0x73, 0xe7, 0xfb, 0xfb, 0x32, 0xda, 0x45, 0xf9, 0x31, 0xb4, 0xff,
	0xa8, 0xd7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xc2, 0xf3, 0xd8, 0xbf, 0x04, 0x00, 0x00,
}
