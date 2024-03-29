// Code generated by protoc-gen-go. DO NOT EDIT.
// source: company_cms_details.proto

package companycmspb

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

type CompanyCmsDetails struct {
	CompanyName          string   `protobuf:"bytes,1,opt,name=CompanyName,proto3" json:"CompanyName,omitempty"`
	CompanyWebsite       string   `protobuf:"bytes,2,opt,name=CompanyWebsite,proto3" json:"CompanyWebsite,omitempty"`
	WantedDepartments    string   `protobuf:"bytes,3,opt,name=WantedDepartments,proto3" json:"WantedDepartments,omitempty"`
	WantedLocations      string   `protobuf:"bytes,4,opt,name=WantedLocations,proto3" json:"WantedLocations,omitempty"`
	GreenHouse           bool     `protobuf:"varint,5,opt,name=GreenHouse,proto3" json:"GreenHouse,omitempty"`
	Lever                bool     `protobuf:"varint,6,opt,name=Lever,proto3" json:"Lever,omitempty"`
	Other                bool     `protobuf:"varint,7,opt,name=Other,proto3" json:"Other,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompanyCmsDetails) Reset()         { *m = CompanyCmsDetails{} }
func (m *CompanyCmsDetails) String() string { return proto.CompactTextString(m) }
func (*CompanyCmsDetails) ProtoMessage()    {}
func (*CompanyCmsDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_93271db20c458770, []int{0}
}

func (m *CompanyCmsDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompanyCmsDetails.Unmarshal(m, b)
}
func (m *CompanyCmsDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompanyCmsDetails.Marshal(b, m, deterministic)
}
func (m *CompanyCmsDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompanyCmsDetails.Merge(m, src)
}
func (m *CompanyCmsDetails) XXX_Size() int {
	return xxx_messageInfo_CompanyCmsDetails.Size(m)
}
func (m *CompanyCmsDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CompanyCmsDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CompanyCmsDetails proto.InternalMessageInfo

func (m *CompanyCmsDetails) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *CompanyCmsDetails) GetCompanyWebsite() string {
	if m != nil {
		return m.CompanyWebsite
	}
	return ""
}

func (m *CompanyCmsDetails) GetWantedDepartments() string {
	if m != nil {
		return m.WantedDepartments
	}
	return ""
}

func (m *CompanyCmsDetails) GetWantedLocations() string {
	if m != nil {
		return m.WantedLocations
	}
	return ""
}

func (m *CompanyCmsDetails) GetGreenHouse() bool {
	if m != nil {
		return m.GreenHouse
	}
	return false
}

func (m *CompanyCmsDetails) GetLever() bool {
	if m != nil {
		return m.Lever
	}
	return false
}

func (m *CompanyCmsDetails) GetOther() bool {
	if m != nil {
		return m.Other
	}
	return false
}

func init() {
	proto.RegisterType((*CompanyCmsDetails)(nil), "CompanyCmsDetails")
}

func init() {
	proto.RegisterFile("company_cms_details.proto", fileDescriptor_93271db20c458770)
}

var fileDescriptor_93271db20c458770 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xd0, 0x3f, 0x4b, 0xc4, 0x40,
	0x10, 0x05, 0x70, 0x72, 0x7a, 0xa7, 0x8e, 0x72, 0x92, 0xc5, 0x62, 0x6d, 0x24, 0x58, 0x48, 0x0a,
	0xb1, 0xf1, 0x1b, 0x98, 0x80, 0x16, 0x41, 0x21, 0x4d, 0xc0, 0x26, 0x6c, 0x36, 0x03, 0x06, 0xdc,
	0x3f, 0xec, 0x8c, 0x82, 0xa5, 0xdf, 0x5c, 0xdc, 0x4d, 0x11, 0x72, 0xe5, 0xfb, 0xbd, 0x57, 0x3d,
	0xb8, 0xd6, 0xce, 0x78, 0x65, 0x7f, 0x7a, 0x6d, 0xa8, 0x1f, 0x91, 0xd5, 0xf4, 0x49, 0x0f, 0x3e,
	0x38, 0x76, 0xb7, 0xbf, 0x1b, 0xc8, 0xab, 0xd4, 0x56, 0x86, 0xea, 0xd4, 0x89, 0x02, 0xce, 0x67,
	0x7c, 0x55, 0x06, 0x65, 0x56, 0x64, 0xe5, 0x59, 0xbb, 0x24, 0x71, 0x07, 0xfb, 0x39, 0x76, 0x38,
	0xd0, 0xc4, 0x28, 0x37, 0x71, 0xb4, 0x52, 0x71, 0x0f, 0x79, 0xa7, 0x2c, 0xe3, 0x58, 0xa3, 0x57,
	0x81, 0x0d, 0x5a, 0x26, 0x79, 0x14, 0xa7, 0x87, 0x85, 0x28, 0xe1, 0x32, 0x61, 0xe3, 0xb4, 0xe2,
	0xc9, 0x59, 0x92, 0xc7, 0x71, 0xbb, 0x66, 0x71, 0x03, 0xf0, 0x1c, 0x10, 0xed, 0x8b, 0xfb, 0x22,
	0x94, 0xdb, 0x22, 0x2b, 0x4f, 0xdb, 0x85, 0x88, 0x2b, 0xd8, 0x36, 0xf8, 0x8d, 0x41, 0xee, 0x62,
	0x95, 0xc2, 0xbf, 0xbe, 0xf1, 0x07, 0x06, 0x79, 0x92, 0x34, 0x86, 0xa7, 0xfd, 0xfb, 0xc5, 0x7c,
	0x90, 0x36, 0xe4, 0x87, 0x61, 0x17, 0xaf, 0x79, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xed,
	0x22, 0x7e, 0x37, 0x01, 0x00, 0x00,
}
