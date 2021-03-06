// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/cfgmgmt/response/root.proto

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

type VersionInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty" toml:"version,omitempty" mapstructure:"version,omitempty"`
	SHA                  string   `protobuf:"bytes,3,opt,name=SHA,proto3" json:"SHA,omitempty" toml:"SHA,omitempty" mapstructure:"SHA,omitempty"`
	Built                string   `protobuf:"bytes,4,opt,name=built,proto3" json:"built,omitempty" toml:"built,omitempty" mapstructure:"built,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *VersionInfo) Reset()         { *m = VersionInfo{} }
func (m *VersionInfo) String() string { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()    {}
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b42b0f1f6d7b8d4, []int{0}
}

func (m *VersionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionInfo.Unmarshal(m, b)
}
func (m *VersionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionInfo.Marshal(b, m, deterministic)
}
func (m *VersionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionInfo.Merge(m, src)
}
func (m *VersionInfo) XXX_Size() int {
	return xxx_messageInfo_VersionInfo.Size(m)
}
func (m *VersionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VersionInfo proto.InternalMessageInfo

func (m *VersionInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VersionInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionInfo) GetSHA() string {
	if m != nil {
		return m.SHA
	}
	return ""
}

func (m *VersionInfo) GetBuilt() string {
	if m != nil {
		return m.Built
	}
	return ""
}

// Health message
//
// The config-mgmt-service health is constructed with:
// * Status:
//            => ok:             Everything is alright
//            => initialization: The service is in its initialization process
//            => warning:        Something might be wrong?
//            => critical:       Something is wrong!
//
// @afiune: Here we can add more health information to the response
type Health struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" toml:"status,omitempty" mapstructure:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Health) Reset()         { *m = Health{} }
func (m *Health) String() string { return proto.CompactTextString(m) }
func (*Health) ProtoMessage()    {}
func (*Health) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b42b0f1f6d7b8d4, []int{1}
}

func (m *Health) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Health.Unmarshal(m, b)
}
func (m *Health) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Health.Marshal(b, m, deterministic)
}
func (m *Health) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Health.Merge(m, src)
}
func (m *Health) XXX_Size() int {
	return xxx_messageInfo_Health.Size(m)
}
func (m *Health) XXX_DiscardUnknown() {
	xxx_messageInfo_Health.DiscardUnknown(m)
}

var xxx_messageInfo_Health proto.InternalMessageInfo

func (m *Health) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Organizations struct {
	Organizations        []string `protobuf:"bytes,1,rep,name=organizations,proto3" json:"organizations,omitempty" toml:"organizations,omitempty" mapstructure:"organizations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Organizations) Reset()         { *m = Organizations{} }
func (m *Organizations) String() string { return proto.CompactTextString(m) }
func (*Organizations) ProtoMessage()    {}
func (*Organizations) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b42b0f1f6d7b8d4, []int{2}
}

func (m *Organizations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Organizations.Unmarshal(m, b)
}
func (m *Organizations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Organizations.Marshal(b, m, deterministic)
}
func (m *Organizations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Organizations.Merge(m, src)
}
func (m *Organizations) XXX_Size() int {
	return xxx_messageInfo_Organizations.Size(m)
}
func (m *Organizations) XXX_DiscardUnknown() {
	xxx_messageInfo_Organizations.DiscardUnknown(m)
}

var xxx_messageInfo_Organizations proto.InternalMessageInfo

func (m *Organizations) GetOrganizations() []string {
	if m != nil {
		return m.Organizations
	}
	return nil
}

