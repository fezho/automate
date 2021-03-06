// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/authz/common/subject_purge.proto

package common

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type PurgeSubjectFromPoliciesReq struct {
	// Q: Right now, this allows purging subject wildcards (like "user:*").
	// -- Do we want to restrict this to only complete subjects?
	Subject              string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty" toml:"subject,omitempty" mapstructure:"subject,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *PurgeSubjectFromPoliciesReq) Reset()         { *m = PurgeSubjectFromPoliciesReq{} }
func (m *PurgeSubjectFromPoliciesReq) String() string { return proto.CompactTextString(m) }
func (*PurgeSubjectFromPoliciesReq) ProtoMessage()    {}
func (*PurgeSubjectFromPoliciesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_78588a1a247c8462, []int{0}
}

func (m *PurgeSubjectFromPoliciesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurgeSubjectFromPoliciesReq.Unmarshal(m, b)
}
func (m *PurgeSubjectFromPoliciesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurgeSubjectFromPoliciesReq.Marshal(b, m, deterministic)
}
func (m *PurgeSubjectFromPoliciesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurgeSubjectFromPoliciesReq.Merge(m, src)
}
func (m *PurgeSubjectFromPoliciesReq) XXX_Size() int {
	return xxx_messageInfo_PurgeSubjectFromPoliciesReq.Size(m)
}
func (m *PurgeSubjectFromPoliciesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PurgeSubjectFromPoliciesReq.DiscardUnknown(m)
}

var xxx_messageInfo_PurgeSubjectFromPoliciesReq proto.InternalMessageInfo

func (m *PurgeSubjectFromPoliciesReq) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

