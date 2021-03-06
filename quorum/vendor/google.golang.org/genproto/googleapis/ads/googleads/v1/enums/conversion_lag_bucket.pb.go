// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/conversion_lag_bucket.proto

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

// Enum representing the number of days between impression and conversion.
type ConversionLagBucketEnum_ConversionLagBucket int32

const (
	// Not specified.
	ConversionLagBucketEnum_UNSPECIFIED ConversionLagBucketEnum_ConversionLagBucket = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionLagBucketEnum_UNKNOWN ConversionLagBucketEnum_ConversionLagBucket = 1
	// Conversion lag bucket from 0 to 1 day. 0 day is included, 1 day is not.
	ConversionLagBucketEnum_LESS_THAN_ONE_DAY ConversionLagBucketEnum_ConversionLagBucket = 2
	// Conversion lag bucket from 1 to 2 days. 1 day is included, 2 days is not.
	ConversionLagBucketEnum_ONE_TO_TWO_DAYS ConversionLagBucketEnum_ConversionLagBucket = 3
	// Conversion lag bucket from 2 to 3 days. 2 days is included,
	// 3 days is not.
	ConversionLagBucketEnum_TWO_TO_THREE_DAYS ConversionLagBucketEnum_ConversionLagBucket = 4
	// Conversion lag bucket from 3 to 4 days. 3 days is included,
	// 4 days is not.
	ConversionLagBucketEnum_THREE_TO_FOUR_DAYS ConversionLagBucketEnum_ConversionLagBucket = 5
	// Conversion lag bucket from 4 to 5 days. 4 days is included,
	// 5 days is not.
	ConversionLagBucketEnum_FOUR_TO_FIVE_DAYS ConversionLagBucketEnum_ConversionLagBucket = 6
	// Conversion lag bucket from 5 to 6 days. 5 days is included,
	// 6 days is not.
	ConversionLagBucketEnum_FIVE_TO_SIX_DAYS ConversionLagBucketEnum_ConversionLagBucket = 7
	// Conversion lag bucket from 6 to 7 days. 6 days is included,
	// 7 days is not.
	ConversionLagBucketEnum_SIX_TO_SEVEN_DAYS ConversionLagBucketEnum_ConversionLagBucket = 8
	// Conversion lag bucket from 7 to 8 days. 7 days is included,
	// 8 days is not.
	ConversionLagBucketEnum_SEVEN_TO_EIGHT_DAYS ConversionLagBucketEnum_ConversionLagBucket = 9
	// Conversion lag bucket from 8 to 9 days. 8 days is included,
	// 9 days is not.
	ConversionLagBucketEnum_EIGHT_TO_NINE_DAYS ConversionLagBucketEnum_ConversionLagBucket = 10
	// Conversion lag bucket from 9 to 10 days. 9 days is included,
	// 10 days is not.
	ConversionLagBucketEnum_NINE_TO_TEN_DAYS ConversionLagBucketEnum_ConversionLagBucket = 11
	// Conversion lag bucket from 10 to 11 days. 10 days is included,
	// 11 days is not.
	ConversionLagBucketEnum_TEN_TO_ELEVEN_DAYS ConversionLagBucketEnum_ConversionLagBucket = 12
	// Conversion lag bucket from 11 to 12 days. 11 days is included,
	// 12 days is not.
	ConversionLagBucketEnum_ELEVEN_TO_TWELVE_DAYS ConversionLagBucketEnum_ConversionLagBucket = 13
	// Conversion lag bucket from 12 to 13 days. 12 days is included,
	// 13 days is not.
	ConversionLagBucketEnum_TWELVE_TO_THIRTEEN_DAYS ConversionLagBucketEnum_ConversionLagBucket = 14
	// Conversion lag bucket from 13 to 14 days. 13 days is included,
	// 14 days is not.
	ConversionLagBucketEnum_THIRTEEN_TO_FOURTEEN_DAYS ConversionLagBucketEnum_ConversionLagBucket = 15
	// Conversion lag bucket from 14 to 21 days. 14 days is included,
	// 21 days is not.
	ConversionLagBucketEnum_FOURTEEN_TO_TWENTY_ONE_DAYS ConversionLagBucketEnum_ConversionLagBucket = 16
	// Conversion lag bucket from 21 to 30 days. 21 days is included,
	// 30 days is not.
	ConversionLagBucketEnum_TWENTY_ONE_TO_THIRTY_DAYS ConversionLagBucketEnum_ConversionLagBucket = 17
	// Conversion lag bucket from 30 to 45 days. 30 days is included,
	// 45 days is not.
	ConversionLagBucketEnum_THIRTY_TO_FORTY_FIVE_DAYS ConversionLagBucketEnum_ConversionLagBucket = 18
	// Conversion lag bucket from 45 to 60 days. 45 days is included,
	// 60 days is not.
	ConversionLagBucketEnum_FORTY_FIVE_TO_SIXTY_DAYS ConversionLagBucketEnum_ConversionLagBucket = 19
	// Conversion lag bucket from 60 to 90 days. 60 days is included,
	// 90 days is not.
	ConversionLagBucketEnum_SIXTY_TO_NINETY_DAYS ConversionLagBucketEnum_ConversionLagBucket = 20
)

