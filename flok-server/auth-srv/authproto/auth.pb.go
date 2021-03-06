// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package authproto

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
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	AccountType          string   `protobuf:"bytes,4,opt,name=accountType,proto3" json:"accountType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_80c70f8f05f55273, []int{0}
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

func (m *User) GetAccountType() string {
	if m != nil {
		return m.AccountType
	}
	return ""
}

type CreateRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_80c70f8f05f55273, []int{1}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_80c70f8f05f55273, []int{2}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (dst *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(dst, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

type UpdateRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_80c70f8f05f55273, []int{3}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(dst, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_80c70f8f05f55273, []int{4}
}
func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(dst, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_80c70f8f05f55273, []int{5}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(dst, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_80c70f8f05f55273, []int{6}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(dst, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type VerifyUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyUserRequest) Reset()         { *m = VerifyUserRequest{} }
func (m *VerifyUserRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyUserRequest) ProtoMessage()    {}
func (*VerifyUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_80c70f8f05f55273, []int{7}
}
func (m *VerifyUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyUserRequest.Unmarshal(m, b)
}
func (m *VerifyUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyUserRequest.Marshal(b, m, deterministic)
}
func (dst *VerifyUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyUserRequest.Merge(dst, src)
}
func (m *VerifyUserRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyUserRequest.Size(m)
}
func (m *VerifyUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyUserRequest proto.InternalMessageInfo

func (m *VerifyUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *VerifyUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type VerifyUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyUserResponse) Reset()         { *m = VerifyUserResponse{} }
func (m *VerifyUserResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyUserResponse) ProtoMessage()    {}
func (*VerifyUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_80c70f8f05f55273, []int{8}
}
func (m *VerifyUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyUserResponse.Unmarshal(m, b)
}
func (m *VerifyUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyUserResponse.Marshal(b, m, deterministic)
}
func (dst *VerifyUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyUserResponse.Merge(dst, src)
}
func (m *VerifyUserResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyUserResponse.Size(m)
}
func (m *VerifyUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyUserResponse proto.InternalMessageInfo

func (m *VerifyUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "authproto.User")
	proto.RegisterType((*CreateRequest)(nil), "authproto.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "authproto.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "authproto.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "authproto.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "authproto.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "authproto.DeleteResponse")
	proto.RegisterType((*VerifyUserRequest)(nil), "authproto.VerifyUserRequest")
	proto.RegisterType((*VerifyUserResponse)(nil), "authproto.VerifyUserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	VerifyUser(ctx context.Context, in *VerifyUserRequest, opts ...grpc.CallOption) (*VerifyUserResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/authproto.AuthService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/authproto.AuthService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/authproto.AuthService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyUser(ctx context.Context, in *VerifyUserRequest, opts ...grpc.CallOption) (*VerifyUserResponse, error) {
	out := new(VerifyUserResponse)
	err := c.cc.Invoke(ctx, "/authproto.AuthService/VerifyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	VerifyUser(context.Context, *VerifyUserRequest) (*VerifyUserResponse, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthService/VerifyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyUser(ctx, req.(*VerifyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authproto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AuthService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AuthService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AuthService_Update_Handler,
		},
		{
			MethodName: "VerifyUser",
			Handler:    _AuthService_VerifyUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_auth_80c70f8f05f55273) }

var fileDescriptor_auth_80c70f8f05f55273 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0x9b, 0x7c, 0xf9, 0x2a, 0x72, 0xa2, 0x16, 0x38, 0x62, 0x08, 0x11, 0x88, 0x28, 0x2c,
	0x9d, 0x32, 0x14, 0x16, 0x26, 0xc4, 0xdf, 0xc4, 0x16, 0x28, 0x43, 0xb7, 0x90, 0x1c, 0xd4, 0x48,
	0xa5, 0x09, 0xb6, 0x03, 0xea, 0xc5, 0x70, 0xaf, 0x28, 0x76, 0xfe, 0x5c, 0xa8, 0x84, 0x98, 0x6a,
	0xf7, 0xf5, 0xfb, 0xc8, 0xe7, 0x71, 0x00, 0xe2, 0x52, 0x2c, 0xc2, 0x82, 0xe5, 0x22, 0x47, 0xbb,
	0x5a, 0xcb, 0x65, 0x30, 0x07, 0x6b, 0xc6, 0x89, 0xe1, 0x01, 0xfc, 0xa7, 0xd7, 0x38, 0x5b, 0xba,
	0xa6, 0x6f, 0x4c, 0xec, 0x48, 0x6d, 0xd0, 0x83, 0x9d, 0x22, 0xe6, 0xfc, 0x23, 0x67, 0xa9, 0xfb,
	0x4f, 0x06, 0xed, 0x1e, 0x7d, 0x70, 0xe2, 0x24, 0xc9, 0xcb, 0x95, 0x78, 0x5c, 0x17, 0xe4, 0x5a,
	0x32, 0xee, 0xff, 0x15, 0x9c, 0xc3, 0xe8, 0x86, 0x51, 0x2c, 0x28, 0xa2, 0xb7, 0x92, 0xb8, 0xc0,
	0x53, 0xb0, 0x4a, 0x4e, 0xcc, 0x35, 0x7c, 0x63, 0xe2, 0x4c, 0x77, 0xc3, 0xf6, 0x1a, 0x61, 0x75,
	0x87, 0x48, 0x86, 0xc1, 0x1e, 0x8c, 0x9b, 0x16, 0x2f, 0xf2, 0x15, 0x97, 0x9c, 0x59, 0x91, 0xfe,
	0x81, 0xd3, 0xb4, 0x6a, 0xce, 0x09, 0x8c, 0x6e, 0x69, 0x49, 0x1d, 0x67, 0x0c, 0x66, 0x96, 0x4a,
	0x8a, 0x1d, 0x99, 0x59, 0x5a, 0x55, 0x9a, 0x03, 0x75, 0xe5, 0x0e, 0xf6, 0x9f, 0x88, 0x65, 0x2f,
	0x6b, 0x09, 0xae, 0x6b, 0xad, 0x2b, 0x63, 0x9b, 0x2b, 0x53, 0x77, 0x15, 0x5c, 0x00, 0xf6, 0x31,
	0x0a, 0xfe, 0xab, 0x31, 0xa6, 0x9f, 0x26, 0x38, 0x57, 0xa5, 0x58, 0x3c, 0x10, 0x7b, 0xcf, 0x12,
	0xc2, 0x4b, 0x18, 0x2a, 0x3d, 0xe8, 0xf6, 0x0a, 0x9a, 0x67, 0xef, 0xf0, 0x87, 0xa4, 0x1e, 0x68,
	0x50, 0x01, 0xd4, 0x90, 0x1a, 0x40, 0x13, 0xa3, 0x01, 0x36, 0x8c, 0x48, 0x80, 0x12, 0xab, 0x01,
	0xb4, 0x17, 0xd2, 0x00, 0x1b, 0xaf, 0x30, 0xc0, 0x7b, 0x80, 0xce, 0x06, 0x1e, 0xf5, 0x8e, 0x7e,
	0x73, 0xed, 0x1d, 0x6f, 0x49, 0x1b, 0xd8, 0xb5, 0x33, 0xef, 0xbe, 0xe6, 0xe7, 0xa1, 0xfc, 0x39,
	0xfb, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x62, 0x31, 0xa8, 0xed, 0x02, 0x00, 0x00,
}