type SourceFQDNS struct {
	SourceFqdns          []string `protobuf:"bytes,1,rep,name=source_fqdns,json=sourceFqdns,proto3" json:"source_fqdns,omitempty" toml:"source_fqdns,omitempty" mapstructure:"source_fqdns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *SourceFQDNS) Reset()         { *m = SourceFQDNS{} }
func (m *SourceFQDNS) String() string { return proto.CompactTextString(m) }
func (*SourceFQDNS) ProtoMessage()    {}
func (*SourceFQDNS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b42b0f1f6d7b8d4, []int{3}
}

func (m *SourceFQDNS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceFQDNS.Unmarshal(m, b)
}
func (m *SourceFQDNS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceFQDNS.Marshal(b, m, deterministic)
}
func (m *SourceFQDNS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceFQDNS.Merge(m, src)
}
func (m *SourceFQDNS) XXX_Size() int {
	return xxx_messageInfo_SourceFQDNS.Size(m)
}
func (m *SourceFQDNS) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceFQDNS.DiscardUnknown(m)
}

var xxx_messageInfo_SourceFQDNS proto.InternalMessageInfo

func (m *SourceFQDNS) GetSourceFqdns() []string {
	if m != nil {
		return m.SourceFqdns
	}
	return nil
}

type ExportData struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty" toml:"content,omitempty" mapstructure:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ExportData) Reset()         { *m = ExportData{} }
func (m *ExportData) String() string { return proto.CompactTextString(m) }
func (*ExportData) ProtoMessage()    {}
func (*ExportData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b42b0f1f6d7b8d4, []int{4}
}

func (m *ExportData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportData.Unmarshal(m, b)
}
func (m *ExportData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportData.Marshal(b, m, deterministic)
}
func (m *ExportData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportData.Merge(m, src)
}
func (m *ExportData) XXX_Size() int {
	return xxx_messageInfo_ExportData.Size(m)
}
func (m *ExportData) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportData.DiscardUnknown(m)
}

var xxx_messageInfo_ExportData proto.InternalMessageInfo

func (m *ExportData) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type ReportExportData struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty" toml:"content,omitempty" mapstructure:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ReportExportData) Reset()         { *m = ReportExportData{} }
func (m *ReportExportData) String() string { return proto.CompactTextString(m) }
func (*ReportExportData) ProtoMessage()    {}
func (*ReportExportData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b42b0f1f6d7b8d4, []int{5}
}

func (m *ReportExportData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportExportData.Unmarshal(m, b)
}
func (m *ReportExportData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportExportData.Marshal(b, m, deterministic)
}
func (m *ReportExportData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportExportData.Merge(m, src)
}
func (m *ReportExportData) XXX_Size() int {
	return xxx_messageInfo_ReportExportData.Size(m)
}
func (m *ReportExportData) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportExportData.DiscardUnknown(m)
}

var xxx_messageInfo_ReportExportData proto.InternalMessageInfo

func (m *ReportExportData) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionInfo)(nil), "chef.automate.domain.cfgmgmt.response.VersionInfo")
	proto.RegisterType((*Health)(nil), "chef.automate.domain.cfgmgmt.response.Health")
	proto.RegisterType((*Organizations)(nil), "chef.automate.domain.cfgmgmt.response.Organizations")
	proto.RegisterType((*SourceFQDNS)(nil), "chef.automate.domain.cfgmgmt.response.SourceFQDNS")
	proto.RegisterType((*ExportData)(nil), "chef.automate.domain.cfgmgmt.response.ExportData")
	proto.RegisterType((*ReportExportData)(nil), "chef.automate.domain.cfgmgmt.response.ReportExportData")
}

func init() {
	proto.RegisterFile("api/interservice/cfgmgmt/response/root.proto", fileDescriptor_1b42b0f1f6d7b8d4)
}

var fileDescriptor_1b42b0f1f6d7b8d4 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4b, 0xeb, 0x40,
	0x10, 0xc6, 0xe9, 0x6b, 0x5f, 0x1f, 0x9d, 0xb6, 0x50, 0x96, 0xc7, 0x23, 0xc7, 0xbe, 0xa0, 0xd2,
	0x43, 0xc9, 0x0a, 0xe2, 0x45, 0xbc, 0x28, 0xb5, 0xd4, 0x8b, 0x62, 0x0a, 0x1e, 0xbc, 0xc8, 0x76,
	0x3b, 0x49, 0x17, 0x9a, 0x9d, 0xb8, 0x3b, 0x29, 0xe2, 0x5f, 0x2f, 0xd9, 0x36, 0xa2, 0x27, 0xbd,
	0xcd, 0xef, 0xc7, 0xf7, 0x91, 0xec, 0x0c, 0x4c, 0x55, 0x69, 0xa4, 0xb1, 0x8c, 0xce, 0xa3, 0xdb,
	0x19, 0x8d, 0x52, 0x67, 0x79, 0x91, 0x17, 0x2c, 0x1d, 0xfa, 0x92, 0xac, 0x47, 0xe9, 0x88, 0x38,
	0x29, 0x1d, 0x31, 0x89, 0x63, 0xbd, 0xc1, 0x2c, 0x51, 0x15, 0x53, 0xa1, 0x18, 0x93, 0x35, 0x15,
	0xca, 0xd8, 0xe4, 0xd0, 0x48, 0x9a, 0x46, 0xac, 0xa1, 0xff, 0x88, 0xce, 0x1b, 0xb2, 0xb7, 0x36,
	0x23, 0x21, 0xa0, 0x63, 0x55, 0x81, 0x51, 0x6b, 0xdc, 0x9a, 0xf4, 0xd2, 0x30, 0x8b, 0x08, 0xfe,
	0xec, 0xf6, 0x91, 0xe8, 0x57, 0xd0, 0x0d, 0x8a, 0x11, 0xb4, 0x97, 0x8b, 0xab, 0xa8, 0x1d, 0x6c,
	0x3d, 0x8a, 0xbf, 0xf0, 0x7b, 0x55, 0x99, 0x2d, 0x47, 0x9d, 0xe0, 0xf6, 0x10, 0x8f, 0xa1, 0xbb,
	0x40, 0xb5, 0xe5, 0x8d, 0xf8, 0x07, 0x5d, 0xcf, 0x8a, 0x2b, 0x7f, 0xf8, 0xc2, 0x81, 0xe2, 0x73,
	0x18, 0xde, 0xbb, 0x5c, 0x59, 0xf3, 0xa6, 0xd8, 0x90, 0xf5, 0xe2, 0x08, 0x86, 0xf4, 0x59, 0x44,
	0xad, 0x71, 0x7b, 0xd2, 0x4b, 0xbf, 0xca, 0xf8, 0x14, 0xfa, 0x4b, 0xaa, 0x9c, 0xc6, 0xf9, 0xc3,
	0xec, 0x6e, 0x29, 0xfe, 0xc3, 0xc0, 0x07, 0x7c, 0xce, 0x5e, 0xd6, 0x1f, 0x9d, 0xfe, 0xde, 0xcd,
	0x6b, 0x15, 0x9f, 0x00, 0xdc, 0xbc, 0x96, 0xe4, 0x78, 0xa6, 0x58, 0xd5, 0x4f, 0xd3, 0x64, 0x19,
	0x2d, 0x87, 0xff, 0x19, 0xa4, 0x0d, 0xc6, 0x53, 0x18, 0xa5, 0x58, 0xe7, 0x7e, 0x92, 0xbe, 0xbe,
	0x7c, 0xba, 0xc8, 0x0d, 0x6f, 0xaa, 0x55, 0xa2, 0xa9, 0x90, 0xf5, 0xe6, 0x65, 0xb3, 0x79, 0xf9,
	0xed, 0xd5, 0x56, 0xdd, 0x70, 0xb1, 0xb3, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xe1, 0x3a,
	0xee, 0xe1, 0x01, 0x00, 0x00,
}
