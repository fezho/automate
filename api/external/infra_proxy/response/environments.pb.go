// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/infra_proxy/response/environments.proto

package response

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

type Environments struct {
	// Environments list.
	Environments         []*EnvironmentListItem `protobuf:"bytes,1,rep,name=environments,proto3" json:"environments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Environments) Reset()         { *m = Environments{} }
func (m *Environments) String() string { return proto.CompactTextString(m) }
func (*Environments) ProtoMessage()    {}
func (*Environments) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6cb2a4fff0e747c, []int{0}
}

func (m *Environments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Environments.Unmarshal(m, b)
}
func (m *Environments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Environments.Marshal(b, m, deterministic)
}
func (m *Environments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Environments.Merge(m, src)
}
func (m *Environments) XXX_Size() int {
	return xxx_messageInfo_Environments.Size(m)
}
func (m *Environments) XXX_DiscardUnknown() {
	xxx_messageInfo_Environments.DiscardUnknown(m)
}

var xxx_messageInfo_Environments proto.InternalMessageInfo

func (m *Environments) GetEnvironments() []*EnvironmentListItem {
	if m != nil {
		return m.Environments
	}
	return nil
}

type EnvironmentListItem struct {
	// Environment name.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvironmentListItem) Reset()         { *m = EnvironmentListItem{} }
func (m *EnvironmentListItem) String() string { return proto.CompactTextString(m) }
func (*EnvironmentListItem) ProtoMessage()    {}
func (*EnvironmentListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6cb2a4fff0e747c, []int{1}
}

func (m *EnvironmentListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvironmentListItem.Unmarshal(m, b)
}
func (m *EnvironmentListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvironmentListItem.Marshal(b, m, deterministic)
}
func (m *EnvironmentListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvironmentListItem.Merge(m, src)
}
func (m *EnvironmentListItem) XXX_Size() int {
	return xxx_messageInfo_EnvironmentListItem.Size(m)
}
func (m *EnvironmentListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvironmentListItem.DiscardUnknown(m)
}

var xxx_messageInfo_EnvironmentListItem proto.InternalMessageInfo

func (m *EnvironmentListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Environment struct {
	// Environment name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Object type.
	ChefType string `protobuf:"bytes,2,opt,name=chef_type,json=chefType,proto3" json:"chef_type,omitempty"`
	// Node description.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Class name.
	JsonClass string `protobuf:"bytes,5,opt,name=json_class,json=jsonClass,proto3" json:"json_class,omitempty"`
	// Versioned cookbooks list.
	CookbookVersions []string `protobuf:"bytes,6,rep,name=cookbook_versions,json=cookbookVersions,proto3" json:"cookbook_versions,omitempty"`
	// Stringified json of the default attributes.
	DefaultAttributes string `protobuf:"bytes,7,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty"`
	// Stringified json of the override attributes.
	OverrideAttributes   string   `protobuf:"bytes,8,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Environment) Reset()         { *m = Environment{} }
func (m *Environment) String() string { return proto.CompactTextString(m) }
func (*Environment) ProtoMessage()    {}
func (*Environment) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6cb2a4fff0e747c, []int{2}
}

func (m *Environment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Environment.Unmarshal(m, b)
}
func (m *Environment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Environment.Marshal(b, m, deterministic)
}
func (m *Environment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Environment.Merge(m, src)
}
func (m *Environment) XXX_Size() int {
	return xxx_messageInfo_Environment.Size(m)
}
func (m *Environment) XXX_DiscardUnknown() {
	xxx_messageInfo_Environment.DiscardUnknown(m)
}

var xxx_messageInfo_Environment proto.InternalMessageInfo

func (m *Environment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Environment) GetChefType() string {
	if m != nil {
		return m.ChefType
	}
	return ""
}

func (m *Environment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Environment) GetJsonClass() string {
	if m != nil {
		return m.JsonClass
	}
	return ""
}

func (m *Environment) GetCookbookVersions() []string {
	if m != nil {
		return m.CookbookVersions
	}
	return nil
}

func (m *Environment) GetDefaultAttributes() string {
	if m != nil {
		return m.DefaultAttributes
	}
	return ""
}

func (m *Environment) GetOverrideAttributes() string {
	if m != nil {
		return m.OverrideAttributes
	}
	return ""
}

func init() {
	proto.RegisterType((*Environments)(nil), "chef.automate.api.infra_proxy.response.Environments")
	proto.RegisterType((*EnvironmentListItem)(nil), "chef.automate.api.infra_proxy.response.EnvironmentListItem")
	proto.RegisterType((*Environment)(nil), "chef.automate.api.infra_proxy.response.Environment")
}

