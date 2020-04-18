// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/infra_proxy/request/servers.proto

package request

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CreateServer struct {
	// Chef infra server ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Chef infra server name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Chef infra server FQDN.
	Fqdn string `protobuf:"bytes,3,opt,name=fqdn,proto3" json:"fqdn,omitempty" toml:"fqdn,omitempty" mapstructure:"fqdn,omitempty"`
	// Chef infra server IP address.
	IpAddress            string   `protobuf:"bytes,4,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty" toml:"ip_address,omitempty" mapstructure:"ip_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *CreateServer) Reset()         { *m = CreateServer{} }
func (m *CreateServer) String() string { return proto.CompactTextString(m) }
func (*CreateServer) ProtoMessage()    {}
func (*CreateServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_64390fbabed0d9db, []int{0}
}

func (m *CreateServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServer.Unmarshal(m, b)
}
func (m *CreateServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServer.Marshal(b, m, deterministic)
}
func (m *CreateServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServer.Merge(m, src)
}
func (m *CreateServer) XXX_Size() int {
	return xxx_messageInfo_CreateServer.Size(m)
}
func (m *CreateServer) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServer.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServer proto.InternalMessageInfo

func (m *CreateServer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateServer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateServer) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *CreateServer) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

type UpdateServer struct {
	// Chef infra server ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Chef infra server name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Chef infra server FQDN.
	Fqdn string `protobuf:"bytes,3,opt,name=fqdn,proto3" json:"fqdn,omitempty" toml:"fqdn,omitempty" mapstructure:"fqdn,omitempty"`
	// Chef infra server IP address.
	IpAddress            string   `protobuf:"bytes,4,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty" toml:"ip_address,omitempty" mapstructure:"ip_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *UpdateServer) Reset()         { *m = UpdateServer{} }
func (m *UpdateServer) String() string { return proto.CompactTextString(m) }
func (*UpdateServer) ProtoMessage()    {}
func (*UpdateServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_64390fbabed0d9db, []int{1}
}

func (m *UpdateServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateServer.Unmarshal(m, b)
}
func (m *UpdateServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateServer.Marshal(b, m, deterministic)
}
func (m *UpdateServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateServer.Merge(m, src)
}
func (m *UpdateServer) XXX_Size() int {
	return xxx_messageInfo_UpdateServer.Size(m)
}
func (m *UpdateServer) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateServer.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateServer proto.InternalMessageInfo

func (m *UpdateServer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateServer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateServer) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *UpdateServer) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

type DeleteServer struct {
	// Chef infra server ID.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DeleteServer) Reset()         { *m = DeleteServer{} }
func (m *DeleteServer) String() string { return proto.CompactTextString(m) }
func (*DeleteServer) ProtoMessage()    {}
func (*DeleteServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_64390fbabed0d9db, []int{2}
}

func (m *DeleteServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteServer.Unmarshal(m, b)
}
func (m *DeleteServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteServer.Marshal(b, m, deterministic)
}
func (m *DeleteServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteServer.Merge(m, src)
}
func (m *DeleteServer) XXX_Size() int {
	return xxx_messageInfo_DeleteServer.Size(m)
}
func (m *DeleteServer) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteServer.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteServer proto.InternalMessageInfo

func (m *DeleteServer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetServers struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *GetServers) Reset()         { *m = GetServers{} }
func (m *GetServers) String() string { return proto.CompactTextString(m) }
func (*GetServers) ProtoMessage()    {}
func (*GetServers) Descriptor() ([]byte, []int) {
	return fileDescriptor_64390fbabed0d9db, []int{3}
}

func (m *GetServers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServers.Unmarshal(m, b)
}
func (m *GetServers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServers.Marshal(b, m, deterministic)
}
func (m *GetServers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServers.Merge(m, src)
}
func (m *GetServers) XXX_Size() int {
	return xxx_messageInfo_GetServers.Size(m)
}
func (m *GetServers) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServers.DiscardUnknown(m)
}

var xxx_messageInfo_GetServers proto.InternalMessageInfo

