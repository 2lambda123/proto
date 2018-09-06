// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iam/grpcapi.proto

package grpcapi

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

// Input & output messages
type GetHierarchyRelationsInput struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHierarchyRelationsInput) Reset()         { *m = GetHierarchyRelationsInput{} }
func (m *GetHierarchyRelationsInput) String() string { return proto.CompactTextString(m) }
func (*GetHierarchyRelationsInput) ProtoMessage()    {}
func (*GetHierarchyRelationsInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_fb9b75323b01a2e3, []int{0}
}
func (m *GetHierarchyRelationsInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHierarchyRelationsInput.Unmarshal(m, b)
}
func (m *GetHierarchyRelationsInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHierarchyRelationsInput.Marshal(b, m, deterministic)
}
func (dst *GetHierarchyRelationsInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHierarchyRelationsInput.Merge(dst, src)
}
func (m *GetHierarchyRelationsInput) XXX_Size() int {
	return xxx_messageInfo_GetHierarchyRelationsInput.Size(m)
}
func (m *GetHierarchyRelationsInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHierarchyRelationsInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetHierarchyRelationsInput proto.InternalMessageInfo

func (m *GetHierarchyRelationsInput) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetHierarchyRelationsOutput struct {
	NodeIds              []string `protobuf:"bytes,1,rep,name=node_ids,json=nodeIds,proto3" json:"node_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHierarchyRelationsOutput) Reset()         { *m = GetHierarchyRelationsOutput{} }
func (m *GetHierarchyRelationsOutput) String() string { return proto.CompactTextString(m) }
func (*GetHierarchyRelationsOutput) ProtoMessage()    {}
func (*GetHierarchyRelationsOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_fb9b75323b01a2e3, []int{1}
}
func (m *GetHierarchyRelationsOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHierarchyRelationsOutput.Unmarshal(m, b)
}
func (m *GetHierarchyRelationsOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHierarchyRelationsOutput.Marshal(b, m, deterministic)
}
func (dst *GetHierarchyRelationsOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHierarchyRelationsOutput.Merge(dst, src)
}
func (m *GetHierarchyRelationsOutput) XXX_Size() int {
	return xxx_messageInfo_GetHierarchyRelationsOutput.Size(m)
}
func (m *GetHierarchyRelationsOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHierarchyRelationsOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GetHierarchyRelationsOutput proto.InternalMessageInfo

func (m *GetHierarchyRelationsOutput) GetNodeIds() []string {
	if m != nil {
		return m.NodeIds
	}
	return nil
}

type CheckAuthenticationInput struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	MethodArn            string   `protobuf:"bytes,2,opt,name=methodArn,proto3" json:"methodArn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAuthenticationInput) Reset()         { *m = CheckAuthenticationInput{} }
func (m *CheckAuthenticationInput) String() string { return proto.CompactTextString(m) }
func (*CheckAuthenticationInput) ProtoMessage()    {}
func (*CheckAuthenticationInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_fb9b75323b01a2e3, []int{2}
}
func (m *CheckAuthenticationInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAuthenticationInput.Unmarshal(m, b)
}
func (m *CheckAuthenticationInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAuthenticationInput.Marshal(b, m, deterministic)
}
func (dst *CheckAuthenticationInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAuthenticationInput.Merge(dst, src)
}
func (m *CheckAuthenticationInput) XXX_Size() int {
	return xxx_messageInfo_CheckAuthenticationInput.Size(m)
}
func (m *CheckAuthenticationInput) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAuthenticationInput.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAuthenticationInput proto.InternalMessageInfo

func (m *CheckAuthenticationInput) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CheckAuthenticationInput) GetMethodArn() string {
	if m != nil {
		return m.MethodArn
	}
	return ""
}

type GetEventRecordsInput struct {
	Since                int64           `protobuf:"varint,1,opt,name=since,proto3" json:"since,omitempty"`
	Limit                *PrimitiveInt32 `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetEventRecordsInput) Reset()         { *m = GetEventRecordsInput{} }
func (m *GetEventRecordsInput) String() string { return proto.CompactTextString(m) }
func (*GetEventRecordsInput) ProtoMessage()    {}
func (*GetEventRecordsInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_fb9b75323b01a2e3, []int{3}
}
func (m *GetEventRecordsInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventRecordsInput.Unmarshal(m, b)
}
func (m *GetEventRecordsInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventRecordsInput.Marshal(b, m, deterministic)
}
func (dst *GetEventRecordsInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventRecordsInput.Merge(dst, src)
}
func (m *GetEventRecordsInput) XXX_Size() int {
	return xxx_messageInfo_GetEventRecordsInput.Size(m)
}
func (m *GetEventRecordsInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventRecordsInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventRecordsInput proto.InternalMessageInfo

func (m *GetEventRecordsInput) GetSince() int64 {
	if m != nil {
		return m.Since
	}
	return 0
}

func (m *GetEventRecordsInput) GetLimit() *PrimitiveInt32 {
	if m != nil {
		return m.Limit
	}
	return nil
}

type GetEventRecordsOutput struct {
	Records              []byte   `protobuf:"bytes,1,opt,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventRecordsOutput) Reset()         { *m = GetEventRecordsOutput{} }
func (m *GetEventRecordsOutput) String() string { return proto.CompactTextString(m) }
func (*GetEventRecordsOutput) ProtoMessage()    {}
func (*GetEventRecordsOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_fb9b75323b01a2e3, []int{4}
}
func (m *GetEventRecordsOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventRecordsOutput.Unmarshal(m, b)
}
func (m *GetEventRecordsOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventRecordsOutput.Marshal(b, m, deterministic)
}
func (dst *GetEventRecordsOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventRecordsOutput.Merge(dst, src)
}
func (m *GetEventRecordsOutput) XXX_Size() int {
	return xxx_messageInfo_GetEventRecordsOutput.Size(m)
}
func (m *GetEventRecordsOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventRecordsOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventRecordsOutput proto.InternalMessageInfo

func (m *GetEventRecordsOutput) GetRecords() []byte {
	if m != nil {
		return m.Records
	}
	return nil
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	EulaAgreedDate       string   `protobuf:"bytes,3,opt,name=eulaAgreedDate,proto3" json:"eulaAgreedDate,omitempty"`
	ValidEula            string   `protobuf:"bytes,4,opt,name=validEula,proto3" json:"validEula,omitempty"`
	Username             string   `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	UserStatus           string   `protobuf:"bytes,6,opt,name=userStatus,proto3" json:"userStatus,omitempty"`
	UserRoles            string   `protobuf:"bytes,7,opt,name=userRoles,proto3" json:"userRoles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_fb9b75323b01a2e3, []int{5}
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetEulaAgreedDate() string {
	if m != nil {
		return m.EulaAgreedDate
	}
	return ""
}

func (m *User) GetValidEula() string {
	if m != nil {
		return m.ValidEula
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetUserStatus() string {
	if m != nil {
		return m.UserStatus
	}
	return ""
}

func (m *User) GetUserRoles() string {
	if m != nil {
		return m.UserRoles
	}
	return ""
}

// Primitive types
type PrimitiveString struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimitiveString) Reset()         { *m = PrimitiveString{} }
func (m *PrimitiveString) String() string { return proto.CompactTextString(m) }
func (*PrimitiveString) ProtoMessage()    {}
func (*PrimitiveString) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_fb9b75323b01a2e3, []int{6}
}
func (m *PrimitiveString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimitiveString.Unmarshal(m, b)
}
func (m *PrimitiveString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimitiveString.Marshal(b, m, deterministic)
}
func (dst *PrimitiveString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveString.Merge(dst, src)
}
func (m *PrimitiveString) XXX_Size() int {
	return xxx_messageInfo_PrimitiveString.Size(m)
}
func (m *PrimitiveString) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveString.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveString proto.InternalMessageInfo

func (m *PrimitiveString) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type PrimitiveInt32 struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimitiveInt32) Reset()         { *m = PrimitiveInt32{} }
func (m *PrimitiveInt32) String() string { return proto.CompactTextString(m) }
func (*PrimitiveInt32) ProtoMessage()    {}
func (*PrimitiveInt32) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_fb9b75323b01a2e3, []int{7}
}
func (m *PrimitiveInt32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimitiveInt32.Unmarshal(m, b)
}
func (m *PrimitiveInt32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimitiveInt32.Marshal(b, m, deterministic)
}
func (dst *PrimitiveInt32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveInt32.Merge(dst, src)
}
func (m *PrimitiveInt32) XXX_Size() int {
	return xxx_messageInfo_PrimitiveInt32.Size(m)
}
func (m *PrimitiveInt32) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveInt32.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveInt32 proto.InternalMessageInfo

func (m *PrimitiveInt32) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type PrimitiveBool struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimitiveBool) Reset()         { *m = PrimitiveBool{} }
func (m *PrimitiveBool) String() string { return proto.CompactTextString(m) }
func (*PrimitiveBool) ProtoMessage()    {}
func (*PrimitiveBool) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_fb9b75323b01a2e3, []int{8}
}
func (m *PrimitiveBool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimitiveBool.Unmarshal(m, b)
}
func (m *PrimitiveBool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimitiveBool.Marshal(b, m, deterministic)
}
func (dst *PrimitiveBool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveBool.Merge(dst, src)
}
func (m *PrimitiveBool) XXX_Size() int {
	return xxx_messageInfo_PrimitiveBool.Size(m)
}
func (m *PrimitiveBool) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveBool.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveBool proto.InternalMessageInfo

