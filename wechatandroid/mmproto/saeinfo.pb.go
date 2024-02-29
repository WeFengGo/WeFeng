// Code generated by protoc-gen-go. DO NOT EDIT.
// source: saeinfo.proto

package mmproto

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

type SaeInfo struct {
	Ver                  *string  `protobuf:"bytes,1,req,name=ver" json:"ver,omitempty"`
	InitKey              []byte   `protobuf:"bytes,2,req,name=initKey" json:"initKey,omitempty"`
	TotalSize            *int32   `protobuf:"varint,3,req,name=totalSize" json:"totalSize,omitempty"`
	XorKey1              []byte   `protobuf:"bytes,9,req,name=xorKey1" json:"xorKey1,omitempty"`
	Key1                 []byte   `protobuf:"bytes,10,req,name=key1" json:"key1,omitempty"`
	XorKey2              []byte   `protobuf:"bytes,11,req,name=xorKey2" json:"xorKey2,omitempty"`
	Key2                 []byte   `protobuf:"bytes,12,req,name=key2" json:"key2,omitempty"`
	Key3                 []byte   `protobuf:"bytes,18,req,name=key3" json:"key3,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaeInfo) Reset()         { *m = SaeInfo{} }
func (m *SaeInfo) String() string { return proto.CompactTextString(m) }
func (*SaeInfo) ProtoMessage()    {}
func (*SaeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_35810a54fd6e608d, []int{0}
}

func (m *SaeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaeInfo.Unmarshal(m, b)
}
func (m *SaeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaeInfo.Marshal(b, m, deterministic)
}
func (m *SaeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaeInfo.Merge(m, src)
}
func (m *SaeInfo) XXX_Size() int {
	return xxx_messageInfo_SaeInfo.Size(m)
}
func (m *SaeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SaeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SaeInfo proto.InternalMessageInfo

func (m *SaeInfo) GetVer() string {
	if m != nil && m.Ver != nil {
		return *m.Ver
	}
	return ""
}

func (m *SaeInfo) GetInitKey() []byte {
	if m != nil {
		return m.InitKey
	}
	return nil
}

func (m *SaeInfo) GetTotalSize() int32 {
	if m != nil && m.TotalSize != nil {
		return *m.TotalSize
	}
	return 0
}

func (m *SaeInfo) GetXorKey1() []byte {
	if m != nil {
		return m.XorKey1
	}
	return nil
}

func (m *SaeInfo) GetKey1() []byte {
	if m != nil {
		return m.Key1
	}
	return nil
}

func (m *SaeInfo) GetXorKey2() []byte {
	if m != nil {
		return m.XorKey2
	}
	return nil
}

func (m *SaeInfo) GetKey2() []byte {
	if m != nil {
		return m.Key2
	}
	return nil
}

func (m *SaeInfo) GetKey3() []byte {
	if m != nil {
		return m.Key3
	}
	return nil
}

func init() {
	proto.RegisterType((*SaeInfo)(nil), "mmproto.SaeInfo")
}

func init() {
	proto.RegisterFile("saeinfo.proto", fileDescriptor_35810a54fd6e608d)
}

var fileDescriptor_35810a54fd6e608d = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x4c, 0xcd,
	0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0xcd, 0x05, 0x33, 0x94,
	0x0e, 0x33, 0x72, 0xb1, 0x07, 0x27, 0xa6, 0x7a, 0xe6, 0xa5, 0xe5, 0x0b, 0x09, 0x70, 0x31, 0x97,
	0xa5, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x69, 0x70, 0x06, 0x81, 0x98, 0x42, 0x12, 0x5c, 0xec, 0x99,
	0x79, 0x99, 0x25, 0xde, 0xa9, 0x95, 0x12, 0x4c, 0x0a, 0x4c, 0x1a, 0x3c, 0x41, 0x30, 0xae, 0x90,
	0x0c, 0x17, 0x67, 0x49, 0x7e, 0x49, 0x62, 0x4e, 0x70, 0x66, 0x55, 0xaa, 0x04, 0xb3, 0x02, 0x93,
	0x06, 0x6b, 0x10, 0x42, 0x00, 0xa4, 0xaf, 0x22, 0xbf, 0xc8, 0x3b, 0xb5, 0xd2, 0x50, 0x82, 0x13,
	0xa2, 0x0f, 0xca, 0x15, 0x12, 0xe2, 0x62, 0xc9, 0x06, 0x09, 0x73, 0x81, 0x85, 0xc1, 0x6c, 0x84,
	0x6a, 0x23, 0x09, 0x6e, 0x64, 0xd5, 0x46, 0x50, 0xd5, 0x46, 0x12, 0x3c, 0x70, 0xd5, 0x30, 0x31,
	0x63, 0x09, 0x21, 0xb8, 0x98, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x7e, 0xa4, 0xed, 0xde,
	0x00, 0x00, 0x00,
}