type GetServer struct {
	// Chef infra server ID.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *GetServer) Reset()         { *m = GetServer{} }
func (m *GetServer) String() string { return proto.CompactTextString(m) }
func (*GetServer) ProtoMessage()    {}
func (*GetServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_64390fbabed0d9db, []int{4}
}

func (m *GetServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServer.Unmarshal(m, b)
}
func (m *GetServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServer.Marshal(b, m, deterministic)
}
func (m *GetServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServer.Merge(m, src)
}
func (m *GetServer) XXX_Size() int {
	return xxx_messageInfo_GetServer.Size(m)
}
func (m *GetServer) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServer.DiscardUnknown(m)
}

var xxx_messageInfo_GetServer proto.InternalMessageInfo

func (m *GetServer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateServer)(nil), "chef.automate.domain.infra_proxy.request.CreateServer")
	proto.RegisterType((*UpdateServer)(nil), "chef.automate.domain.infra_proxy.request.UpdateServer")
	proto.RegisterType((*DeleteServer)(nil), "chef.automate.domain.infra_proxy.request.DeleteServer")
	proto.RegisterType((*GetServers)(nil), "chef.automate.domain.infra_proxy.request.GetServers")
	proto.RegisterType((*GetServer)(nil), "chef.automate.domain.infra_proxy.request.GetServer")
}

func init() {
	proto.RegisterFile("api/interservice/infra_proxy/request/servers.proto", fileDescriptor_64390fbabed0d9db)
}

var fileDescriptor_64390fbabed0d9db = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0x3f, 0x4b, 0x04, 0x31,
	0x10, 0xc5, 0xb9, 0xf5, 0x10, 0x76, 0x58, 0x2c, 0x52, 0x2d, 0x88, 0x22, 0x5b, 0x5d, 0x95, 0x80,
	0xd6, 0x22, 0xfe, 0x01, 0x7b, 0xc5, 0xc6, 0xe6, 0xc8, 0x6d, 0x66, 0xbd, 0x01, 0x37, 0xc9, 0x4d,
	0x66, 0x45, 0xbf, 0xbd, 0x24, 0xab, 0x22, 0x82, 0x60, 0x63, 0x37, 0xf9, 0xbd, 0xcc, 0x3c, 0x78,
	0x0f, 0x4e, 0x6d, 0x24, 0x43, 0x5e, 0x90, 0x13, 0xf2, 0x0b, 0xf5, 0x68, 0xc8, 0x0f, 0x6c, 0xd7,
	0x91, 0xc3, 0xeb, 0x9b, 0x61, 0xdc, 0x4d, 0x98, 0xc4, 0x64, 0x0d, 0x39, 0xe9, 0xc8, 0x41, 0x82,
	0x5a, 0xf5, 0x5b, 0x1c, 0xb4, 0x9d, 0x24, 0x8c, 0x56, 0x50, 0xbb, 0x30, 0x5a, 0xf2, 0xfa, 0xdb,
	0x9e, 0xfe, 0xd8, 0xeb, 0x10, 0x9a, 0x6b, 0x46, 0x2b, 0x78, 0x5f, 0x0e, 0xa8, 0x03, 0xa8, 0xc8,
	0xb5, 0x8b, 0x93, 0xc5, 0xaa, 0xbe, 0xab, 0xc8, 0x29, 0x05, 0x4b, 0x6f, 0x47, 0x6c, 0xab, 0x42,
	0xca, 0x9c, 0xd9, 0xb0, 0x73, 0xbe, 0xdd, 0x9b, 0x59, 0x9e, 0xd5, 0x11, 0x00, 0xc5, 0xb5, 0x75,
	0x8e, 0x31, 0xa5, 0x76, 0x59, 0x94, 0x9a, 0xe2, 0xe5, 0x0c, 0xb2, 0xcd, 0x43, 0x74, 0xff, 0x6e,
	0x73, 0x0c, 0xcd, 0x0d, 0x3e, 0xe3, 0x6f, 0x36, 0x5d, 0x03, 0x70, 0x8b, 0x32, 0x8b, 0xa9, 0x3b,
	0x84, 0xfa, 0xeb, 0xf5, 0xf3, 0xeb, 0xd5, 0xc5, 0xe3, 0xf9, 0x13, 0xc9, 0x76, 0xda, 0xe8, 0x3e,
	0x8c, 0x26, 0xe7, 0x69, 0x3e, 0xf3, 0x34, 0x7f, 0x69, 0x64, 0xb3, 0x5f, 0xaa, 0x38, 0x7b, 0x0f,
	0x00, 0x00, 0xff, 0xff, 0x4b, 0x44, 0xdb, 0x2f, 0xc0, 0x01, 0x00, 0x00,
}
