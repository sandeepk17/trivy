// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/detector/service.proto

package detector

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

type OSDetectRequest struct {
	OsFamily             string     `protobuf:"bytes,1,opt,name=os_family,json=osFamily,proto3" json:"os_family,omitempty"`
	OsName               string     `protobuf:"bytes,2,opt,name=os_name,json=osName,proto3" json:"os_name,omitempty"`
	Packages             []*Package `protobuf:"bytes,3,rep,name=packages,proto3" json:"packages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OSDetectRequest) Reset()         { *m = OSDetectRequest{} }
func (m *OSDetectRequest) String() string { return proto.CompactTextString(m) }
func (*OSDetectRequest) ProtoMessage()    {}
func (*OSDetectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_93e16dbd737b8924, []int{0}
}

func (m *OSDetectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OSDetectRequest.Unmarshal(m, b)
}
func (m *OSDetectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OSDetectRequest.Marshal(b, m, deterministic)
}
func (m *OSDetectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OSDetectRequest.Merge(m, src)
}
func (m *OSDetectRequest) XXX_Size() int {
	return xxx_messageInfo_OSDetectRequest.Size(m)
}
func (m *OSDetectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OSDetectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OSDetectRequest proto.InternalMessageInfo

func (m *OSDetectRequest) GetOsFamily() string {
	if m != nil {
		return m.OsFamily
	}
	return ""
}

func (m *OSDetectRequest) GetOsName() string {
	if m != nil {
		return m.OsName
	}
	return ""
}

func (m *OSDetectRequest) GetPackages() []*Package {
	if m != nil {
		return m.Packages
	}
	return nil
}

type DetectResponse struct {
	DetectedVulnerabilities []*DetectedVulnerability `protobuf:"bytes,1,rep,name=detected_vulnerabilities,json=detectedVulnerabilities,proto3" json:"detected_vulnerabilities,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *DetectResponse) Reset()         { *m = DetectResponse{} }
func (m *DetectResponse) String() string { return proto.CompactTextString(m) }
func (*DetectResponse) ProtoMessage()    {}
func (*DetectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_93e16dbd737b8924, []int{1}
}

func (m *DetectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectResponse.Unmarshal(m, b)
}
func (m *DetectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectResponse.Marshal(b, m, deterministic)
}
func (m *DetectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectResponse.Merge(m, src)
}
func (m *DetectResponse) XXX_Size() int {
	return xxx_messageInfo_DetectResponse.Size(m)
}
func (m *DetectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DetectResponse proto.InternalMessageInfo

func (m *DetectResponse) GetDetectedVulnerabilities() []*DetectedVulnerability {
	if m != nil {
		return m.DetectedVulnerabilities
	}
	return nil
}

type Package struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Release              string   `protobuf:"bytes,3,opt,name=release,proto3" json:"release,omitempty"`
	Epoch                int32    `protobuf:"varint,4,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Arch                 string   `protobuf:"bytes,5,opt,name=arch,proto3" json:"arch,omitempty"`
	SrcName              string   `protobuf:"bytes,6,opt,name=src_name,json=srcName,proto3" json:"src_name,omitempty"`
	SrcVersion           string   `protobuf:"bytes,7,opt,name=src_version,json=srcVersion,proto3" json:"src_version,omitempty"`
	SrcRelease           string   `protobuf:"bytes,8,opt,name=src_release,json=srcRelease,proto3" json:"src_release,omitempty"`
	SrcEpoch             int32    `protobuf:"varint,9,opt,name=src_epoch,json=srcEpoch,proto3" json:"src_epoch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Package) Reset()         { *m = Package{} }
func (m *Package) String() string { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()    {}
func (*Package) Descriptor() ([]byte, []int) {
	return fileDescriptor_93e16dbd737b8924, []int{2}
}

func (m *Package) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Package.Unmarshal(m, b)
}
func (m *Package) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Package.Marshal(b, m, deterministic)
}
func (m *Package) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Package.Merge(m, src)
}
func (m *Package) XXX_Size() int {
	return xxx_messageInfo_Package.Size(m)
}
func (m *Package) XXX_DiscardUnknown() {
	xxx_messageInfo_Package.DiscardUnknown(m)
}

var xxx_messageInfo_Package proto.InternalMessageInfo

func (m *Package) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Package) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Package) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *Package) GetEpoch() int32 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *Package) GetArch() string {
	if m != nil {
		return m.Arch
	}
	return ""
}

