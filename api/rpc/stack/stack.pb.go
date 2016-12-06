// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/stack/stack.proto
// DO NOT EDIT!

/*
Package stack is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/stack/stack.proto

It has these top-level messages:
	StackFileRequest
	StackRequest
	RemoveRequest
	StackReply
	ListRequest
	ListReply
	TasksRequest
	TasksReply
	StackInfo
	StackID
	CustomNetwork
	IdList
	NetworkSpec
	NetworkIPAM
	NetworkIPAMConfig
	Stack
*/
package stack

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import service "github.com/appcelerator/amp/api/rpc/service"

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

type StackState int32

const (
	StackState_Stopped     StackState = 0
	StackState_Starting    StackState = 1
	StackState_Running     StackState = 2
	StackState_Redeploying StackState = 3
)

var StackState_name = map[int32]string{
	0: "Stopped",
	1: "Starting",
	2: "Running",
	3: "Redeploying",
}
var StackState_value = map[string]int32{
	"Stopped":     0,
	"Starting":    1,
	"Running":     2,
	"Redeploying": 3,
}

func (x StackState) String() string {
	return proto.EnumName(StackState_name, int32(x))
}
func (StackState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// struct for stackfile request
type StackFileRequest struct {
	Stack *Stack `protobuf:"bytes,1,opt,name=stack" json:"stack,omitempty"`
}

func (m *StackFileRequest) Reset()                    { *m = StackFileRequest{} }
func (m *StackFileRequest) String() string            { return proto.CompactTextString(m) }
func (*StackFileRequest) ProtoMessage()               {}
func (*StackFileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StackFileRequest) GetStack() *Stack {
	if m != nil {
		return m.Stack
	}
	return nil
}

// struct stack name/id based requests
type StackRequest struct {
	StackIdent string `protobuf:"bytes,1,opt,name=stack_ident,json=stackIdent" json:"stack_ident,omitempty"`
}

func (m *StackRequest) Reset()                    { *m = StackRequest{} }
func (m *StackRequest) String() string            { return proto.CompactTextString(m) }
func (*StackRequest) ProtoMessage()               {}
func (*StackRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StackRequest) GetStackIdent() string {
	if m != nil {
		return m.StackIdent
	}
	return ""
}

// struct for remove request function
type RemoveRequest struct {
	StackIdent string `protobuf:"bytes,1,opt,name=stack_ident,json=stackIdent" json:"stack_ident,omitempty"`
	Force      bool   `protobuf:"varint,2,opt,name=force" json:"force,omitempty"`
}

func (m *RemoveRequest) Reset()                    { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()               {}
func (*RemoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RemoveRequest) GetStackIdent() string {
	if m != nil {
		return m.StackIdent
	}
	return ""
}

func (m *RemoveRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

// struct for stack id responses
type StackReply struct {
	StackId string `protobuf:"bytes,1,opt,name=stack_id,json=stackId" json:"stack_id,omitempty"`
}

func (m *StackReply) Reset()                    { *m = StackReply{} }
func (m *StackReply) String() string            { return proto.CompactTextString(m) }
func (*StackReply) ProtoMessage()               {}
func (*StackReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StackReply) GetStackId() string {
	if m != nil {
		return m.StackId
	}
	return ""
}

// struct for list request function
type ListRequest struct {
	All   bool  `protobuf:"varint,1,opt,name=all" json:"all,omitempty"`
	Limit int64 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListRequest) GetAll() bool {
	if m != nil {
		return m.All
	}
	return false
}

func (m *ListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// struct for list reply function
type ListReply struct {
	List []*StackInfo `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListReply) GetList() []*StackInfo {
	if m != nil {
		return m.List
	}
	return nil
}

// struct for tasks request function
type TasksRequest struct {
	StackIdent string `protobuf:"bytes,1,opt,name=stack_ident,json=stackIdent" json:"stack_ident,omitempty"`
}

func (m *TasksRequest) Reset()                    { *m = TasksRequest{} }
func (m *TasksRequest) String() string            { return proto.CompactTextString(m) }
func (*TasksRequest) ProtoMessage()               {}
func (*TasksRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TasksRequest) GetStackIdent() string {
	if m != nil {
		return m.StackIdent
	}
	return ""
}

// struct for tasks reply function
type TasksReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *TasksReply) Reset()                    { *m = TasksReply{} }
func (m *TasksReply) String() string            { return proto.CompactTextString(m) }
func (*TasksReply) ProtoMessage()               {}
func (*TasksReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TasksReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// struct part of ListReply Struct
type StackInfo struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	State string `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
}

func (m *StackInfo) Reset()                    { *m = StackInfo{} }
func (m *StackInfo) String() string            { return proto.CompactTextString(m) }
func (*StackInfo) ProtoMessage()               {}
func (*StackInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *StackInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StackInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StackInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

// struct to store Stack id in ETCD
type StackID struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *StackID) Reset()                    { *m = StackID{} }
func (m *StackID) String() string            { return proto.CompactTextString(m) }
func (*StackID) ProtoMessage()               {}
func (*StackID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *StackID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// struct to store network info in ETCD
type CustomNetwork struct {
	Id          string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OwnerNumber int32        `protobuf:"varint,2,opt,name=owner_number,json=ownerNumber" json:"owner_number,omitempty"`
	Data        *NetworkSpec `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *CustomNetwork) Reset()                    { *m = CustomNetwork{} }
func (m *CustomNetwork) String() string            { return proto.CompactTextString(m) }
func (*CustomNetwork) ProtoMessage()               {}
func (*CustomNetwork) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CustomNetwork) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CustomNetwork) GetOwnerNumber() int32 {
	if m != nil {
		return m.OwnerNumber
	}
	return 0
}

func (m *CustomNetwork) GetData() *NetworkSpec {
	if m != nil {
		return m.Data
	}
	return nil
}

// struct to store service id list in ETCD
type IdList struct {
	List []string `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *IdList) Reset()                    { *m = IdList{} }
func (m *IdList) String() string            { return proto.CompactTextString(m) }
func (*IdList) ProtoMessage()               {}
func (*IdList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *IdList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type NetworkSpec struct {
	Name       string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Driver     string            `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	EnableIpv6 bool              `protobuf:"varint,3,opt,name=enable_ipv6,json=enableIpv6" json:"enable_ipv6,omitempty"`
	Ipam       *NetworkIPAM      `protobuf:"bytes,4,opt,name=ipam" json:"ipam,omitempty"`
	Internal   bool              `protobuf:"varint,5,opt,name=internal" json:"internal,omitempty"`
	Options    map[string]string `protobuf:"bytes,6,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Labels     map[string]string `protobuf:"bytes,7,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	External   string            `protobuf:"bytes,8,opt,name=external" json:"external,omitempty"`
}

func (m *NetworkSpec) Reset()                    { *m = NetworkSpec{} }
func (m *NetworkSpec) String() string            { return proto.CompactTextString(m) }
func (*NetworkSpec) ProtoMessage()               {}
func (*NetworkSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *NetworkSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkSpec) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *NetworkSpec) GetEnableIpv6() bool {
	if m != nil {
		return m.EnableIpv6
	}
	return false
}

func (m *NetworkSpec) GetIpam() *NetworkIPAM {
	if m != nil {
		return m.Ipam
	}
	return nil
}

func (m *NetworkSpec) GetInternal() bool {
	if m != nil {
		return m.Internal
	}
	return false
}

func (m *NetworkSpec) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *NetworkSpec) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *NetworkSpec) GetExternal() string {
	if m != nil {
		return m.External
	}
	return ""
}

type NetworkIPAM struct {
	Driver  string               `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	Options map[string]string    `protobuf:"bytes,2,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Config  []*NetworkIPAMConfig `protobuf:"bytes,3,rep,name=config" json:"config,omitempty"`
}

func (m *NetworkIPAM) Reset()                    { *m = NetworkIPAM{} }
func (m *NetworkIPAM) String() string            { return proto.CompactTextString(m) }
func (*NetworkIPAM) ProtoMessage()               {}
func (*NetworkIPAM) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *NetworkIPAM) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *NetworkIPAM) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *NetworkIPAM) GetConfig() []*NetworkIPAMConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type NetworkIPAMConfig struct {
	Subnet     string            `protobuf:"bytes,1,opt,name=subnet" json:"subnet,omitempty"`
	IpRange    string            `protobuf:"bytes,2,opt,name=ip_range,json=ipRange" json:"ip_range,omitempty"`
	Gateway    string            `protobuf:"bytes,3,opt,name=gateway" json:"gateway,omitempty"`
	AuxAddress map[string]string `protobuf:"bytes,4,rep,name=aux_address,json=auxAddress" json:"aux_address,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NetworkIPAMConfig) Reset()                    { *m = NetworkIPAMConfig{} }
func (m *NetworkIPAMConfig) String() string            { return proto.CompactTextString(m) }
func (*NetworkIPAMConfig) ProtoMessage()               {}
func (*NetworkIPAMConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *NetworkIPAMConfig) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

func (m *NetworkIPAMConfig) GetIpRange() string {
	if m != nil {
		return m.IpRange
	}
	return ""
}

func (m *NetworkIPAMConfig) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *NetworkIPAMConfig) GetAuxAddress() map[string]string {
	if m != nil {
		return m.AuxAddress
	}
	return nil
}

// Stack struct
type Stack struct {
	Name     string                 `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id       string                 `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Services []*service.ServiceSpec `protobuf:"bytes,3,rep,name=services" json:"services,omitempty"`
	Networks []*NetworkSpec         `protobuf:"bytes,4,rep,name=networks" json:"networks,omitempty"`
	IsPublic bool                   `protobuf:"varint,5,opt,name=is_public,json=isPublic" json:"is_public,omitempty"`
}

func (m *Stack) Reset()                    { *m = Stack{} }
func (m *Stack) String() string            { return proto.CompactTextString(m) }
func (*Stack) ProtoMessage()               {}
func (*Stack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *Stack) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stack) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Stack) GetServices() []*service.ServiceSpec {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *Stack) GetNetworks() []*NetworkSpec {
	if m != nil {
		return m.Networks
	}
	return nil
}

func (m *Stack) GetIsPublic() bool {
	if m != nil {
		return m.IsPublic
	}
	return false
}

func init() {
	proto.RegisterType((*StackFileRequest)(nil), "stack.StackFileRequest")
	proto.RegisterType((*StackRequest)(nil), "stack.StackRequest")
	proto.RegisterType((*RemoveRequest)(nil), "stack.RemoveRequest")
	proto.RegisterType((*StackReply)(nil), "stack.StackReply")
	proto.RegisterType((*ListRequest)(nil), "stack.ListRequest")
	proto.RegisterType((*ListReply)(nil), "stack.ListReply")
	proto.RegisterType((*TasksRequest)(nil), "stack.TasksRequest")
	proto.RegisterType((*TasksReply)(nil), "stack.TasksReply")
	proto.RegisterType((*StackInfo)(nil), "stack.StackInfo")
	proto.RegisterType((*StackID)(nil), "stack.StackID")
	proto.RegisterType((*CustomNetwork)(nil), "stack.CustomNetwork")
	proto.RegisterType((*IdList)(nil), "stack.IdList")
	proto.RegisterType((*NetworkSpec)(nil), "stack.NetworkSpec")
	proto.RegisterType((*NetworkIPAM)(nil), "stack.NetworkIPAM")
	proto.RegisterType((*NetworkIPAMConfig)(nil), "stack.NetworkIPAMConfig")
	proto.RegisterType((*Stack)(nil), "stack.Stack")
	proto.RegisterEnum("stack.StackState", StackState_name, StackState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StackService service

type StackServiceClient interface {
	Up(ctx context.Context, in *StackFileRequest, opts ...grpc.CallOption) (*StackReply, error)
	Create(ctx context.Context, in *StackFileRequest, opts ...grpc.CallOption) (*StackReply, error)
	Start(ctx context.Context, in *StackRequest, opts ...grpc.CallOption) (*StackReply, error)
	Stop(ctx context.Context, in *StackRequest, opts ...grpc.CallOption) (*StackReply, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*StackReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	Tasks(ctx context.Context, in *TasksRequest, opts ...grpc.CallOption) (*TasksReply, error)
}

type stackServiceClient struct {
	cc *grpc.ClientConn
}

func NewStackServiceClient(cc *grpc.ClientConn) StackServiceClient {
	return &stackServiceClient{cc}
}

func (c *stackServiceClient) Up(ctx context.Context, in *StackFileRequest, opts ...grpc.CallOption) (*StackReply, error) {
	out := new(StackReply)
	err := grpc.Invoke(ctx, "/stack.StackService/Up", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackServiceClient) Create(ctx context.Context, in *StackFileRequest, opts ...grpc.CallOption) (*StackReply, error) {
	out := new(StackReply)
	err := grpc.Invoke(ctx, "/stack.StackService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackServiceClient) Start(ctx context.Context, in *StackRequest, opts ...grpc.CallOption) (*StackReply, error) {
	out := new(StackReply)
	err := grpc.Invoke(ctx, "/stack.StackService/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackServiceClient) Stop(ctx context.Context, in *StackRequest, opts ...grpc.CallOption) (*StackReply, error) {
	out := new(StackReply)
	err := grpc.Invoke(ctx, "/stack.StackService/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackServiceClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*StackReply, error) {
	out := new(StackReply)
	err := grpc.Invoke(ctx, "/stack.StackService/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/stack.StackService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackServiceClient) Tasks(ctx context.Context, in *TasksRequest, opts ...grpc.CallOption) (*TasksReply, error) {
	out := new(TasksReply)
	err := grpc.Invoke(ctx, "/stack.StackService/Tasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StackService service

type StackServiceServer interface {
	Up(context.Context, *StackFileRequest) (*StackReply, error)
	Create(context.Context, *StackFileRequest) (*StackReply, error)
	Start(context.Context, *StackRequest) (*StackReply, error)
	Stop(context.Context, *StackRequest) (*StackReply, error)
	Remove(context.Context, *RemoveRequest) (*StackReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	Tasks(context.Context, *TasksRequest) (*TasksReply, error)
}

func RegisterStackServiceServer(s *grpc.Server, srv StackServiceServer) {
	s.RegisterService(&_StackService_serviceDesc, srv)
}

func _StackService_Up_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).Up(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/Up",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).Up(ctx, req.(*StackFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).Create(ctx, req.(*StackFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).Start(ctx, req.(*StackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).Stop(ctx, req.(*StackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackService_Tasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).Tasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/Tasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).Tasks(ctx, req.(*TasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StackService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stack.StackService",
	HandlerType: (*StackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Up",
			Handler:    _StackService_Up_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _StackService_Create_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _StackService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _StackService_Stop_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _StackService_Remove_Handler,
		},
		{
			MethodName: "List",
			Handler:    _StackService_List_Handler,
		},
		{
			MethodName: "Tasks",
			Handler:    _StackService_Tasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/stack/stack.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/stack/stack.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1008 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0x5f, 0x6f, 0x1b, 0x45,
	0x10, 0xe7, 0x6c, 0xc7, 0x7f, 0xe6, 0xdc, 0xd6, 0x59, 0x22, 0x7a, 0x31, 0xa1, 0x29, 0xa7, 0x52,
	0xa2, 0x3c, 0xd8, 0x6d, 0x50, 0x23, 0x12, 0x09, 0xa1, 0x28, 0xa4, 0x92, 0x51, 0x29, 0xd5, 0xa6,
	0xf0, 0x6a, 0xad, 0x7d, 0x1b, 0xb3, 0xe4, 0x7c, 0xb7, 0xdc, 0xed, 0x39, 0xb1, 0x10, 0x2f, 0x7c,
	0x05, 0xbe, 0x02, 0x4f, 0x7c, 0x0a, 0x3e, 0x01, 0x2f, 0x7c, 0x05, 0x9e, 0xf8, 0x14, 0x68, 0x67,
	0xf7, 0x2e, 0xe7, 0xc6, 0x96, 0x1c, 0xf1, 0x12, 0xdf, 0xcc, 0xce, 0xef, 0x37, 0x33, 0xbf, 0xbb,
	0x9d, 0x09, 0xbc, 0x98, 0x08, 0xf5, 0x43, 0x36, 0xea, 0x8d, 0xe3, 0x69, 0x9f, 0x49, 0x39, 0xe6,
	0x21, 0x4f, 0x98, 0x8a, 0x93, 0x3e, 0x9b, 0xca, 0x3e, 0x93, 0xa2, 0x9f, 0xc8, 0x71, 0x3f, 0x55,
	0x6c, 0x7c, 0x69, 0xfe, 0xf6, 0x64, 0x12, 0xab, 0x98, 0x6c, 0xa0, 0xd1, 0xdd, 0x99, 0xc4, 0xf1,
	0x24, 0xe4, 0x18, 0xc8, 0xa2, 0x28, 0x56, 0x4c, 0x89, 0x38, 0x4a, 0x4d, 0x50, 0xf7, 0x68, 0x2d,
	0x6e, 0x9e, 0xcc, 0xc4, 0x98, 0xe7, 0xbf, 0x06, 0xea, 0x1f, 0x42, 0xe7, 0x5c, 0x67, 0x78, 0x29,
	0x42, 0x4e, 0xf9, 0x4f, 0x19, 0x4f, 0x15, 0xf1, 0xc1, 0x64, 0xf5, 0x9c, 0xc7, 0xce, 0x9e, 0x7b,
	0xd0, 0xee, 0x99, 0x82, 0x30, 0x8e, 0x9a, 0x23, 0xbf, 0x0f, 0x6d, 0x63, 0x5b, 0xcc, 0x2e, 0xb8,
	0x78, 0x30, 0x14, 0x01, 0x8f, 0x14, 0x22, 0x5b, 0x14, 0xd0, 0x35, 0xd0, 0x1e, 0xff, 0x25, 0xdc,
	0xa3, 0x7c, 0x1a, 0xcf, 0xf8, 0xba, 0x08, 0xb2, 0x05, 0x1b, 0x17, 0x71, 0x32, 0xe6, 0x5e, 0xe5,
	0xb1, 0xb3, 0xd7, 0xa4, 0xc6, 0xf0, 0x3f, 0x05, 0xb0, 0x89, 0x65, 0x38, 0x27, 0xdb, 0xd0, 0xcc,
	0x49, 0x2c, 0x43, 0xc3, 0x32, 0xf8, 0x2f, 0xc0, 0x7d, 0x25, 0x52, 0x95, 0xa7, 0xeb, 0x40, 0x95,
	0x85, 0x21, 0x06, 0x35, 0xa9, 0x7e, 0xd4, 0xfc, 0xa1, 0x98, 0x0a, 0x85, 0xfc, 0x55, 0x6a, 0x0c,
	0xff, 0x39, 0xb4, 0x0c, 0x4c, 0xd3, 0x3f, 0x81, 0x5a, 0x28, 0x52, 0x5d, 0x5c, 0x75, 0xcf, 0x3d,
	0xe8, 0x94, 0x85, 0x18, 0x44, 0x17, 0x31, 0xc5, 0x53, 0xad, 0xc5, 0x5b, 0x96, 0x5e, 0xa6, 0x6b,
	0x6b, 0xf1, 0x14, 0xc0, 0x02, 0x74, 0x12, 0x0f, 0x1a, 0x53, 0x9e, 0xa6, 0x6c, 0xc2, 0xf3, 0x16,
	0xac, 0xe9, 0x9f, 0x41, 0xab, 0xc8, 0x45, 0x08, 0xd4, 0x22, 0x36, 0xcd, 0x63, 0xf0, 0x99, 0xdc,
	0x87, 0x8a, 0x08, 0xb0, 0xfe, 0x16, 0xad, 0x88, 0x40, 0xb7, 0x94, 0x2a, 0xa6, 0xb8, 0x57, 0x45,
	0x97, 0x31, 0xfc, 0x6d, 0x68, 0x18, 0x9a, 0xaf, 0x2c, 0xc0, 0xc9, 0x01, 0xfe, 0x8f, 0x70, 0xef,
	0x34, 0x4b, 0x55, 0x3c, 0x7d, 0xcd, 0xd5, 0x55, 0x9c, 0x5c, 0xbe, 0x1b, 0x40, 0x3e, 0x86, 0x76,
	0x7c, 0x15, 0xf1, 0x64, 0x18, 0x65, 0xd3, 0x11, 0x4f, 0x30, 0xd7, 0x06, 0x75, 0xd1, 0xf7, 0x1a,
	0x5d, 0xe4, 0x29, 0xd4, 0x02, 0xa6, 0x18, 0xe6, 0x74, 0x0f, 0x88, 0x15, 0xc9, 0x12, 0x9e, 0x4b,
	0x3e, 0xa6, 0x78, 0xee, 0xef, 0x40, 0x7d, 0x10, 0x68, 0x6d, 0x75, 0x2b, 0x85, 0xac, 0x2d, 0x2b,
	0xe2, 0xef, 0x55, 0x70, 0x4b, 0x98, 0xa5, 0xed, 0x7e, 0x00, 0xf5, 0x20, 0x11, 0x33, 0x5b, 0x46,
	0x8b, 0x5a, 0x4b, 0x0b, 0xce, 0x23, 0x36, 0x0a, 0xf9, 0x50, 0xc8, 0xd9, 0x21, 0x16, 0xd2, 0xa4,
	0x60, 0x5c, 0x03, 0x39, 0x3b, 0xd4, 0x25, 0x0a, 0xc9, 0xa6, 0x5e, 0x6d, 0x59, 0x89, 0x83, 0x37,
	0x27, 0xdf, 0x50, 0x3c, 0x27, 0x5d, 0x68, 0x8a, 0x48, 0xf1, 0x24, 0x62, 0xa1, 0xb7, 0x81, 0x2c,
	0x85, 0x4d, 0x8e, 0xa0, 0x11, 0x4b, 0xbc, 0x75, 0x5e, 0x1d, 0x3f, 0x87, 0xdd, 0xdb, 0x9d, 0xf6,
	0xbe, 0x35, 0x11, 0x67, 0x91, 0x4a, 0xe6, 0x34, 0x8f, 0x27, 0x87, 0x50, 0x0f, 0xd9, 0x88, 0x87,
	0xa9, 0xd7, 0x40, 0xe4, 0xa3, 0x25, 0xc8, 0x57, 0x18, 0x60, 0x80, 0x36, 0x5a, 0x97, 0xc3, 0xaf,
	0x6d, 0x39, 0x4d, 0xec, 0xb8, 0xb0, 0xbb, 0xc7, 0xd0, 0x2e, 0x27, 0xd3, 0xdf, 0xf7, 0x25, 0x9f,
	0x5b, 0xb9, 0xf4, 0xa3, 0xfe, 0x18, 0x66, 0x2c, 0xcc, 0xb8, 0x15, 0xcb, 0x18, 0xc7, 0x95, 0xcf,
	0x9d, 0xee, 0x11, 0xb8, 0xa5, 0x74, 0x77, 0x81, 0xfa, 0x7f, 0x39, 0xc5, 0x6b, 0xd2, 0xba, 0x95,
	0x5e, 0x89, 0xb3, 0xf0, 0x4a, 0x4a, 0x6a, 0x55, 0x96, 0xa9, 0xa5, 0xc1, 0x2b, 0xd4, 0x7a, 0x06,
	0xf5, 0x71, 0x1c, 0x5d, 0x88, 0x89, 0x57, 0x45, 0xa4, 0x77, 0x1b, 0x79, 0x8a, 0xe7, 0xd4, 0xc6,
	0xfd, 0x1f, 0x2d, 0xfc, 0x7f, 0x1d, 0xd8, 0xbc, 0xc5, 0xac, 0xdb, 0x4a, 0xb3, 0x51, 0xc4, 0xf3,
	0xdb, 0x6b, 0x2d, 0x3d, 0x6f, 0x84, 0x1c, 0x26, 0x2c, 0x9a, 0xe4, 0x54, 0x0d, 0x21, 0xa9, 0x36,
	0xf5, 0x35, 0x9e, 0x30, 0xc5, 0xaf, 0xd8, 0xdc, 0xde, 0xbe, 0xdc, 0x24, 0x03, 0x70, 0x59, 0x76,
	0x3d, 0x64, 0x41, 0x90, 0xf0, 0x34, 0xf5, 0x6a, 0xd8, 0xd5, 0xde, 0xaa, 0xae, 0x7a, 0x27, 0xd9,
	0xf5, 0x89, 0x09, 0x35, 0xc2, 0x00, 0x2b, 0x1c, 0xdd, 0x2f, 0xe0, 0xc1, 0x3b, 0xc7, 0x77, 0x6a,
	0xf6, 0x0f, 0x07, 0x36, 0x70, 0x14, 0xac, 0x35, 0x4d, 0x9e, 0x41, 0xd3, 0x2e, 0x8b, 0xd4, 0xbe,
	0x8a, 0xad, 0x5e, 0xbe, 0x3d, 0xce, 0xcd, 0x2f, 0x5e, 0xef, 0x22, 0x8a, 0xf4, 0xa0, 0x19, 0x99,
	0x7e, 0xf2, 0x36, 0x97, 0x8d, 0x83, 0x22, 0x86, 0x7c, 0x08, 0x2d, 0x91, 0x0e, 0x65, 0x36, 0x0a,
	0xc5, 0xb8, 0xb8, 0x70, 0xe9, 0x1b, 0xb4, 0xf7, 0xcf, 0xec, 0xa4, 0x3f, 0xd7, 0x43, 0x8c, 0xb8,
	0x7a, 0x88, 0xc5, 0x52, 0xf2, 0xa0, 0xf3, 0x1e, 0x69, 0x43, 0xf3, 0x5c, 0xb1, 0x44, 0x89, 0x68,
	0xd2, 0x71, 0xf4, 0x11, 0xcd, 0xa2, 0x48, 0x1b, 0x15, 0xf2, 0x00, 0x5c, 0xca, 0x03, 0x2e, 0xc3,
	0x78, 0xae, 0x1d, 0xd5, 0x83, 0x3f, 0x6b, 0x76, 0x55, 0xd9, 0x92, 0xc9, 0x00, 0x2a, 0xdf, 0x49,
	0xf2, 0xb0, 0x3c, 0xcc, 0x4b, 0xdb, 0xaf, 0xbb, 0xb9, 0xb0, 0xee, 0xf4, 0x84, 0xf6, 0x1f, 0xfe,
	0xfa, 0xf7, 0x3f, 0xbf, 0x55, 0x36, 0xfd, 0x76, 0x7f, 0xf6, 0xdc, 0xee, 0xe8, 0x4c, 0x1e, 0x3b,
	0xfb, 0xe4, 0x6b, 0xa8, 0x9f, 0x26, 0x5c, 0x97, 0x77, 0x17, 0xba, 0x2d, 0xa4, 0xbb, 0xef, 0xb7,
	0x0a, 0x3a, 0xcd, 0xf5, 0x3d, 0xbe, 0x9a, 0x44, 0x91, 0xf7, 0x17, 0x11, 0x2b, 0x69, 0x3e, 0x41,
	0x9a, 0x5d, 0xff, 0xa3, 0x9b, 0xaa, 0x7e, 0x2e, 0xed, 0x9d, 0x5f, 0xb4, 0x2f, 0x51, 0xe4, 0x2d,
	0xd4, 0xb4, 0x70, 0x6b, 0xd3, 0x3e, 0x41, 0xda, 0x47, 0xfe, 0xce, 0x6a, 0xda, 0x58, 0x12, 0x0a,
	0x75, 0xb3, 0xce, 0xc9, 0x96, 0xa5, 0x58, 0xd8, 0xee, 0xcb, 0x88, 0x77, 0x91, 0x78, 0x7b, 0xff,
	0xe1, 0x0a, 0x62, 0xf2, 0x25, 0xd4, 0xcc, 0x7a, 0xb0, 0xd8, 0xd2, 0xfa, 0xee, 0x76, 0x16, 0x7c,
	0x9a, 0x6e, 0x13, 0xe9, 0x5c, 0x72, 0xa3, 0xa2, 0x96, 0x10, 0xf7, 0x6a, 0xd1, 0x6b, 0x79, 0x2d,
	0x17, 0x25, 0xdd, 0xac, 0xde, 0x5c, 0x42, 0xb2, 0x52, 0x42, 0xa5, 0x63, 0x47, 0x75, 0xfc, 0x5f,
	0xe9, 0xb3, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x45, 0x40, 0x3c, 0x4a, 0xc4, 0x09, 0x00, 0x00,
}