func (m *PrimitiveBool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type PrimitiveBytes struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimitiveBytes) Reset()         { *m = PrimitiveBytes{} }
func (m *PrimitiveBytes) String() string { return proto.CompactTextString(m) }
func (*PrimitiveBytes) ProtoMessage()    {}
func (*PrimitiveBytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_fb9b75323b01a2e3, []int{9}
}
func (m *PrimitiveBytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimitiveBytes.Unmarshal(m, b)
}
func (m *PrimitiveBytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimitiveBytes.Marshal(b, m, deterministic)
}
func (dst *PrimitiveBytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveBytes.Merge(dst, src)
}
func (m *PrimitiveBytes) XXX_Size() int {
	return xxx_messageInfo_PrimitiveBytes.Size(m)
}
func (m *PrimitiveBytes) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveBytes.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveBytes proto.InternalMessageInfo

func (m *PrimitiveBytes) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type PrimitiveVoid struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimitiveVoid) Reset()         { *m = PrimitiveVoid{} }
func (m *PrimitiveVoid) String() string { return proto.CompactTextString(m) }
func (*PrimitiveVoid) ProtoMessage()    {}
func (*PrimitiveVoid) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_fb9b75323b01a2e3, []int{10}
}
func (m *PrimitiveVoid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimitiveVoid.Unmarshal(m, b)
}
func (m *PrimitiveVoid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimitiveVoid.Marshal(b, m, deterministic)
}
func (dst *PrimitiveVoid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveVoid.Merge(dst, src)
}
func (m *PrimitiveVoid) XXX_Size() int {
	return xxx_messageInfo_PrimitiveVoid.Size(m)
}
func (m *PrimitiveVoid) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveVoid.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveVoid proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetHierarchyRelationsInput)(nil), "grpcapi.GetHierarchyRelationsInput")
	proto.RegisterType((*GetHierarchyRelationsOutput)(nil), "grpcapi.GetHierarchyRelationsOutput")
	proto.RegisterType((*CheckAuthenticationInput)(nil), "grpcapi.CheckAuthenticationInput")
	proto.RegisterType((*GetEventRecordsInput)(nil), "grpcapi.GetEventRecordsInput")
	proto.RegisterType((*GetEventRecordsOutput)(nil), "grpcapi.GetEventRecordsOutput")
	proto.RegisterType((*User)(nil), "grpcapi.User")
	proto.RegisterType((*PrimitiveString)(nil), "grpcapi.PrimitiveString")
	proto.RegisterType((*PrimitiveInt32)(nil), "grpcapi.PrimitiveInt32")
	proto.RegisterType((*PrimitiveBool)(nil), "grpcapi.PrimitiveBool")
	proto.RegisterType((*PrimitiveBytes)(nil), "grpcapi.PrimitiveBytes")
	proto.RegisterType((*PrimitiveVoid)(nil), "grpcapi.PrimitiveVoid")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IAMClient is the client API for IAM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IAMClient interface {
	DeepPing(ctx context.Context, in *PrimitiveVoid, opts ...grpc.CallOption) (*PrimitiveString, error)
	CheckAuthentication(ctx context.Context, in *CheckAuthenticationInput, opts ...grpc.CallOption) (*User, error)
	GetHierarchyRelations(ctx context.Context, in *GetHierarchyRelationsInput, opts ...grpc.CallOption) (*GetHierarchyRelationsOutput, error)
	GetEventRecords(ctx context.Context, in *GetEventRecordsInput, opts ...grpc.CallOption) (*GetEventRecordsOutput, error)
}

