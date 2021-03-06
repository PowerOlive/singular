// Code generated by protoc-gen-go.
// source: singular.proto
// DO NOT EDIT!

/*
Package singular is a generated protocol buffer package.

It is generated from these files:
	singular.proto

It has these top-level messages:
	Request
	Data
*/
package singular

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request_Meta int32

const (
	Request_NewProxy Request_Meta = 0
	Request_Assign   Request_Meta = 1
)

var Request_Meta_name = map[int32]string{
	0: "NewProxy",
	1: "Assign",
}
var Request_Meta_value = map[string]int32{
	"NewProxy": 0,
	"Assign":   1,
}

func (x Request_Meta) String() string {
	return proto.EnumName(Request_Meta_name, int32(x))
}
func (Request_Meta) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Request struct {
	Meta    Request_Meta `protobuf:"varint,1,opt,name=meta,enum=Request_Meta" json:"meta,omitempty"`
	Payload string       `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Data struct {
	RequestId []byte `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Payload   []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Data)(nil), "Data")
	proto.RegisterEnum("Request_Meta", Request_Meta_name, Request_Meta_value)
}

func init() { proto.RegisterFile("singular.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xce, 0xcc, 0x4b,
	0x2f, 0xcd, 0x49, 0x2c, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xca, 0xe0, 0x62, 0x0f, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x52, 0xe4, 0x62, 0xc9, 0x4d, 0x2d, 0x49, 0x94, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x33, 0xe2, 0xd5, 0x83, 0x8a, 0xeb, 0xf9, 0x02, 0x05, 0x83, 0xc0, 0x52, 0x42,
	0x12, 0x5c, 0xec, 0x05, 0x89, 0x95, 0x39, 0xf9, 0x89, 0x29, 0x12, 0x4c, 0x40, 0x55, 0x9c, 0x41,
	0x30, 0xae, 0x92, 0x02, 0x17, 0x0b, 0x48, 0x9d, 0x10, 0x0f, 0x17, 0x87, 0x5f, 0x6a, 0x79, 0x40,
	0x51, 0x7e, 0x45, 0xa5, 0x00, 0x83, 0x10, 0x17, 0x17, 0x9b, 0x63, 0x71, 0x71, 0x66, 0x7a, 0x9e,
	0x00, 0xa3, 0x92, 0x1d, 0x17, 0x8b, 0x4b, 0x22, 0x50, 0x85, 0x0c, 0x17, 0x67, 0x11, 0xc4, 0x64,
	0xcf, 0x14, 0xb0, 0x5d, 0x3c, 0x41, 0x08, 0x01, 0x74, 0x1b, 0x78, 0xe0, 0x36, 0x24, 0xb1, 0x81,
	0x1d, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x78, 0x40, 0xd7, 0x49, 0xc2, 0x00, 0x00, 0x00,
}
