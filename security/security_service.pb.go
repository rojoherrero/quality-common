// Code generated by protoc-gen-go. DO NOT EDIT.
// source: security_service.proto

package security

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

type LogInRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInRequest) Reset()         { *m = LogInRequest{} }
func (m *LogInRequest) String() string { return proto.CompactTextString(m) }
func (*LogInRequest) ProtoMessage()    {}
func (*LogInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_security_service_25032f165235892b, []int{0}
}
func (m *LogInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInRequest.Unmarshal(m, b)
}
func (m *LogInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInRequest.Marshal(b, m, deterministic)
}
func (dst *LogInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInRequest.Merge(dst, src)
}
func (m *LogInRequest) XXX_Size() int {
	return xxx_messageInfo_LogInRequest.Size(m)
}
func (m *LogInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogInRequest proto.InternalMessageInfo

func (m *LogInRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LogInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LogInResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Fullname             string   `protobuf:"bytes,3,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Departments          []string `protobuf:"bytes,4,rep,name=departments,proto3" json:"departments,omitempty"`
	Roles                []string `protobuf:"bytes,5,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInResponse) Reset()         { *m = LogInResponse{} }
func (m *LogInResponse) String() string { return proto.CompactTextString(m) }
func (*LogInResponse) ProtoMessage()    {}
func (*LogInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_security_service_25032f165235892b, []int{1}
}
func (m *LogInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInResponse.Unmarshal(m, b)
}
func (m *LogInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInResponse.Marshal(b, m, deterministic)
}
func (dst *LogInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInResponse.Merge(dst, src)
}
func (m *LogInResponse) XXX_Size() int {
	return xxx_messageInfo_LogInResponse.Size(m)
}
func (m *LogInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogInResponse proto.InternalMessageInfo

func (m *LogInResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *LogInResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LogInResponse) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *LogInResponse) GetDepartments() []string {
	if m != nil {
		return m.Departments
	}
	return nil
}

func (m *LogInResponse) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*LogInRequest)(nil), "security.LogInRequest")
	proto.RegisterType((*LogInResponse)(nil), "security.LogInResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecurityClient is the client API for Security service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecurityClient interface {
	GetUserInfo(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error)
}

type securityClient struct {
	cc *grpc.ClientConn
}

func NewSecurityClient(cc *grpc.ClientConn) SecurityClient {
	return &securityClient{cc}
}

func (c *securityClient) GetUserInfo(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error) {
	out := new(LogInResponse)
	err := c.cc.Invoke(ctx, "/security.Security/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityServer is the server API for Security service.
type SecurityServer interface {
	GetUserInfo(context.Context, *LogInRequest) (*LogInResponse, error)
}

func RegisterSecurityServer(s *grpc.Server, srv SecurityServer) {
	s.RegisterService(&_Security_serviceDesc, srv)
}

func _Security_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.Security/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServer).GetUserInfo(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Security_serviceDesc = grpc.ServiceDesc{
	ServiceName: "security.Security",
	HandlerType: (*SecurityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _Security_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "security_service.proto",
}

func init() {
	proto.RegisterFile("security_service.proto", fileDescriptor_security_service_25032f165235892b)
}

var fileDescriptor_security_service_25032f165235892b = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0xc7, 0x49, 0xf3, 0xb4, 0xa4, 0xd3, 0x47, 0x84, 0x45, 0x62, 0xa8, 0x97, 0xd0, 0x53, 0x4f,
	0x8b, 0xe8, 0xdd, 0x43, 0x0f, 0x96, 0x8a, 0x87, 0x12, 0xf1, 0x2c, 0x6b, 0x32, 0x0d, 0x81, 0x64,
	0x77, 0x9d, 0xd9, 0xf5, 0xe5, 0x93, 0xf8, 0x75, 0x25, 0x6f, 0x52, 0xe9, 0x6d, 0x7e, 0xff, 0x1f,
	0xcc, 0x30, 0x7f, 0x88, 0x19, 0x73, 0x4f, 0x95, 0xfb, 0x7a, 0x61, 0xa4, 0xf7, 0x2a, 0x47, 0x69,
	0xc9, 0x38, 0x23, 0xa2, 0x31, 0x5f, 0xdd, 0xc3, 0xff, 0x47, 0x53, 0xee, 0x74, 0x86, 0x6f, 0x1e,
	0xd9, 0x89, 0x25, 0x44, 0x9e, 0x91, 0xb4, 0x6a, 0x30, 0x09, 0xd2, 0x60, 0x3d, 0xcf, 0x7e, 0xb9,
	0x75, 0x56, 0x31, 0x7f, 0x18, 0x2a, 0x92, 0x49, 0xef, 0x46, 0x5e, 0x7d, 0x07, 0x70, 0x36, 0x2c,
	0x62, 0x6b, 0x34, 0xa3, 0x88, 0x61, 0xc6, 0x4e, 0x39, 0xcf, 0xdd, 0x9e, 0x30, 0x1b, 0xe8, 0xcf,
	0x85, 0xc9, 0xe9, 0x85, 0x83, 0xaf, 0xeb, 0xce, 0x85, 0xbd, 0x1b, 0x59, 0xa4, 0xb0, 0x28, 0xd0,
	0x2a, 0x72, 0x0d, 0x6a, 0xc7, 0xc9, 0xbf, 0x34, 0x5c, 0xcf, 0xb3, 0xe3, 0x48, 0x5c, 0xc0, 0x94,
	0x4c, 0x8d, 0x9c, 0x4c, 0x3b, 0xd7, 0xc3, 0xcd, 0x03, 0x44, 0x4f, 0xc3, 0xb7, 0xe2, 0x0e, 0x16,
	0x5b, 0x74, 0xcf, 0x8c, 0xb4, 0xd3, 0x07, 0x23, 0x62, 0x39, 0xf6, 0x20, 0x8f, 0x4b, 0x58, 0x5e,
	0x9e, 0xe4, 0xfd, 0x4f, 0x9b, 0x6b, 0xb8, 0xaa, 0x8c, 0x2c, 0xc9, 0xe6, 0x12, 0x3f, 0x55, 0x63,
	0x6b, 0x64, 0x49, 0xc6, 0x3b, 0x2c, 0x7d, 0x55, 0xe0, 0xe6, 0x3c, 0x6b, 0xe7, 0x6d, 0x3b, 0xef,
	0xdb, 0x9e, 0xf7, 0xc1, 0xeb, 0xac, 0x2b, 0xfc, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x24, 0x7c,
	0x7a, 0xf2, 0x8a, 0x01, 0x00, 0x00,
}
