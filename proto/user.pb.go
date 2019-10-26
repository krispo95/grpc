// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Status int32

const (
	Status_Unknown Status = 0
	Status_Active  Status = 1
	Status_Blocked Status = 2
)

var Status_name = map[int32]string{
	0: "Unknown",
	1: "Active",
	2: "Blocked",
}

var Status_value = map[string]int32{
	"Unknown": 0,
	"Active":  1,
	"Blocked": 2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

type UserMessage struct {
	Id                   *UserId    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  uint32     `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	ChildCount           *NatNumber `protobuf:"bytes,4,opt,name=childCount,proto3" json:"childCount,omitempty"`
	CatsCount            *NatNumber `protobuf:"bytes,5,opt,name=catsCount,proto3" json:"catsCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UserMessage) Reset()         { *m = UserMessage{} }
func (m *UserMessage) String() string { return proto.CompactTextString(m) }
func (*UserMessage) ProtoMessage()    {}
func (*UserMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserMessage.Unmarshal(m, b)
}
func (m *UserMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserMessage.Marshal(b, m, deterministic)
}
func (m *UserMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMessage.Merge(m, src)
}
func (m *UserMessage) XXX_Size() int {
	return xxx_messageInfo_UserMessage.Size(m)
}
func (m *UserMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMessage.DiscardUnknown(m)
}

var xxx_messageInfo_UserMessage proto.InternalMessageInfo

func (m *UserMessage) GetId() *UserId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *UserMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserMessage) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserMessage) GetChildCount() *NatNumber {
	if m != nil {
		return m.ChildCount
	}
	return nil
}

func (m *UserMessage) GetCatsCount() *NatNumber {
	if m != nil {
		return m.CatsCount
	}
	return nil
}

type UserId struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserId) Reset()         { *m = UserId{} }
func (m *UserId) String() string { return proto.CompactTextString(m) }
func (*UserId) ProtoMessage()    {}
func (*UserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserId.Unmarshal(m, b)
}
func (m *UserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserId.Marshal(b, m, deterministic)
}
func (m *UserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserId.Merge(m, src)
}
func (m *UserId) XXX_Size() int {
	return xxx_messageInfo_UserId.Size(m)
}
func (m *UserId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserId.DiscardUnknown(m)
}

var xxx_messageInfo_UserId proto.InternalMessageInfo

func (m *UserId) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type NatNumber struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatNumber) Reset()         { *m = NatNumber{} }
func (m *NatNumber) String() string { return proto.CompactTextString(m) }
func (*NatNumber) ProtoMessage()    {}
func (*NatNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *NatNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatNumber.Unmarshal(m, b)
}
func (m *NatNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatNumber.Marshal(b, m, deterministic)
}
func (m *NatNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatNumber.Merge(m, src)
}
func (m *NatNumber) XXX_Size() int {
	return xxx_messageInfo_NatNumber.Size(m)
}
func (m *NatNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_NatNumber.DiscardUnknown(m)
}

var xxx_messageInfo_NatNumber proto.InternalMessageInfo

func (m *NatNumber) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterEnum("user.Status", Status_name, Status_value)
	proto.RegisterType((*UserMessage)(nil), "user.UserMessage")
	proto.RegisterType((*UserId)(nil), "user.UserId")
	proto.RegisterType((*NatNumber)(nil), "user.NatNumber")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0x4d, 0xdb, 0x55, 0x7a, 0xab, 0x58, 0x2f, 0x3e, 0x94, 0x21, 0x52, 0xfb, 0x20, 0x45,
	0xb0, 0xc2, 0xfc, 0x05, 0xea, 0x83, 0xf8, 0xe0, 0x1e, 0x3a, 0xe6, 0x7b, 0xd6, 0x5e, 0x66, 0x59,
	0x97, 0x48, 0x93, 0xcc, 0xdf, 0xe5, 0x3f, 0x94, 0xa4, 0xe2, 0x1c, 0xdb, 0xdb, 0x49, 0xee, 0x77,
	0x4f, 0x0e, 0x27, 0x00, 0x46, 0x51, 0x5f, 0x7e, 0xf6, 0x52, 0x4b, 0x0c, 0xac, 0xce, 0xbf, 0x19,
	0xc4, 0x73, 0x45, 0xfd, 0x1b, 0x29, 0xc5, 0x97, 0x84, 0x97, 0xe0, 0xb5, 0x4d, 0xca, 0x32, 0x56,
	0xc4, 0x93, 0x93, 0xd2, 0xe1, 0x76, 0xfc, 0xda, 0x54, 0x5e, 0xdb, 0x20, 0x42, 0x20, 0xf8, 0x9a,
	0x52, 0x2f, 0x63, 0x45, 0x54, 0x39, 0x8d, 0x09, 0xf8, 0x7c, 0x49, 0xa9, 0x9f, 0xb1, 0xe2, 0xb4,
	0xb2, 0x12, 0xef, 0x01, 0xea, 0x8f, 0xb6, 0x6b, 0x9e, 0xa5, 0x11, 0x3a, 0x0d, 0x9c, 0xd7, 0xd9,
	0xe0, 0x35, 0xe5, 0x7a, 0x6a, 0xd6, 0x0b, 0xea, 0xab, 0x7f, 0x08, 0xde, 0x41, 0x54, 0x73, 0xad,
	0x06, 0x7e, 0x74, 0x98, 0xdf, 0x12, 0xf9, 0x15, 0x84, 0x43, 0x26, 0xbc, 0x80, 0xd1, 0x86, 0x77,
	0x86, 0x5c, 0xe0, 0xa0, 0x1a, 0x0e, 0xf9, 0x35, 0x44, 0x7f, 0x7b, 0x87, 0x91, 0xdb, 0x12, 0xc2,
	0x99, 0xe6, 0xda, 0x28, 0x8c, 0xe1, 0x78, 0x2e, 0x56, 0x42, 0x7e, 0x89, 0xe4, 0x08, 0x01, 0xc2,
	0xc7, 0x5a, 0xb7, 0x1b, 0x4a, 0x98, 0x1d, 0x3c, 0x75, 0xb2, 0x5e, 0x51, 0x93, 0x78, 0x93, 0x77,
	0x08, 0xec, 0x93, 0x78, 0x03, 0xfe, 0x8c, 0x34, 0x9e, 0x6f, 0x9b, 0xf9, 0x2d, 0x6e, 0xbc, 0x53,
	0x96, 0xe5, 0x5e, 0x48, 0xe3, 0xce, 0xe5, 0x78, 0x7f, 0x6b, 0x11, 0xba, 0xbf, 0x78, 0xf8, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x88, 0x95, 0x33, 0xc1, 0x99, 0x01, 0x00, 0x00,
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
	Set(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*UserId, error)
	Get(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserMessage, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Set(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*UserId, error) {
	out := new(UserId)
	err := c.cc.Invoke(ctx, "/user.User/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Get(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserMessage, error) {
	out := new(UserMessage)
	err := c.cc.Invoke(ctx, "/user.User/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Set(context.Context, *UserMessage) (*UserId, error)
	Get(context.Context, *UserId) (*UserMessage, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) Set(ctx context.Context, req *UserMessage) (*UserId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedUserServer) Get(ctx context.Context, req *UserId) (*UserMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Set(ctx, req.(*UserMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Get(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _User_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _User_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}