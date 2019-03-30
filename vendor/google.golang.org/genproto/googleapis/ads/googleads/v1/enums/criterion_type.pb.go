// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/criterion_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible criterion types.
type CriterionTypeEnum_CriterionType int32

const (
	// Not specified.
	CriterionTypeEnum_UNSPECIFIED CriterionTypeEnum_CriterionType = 0
	// Used for return value only. Represents value unknown in this version.
	CriterionTypeEnum_UNKNOWN CriterionTypeEnum_CriterionType = 1
	// Keyword. e.g. 'mars cruise'.
	CriterionTypeEnum_KEYWORD CriterionTypeEnum_CriterionType = 2
	// Placement, aka Website. e.g. 'www.flowers4sale.com'
	CriterionTypeEnum_PLACEMENT CriterionTypeEnum_CriterionType = 3
	// Mobile application categories to target.
	CriterionTypeEnum_MOBILE_APP_CATEGORY CriterionTypeEnum_CriterionType = 4
	// Mobile applications to target.
	CriterionTypeEnum_MOBILE_APPLICATION CriterionTypeEnum_CriterionType = 5
	// Devices to target.
	CriterionTypeEnum_DEVICE CriterionTypeEnum_CriterionType = 6
	// Locations to target.
	CriterionTypeEnum_LOCATION CriterionTypeEnum_CriterionType = 7
	// Listing groups to target.
	CriterionTypeEnum_LISTING_GROUP CriterionTypeEnum_CriterionType = 8
	// Ad Schedule.
	CriterionTypeEnum_AD_SCHEDULE CriterionTypeEnum_CriterionType = 9
	// Age range.
	CriterionTypeEnum_AGE_RANGE CriterionTypeEnum_CriterionType = 10
	// Gender.
	CriterionTypeEnum_GENDER CriterionTypeEnum_CriterionType = 11
	// Income Range.
	CriterionTypeEnum_INCOME_RANGE CriterionTypeEnum_CriterionType = 12
	// Parental status.
	CriterionTypeEnum_PARENTAL_STATUS CriterionTypeEnum_CriterionType = 13
	// YouTube Video.
	CriterionTypeEnum_YOUTUBE_VIDEO CriterionTypeEnum_CriterionType = 14
	// YouTube Channel.
	CriterionTypeEnum_YOUTUBE_CHANNEL CriterionTypeEnum_CriterionType = 15
	// User list.
	CriterionTypeEnum_USER_LIST CriterionTypeEnum_CriterionType = 16
	// Proximity.
	CriterionTypeEnum_PROXIMITY CriterionTypeEnum_CriterionType = 17
	// A topic target on the display network (e.g. "Pets & Animals").
	CriterionTypeEnum_TOPIC CriterionTypeEnum_CriterionType = 18
	// Listing scope to target.
	CriterionTypeEnum_LISTING_SCOPE CriterionTypeEnum_CriterionType = 19
	// Language.
	CriterionTypeEnum_LANGUAGE CriterionTypeEnum_CriterionType = 20
	// IpBlock.
	CriterionTypeEnum_IP_BLOCK CriterionTypeEnum_CriterionType = 21
	// Content Label for category exclusion.
	CriterionTypeEnum_CONTENT_LABEL CriterionTypeEnum_CriterionType = 22
	// Carrier.
	CriterionTypeEnum_CARRIER CriterionTypeEnum_CriterionType = 23
	// A category the user is interested in.
	CriterionTypeEnum_USER_INTEREST CriterionTypeEnum_CriterionType = 24
	// Webpage criterion for dynamic search ads.
	CriterionTypeEnum_WEBPAGE CriterionTypeEnum_CriterionType = 25
	// Operating system version.
	CriterionTypeEnum_OPERATING_SYSTEM_VERSION CriterionTypeEnum_CriterionType = 26
	// App payment model.
	CriterionTypeEnum_APP_PAYMENT_MODEL CriterionTypeEnum_CriterionType = 27
	// Mobile device.
	CriterionTypeEnum_MOBILE_DEVICE CriterionTypeEnum_CriterionType = 28
	// Custom affinity.
	CriterionTypeEnum_CUSTOM_AFFINITY CriterionTypeEnum_CriterionType = 29
	// Custom intent.
	CriterionTypeEnum_CUSTOM_INTENT CriterionTypeEnum_CriterionType = 30
)