func (m *Package) GetSrcName() string {
	if m != nil {
		return m.SrcName
	}
	return ""
}

func (m *Package) GetSrcVersion() string {
	if m != nil {
		return m.SrcVersion
	}
	return ""
}

func (m *Package) GetSrcRelease() string {
	if m != nil {
		return m.SrcRelease
	}
	return ""
}

func (m *Package) GetSrcEpoch() int32 {
	if m != nil {
		return m.SrcEpoch
	}
	return 0
}

type LibDetectRequest struct {
	FilePath             string     `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	Libraries            []*Library `protobuf:"bytes,2,rep,name=libraries,proto3" json:"libraries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LibDetectRequest) Reset()         { *m = LibDetectRequest{} }
func (m *LibDetectRequest) String() string { return proto.CompactTextString(m) }
func (*LibDetectRequest) ProtoMessage()    {}
func (*LibDetectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_93e16dbd737b8924, []int{3}
}

func (m *LibDetectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LibDetectRequest.Unmarshal(m, b)
}
func (m *LibDetectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LibDetectRequest.Marshal(b, m, deterministic)
}
func (m *LibDetectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LibDetectRequest.Merge(m, src)
}
func (m *LibDetectRequest) XXX_Size() int {
	return xxx_messageInfo_LibDetectRequest.Size(m)
}
func (m *LibDetectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LibDetectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LibDetectRequest proto.InternalMessageInfo

func (m *LibDetectRequest) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *LibDetectRequest) GetLibraries() []*Library {
	if m != nil {
		return m.Libraries
	}
	return nil
}

type Library struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Library) Reset()         { *m = Library{} }
func (m *Library) String() string { return proto.CompactTextString(m) }
func (*Library) ProtoMessage()    {}
func (*Library) Descriptor() ([]byte, []int) {
	return fileDescriptor_93e16dbd737b8924, []int{4}
}

func (m *Library) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Library.Unmarshal(m, b)
}
func (m *Library) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Library.Marshal(b, m, deterministic)
}
func (m *Library) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Library.Merge(m, src)
}
func (m *Library) XXX_Size() int {
	return xxx_messageInfo_Library.Size(m)
}
func (m *Library) XXX_DiscardUnknown() {
	xxx_messageInfo_Library.DiscardUnknown(m)
}

var xxx_messageInfo_Library proto.InternalMessageInfo

func (m *Library) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Library) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type DetectedVulnerability struct {
	VulnerabilityId      string   `protobuf:"bytes,1,opt,name=vulnerability_id,json=vulnerabilityId,proto3" json:"vulnerability_id,omitempty"`
	PkgName              string   `protobuf:"bytes,2,opt,name=pkg_name,json=pkgName,proto3" json:"pkg_name,omitempty"`
	InstalledVersion     string   `protobuf:"bytes,3,opt,name=installed_version,json=installedVersion,proto3" json:"installed_version,omitempty"`
	FixedVersion         string   `protobuf:"bytes,4,opt,name=fixed_version,json=fixedVersion,proto3" json:"fixed_version,omitempty"`
	Title                string   `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Severity             string   `protobuf:"bytes,7,opt,name=severity,proto3" json:"severity,omitempty"`
	References           []string `protobuf:"bytes,8,rep,name=references,proto3" json:"references,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetectedVulnerability) Reset()         { *m = DetectedVulnerability{} }
func (m *DetectedVulnerability) String() string { return proto.CompactTextString(m) }
func (*DetectedVulnerability) ProtoMessage()    {}
func (*DetectedVulnerability) Descriptor() ([]byte, []int) {
	return fileDescriptor_93e16dbd737b8924, []int{5}
}

func (m *DetectedVulnerability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectedVulnerability.Unmarshal(m, b)
}
func (m *DetectedVulnerability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectedVulnerability.Marshal(b, m, deterministic)
}
func (m *DetectedVulnerability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectedVulnerability.Merge(m, src)
}
func (m *DetectedVulnerability) XXX_Size() int {
	return xxx_messageInfo_DetectedVulnerability.Size(m)
}
func (m *DetectedVulnerability) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectedVulnerability.DiscardUnknown(m)
}

var xxx_messageInfo_DetectedVulnerability proto.InternalMessageInfo

func (m *DetectedVulnerability) GetVulnerabilityId() string {
	if m != nil {
		return m.VulnerabilityId
	}
	return ""
}

func (m *DetectedVulnerability) GetPkgName() string {
	if m != nil {
		return m.PkgName
	}
	return ""
}