func init() {
	proto.RegisterFile("api/external/infra_proxy/response/environments.proto", fileDescriptor_b6cb2a4fff0e747c)
}

var fileDescriptor_b6cb2a4fff0e747c = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xeb, 0x40,
	0x14, 0x85, 0x69, 0xfb, 0x5e, 0x5f, 0x33, 0xed, 0xe2, 0x75, 0xba, 0x09, 0x88, 0x10, 0xba, 0x90,
	0x8a, 0x38, 0x03, 0xea, 0x4a, 0xdd, 0xa8, 0xb8, 0x10, 0x5c, 0x15, 0x71, 0xe1, 0x26, 0x4c, 0xd2,
	0x5b, 0x3b, 0xb6, 0x99, 0x3b, 0xcc, 0xdc, 0x94, 0xf6, 0x3f, 0xf8, 0xa3, 0x25, 0xa1, 0xd1, 0x11,
	0x0a, 0xba, 0x0b, 0xe7, 0xfb, 0xce, 0x21, 0x21, 0x97, 0x5d, 0x28, 0xab, 0x25, 0x6c, 0x08, 0x9c,
	0x51, 0x2b, 0xa9, 0xcd, 0xdc, 0xa9, 0xd4, 0x3a, 0xdc, 0x6c, 0xa5, 0x03, 0x6f, 0xd1, 0x78, 0x90,
	0x60, 0xd6, 0xda, 0xa1, 0x29, 0xc0, 0x90, 0x17, 0xd6, 0x21, 0x21, 0x3f, 0xca, 0x17, 0x30, 0x17,
	0xaa, 0x24, 0x2c, 0x14, 0x81, 0x50, 0x56, 0x8b, 0xa0, 0x2a, 0x9a, 0xea, 0x18, 0xd9, 0xe0, 0x3e,
	0x68, 0xf3, 0x94, 0x0d, 0xc2, 0xb5, 0xb8, 0x95, 0x74, 0x26, 0xfd, 0xb3, 0x2b, 0xf1, 0xbb, 0x39,
	0x11, 0x6c, 0x3d, 0x6a, 0x4f, 0x0f, 0x04, 0xc5, 0xf4, 0xdb, 0xe0, 0xf8, 0x98, 0x8d, 0xf6, 0x48,
	0x9c, 0xb3, 0x3f, 0x46, 0x15, 0x10, 0xb7, 0x92, 0xd6, 0x24, 0x9a, 0xd6, 0xcf, 0xe3, 0xf7, 0x36,
	0xeb, 0x07, 0xee, 0x3e, 0x87, 0x1f, 0xb0, 0xa8, 0x7a, 0xb5, 0x94, 0xb6, 0x16, 0xe2, 0x76, 0x0d,
	0x7a, 0x55, 0xf0, 0xb4, 0xb5, 0xc0, 0x13, 0xd6, 0x9f, 0x81, 0xcf, 0x9d, 0xb6, 0xa4, 0xd1, 0xc4,
	0x9d, 0x1a, 0x87, 0x11, 0x3f, 0x64, 0xec, 0xcd, 0xa3, 0x49, 0xf3, 0x95, 0xf2, 0x3e, 0xfe, 0x5b,
	0x0b, 0x51, 0x95, 0xdc, 0x55, 0x01, 0x3f, 0x61, 0xc3, 0x1c, 0x71, 0x99, 0x21, 0x2e, 0xd3, 0x35,
	0x38, 0xaf, 0xd1, 0xf8, 0xb8, 0x9b, 0x74, 0x26, 0xd1, 0xf4, 0x7f, 0x03, 0x9e, 0x77, 0x39, 0x3f,
	0x65, 0x7c, 0x06, 0x73, 0x55, 0xae, 0x28, 0x55, 0x44, 0x4e, 0x67, 0x25, 0x81, 0x8f, 0xff, 0xd5,
	0x9b, 0xc3, 0x1d, 0xb9, 0xf9, 0x04, 0x5c, 0xb2, 0x11, 0xae, 0xc1, 0x39, 0x3d, 0x83, 0xd0, 0xef,
	0xd5, 0x3e, 0x6f, 0xd0, 0x57, 0xe1, 0xf6, 0xfa, 0xe5, 0xf2, 0x55, 0xd3, 0xa2, 0xcc, 0x44, 0x8e,
	0x85, 0xac, 0x3e, 0x52, 0x36, 0x3f, 0x44, 0xfe, 0x78, 0x23, 0x59, 0xb7, 0xbe, 0x8b, 0xf3, 0x8f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x5b, 0xb3, 0xc8, 0x4f, 0x02, 0x00, 0x00,
}