var CriterionTypeEnum_CriterionType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "KEYWORD",
	3:  "PLACEMENT",
	4:  "MOBILE_APP_CATEGORY",
	5:  "MOBILE_APPLICATION",
	6:  "DEVICE",
	7:  "LOCATION",
	8:  "LISTING_GROUP",
	9:  "AD_SCHEDULE",
	10: "AGE_RANGE",
	11: "GENDER",
	12: "INCOME_RANGE",
	13: "PARENTAL_STATUS",
	14: "YOUTUBE_VIDEO",
	15: "YOUTUBE_CHANNEL",
	16: "USER_LIST",
	17: "PROXIMITY",
	18: "TOPIC",
	19: "LISTING_SCOPE",
	20: "LANGUAGE",
	21: "IP_BLOCK",
	22: "CONTENT_LABEL",
	23: "CARRIER",
	24: "USER_INTEREST",
	25: "WEBPAGE",
	26: "OPERATING_SYSTEM_VERSION",
	27: "APP_PAYMENT_MODEL",
	28: "MOBILE_DEVICE",
	29: "CUSTOM_AFFINITY",
	30: "CUSTOM_INTENT",
}
var CriterionTypeEnum_CriterionType_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"KEYWORD":                  2,
	"PLACEMENT":                3,
	"MOBILE_APP_CATEGORY":      4,
	"MOBILE_APPLICATION":       5,
	"DEVICE":                   6,
	"LOCATION":                 7,
	"LISTING_GROUP":            8,
	"AD_SCHEDULE":              9,
	"AGE_RANGE":                10,
	"GENDER":                   11,
	"INCOME_RANGE":             12,
	"PARENTAL_STATUS":          13,
	"YOUTUBE_VIDEO":            14,
	"YOUTUBE_CHANNEL":          15,
	"USER_LIST":                16,
	"PROXIMITY":                17,
	"TOPIC":                    18,
	"LISTING_SCOPE":            19,
	"LANGUAGE":                 20,
	"IP_BLOCK":                 21,
	"CONTENT_LABEL":            22,
	"CARRIER":                  23,
	"USER_INTEREST":            24,
	"WEBPAGE":                  25,
	"OPERATING_SYSTEM_VERSION": 26,
	"APP_PAYMENT_MODEL":        27,
	"MOBILE_DEVICE":            28,
	"CUSTOM_AFFINITY":          29,
	"CUSTOM_INTENT":            30,
}

func (x CriterionTypeEnum_CriterionType) String() string {
	return proto.EnumName(CriterionTypeEnum_CriterionType_name, int32(x))
}
func (CriterionTypeEnum_CriterionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_criterion_type_70e52fe08b7ff31c, []int{0, 0}
}

// The possible types of a criterion.
type CriterionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CriterionTypeEnum) Reset()         { *m = CriterionTypeEnum{} }
func (m *CriterionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*CriterionTypeEnum) ProtoMessage()    {}
func (*CriterionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_criterion_type_70e52fe08b7ff31c, []int{0}
}
func (m *CriterionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriterionTypeEnum.Unmarshal(m, b)
}
func (m *CriterionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriterionTypeEnum.Marshal(b, m, deterministic)
}
func (dst *CriterionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriterionTypeEnum.Merge(dst, src)
}
func (m *CriterionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_CriterionTypeEnum.Size(m)
}
func (m *CriterionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CriterionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CriterionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CriterionTypeEnum)(nil), "google.ads.googleads.v1.enums.CriterionTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.CriterionTypeEnum_CriterionType", CriterionTypeEnum_CriterionType_name, CriterionTypeEnum_CriterionType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/criterion_type.proto", fileDescriptor_criterion_type_70e52fe08b7ff31c)
}

