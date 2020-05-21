// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/company/updatecompanypb/update_company_details.proto

package updatecompanypb

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

type UpdateCompanyDetails struct {
	CompanyUUID          string   `protobuf:"bytes,1,opt,name=CompanyUUID,proto3" json:"CompanyUUID,omitempty"`
	CompanyName          string   `protobuf:"bytes,2,opt,name=CompanyName,proto3" json:"CompanyName,omitempty"`
	CompanyWebsite       string   `protobuf:"bytes,3,opt,name=CompanyWebsite,proto3" json:"CompanyWebsite,omitempty"`
	GreenHouse           bool     `protobuf:"varint,4,opt,name=GreenHouse,proto3" json:"GreenHouse,omitempty"`
	Lever                bool     `protobuf:"varint,5,opt,name=Lever,proto3" json:"Lever,omitempty"`
	Other                bool     `protobuf:"varint,6,opt,name=Other,proto3" json:"Other,omitempty"`
	WantedDepartments    string   `protobuf:"bytes,7,opt,name=WantedDepartments,proto3" json:"WantedDepartments,omitempty"`
	WantedLocations      string   `protobuf:"bytes,8,opt,name=WantedLocations,proto3" json:"WantedLocations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCompanyDetails) Reset()         { *m = UpdateCompanyDetails{} }
func (m *UpdateCompanyDetails) String() string { return proto.CompactTextString(m) }
func (*UpdateCompanyDetails) ProtoMessage()    {}
func (*UpdateCompanyDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_97c904546f7e3717, []int{0}
}

func (m *UpdateCompanyDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCompanyDetails.Unmarshal(m, b)
}
func (m *UpdateCompanyDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCompanyDetails.Marshal(b, m, deterministic)
}
func (m *UpdateCompanyDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCompanyDetails.Merge(m, src)
}
func (m *UpdateCompanyDetails) XXX_Size() int {
	return xxx_messageInfo_UpdateCompanyDetails.Size(m)
}
func (m *UpdateCompanyDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCompanyDetails.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCompanyDetails proto.InternalMessageInfo

func (m *UpdateCompanyDetails) GetCompanyUUID() string {
	if m != nil {
		return m.CompanyUUID
	}
	return ""
}

func (m *UpdateCompanyDetails) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *UpdateCompanyDetails) GetCompanyWebsite() string {
	if m != nil {
		return m.CompanyWebsite
	}
	return ""
}

func (m *UpdateCompanyDetails) GetGreenHouse() bool {
	if m != nil {
		return m.GreenHouse
	}
	return false
}

func (m *UpdateCompanyDetails) GetLever() bool {
	if m != nil {
		return m.Lever
	}
	return false
}

func (m *UpdateCompanyDetails) GetOther() bool {
	if m != nil {
		return m.Other
	}
	return false
}

func (m *UpdateCompanyDetails) GetWantedDepartments() string {
	if m != nil {
		return m.WantedDepartments
	}
	return ""
}

func (m *UpdateCompanyDetails) GetWantedLocations() string {
	if m != nil {
		return m.WantedLocations
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateCompanyDetails)(nil), "UpdateCompanyDetails")
}

func init() {
	proto.RegisterFile("pb/company/updatecompanypb/update_company_details.proto", fileDescriptor_97c904546f7e3717)
}

var fileDescriptor_97c904546f7e3717 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xd0, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x05, 0x60, 0x52, 0x6d, 0xad, 0x23, 0x58, 0xba, 0xf4, 0x30, 0x27, 0x09, 0x1e, 0x24, 0x07,
	0xb1, 0x07, 0x0f, 0xde, 0x35, 0xa0, 0x42, 0x51, 0x28, 0x84, 0x82, 0x97, 0xb2, 0xdb, 0x0c, 0x18,
	0x30, 0xbb, 0xcb, 0xee, 0x44, 0xf0, 0x0f, 0xf9, 0x3b, 0x25, 0xbb, 0x7b, 0x88, 0xe9, 0xf1, 0x7d,
	0xef, 0xc1, 0xc0, 0xc0, 0x83, 0x55, 0xeb, 0x83, 0x69, 0xad, 0xd4, 0x3f, 0xeb, 0xce, 0xd6, 0x92,
	0x29, 0x25, 0xab, 0x52, 0xde, 0x27, 0xd8, 0xd7, 0xc4, 0xb2, 0xf9, 0xf2, 0x77, 0xd6, 0x19, 0x36,
	0xd7, 0xbf, 0x13, 0x58, 0x55, 0x61, 0xf0, 0x14, 0xfb, 0x32, 0xd6, 0x22, 0x87, 0x8b, 0x24, 0x55,
	0xf5, 0x5a, 0x62, 0x96, 0x67, 0xc5, 0xf9, 0x76, 0x48, 0x83, 0xc5, 0x9b, 0x6c, 0x09, 0x27, 0xff,
	0x16, 0x3d, 0x89, 0x1b, 0xb8, 0x4c, 0x71, 0x47, 0xca, 0x37, 0x4c, 0x78, 0x12, 0x46, 0x23, 0x15,
	0x57, 0x00, 0xcf, 0x8e, 0x48, 0xbf, 0x98, 0xce, 0x13, 0x9e, 0xe6, 0x59, 0x31, 0xdf, 0x0e, 0x44,
	0xac, 0x60, 0xba, 0xa1, 0x6f, 0x72, 0x38, 0x0d, 0x55, 0x0c, 0xbd, 0xbe, 0xf3, 0x27, 0x39, 0x9c,
	0x45, 0x0d, 0x41, 0xdc, 0xc2, 0x72, 0x27, 0x35, 0x53, 0x5d, 0x92, 0x95, 0x8e, 0x5b, 0xd2, 0xec,
	0xf1, 0x2c, 0x9c, 0x3d, 0x2e, 0x44, 0x01, 0x8b, 0x88, 0x1b, 0x73, 0x90, 0xdc, 0x18, 0xed, 0x71,
	0x1e, 0xb6, 0x63, 0x7e, 0x5c, 0x7e, 0x2c, 0x46, 0x8f, 0x55, 0xb3, 0xf0, 0xc2, 0xfb, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x37, 0x8a, 0xab, 0x28, 0x7d, 0x01, 0x00, 0x00,
}