var ConversionLagBucketEnum_ConversionLagBucket_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "LESS_THAN_ONE_DAY",
	3:  "ONE_TO_TWO_DAYS",
	4:  "TWO_TO_THREE_DAYS",
	5:  "THREE_TO_FOUR_DAYS",
	6:  "FOUR_TO_FIVE_DAYS",
	7:  "FIVE_TO_SIX_DAYS",
	8:  "SIX_TO_SEVEN_DAYS",
	9:  "SEVEN_TO_EIGHT_DAYS",
	10: "EIGHT_TO_NINE_DAYS",
	11: "NINE_TO_TEN_DAYS",
	12: "TEN_TO_ELEVEN_DAYS",
	13: "ELEVEN_TO_TWELVE_DAYS",
	14: "TWELVE_TO_THIRTEEN_DAYS",
	15: "THIRTEEN_TO_FOURTEEN_DAYS",
	16: "FOURTEEN_TO_TWENTY_ONE_DAYS",
	17: "TWENTY_ONE_TO_THIRTY_DAYS",
	18: "THIRTY_TO_FORTY_FIVE_DAYS",
	19: "FORTY_FIVE_TO_SIXTY_DAYS",
	20: "SIXTY_TO_NINETY_DAYS",
}
var ConversionLagBucketEnum_ConversionLagBucket_value = map[string]int32{
	"UNSPECIFIED":                 0,
	"UNKNOWN":                     1,
	"LESS_THAN_ONE_DAY":           2,
	"ONE_TO_TWO_DAYS":             3,
	"TWO_TO_THREE_DAYS":           4,
	"THREE_TO_FOUR_DAYS":          5,
	"FOUR_TO_FIVE_DAYS":           6,
	"FIVE_TO_SIX_DAYS":            7,
	"SIX_TO_SEVEN_DAYS":           8,
	"SEVEN_TO_EIGHT_DAYS":         9,
	"EIGHT_TO_NINE_DAYS":          10,
	"NINE_TO_TEN_DAYS":            11,
	"TEN_TO_ELEVEN_DAYS":          12,
	"ELEVEN_TO_TWELVE_DAYS":       13,
	"TWELVE_TO_THIRTEEN_DAYS":     14,
	"THIRTEEN_TO_FOURTEEN_DAYS":   15,
	"FOURTEEN_TO_TWENTY_ONE_DAYS": 16,
	"TWENTY_ONE_TO_THIRTY_DAYS":   17,
	"THIRTY_TO_FORTY_FIVE_DAYS":   18,
	"FORTY_FIVE_TO_SIXTY_DAYS":    19,
	"SIXTY_TO_NINETY_DAYS":        20,
}

func (x ConversionLagBucketEnum_ConversionLagBucket) String() string {
	return proto.EnumName(ConversionLagBucketEnum_ConversionLagBucket_name, int32(x))
}
func (ConversionLagBucketEnum_ConversionLagBucket) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_conversion_lag_bucket_3d9824be1641c408, []int{0, 0}
}