var fileDescriptor_criterion_type_70e52fe08b7ff31c = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcb, 0x6e, 0xdb, 0x3a,
	0x10, 0xbd, 0x71, 0xde, 0x4c, 0x7c, 0x43, 0x33, 0x37, 0x8f, 0x9b, 0x26, 0x05, 0x92, 0x0f, 0x90,
	0xe1, 0x76, 0xa7, 0xae, 0x28, 0x6a, 0xa2, 0x10, 0x91, 0x48, 0x82, 0xa2, 0x9c, 0xba, 0x30, 0x40,
	0xb8, 0xb1, 0x61, 0x18, 0x48, 0x24, 0xc3, 0x72, 0x02, 0xe4, 0x77, 0xba, 0xcc, 0x57, 0x74, 0xdd,
	0x1f, 0x29, 0xd0, 0xaf, 0x28, 0x28, 0xd9, 0x69, 0xb3, 0x68, 0x37, 0xc2, 0x70, 0xce, 0xcc, 0x99,
	0x33, 0xe4, 0x11, 0x7a, 0x37, 0x2e, 0x8a, 0xf1, 0xdd, 0xa8, 0x3d, 0x18, 0x96, 0xed, 0x3a, 0x74,
	0xd1, 0x63, 0xa7, 0x3d, 0xca, 0x1f, 0xee, 0xcb, 0xf6, 0xed, 0x6c, 0x32, 0x1f, 0xcd, 0x26, 0x45,
	0x6e, 0xe7, 0x4f, 0xd3, 0x91, 0x37, 0x9d, 0x15, 0xf3, 0x82, 0x9c, 0xd5, 0x85, 0xde, 0x60, 0x58,
	0x7a, 0x2f, 0x3d, 0xde, 0x63, 0xc7, 0xab, 0x7a, 0x4e, 0x4e, 0x97, 0x94, 0xd3, 0x49, 0x7b, 0x90,
	0xe7, 0xc5, 0x7c, 0x30, 0x9f, 0x14, 0x79, 0x59, 0x37, 0x5f, 0x7c, 0x5d, 0x43, 0x2d, 0xb6, 0x64,
	0x35, 0x4f, 0xd3, 0x11, 0xe4, 0x0f, 0xf7, 0x17, 0xcf, 0x6b, 0xa8, 0xf9, 0x2a, 0x4b, 0xf6, 0xd0,
	0x4e, 0x26, 0x52, 0x05, 0x8c, 0x5f, 0x72, 0x08, 0xf1, 0x3f, 0x64, 0x07, 0x6d, 0x66, 0xe2, 0x5a,
	0xc8, 0x1b, 0x81, 0x57, 0xdc, 0xe1, 0x1a, 0x7a, 0x37, 0x52, 0x87, 0xb8, 0x41, 0x9a, 0x68, 0x5b,
	0xc5, 0x94, 0x41, 0x02, 0xc2, 0xe0, 0x55, 0x72, 0x84, 0xf6, 0x13, 0x19, 0xf0, 0x18, 0x2c, 0x55,
	0xca, 0x32, 0x6a, 0x20, 0x92, 0xba, 0x87, 0xd7, 0xc8, 0x21, 0x22, 0xbf, 0x80, 0x98, 0x33, 0x6a,
	0xb8, 0x14, 0x78, 0x9d, 0x20, 0xb4, 0x11, 0x42, 0x97, 0x33, 0xc0, 0x1b, 0x64, 0x17, 0x6d, 0xc5,
	0x72, 0x81, 0x6c, 0x92, 0x16, 0x6a, 0xc6, 0x3c, 0x35, 0x5c, 0x44, 0x36, 0xd2, 0x32, 0x53, 0x78,
	0xcb, 0xe9, 0xa2, 0xa1, 0x4d, 0xd9, 0x15, 0x84, 0x59, 0x0c, 0x78, 0xdb, 0x4d, 0xa7, 0x11, 0x58,
	0x4d, 0x45, 0x04, 0x18, 0x39, 0xb2, 0x08, 0x44, 0x08, 0x1a, 0xef, 0x10, 0x8c, 0x76, 0xb9, 0x60,
	0x32, 0x59, 0xa2, 0xbb, 0x64, 0x1f, 0xed, 0x29, 0xaa, 0x41, 0x18, 0x1a, 0xdb, 0xd4, 0x50, 0x93,
	0xa5, 0xb8, 0xe9, 0xa6, 0xf4, 0x64, 0x66, 0xb2, 0x00, 0x6c, 0x97, 0x87, 0x20, 0xf1, 0xbf, 0xae,
	0x6e, 0x99, 0x62, 0x57, 0x54, 0x08, 0x88, 0xf1, 0x9e, 0x9b, 0x94, 0xa5, 0xa0, 0xad, 0x93, 0x84,
	0x71, 0xb5, 0xb6, 0x96, 0x1f, 0x79, 0xc2, 0x4d, 0x0f, 0xb7, 0xc8, 0x36, 0x5a, 0x37, 0x52, 0x71,
	0x86, 0xc9, 0xef, 0xb2, 0x53, 0x26, 0x15, 0xe0, 0xfd, 0x6a, 0x2f, 0x2a, 0xa2, 0x8c, 0x46, 0x80,
	0xff, 0x73, 0x27, 0xae, 0x6c, 0x10, 0x4b, 0x76, 0x8d, 0x0f, 0x5c, 0x39, 0x93, 0xc2, 0x80, 0x30,
	0x36, 0xa6, 0x01, 0xc4, 0xf8, 0xd0, 0xdd, 0x2f, 0xa3, 0x5a, 0x73, 0xd0, 0xf8, 0xc8, 0xe1, 0xd5,
	0x5c, 0x2e, 0x0c, 0x68, 0x48, 0x0d, 0x3e, 0x76, 0xf8, 0x0d, 0x04, 0xca, 0xb1, 0xfd, 0x4f, 0x4e,
	0xd1, 0xb1, 0x54, 0xa0, 0x69, 0x3d, 0xb0, 0x97, 0x1a, 0x48, 0x6c, 0x17, 0x74, 0xea, 0xee, 0xf0,
	0x84, 0x1c, 0xa0, 0x96, 0x7b, 0x07, 0x45, 0x7b, 0xee, 0x7d, 0x6c, 0x22, 0x43, 0x88, 0xf1, 0x1b,
	0x47, 0xba, 0x78, 0x8c, 0xc5, 0xdd, 0x9f, 0xba, 0xa5, 0x59, 0x96, 0x1a, 0x99, 0x58, 0x7a, 0x79,
	0xc9, 0x85, 0x5b, 0xeb, 0xac, 0x12, 0x57, 0x27, 0x79, 0x25, 0x11, 0xbf, 0x0d, 0xbe, 0xaf, 0xa0,
	0xf3, 0xdb, 0xe2, 0xde, 0xfb, 0xab, 0x0d, 0x03, 0xf2, 0xca, 0x4f, 0xca, 0x99, 0x4f, 0xad, 0x7c,
	0x0a, 0x16, 0x4d, 0xe3, 0xe2, 0x6e, 0x90, 0x8f, 0xbd, 0x62, 0x36, 0x6e, 0x8f, 0x47, 0x79, 0x65,
	0xcd, 0xa5, 0xff, 0xa7, 0x93, 0xf2, 0x0f, 0xbf, 0xc3, 0x87, 0xea, 0xfb, 0xa5, 0xb1, 0x1a, 0x51,
	0xfa, 0xdc, 0x38, 0x8b, 0x6a, 0x2a, 0x3a, 0x2c, 0xbd, 0x3a, 0x74, 0x51, 0xb7, 0xe3, 0x39, 0x47,
	0x97, 0xdf, 0x96, 0x78, 0x9f, 0x0e, 0xcb, 0xfe, 0x0b, 0xde, 0xef, 0x76, 0xfa, 0x15, 0xfe, 0xa3,
	0x71, 0x5e, 0x27, 0x7d, 0x9f, 0x0e, 0x4b, 0xdf, 0x7f, 0xa9, 0xf0, 0xfd, 0x6e, 0xc7, 0xf7, 0xab,
	0x9a, 0xcf, 0x1b, 0x95, 0xb0, 0xf7, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x06, 0x96, 0x55,
	0xa6, 0x03, 0x00, 0x00,
}
