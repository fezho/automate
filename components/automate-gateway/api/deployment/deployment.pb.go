// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/deployment/deployment.proto

package deployment

import (
	context "context"
	fmt "fmt"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Version message
//
// The manifest version constructed with:
// * build_timestamp
type Version struct {
	BuildTimestamp       string   `protobuf:"bytes,1,opt,name=build_timestamp,json=buildTimestamp,proto3" json:"build_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4362896e7bedc82, []int{0}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetBuildTimestamp() string {
	if m != nil {
		return m.BuildTimestamp
	}
	return ""
}

type ServiceVersionsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceVersionsRequest) Reset()         { *m = ServiceVersionsRequest{} }
func (m *ServiceVersionsRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceVersionsRequest) ProtoMessage()    {}
func (*ServiceVersionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4362896e7bedc82, []int{1}
}

func (m *ServiceVersionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceVersionsRequest.Unmarshal(m, b)
}
func (m *ServiceVersionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceVersionsRequest.Marshal(b, m, deterministic)
}
func (m *ServiceVersionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceVersionsRequest.Merge(m, src)
}
func (m *ServiceVersionsRequest) XXX_Size() int {
	return xxx_messageInfo_ServiceVersionsRequest.Size(m)
}
func (m *ServiceVersionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceVersionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceVersionsRequest proto.InternalMessageInfo

type ServiceVersionsResponse struct {
	Services             []*ServiceVersion `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServiceVersionsResponse) Reset()         { *m = ServiceVersionsResponse{} }
func (m *ServiceVersionsResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceVersionsResponse) ProtoMessage()    {}
func (*ServiceVersionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4362896e7bedc82, []int{2}
}

func (m *ServiceVersionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceVersionsResponse.Unmarshal(m, b)
}
func (m *ServiceVersionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceVersionsResponse.Marshal(b, m, deterministic)
}
func (m *ServiceVersionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceVersionsResponse.Merge(m, src)
}
func (m *ServiceVersionsResponse) XXX_Size() int {
	return xxx_messageInfo_ServiceVersionsResponse.Size(m)
}
func (m *ServiceVersionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceVersionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceVersionsResponse proto.InternalMessageInfo

func (m *ServiceVersionsResponse) GetServices() []*ServiceVersion {
	if m != nil {
		return m.Services
	}
	return nil
}

type ServiceVersion struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Origin               string   `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Release              string   `protobuf:"bytes,4,opt,name=release,proto3" json:"release,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceVersion) Reset()         { *m = ServiceVersion{} }
func (m *ServiceVersion) String() string { return proto.CompactTextString(m) }
func (*ServiceVersion) ProtoMessage()    {}
func (*ServiceVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4362896e7bedc82, []int{3}
}

func (m *ServiceVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceVersion.Unmarshal(m, b)
}
func (m *ServiceVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceVersion.Marshal(b, m, deterministic)
}
func (m *ServiceVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceVersion.Merge(m, src)
}
func (m *ServiceVersion) XXX_Size() int {
	return xxx_messageInfo_ServiceVersion.Size(m)
}
func (m *ServiceVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceVersion.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceVersion proto.InternalMessageInfo

func (m *ServiceVersion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceVersion) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *ServiceVersion) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServiceVersion) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func init() {
	proto.RegisterType((*Version)(nil), "chef.automate.api.deployment.Version")
	proto.RegisterType((*ServiceVersionsRequest)(nil), "chef.automate.api.deployment.ServiceVersionsRequest")
	proto.RegisterType((*ServiceVersionsResponse)(nil), "chef.automate.api.deployment.ServiceVersionsResponse")
	proto.RegisterType((*ServiceVersion)(nil), "chef.automate.api.deployment.ServiceVersion")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/deployment/deployment.proto", fileDescriptor_e4362896e7bedc82)
}

var fileDescriptor_e4362896e7bedc82 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x3d, 0x6f, 0xd4, 0x40,
	0x10, 0x95, 0x73, 0x51, 0x12, 0x06, 0x29, 0x91, 0xb6, 0x38, 0x16, 0xe7, 0x84, 0x22, 0x4b, 0x88,
	0x14, 0xdc, 0x5a, 0x3a, 0xa0, 0x39, 0x1a, 0x04, 0x44, 0x20, 0xd1, 0x1d, 0x88, 0x82, 0x26, 0xac,
	0x9d, 0x89, 0xb3, 0x92, 0xf7, 0x03, 0xef, 0x38, 0xc8, 0x2d, 0xe5, 0xb5, 0x34, 0xfc, 0x91, 0xfc,
	0x12, 0xfe, 0x00, 0x05, 0x3f, 0x04, 0x9d, 0xbd, 0x3e, 0x72, 0x24, 0x0a, 0xb9, 0x6e, 0x67, 0xde,
	0xbc, 0xf1, 0x7b, 0xfb, 0xbc, 0xf0, 0x3c, 0xb7, 0xda, 0x59, 0x83, 0x86, 0x7c, 0x2a, 0x6b, 0xb2,
	0x5a, 0x12, 0x8e, 0x0b, 0x49, 0xf8, 0x55, 0x36, 0xa9, 0x74, 0x2a, 0x3d, 0x41, 0x57, 0xda, 0x46,
	0xa3, 0xa1, 0x4b, 0x47, 0xe1, 0x2a, 0x4b, 0x96, 0x8d, 0xf2, 0x33, 0x3c, 0x15, 0x3d, 0x4d, 0x48,
	0xa7, 0xc4, 0xdf, 0x99, 0x78, 0x54, 0x58, 0x5b, 0x94, 0xd8, 0x6e, 0x91, 0xc6, 0x58, 0x92, 0xa4,
	0xac, 0xf1, 0x1d, 0x37, 0xde, 0x0f, 0x68, 0x5b, 0x65, 0xf5, 0x69, 0x8a, 0xda, 0x51, 0x13, 0xc0,
	0x17, 0xd7, 0xaa, 0xaa, 0x5c, 0xde, 0x8d, 0xe7, 0xe3, 0x02, 0xcd, 0xd8, 0xd9, 0x52, 0xe5, 0x4d,
	0xaa, 0xa4, 0xbe, 0xba, 0x3e, 0x99, 0xc0, 0xf6, 0x47, 0xac, 0xbc, 0xb2, 0x86, 0x3d, 0x82, 0xbd,
	0xac, 0x56, 0xe5, 0xc9, 0x31, 0x29, 0x8d, 0x9e, 0xa4, 0x76, 0x3c, 0x3a, 0x88, 0x0e, 0xef, 0xcc,
	0x76, 0xdb, 0xf6, 0x87, 0xbe, 0x9b, 0x70, 0x18, 0xbe, 0xc7, 0xea, 0x5c, 0xe5, 0x18, 0xa8, 0x7e,
	0x86, 0x5f, 0x6a, 0xf4, 0x94, 0xe4, 0x70, 0xef, 0x0a, 0xe2, 0x9d, 0x35, 0x1e, 0xd9, 0x5b, 0xd8,
	0xf1, 0x1d, 0xe4, 0x79, 0x74, 0x30, 0x38, 0xbc, 0x3b, 0x79, 0x2c, 0x6e, 0xba, 0x16, 0xb1, 0xba,
	0x68, 0xb6, 0x64, 0x27, 0x0e, 0x76, 0x57, 0x31, 0xc6, 0x60, 0xd3, 0x48, 0x8d, 0x41, 0x6e, 0x7b,
	0x66, 0x43, 0xd8, 0xb2, 0x95, 0x2a, 0x94, 0xe1, 0x1b, 0x6d, 0x37, 0x54, 0x8c, 0xc3, 0xf6, 0x79,
	0x47, 0xe3, 0x83, 0x16, 0xe8, 0xcb, 0x05, 0x52, 0x61, 0x89, 0xd2, 0x23, 0xdf, 0xec, 0x90, 0x50,
	0x4e, 0xe6, 0x03, 0x80, 0xd7, 0x4b, 0x65, 0xec, 0x47, 0x04, 0xf0, 0x06, 0xa9, 0xff, 0xfa, 0x50,
	0x74, 0x11, 0x89, 0x3e, 0x22, 0x71, 0xb4, 0x88, 0x28, 0x7e, 0x78, 0xb3, 0xbf, 0x40, 0x4f, 0xde,
	0xcd, 0x2f, 0x38, 0x87, 0xa1, 0x6f, 0x3c, 0xa1, 0x9e, 0x06, 0x97, 0xd3, 0x20, 0x6a, 0x7e, 0xc1,
	0xf7, 0xd9, 0xfd, 0x55, 0x2c, 0x10, 0xa7, 0x05, 0xd2, 0xb7, 0x9f, 0xbf, 0xbf, 0x6f, 0x00, 0xdb,
	0x49, 0x7b, 0x0f, 0xbf, 0x22, 0xd8, 0xfb, 0x27, 0x01, 0xf6, 0x74, 0x9d, 0x7b, 0xee, 0xa3, 0x8c,
	0x9f, 0xad, 0xc9, 0xea, 0x62, 0x4e, 0x3e, 0xff, 0xc7, 0xcd, 0x88, 0xc5, 0xd7, 0xbb, 0x29, 0x95,
	0xef, 0xec, 0x3c, 0x60, 0xa3, 0xcb, 0x8f, 0x29, 0x8c, 0x1d, 0x87, 0x15, 0xfe, 0xe5, 0xd1, 0xa7,
	0x57, 0x85, 0xa2, 0xb3, 0x3a, 0x13, 0xb9, 0xd5, 0xe9, 0x42, 0xe4, 0xf2, 0xd7, 0x4f, 0x6f, 0xff,
	0x48, 0xb3, 0xad, 0x36, 0xad, 0x27, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x9b, 0xb3, 0x17,
	0xd9, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeploymentClient is the client API for Deployment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeploymentClient interface {
	GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Version, error)
	ServiceVersions(ctx context.Context, in *ServiceVersionsRequest, opts ...grpc.CallOption) (*ServiceVersionsResponse, error)
}