// Container for enum representing the number of days between impression and
// conversion.
type ConversionLagBucketEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionLagBucketEnum) Reset()         { *m = ConversionLagBucketEnum{} }
func (m *ConversionLagBucketEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionLagBucketEnum) ProtoMessage()    {}
func (*ConversionLagBucketEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_lag_bucket_3d9824be1641c408, []int{0}
}
func (m *ConversionLagBucketEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionLagBucketEnum.Unmarshal(m, b)
}
func (m *ConversionLagBucketEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionLagBucketEnum.Marshal(b, m, deterministic)
}
func (dst *ConversionLagBucketEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionLagBucketEnum.Merge(dst, src)
}
func (m *ConversionLagBucketEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionLagBucketEnum.Size(m)
}
func (m *ConversionLagBucketEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionLagBucketEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionLagBucketEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConversionLagBucketEnum)(nil), "google.ads.googleads.v1.enums.ConversionLagBucketEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.ConversionLagBucketEnum_ConversionLagBucket", ConversionLagBucketEnum_ConversionLagBucket_name, ConversionLagBucketEnum_ConversionLagBucket_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/conversion_lag_bucket.proto", fileDescriptor_conversion_lag_bucket_3d9824be1641c408)
}

var fileDescriptor_conversion_lag_bucket_3d9824be1641c408 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xd1, 0x6e, 0xda, 0x30,
	0x14, 0x86, 0x07, 0x65, 0xed, 0x66, 0xb6, 0x91, 0x1a, 0x3a, 0xe8, 0x5a, 0x34, 0xb5, 0x0f, 0x90,
	0x08, 0xed, 0x6a, 0xd9, 0x55, 0x68, 0x5d, 0x88, 0x86, 0x1c, 0x44, 0x0c, 0x1d, 0x13, 0x92, 0x95,
	0x92, 0x28, 0x42, 0x83, 0x18, 0x61, 0xe0, 0x75, 0x26, 0xed, 0x72, 0xaf, 0xb0, 0x37, 0xd8, 0xa3,
	0x4c, 0xda, 0x3b, 0x4c, 0xf6, 0x89, 0xc3, 0x2e, 0xba, 0xdd, 0x44, 0xe7, 0xff, 0xce, 0x39, 0x7f,
	0xec, 0xa3, 0x63, 0xf4, 0x3e, 0x15, 0x22, 0x5d, 0x26, 0x4e, 0x14, 0x4b, 0x07, 0x42, 0x15, 0xed,
	0x3b, 0x4e, 0x92, 0xed, 0x56, 0xd2, 0x99, 0x8b, 0x6c, 0x9f, 0x6c, 0xe4, 0x42, 0x64, 0x7c, 0x19,
	0xa5, 0xfc, 0x61, 0x37, 0xff, 0x92, 0x6c, 0xed, 0xf5, 0x46, 0x6c, 0x05, 0x6e, 0x43, 0xbd, 0x1d,
	0xc5, 0xd2, 0x2e, 0x5a, 0xed, 0x7d, 0xc7, 0xd6, 0xad, 0x6f, 0x2e, 0x8d, 0xf3, 0x7a, 0xe1, 0x44,
	0x59, 0x26, 0xb6, 0xd1, 0x76, 0x21, 0x32, 0x09, 0xcd, 0xd7, 0x3f, 0x2a, 0xa8, 0x79, 0x53, 0x98,
	0x0f, 0xa2, 0xb4, 0xab, 0xad, 0x49, 0xb6, 0x5b, 0x5d, 0x7f, 0xad, 0xa0, 0xfa, 0x23, 0x39, 0x5c,
	0x43, 0xd5, 0x31, 0x0d, 0x87, 0xe4, 0xc6, 0xbf, 0xf3, 0xc9, 0xad, 0xf5, 0x04, 0x57, 0xd1, 0xc9,
	0x98, 0x7e, 0xa4, 0xc1, 0x3d, 0xb5, 0x4a, 0xf8, 0x0c, 0x9d, 0x0e, 0x48, 0x18, 0x72, 0xd6, 0xf7,
	0x28, 0x0f, 0x28, 0xe1, 0xb7, 0xde, 0xd4, 0x2a, 0xe3, 0x3a, 0xaa, 0x29, 0xc1, 0x02, 0xce, 0xee,
	0x03, 0xc5, 0x42, 0xeb, 0x48, 0xd5, 0x2a, 0xa5, 0x60, 0x7f, 0x44, 0x08, 0xe0, 0x0a, 0x7e, 0x8d,
	0x30, 0x68, 0x16, 0xf0, 0xbb, 0x60, 0x3c, 0x02, 0xfe, 0x54, 0x95, 0x6b, 0xa9, 0xb0, 0x3f, 0xc9,
	0xcb, 0x8f, 0x71, 0x03, 0x59, 0x5a, 0xb2, 0x80, 0x87, 0xfe, 0x27, 0xa0, 0x27, 0xaa, 0x58, 0x29,
	0x05, 0xc9, 0x84, 0x50, 0xc0, 0xcf, 0x70, 0x13, 0xd5, 0x41, 0xb3, 0x80, 0x13, 0xbf, 0xd7, 0x67,
	0x90, 0x78, 0xae, 0x7e, 0x0a, 0x9a, 0x05, 0x9c, 0xfa, 0x34, 0x77, 0x47, 0xca, 0x5d, 0x4b, 0x75,
	0x48, 0x63, 0x53, 0xd5, 0x47, 0xcc, 0x4d, 0x06, 0x07, 0xfb, 0x17, 0xf8, 0x1c, 0x9d, 0xe5, 0x40,
	0xdf, 0x94, 0x0c, 0xcc, 0x31, 0x5f, 0xe2, 0x0b, 0xd4, 0xcc, 0x81, 0xbe, 0xaf, 0x3f, 0x62, 0xc4,
	0xf4, 0xbd, 0xc2, 0x6d, 0x74, 0x5e, 0xa0, 0xfc, 0xd6, 0x87, 0x74, 0x0d, 0xbf, 0x45, 0x17, 0x05,
	0x02, 0x63, 0xca, 0xa6, 0x66, 0xba, 0xa1, 0x65, 0xe9, 0xfe, 0x03, 0x34, 0x3f, 0x98, 0x42, 0xfa,
	0xb4, 0xb0, 0x9f, 0x82, 0xb9, 0x0a, 0x0e, 0x13, 0xc4, 0xf8, 0x12, 0xb5, 0xfe, 0x82, 0x30, 0x47,
	0xd3, 0x5c, 0xc7, 0x2d, 0xd4, 0x00, 0x9d, 0x4f, 0xc6, 0x64, 0x1a, 0xdd, 0xdf, 0x25, 0x74, 0x35,
	0x17, 0x2b, 0xfb, 0xbf, 0x1b, 0xd8, 0x6d, 0x3d, 0xb2, 0x44, 0x43, 0xb5, 0x7d, 0xc3, 0xd2, 0xe7,
	0x6e, 0xde, 0x9a, 0x8a, 0x65, 0x94, 0xa5, 0xb6, 0xd8, 0xa4, 0x4e, 0x9a, 0x64, 0x7a, 0x37, 0xcd,
	0x3b, 0x58, 0x2f, 0xe4, 0x3f, 0x9e, 0xc5, 0x07, 0xfd, 0xfd, 0x56, 0x3e, 0xea, 0x79, 0xde, 0xf7,
	0x72, 0xbb, 0x07, 0x56, 0x5e, 0x2c, 0x6d, 0x08, 0x55, 0x34, 0xe9, 0xd8, 0x6a, 0x99, 0xe5, 0x4f,
	0x93, 0x9f, 0x79, 0xb1, 0x9c, 0x15, 0xf9, 0xd9, 0xa4, 0x33, 0xd3, 0xf9, 0x5f, 0xe5, 0x2b, 0x80,
	0xae, 0xeb, 0xc5, 0xd2, 0x75, 0x8b, 0x0a, 0xd7, 0x9d, 0x74, 0x5c, 0x57, 0xd7, 0x3c, 0x1c, 0xeb,
	0x83, 0xbd, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x29, 0xdd, 0xa7, 0xb6, 0xae, 0x03, 0x00, 0x00,
}
