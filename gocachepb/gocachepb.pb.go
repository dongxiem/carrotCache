// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gocachepb.proto

package gocachepb

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

// Request包含的一些字段：
// group 和 cache这与我们之前定义的接口 /_geecache/<group>/<name> 所需的参数吻合
type Request struct {
	Group                string   `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_889d0a4ad37a0d42, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Request) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Response struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_889d0a4ad37a0d42, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "gocachepb.Request")
	proto.RegisterType((*Response)(nil), "gocachepb.Response")
}

func init() { proto.RegisterFile("gocachepb.proto", fileDescriptor_889d0a4ad37a0d42) }

var fileDescriptor_889d0a4ad37a0d42 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4f, 0x4d, 0x4d,
	0x4e, 0x4c, 0xce, 0x48, 0x2d, 0x48, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x19, 0x72, 0xb1, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x70, 0xb1, 0xa6,
	0x17, 0xe5, 0x97, 0x16, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x02, 0x5c,
	0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x4c, 0x60, 0x31, 0x10, 0x53, 0x49, 0x81, 0x8b, 0x23, 0x28, 0xb5,
	0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15, 0xa4, 0xa7, 0x2c, 0x31, 0xa7, 0x34, 0x15, 0xac, 0x87, 0x27,
	0x08, 0xc2, 0x31, 0xb2, 0xe3, 0xe2, 0x72, 0x07, 0x69, 0x76, 0x06, 0x59, 0x22, 0x64, 0xc0, 0xc5,
	0xec, 0x9e, 0x5a, 0x22, 0x24, 0xac, 0x87, 0xe4, 0x10, 0xa8, 0x9d, 0x52, 0x22, 0xa8, 0x82, 0x10,
	0x53, 0x93, 0xd8, 0xc0, 0xee, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xd5, 0xdd, 0x09,
	0xbb, 0x00, 0x00, 0x00,
}