type iAMClient struct {
	cc *grpc.ClientConn
}

func NewIAMClient(cc *grpc.ClientConn) IAMClient {
	return &iAMClient{cc}
}

func (c *iAMClient) DeepPing(ctx context.Context, in *PrimitiveVoid, opts ...grpc.CallOption) (*PrimitiveString, error) {
	out := new(PrimitiveString)
	err := c.cc.Invoke(ctx, "/grpcapi.IAM/DeepPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMClient) CheckAuthentication(ctx context.Context, in *CheckAuthenticationInput, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/grpcapi.IAM/CheckAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMClient) GetHierarchyRelations(ctx context.Context, in *GetHierarchyRelationsInput, opts ...grpc.CallOption) (*GetHierarchyRelationsOutput, error) {
	out := new(GetHierarchyRelationsOutput)
	err := c.cc.Invoke(ctx, "/grpcapi.IAM/GetHierarchyRelations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMClient) GetEventRecords(ctx context.Context, in *GetEventRecordsInput, opts ...grpc.CallOption) (*GetEventRecordsOutput, error) {
	out := new(GetEventRecordsOutput)
	err := c.cc.Invoke(ctx, "/grpcapi.IAM/GetEventRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IAMServer is the server API for IAM service.
type IAMServer interface {
	DeepPing(context.Context, *PrimitiveVoid) (*PrimitiveString, error)
	CheckAuthentication(context.Context, *CheckAuthenticationInput) (*User, error)
	GetHierarchyRelations(context.Context, *GetHierarchyRelationsInput) (*GetHierarchyRelationsOutput, error)
	GetEventRecords(context.Context, *GetEventRecordsInput) (*GetEventRecordsOutput, error)
}

func RegisterIAMServer(s *grpc.Server, srv IAMServer) {
	s.RegisterService(&_IAM_serviceDesc, srv)
}

func _IAM_DeepPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimitiveVoid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).DeepPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.IAM/DeepPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).DeepPing(ctx, req.(*PrimitiveVoid))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAM_CheckAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAuthenticationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).CheckAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.IAM/CheckAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).CheckAuthentication(ctx, req.(*CheckAuthenticationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAM_GetHierarchyRelations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHierarchyRelationsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).GetHierarchyRelations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.IAM/GetHierarchyRelations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).GetHierarchyRelations(ctx, req.(*GetHierarchyRelationsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAM_GetEventRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRecordsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).GetEventRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.IAM/GetEventRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).GetEventRecords(ctx, req.(*GetEventRecordsInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _IAM_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapi.IAM",
	HandlerType: (*IAMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeepPing",
			Handler:    _IAM_DeepPing_Handler,
		},
		{
			MethodName: "CheckAuthentication",
			Handler:    _IAM_CheckAuthentication_Handler,
		},
		{
			MethodName: "GetHierarchyRelations",
			Handler:    _IAM_GetHierarchyRelations_Handler,
		},
		{
			MethodName: "GetEventRecords",
			Handler:    _IAM_GetEventRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iam/grpcapi.proto",
}

func init() { proto.RegisterFile("iam/grpcapi.proto", fileDescriptor_grpcapi_fb9b75323b01a2e3) }

var fileDescriptor_grpcapi_fb9b75323b01a2e3 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x6d, 0x92, 0xa6, 0x49, 0x87, 0x36, 0x11, 0xdb, 0x42, 0x8d, 0x0b, 0x55, 0x31, 0x50, 0x7a,
	0xc1, 0x88, 0x54, 0x48, 0x88, 0x03, 0x92, 0x43, 0x43, 0xb1, 0x50, 0x21, 0x72, 0x04, 0x17, 0x0e,
	0xd5, 0xe2, 0x1d, 0x92, 0x55, 0xec, 0xdd, 0x68, 0xbd, 0x8e, 0xd4, 0xbf, 0xc4, 0x4f, 0x41, 0xfc,
	0x28, 0xb4, 0x76, 0x6c, 0xf2, 0xc9, 0xcd, 0xef, 0xcd, 0x9b, 0xb7, 0x3b, 0xb3, 0x33, 0x86, 0xbb,
	0x9c, 0xc6, 0x2f, 0x87, 0x6a, 0x12, 0xd2, 0x09, 0x77, 0x27, 0x4a, 0x6a, 0x49, 0x1a, 0x33, 0xe8,
	0xbc, 0x06, 0xfb, 0x0a, 0xf5, 0x47, 0x8e, 0x8a, 0xaa, 0x70, 0x74, 0x1b, 0x60, 0x44, 0x35, 0x97,
	0x22, 0xf1, 0xc5, 0x24, 0xd5, 0xe4, 0x08, 0x1a, 0x69, 0x82, 0xea, 0x86, 0x33, 0xab, 0x72, 0x5a,
	0x39, 0xdf, 0x0d, 0x76, 0x0c, 0xf4, 0x99, 0xf3, 0x06, 0x8e, 0xd7, 0xa6, 0x7d, 0x49, 0xb5, 0xc9,
	0x7b, 0x00, 0x4d, 0x21, 0x19, 0xde, 0x70, 0x96, 0x58, 0x95, 0xd3, 0xda, 0xf9, 0x6e, 0xd0, 0x30,
	0xd8, 0x67, 0x89, 0xf3, 0x19, 0xac, 0xf7, 0x23, 0x0c, 0xc7, 0x5e, 0xaa, 0x47, 0x28, 0x34, 0x0f,
	0xb3, 0xc4, 0xfc, 0xb8, 0x43, 0xa8, 0x6b, 0x39, 0x46, 0x31, 0x3b, 0x2c, 0x07, 0xe4, 0x21, 0xec,
	0xc6, 0xa8, 0x47, 0x92, 0x79, 0x4a, 0x58, 0xd5, 0x2c, 0xf2, 0x8f, 0x70, 0xbe, 0xc3, 0xe1, 0x15,
	0xea, 0xde, 0x14, 0x85, 0x0e, 0x30, 0x94, 0x8a, 0x25, 0xa5, 0x57, 0xc2, 0x45, 0x88, 0x99, 0x57,
	0x2d, 0xc8, 0x01, 0x79, 0x01, 0xf5, 0x88, 0xc7, 0x5c, 0x67, 0x3e, 0x77, 0x3a, 0x47, 0x6e, 0xd1,
	0x96, 0xbe, 0x32, 0x34, 0x9f, 0xa2, 0x2f, 0xf4, 0x45, 0x27, 0xc8, 0x55, 0xce, 0x2b, 0xb8, 0xb7,
	0x64, 0x3e, 0x2b, 0xd0, 0x82, 0x86, 0xca, 0x89, 0xcc, 0x7f, 0x2f, 0x28, 0xa0, 0xf3, 0xa7, 0x02,
	0xdb, 0x5f, 0x13, 0x54, 0xa4, 0x05, 0xd5, 0xb2, 0x6d, 0x55, 0xce, 0xcc, 0x85, 0x30, 0xa6, 0x3c,
	0x9a, 0x95, 0x90, 0x03, 0x72, 0x06, 0x2d, 0x4c, 0x23, 0xea, 0x0d, 0x15, 0x22, 0xbb, 0xa4, 0x1a,
	0xad, 0x5a, 0x16, 0x5e, 0x62, 0x4d, 0x13, 0xa6, 0x34, 0xe2, 0xac, 0x97, 0x46, 0xd4, 0xda, 0xce,
	0x9b, 0x50, 0x12, 0xc4, 0x86, 0xa6, 0x79, 0x18, 0x41, 0x63, 0xb4, 0xea, 0x59, 0xb0, 0xc4, 0xe4,
	0x04, 0xc0, 0x7c, 0x0f, 0x34, 0xd5, 0x69, 0x62, 0xed, 0x64, 0xd1, 0x39, 0xc6, 0x38, 0x1b, 0x14,
	0xc8, 0x08, 0x13, 0xab, 0x91, 0x3b, 0x97, 0x84, 0xf3, 0x1c, 0xda, 0x65, 0x6b, 0x06, 0x5a, 0x71,
	0x31, 0x34, 0x85, 0x4c, 0x69, 0x94, 0x62, 0xf1, 0x4a, 0x19, 0x70, 0xce, 0xa0, 0xb5, 0xd8, 0xc3,
	0x45, 0x5d, 0xbd, 0xd0, 0x3d, 0x83, 0xfd, 0x52, 0xd7, 0x95, 0x32, 0x5a, 0x94, 0x35, 0xd7, 0xd9,
	0x75, 0x6f, 0x35, 0x26, 0x8b, 0xba, 0xbd, 0x42, 0xd7, 0x9e, 0xb3, 0xfb, 0x26, 0x39, 0xeb, 0xfc,
	0xae, 0x42, 0xcd, 0xf7, 0xae, 0xc9, 0x3b, 0x68, 0x5e, 0x22, 0x4e, 0xfa, 0xe6, 0xc6, 0xf7, 0x57,
	0x9f, 0xd9, 0x68, 0x6d, 0x6b, 0x95, 0xcf, 0x6b, 0x74, 0xb6, 0x88, 0x0f, 0x07, 0x6b, 0xe6, 0x94,
	0x3c, 0x2e, 0x53, 0x36, 0x4d, 0xb1, 0xbd, 0x5f, 0x4a, 0xcc, 0x1c, 0x38, 0x5b, 0x84, 0x65, 0x53,
	0xb4, 0xba, 0x2c, 0xe4, 0x49, 0xa9, 0xdc, 0xbc, 0x83, 0xf6, 0xd3, 0xff, 0x8b, 0xf2, 0x81, 0x74,
	0xb6, 0x48, 0x00, 0xed, 0xa5, 0x59, 0x25, 0x8f, 0xe6, 0x53, 0x57, 0x56, 0xc4, 0x3e, 0xd9, 0x14,
	0x2e, 0x3c, 0xbb, 0x6f, 0xe1, 0x38, 0x94, 0xb1, 0x9b, 0x8c, 0x7f, 0xba, 0x28, 0x22, 0x3e, 0x1c,
	0x69, 0x97, 0xd3, 0xb8, 0x48, 0xeb, 0xd6, 0x38, 0x8d, 0xfb, 0x95, 0x5f, 0xd5, 0x83, 0xc1, 0xa7,
	0x0f, 0x6e, 0x6f, 0x26, 0xf0, 0xfa, 0xbe, 0xeb, 0x7b, 0xd7, 0x3f, 0x76, 0xb2, 0x3f, 0xcd, 0xc5,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x4d, 0x12, 0xcb, 0x7e, 0x04, 0x00, 0x00,
}