func (m *DetectedVulnerability) GetInstalledVersion() string {
	if m != nil {
		return m.InstalledVersion
	}
	return ""
}

func (m *DetectedVulnerability) GetFixedVersion() string {
	if m != nil {
		return m.FixedVersion
	}
	return ""
}

func (m *DetectedVulnerability) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *DetectedVulnerability) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DetectedVulnerability) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

func (m *DetectedVulnerability) GetReferences() []string {
	if m != nil {
		return m.References
	}
	return nil
}

func init() {
	proto.RegisterType((*OSDetectRequest)(nil), "trivy.detector.OSDetectRequest")
	proto.RegisterType((*DetectResponse)(nil), "trivy.detector.DetectResponse")
	proto.RegisterType((*Package)(nil), "trivy.detector.Package")
	proto.RegisterType((*LibDetectRequest)(nil), "trivy.detector.LibDetectRequest")
	proto.RegisterType((*Library)(nil), "trivy.detector.Library")
	proto.RegisterType((*DetectedVulnerability)(nil), "trivy.detector.DetectedVulnerability")
}

func init() { proto.RegisterFile("rpc/detector/service.proto", fileDescriptor_93e16dbd737b8924) }

var fileDescriptor_93e16dbd737b8924 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0xdb, 0x4c,
	0x10, 0xc5, 0xf1, 0x8f, 0xe4, 0xf1, 0xf7, 0x25, 0xee, 0x92, 0x62, 0xd5, 0x81, 0x44, 0xa8, 0x14,
	0x5c, 0x0a, 0x0e, 0x38, 0x94, 0xde, 0x97, 0xb4, 0x90, 0x12, 0xda, 0xa0, 0x42, 0x4a, 0x7b, 0xe3,
	0xae, 0xa5, 0xb1, 0xbd, 0x58, 0xd6, 0xaa, 0xbb, 0x1b, 0x53, 0x41, 0xdf, 0xa6, 0xcf, 0xd8, 0xfb,
	0xb2, 0x3f, 0x52, 0x1c, 0xd7, 0x37, 0xb9, 0xdb, 0x99, 0x73, 0x74, 0x66, 0x76, 0xce, 0x8e, 0x60,
	0x28, 0x8a, 0xe4, 0x3c, 0x45, 0x85, 0x89, 0xe2, 0xe2, 0x5c, 0xa2, 0xd8, 0xb0, 0x04, 0xc7, 0x85,
	0xe0, 0x8a, 0x93, 0x43, 0x25, 0xd8, 0xa6, 0x1c, 0x57, 0x68, 0xf4, 0x0b, 0x8e, 0x3e, 0x7d, 0xbe,
	0x34, 0x51, 0x8c, 0x3f, 0xee, 0x50, 0x2a, 0x72, 0x02, 0x5d, 0x2e, 0xa7, 0x73, 0xba, 0x66, 0x59,
	0x19, 0x34, 0xc2, 0xc6, 0xa8, 0x1b, 0xfb, 0x5c, 0xbe, 0x37, 0x31, 0x19, 0x80, 0xc7, 0xe5, 0x34,
	0xa7, 0x6b, 0x0c, 0x0e, 0x0c, 0xd4, 0xe1, 0xf2, 0x23, 0x5d, 0x23, 0xb9, 0x00, 0xbf, 0xa0, 0xc9,
	0x8a, 0x2e, 0x50, 0x06, 0xcd, 0xb0, 0x39, 0xea, 0x4d, 0x06, 0xe3, 0x87, 0xb5, 0xc6, 0x37, 0x16,
	0x8f, 0x6b, 0x62, 0x24, 0xe0, 0xb0, 0xaa, 0x2d, 0x0b, 0x9e, 0x4b, 0x24, 0xdf, 0x21, 0xb0, 0x7c,
	0x4c, 0xa7, 0x9b, 0xbb, 0x2c, 0x47, 0x41, 0x67, 0x2c, 0x63, 0x8a, 0xa1, 0x0c, 0x1a, 0x46, 0xf6,
	0xc5, 0xae, 0xec, 0xa5, 0xe3, 0xdf, 0x6e, 0xd1, 0xcb, 0x78, 0x90, 0xee, 0x49, 0x33, 0x94, 0xd1,
	0x9f, 0x06, 0x78, 0xae, 0x13, 0x42, 0xa0, 0x65, 0xae, 0x62, 0x6f, 0x69, 0xce, 0x24, 0x00, 0x6f,
	0x83, 0x42, 0x32, 0x9e, 0xbb, 0x1b, 0x56, 0xa1, 0x46, 0x04, 0x66, 0x48, 0x25, 0x06, 0x4d, 0x8b,
	0xb8, 0x90, 0x1c, 0x43, 0x1b, 0x0b, 0x9e, 0x2c, 0x83, 0x56, 0xd8, 0x18, 0xb5, 0x63, 0x1b, 0x68,
	0x75, 0x2a, 0x92, 0x65, 0xd0, 0xb6, 0xea, 0xfa, 0x4c, 0x9e, 0x81, 0x2f, 0x45, 0x62, 0x07, 0xd8,
	0xb1, 0x22, 0x52, 0x24, 0x66, 0x82, 0x67, 0xd0, 0xd3, 0x50, 0x55, 0xdc, 0x33, 0x28, 0x48, 0x91,
	0xdc, 0xba, 0xfa, 0x8e, 0x50, 0xf5, 0xe0, 0xd7, 0x84, 0xd8, 0xb5, 0x71, 0x02, 0x5d, 0x4d, 0xb0,
	0xad, 0x74, 0x4d, 0x2b, 0xba, 0xda, 0x3b, 0x1d, 0x47, 0x73, 0xe8, 0x5f, 0xb3, 0xd9, 0x3f, 0x56,
	0xcf, 0x59, 0x86, 0xd3, 0x82, 0xaa, 0x65, 0x65, 0xb5, 0x4e, 0xdc, 0x50, 0xb5, 0x24, 0xaf, 0xa1,
	0x9b, 0xb1, 0x99, 0xa0, 0x42, 0xcf, 0xfe, 0x60, 0xbf, 0xa5, 0xd7, 0x86, 0x50, 0xc6, 0xf7, 0xcc,
	0xe8, 0x0d, 0x78, 0x2e, 0xfb, 0xb8, 0xf1, 0x46, 0xbf, 0x0f, 0xe0, 0xe9, 0x5e, 0x2f, 0xc9, 0x4b,
	0xe8, 0x6f, 0xbf, 0x85, 0x72, 0xca, 0x52, 0xa7, 0x79, 0xf4, 0x20, 0x7f, 0x95, 0xea, 0xf9, 0x16,
	0xab, 0xc5, 0xf6, 0x03, 0xf5, 0x8a, 0xd5, 0xc2, 0xcc, 0xf7, 0x15, 0x3c, 0x61, 0xb9, 0x54, 0x34,
	0xcb, 0xf4, 0xdb, 0x72, 0x3d, 0x58, 0x23, 0xfb, 0x35, 0x50, 0xcd, 0xfa, 0x39, 0xfc, 0x3f, 0x67,
	0x3f, 0xb7, 0x88, 0x2d, 0x43, 0xfc, 0xcf, 0x24, 0x2b, 0xd2, 0x31, 0xb4, 0x15, 0x53, 0x19, 0x3a,
	0x87, 0x6d, 0x40, 0x42, 0xe8, 0xa5, 0x28, 0x13, 0xc1, 0x0a, 0xa5, 0x3f, 0xb4, 0x2e, 0x6f, 0xa7,
	0xc8, 0x10, 0x7c, 0x89, 0x1b, 0x14, 0x4c, 0x95, 0xce, 0xe6, 0x3a, 0x26, 0xa7, 0x00, 0x02, 0xe7,
	0x28, 0x30, 0x4f, 0x50, 0x06, 0x7e, 0xd8, 0xd4, 0x1e, 0xdf, 0x67, 0x26, 0x5f, 0x00, 0xaa, 0x85,
	0xe5, 0x82, 0x5c, 0x41, 0xc7, 0x9e, 0xc9, 0xd9, 0xae, 0x35, 0x3b, 0x6b, 0x3d, 0x3c, 0xdd, 0xbf,
	0x37, 0xd5, 0xe6, 0x4d, 0xbe, 0x42, 0xaf, 0x7e, 0x1f, 0x5c, 0x90, 0x0f, 0xb5, 0x72, 0xb8, 0xc7,
	0xf4, 0x47, 0x49, 0xbf, 0x85, 0x6f, 0x7e, 0x05, 0xcd, 0x3a, 0xe6, 0x3f, 0x74, 0xf1, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x43, 0xf0, 0xb3, 0x38, 0xa5, 0x04, 0x00, 0x00,
}
