// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugin.proto

package plugin

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

//* Represents the plugin-specific configuration string.
type ConfigureRequest struct {
	//* The configuration for the plugin.
	Configuration string `protobuf:"bytes,1,opt,name=configuration,proto3" json:"configuration,omitempty"`
	//* Global configurations.
	GlobalConfig         *ConfigureRequest_GlobalConfig `protobuf:"bytes,2,opt,name=globalConfig,proto3" json:"globalConfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ConfigureRequest) Reset()         { *m = ConfigureRequest{} }
func (m *ConfigureRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigureRequest) ProtoMessage()    {}
func (*ConfigureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{0}
}

func (m *ConfigureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigureRequest.Unmarshal(m, b)
}
func (m *ConfigureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigureRequest.Marshal(b, m, deterministic)
}
func (m *ConfigureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigureRequest.Merge(m, src)
}
func (m *ConfigureRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigureRequest.Size(m)
}
func (m *ConfigureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigureRequest proto.InternalMessageInfo

func (m *ConfigureRequest) GetConfiguration() string {
	if m != nil {
		return m.Configuration
	}
	return ""
}

func (m *ConfigureRequest) GetGlobalConfig() *ConfigureRequest_GlobalConfig {
	if m != nil {
		return m.GlobalConfig
	}
	return nil
}

//* Global configuration nested type.
type ConfigureRequest_GlobalConfig struct {
	TrustDomain          string   `protobuf:"bytes,1,opt,name=trustDomain,proto3" json:"trustDomain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigureRequest_GlobalConfig) Reset()         { *m = ConfigureRequest_GlobalConfig{} }
func (m *ConfigureRequest_GlobalConfig) String() string { return proto.CompactTextString(m) }
func (*ConfigureRequest_GlobalConfig) ProtoMessage()    {}
func (*ConfigureRequest_GlobalConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{0, 0}
}

func (m *ConfigureRequest_GlobalConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigureRequest_GlobalConfig.Unmarshal(m, b)
}
func (m *ConfigureRequest_GlobalConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigureRequest_GlobalConfig.Marshal(b, m, deterministic)
}
func (m *ConfigureRequest_GlobalConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigureRequest_GlobalConfig.Merge(m, src)
}
func (m *ConfigureRequest_GlobalConfig) XXX_Size() int {
	return xxx_messageInfo_ConfigureRequest_GlobalConfig.Size(m)
}
func (m *ConfigureRequest_GlobalConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigureRequest_GlobalConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigureRequest_GlobalConfig proto.InternalMessageInfo

func (m *ConfigureRequest_GlobalConfig) GetTrustDomain() string {
	if m != nil {
		return m.TrustDomain
	}
	return ""
}

//* Represents a list of configuration problems
//found in the configuration string.
type ConfigureResponse struct {
	//* A list of errors
	ErrorList            []string `protobuf:"bytes,1,rep,name=errorList,proto3" json:"errorList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigureResponse) Reset()         { *m = ConfigureResponse{} }
func (m *ConfigureResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigureResponse) ProtoMessage()    {}
func (*ConfigureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{1}
}

func (m *ConfigureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigureResponse.Unmarshal(m, b)
}
func (m *ConfigureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigureResponse.Marshal(b, m, deterministic)
}
func (m *ConfigureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigureResponse.Merge(m, src)
}
func (m *ConfigureResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigureResponse.Size(m)
}
func (m *ConfigureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigureResponse proto.InternalMessageInfo

func (m *ConfigureResponse) GetErrorList() []string {
	if m != nil {
		return m.ErrorList
	}
	return nil
}

//* Represents an empty request.
type GetPluginInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPluginInfoRequest) Reset()         { *m = GetPluginInfoRequest{} }
func (m *GetPluginInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetPluginInfoRequest) ProtoMessage()    {}
func (*GetPluginInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{2}
}

func (m *GetPluginInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginInfoRequest.Unmarshal(m, b)
}
func (m *GetPluginInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetPluginInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginInfoRequest.Merge(m, src)
}
func (m *GetPluginInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetPluginInfoRequest.Size(m)
}
func (m *GetPluginInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginInfoRequest proto.InternalMessageInfo

//* Represents the plugin metadata.
type GetPluginInfoResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Category             string   `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	DateCreated          string   `protobuf:"bytes,5,opt,name=dateCreated,proto3" json:"dateCreated,omitempty"`
	Location             string   `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Version              string   `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	Author               string   `protobuf:"bytes,8,opt,name=author,proto3" json:"author,omitempty"`
	Company              string   `protobuf:"bytes,9,opt,name=company,proto3" json:"company,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPluginInfoResponse) Reset()         { *m = GetPluginInfoResponse{} }
func (m *GetPluginInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetPluginInfoResponse) ProtoMessage()    {}
func (*GetPluginInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{3}
}

func (m *GetPluginInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginInfoResponse.Unmarshal(m, b)
}
func (m *GetPluginInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetPluginInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginInfoResponse.Merge(m, src)
}
func (m *GetPluginInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetPluginInfoResponse.Size(m)
}
func (m *GetPluginInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginInfoResponse proto.InternalMessageInfo

func (m *GetPluginInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetPluginInfoResponse) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *GetPluginInfoResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetPluginInfoResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetPluginInfoResponse) GetDateCreated() string {
	if m != nil {
		return m.DateCreated
	}
	return ""
}

func (m *GetPluginInfoResponse) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *GetPluginInfoResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GetPluginInfoResponse) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *GetPluginInfoResponse) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func init() {
	proto.RegisterType((*ConfigureRequest)(nil), "spire.common.plugin.ConfigureRequest")
	proto.RegisterType((*ConfigureRequest_GlobalConfig)(nil), "spire.common.plugin.ConfigureRequest.GlobalConfig")
	proto.RegisterType((*ConfigureResponse)(nil), "spire.common.plugin.ConfigureResponse")
	proto.RegisterType((*GetPluginInfoRequest)(nil), "spire.common.plugin.GetPluginInfoRequest")
	proto.RegisterType((*GetPluginInfoResponse)(nil), "spire.common.plugin.GetPluginInfoResponse")
}

func init() { proto.RegisterFile("plugin.proto", fileDescriptor_22a625af4bc1cc87) }

var fileDescriptor_22a625af4bc1cc87 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcf, 0x4e, 0xf2, 0x40,
	0x10, 0x4f, 0x81, 0xaf, 0xd0, 0x81, 0x2f, 0xd1, 0x55, 0xc9, 0x86, 0x78, 0x68, 0x1a, 0x0f, 0x9c,
	0x1a, 0xc5, 0x37, 0x10, 0x13, 0x62, 0xe2, 0xc1, 0xf4, 0xe0, 0xc1, 0xdb, 0x52, 0x86, 0xda, 0xa4,
	0xdd, 0xa9, 0xbb, 0x5b, 0x13, 0x9e, 0xc0, 0xf7, 0xf2, 0xc9, 0x4c, 0x77, 0x8b, 0x14, 0xe2, 0x6d,
	0x7f, 0xff, 0xf6, 0x37, 0xd3, 0x2d, 0x4c, 0xaa, 0xa2, 0xce, 0x72, 0x19, 0x57, 0x8a, 0x0c, 0xb1,
	0x0b, 0x5d, 0xe5, 0x0a, 0xe3, 0x94, 0xca, 0x92, 0x64, 0xec, 0xa4, 0xe8, 0xdb, 0x83, 0xb3, 0x25,
	0xc9, 0x6d, 0x9e, 0xd5, 0x0a, 0x13, 0xfc, 0xa8, 0x51, 0x1b, 0x76, 0x03, 0xff, 0xd3, 0x96, 0x13,
	0x26, 0x27, 0xc9, 0xbd, 0xd0, 0x9b, 0x07, 0xc9, 0x31, 0xc9, 0x5e, 0x61, 0x92, 0x15, 0xb4, 0x16,
	0x85, 0xcb, 0xf3, 0x5e, 0xe8, 0xcd, 0xc7, 0x8b, 0x45, 0xfc, 0x47, 0x4d, 0x7c, 0x5a, 0x11, 0xaf,
	0x3a, 0xc9, 0xe4, 0xe8, 0x9e, 0xd9, 0x2d, 0x4c, 0xba, 0x2a, 0x0b, 0x61, 0x6c, 0x54, 0xad, 0xcd,
	0x23, 0x95, 0x22, 0xdf, 0xcf, 0xd2, 0xa5, 0xa2, 0x3b, 0x38, 0xef, 0x14, 0xe8, 0x8a, 0xa4, 0x46,
	0x76, 0x0d, 0x01, 0x2a, 0x45, 0xea, 0x39, 0xd7, 0x86, 0x7b, 0x61, 0x7f, 0x1e, 0x24, 0x07, 0x22,
	0x9a, 0xc2, 0xe5, 0x0a, 0xcd, 0x8b, 0x9d, 0xee, 0x49, 0x6e, 0xa9, 0x9d, 0x2b, 0xfa, 0xea, 0xc1,
	0xd5, 0x89, 0xd0, 0xde, 0xc7, 0x60, 0x20, 0x45, 0x89, 0x6d, 0xbf, 0x3d, 0xb3, 0x19, 0x8c, 0x52,
	0x61, 0x30, 0x23, 0xb5, 0xb3, 0xeb, 0x07, 0xc9, 0x2f, 0x6e, 0xfc, 0x66, 0x57, 0x21, 0xef, 0x3b,
	0x7f, 0x73, 0x6e, 0x56, 0xd9, 0xa0, 0x4e, 0x55, 0x5e, 0xd9, 0xcf, 0x3a, 0x70, 0xab, 0x74, 0x28,
	0xeb, 0x10, 0x06, 0x97, 0x0a, 0x85, 0xc1, 0x0d, 0xff, 0xd7, 0x3a, 0x0e, 0x54, 0xd3, 0x59, 0x50,
	0xea, 0xde, 0xc5, 0x77, 0x9d, 0x7b, 0xcc, 0x38, 0x0c, 0x3f, 0x51, 0xe9, 0x46, 0x1a, 0x5a, 0x69,
	0x0f, 0xd9, 0x14, 0x7c, 0x51, 0x9b, 0x77, 0x52, 0x7c, 0x64, 0x85, 0x16, 0x35, 0x89, 0x94, 0xca,
	0x4a, 0xc8, 0x1d, 0x0f, 0x5c, 0xa2, 0x85, 0x0f, 0xa3, 0x37, 0xdf, 0x3d, 0xde, 0xda, 0xb7, 0xff,
	0xcf, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x44, 0x00, 0x6d, 0xf3, 0x4f, 0x02, 0x00, 0x00,
}