// PurgeSubjectFromPolicies() removes the passed subject from every policy,
// IAM v1 _and_ IAM v2. We thus differentiate the returned "affected" arrays.
type PurgeSubjectFromPoliciesResp struct {
	PoliciesV1           []string `protobuf:"bytes,1,rep,name=policies_v1,json=policiesV1,proto3" json:"policies_v1,omitempty" toml:"policies_v1,omitempty" mapstructure:"policies_v1,omitempty"` // Deprecated: Do not use.
	PoliciesV2           []string `protobuf:"bytes,2,rep,name=policies_v2,json=policiesV2,proto3" json:"policies_v2,omitempty" toml:"policies_v2,omitempty" mapstructure:"policies_v2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *PurgeSubjectFromPoliciesResp) Reset()         { *m = PurgeSubjectFromPoliciesResp{} }
func (m *PurgeSubjectFromPoliciesResp) String() string { return proto.CompactTextString(m) }
func (*PurgeSubjectFromPoliciesResp) ProtoMessage()    {}
func (*PurgeSubjectFromPoliciesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_78588a1a247c8462, []int{1}
}

func (m *PurgeSubjectFromPoliciesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurgeSubjectFromPoliciesResp.Unmarshal(m, b)
}
func (m *PurgeSubjectFromPoliciesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurgeSubjectFromPoliciesResp.Marshal(b, m, deterministic)
}
func (m *PurgeSubjectFromPoliciesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurgeSubjectFromPoliciesResp.Merge(m, src)
}
func (m *PurgeSubjectFromPoliciesResp) XXX_Size() int {
	return xxx_messageInfo_PurgeSubjectFromPoliciesResp.Size(m)
}
func (m *PurgeSubjectFromPoliciesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PurgeSubjectFromPoliciesResp.DiscardUnknown(m)
}

var xxx_messageInfo_PurgeSubjectFromPoliciesResp proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *PurgeSubjectFromPoliciesResp) GetPoliciesV1() []string {
	if m != nil {
		return m.PoliciesV1
	}
	return nil
}

func (m *PurgeSubjectFromPoliciesResp) GetPoliciesV2() []string {
	if m != nil {
		return m.PoliciesV2
	}
	return nil
}

func init() {
	proto.RegisterType((*PurgeSubjectFromPoliciesReq)(nil), "chef.automate.domain.authz.common.PurgeSubjectFromPoliciesReq")
	proto.RegisterType((*PurgeSubjectFromPoliciesResp)(nil), "chef.automate.domain.authz.common.PurgeSubjectFromPoliciesResp")
}

func init() {
	proto.RegisterFile("api/interservice/authz/common/subject_purge.proto", fileDescriptor_78588a1a247c8462)
}

var fileDescriptor_78588a1a247c8462 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xff, 0x69, 0xe1, 0x97, 0x8e, 0xae, 0xb2, 0x31, 0x54, 0xc1, 0x5a, 0xbb, 0x48, 0x2a,
	0x64, 0x68, 0x04, 0x91, 0x2c, 0x0c, 0x74, 0xe1, 0xba, 0x54, 0x70, 0x61, 0x69, 0xcb, 0x34, 0x19,
	0xdb, 0xd1, 0x24, 0x33, 0xce, 0x4c, 0xba, 0x90, 0x79, 0x09, 0x41, 0x70, 0xd5, 0x57, 0x70, 0x2d,
	0xae, 0x7c, 0x0c, 0x5f, 0xc1, 0xb7, 0x90, 0xa4, 0x09, 0x16, 0xa1, 0x75, 0xe1, 0x6e, 0xee, 0xe5,
	0x7c, 0xe7, 0xcc, 0xe5, 0x5e, 0xd8, 0xc1, 0x9c, 0x22, 0x9a, 0x28, 0x22, 0x24, 0x11, 0x73, 0x1a,
	0x10, 0x84, 0x53, 0x35, 0x7b, 0x40, 0x01, 0x8b, 0x63, 0x96, 0x20, 0x99, 0x4e, 0x6e, 0x49, 0xa0,
	0xc6, 0x3c, 0x15, 0x53, 0xe2, 0x70, 0xc1, 0x14, 0x33, 0x0e, 0x83, 0x19, 0xb9, 0x71, 0x70, 0xaa,
	0x58, 0x8c, 0x15, 0x71, 0x42, 0x16, 0x63, 0x9a, 0x38, 0x39, 0xe6, 0x2c, 0xb1, 0xfa, 0xee, 0x1c,
	0x47, 0x34, 0xc4, 0x8a, 0xa0, 0xf2, 0xb1, 0x64, 0x9b, 0x1f, 0x00, 0xee, 0xf5, 0x32, 0xaf, 0xcb,
	0xa5, 0xf1, 0x85, 0x60, 0x71, 0x8f, 0x45, 0x34, 0xa0, 0x44, 0xf6, 0xc9, 0xbd, 0xf1, 0x0a, 0xe0,
	0x56, 0x91, 0x69, 0x82, 0x06, 0xb0, 0x6a, 0xdd, 0x05, 0x78, 0xfb, 0x7c, 0xaf, 0x3e, 0x03, 0xf1,
	0x04, 0xdc, 0x47, 0x30, 0xb2, 0x7c, 0x4f, 0x11, 0x1c, 0xeb, 0x54, 0x12, 0x61, 0x7b, 0x96, 0xef,
	0x45, 0x2c, 0xc0, 0x91, 0x8e, 0x42, 0xcc, 0xb5, 0xc4, 0x71, 0x94, 0xf7, 0x06, 0x23, 0xaf, 0x3d,
	0x3c, 0xd6, 0x83, 0xf6, 0xd0, 0x6e, 0xe9, 0x0c, 0x59, 0xa5, 0xb4, 0x62, 0x77, 0x24, 0xd1, 0xc5,
	0xb8, 0xb6, 0x67, 0xfb, 0x83, 0xf6, 0xb0, 0xa5, 0x47, 0x79, 0xbb, 0x40, 0xb3, 0x32, 0x92, 0x5e,
	0xa1, 0xf9, 0xb6, 0xf4, 0x6c, 0xff, 0x87, 0x7d, 0xbf, 0xfc, 0x6e, 0x33, 0x84, 0xfb, 0xeb, 0x27,
	0x93, 0xdc, 0x38, 0x82, 0xdb, 0xbc, 0xa8, 0xc7, 0xf3, 0x8e, 0x09, 0x1a, 0x55, 0xab, 0xd6, 0xad,
	0x98, 0xa0, 0x0f, 0xcb, 0xf6, 0x55, 0xc7, 0x38, 0x58, 0x15, 0xb9, 0x66, 0x25, 0x13, 0xad, 0x08,
	0x5c, 0xf7, 0x05, 0xc0, 0x9d, 0x22, 0x21, 0x4f, 0x33, 0x16, 0x00, 0x9a, 0xeb, 0x72, 0x8d, 0x73,
	0xe7, 0xd7, 0x5d, 0x39, 0x1b, 0xd6, 0x51, 0xf7, 0xff, 0xc4, 0x4b, 0xde, 0xfc, 0xd7, 0x3d, 0xbb,
	0x3e, 0x9d, 0x52, 0x35, 0x4b, 0x27, 0x99, 0x1e, 0x65, 0x76, 0xa8, 0xb4, 0x43, 0x1b, 0x6f, 0x6f,
	0xf2, 0x3f, 0x3f, 0x99, 0x93, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x3d, 0x4e, 0x7f, 0xa3,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SubjectPurgeClient is the client API for SubjectPurge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubjectPurgeClient interface {
	PurgeSubjectFromPolicies(ctx context.Context, in *PurgeSubjectFromPoliciesReq, opts ...grpc.CallOption) (*PurgeSubjectFromPoliciesResp, error)
}

type subjectPurgeClient struct {
	cc grpc.ClientConnInterface
}

func NewSubjectPurgeClient(cc grpc.ClientConnInterface) SubjectPurgeClient {
	return &subjectPurgeClient{cc}
}

func (c *subjectPurgeClient) PurgeSubjectFromPolicies(ctx context.Context, in *PurgeSubjectFromPoliciesReq, opts ...grpc.CallOption) (*PurgeSubjectFromPoliciesResp, error) {
	out := new(PurgeSubjectFromPoliciesResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authz.common.SubjectPurge/PurgeSubjectFromPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubjectPurgeServer is the server API for SubjectPurge service.
type SubjectPurgeServer interface {
	PurgeSubjectFromPolicies(context.Context, *PurgeSubjectFromPoliciesReq) (*PurgeSubjectFromPoliciesResp, error)
}

// UnimplementedSubjectPurgeServer can be embedded to have forward compatible implementations.
type UnimplementedSubjectPurgeServer struct {
}

func (*UnimplementedSubjectPurgeServer) PurgeSubjectFromPolicies(ctx context.Context, req *PurgeSubjectFromPoliciesReq) (*PurgeSubjectFromPoliciesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurgeSubjectFromPolicies not implemented")
}

func RegisterSubjectPurgeServer(s *grpc.Server, srv SubjectPurgeServer) {
	s.RegisterService(&_SubjectPurge_serviceDesc, srv)
}

func _SubjectPurge_PurgeSubjectFromPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeSubjectFromPoliciesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectPurgeServer).PurgeSubjectFromPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authz.common.SubjectPurge/PurgeSubjectFromPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectPurgeServer).PurgeSubjectFromPolicies(ctx, req.(*PurgeSubjectFromPoliciesReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SubjectPurge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.authz.common.SubjectPurge",
	HandlerType: (*SubjectPurgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurgeSubjectFromPolicies",
			Handler:    _SubjectPurge_PurgeSubjectFromPolicies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interservice/authz/common/subject_purge.proto",
}