type deploymentClient struct {
	cc grpc.ClientConnInterface
}

func NewDeploymentClient(cc grpc.ClientConnInterface) DeploymentClient {
	return &deploymentClient{cc}
}

func (c *deploymentClient) GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/chef.automate.api.deployment.Deployment/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentClient) ServiceVersions(ctx context.Context, in *ServiceVersionsRequest, opts ...grpc.CallOption) (*ServiceVersionsResponse, error) {
	out := new(ServiceVersionsResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.deployment.Deployment/ServiceVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeploymentServer is the server API for Deployment service.
type DeploymentServer interface {
	GetVersion(context.Context, *empty.Empty) (*Version, error)
	ServiceVersions(context.Context, *ServiceVersionsRequest) (*ServiceVersionsResponse, error)
}

// UnimplementedDeploymentServer can be embedded to have forward compatible implementations.
type UnimplementedDeploymentServer struct {
}

func (*UnimplementedDeploymentServer) GetVersion(ctx context.Context, req *empty.Empty) (*Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (*UnimplementedDeploymentServer) ServiceVersions(ctx context.Context, req *ServiceVersionsRequest) (*ServiceVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceVersions not implemented")
}

func RegisterDeploymentServer(s *grpc.Server, srv DeploymentServer) {
	s.RegisterService(&_Deployment_serviceDesc, srv)
}

func _Deployment_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.deployment.Deployment/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServer).GetVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployment_ServiceVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServer).ServiceVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.deployment.Deployment/ServiceVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServer).ServiceVersions(ctx, req.(*ServiceVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Deployment_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.deployment.Deployment",
	HandlerType: (*DeploymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Deployment_GetVersion_Handler,
		},
		{
			MethodName: "ServiceVersions",
			Handler:    _Deployment_ServiceVersions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/deployment/deployment.proto",
}
