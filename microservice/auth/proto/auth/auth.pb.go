// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package mu_micro_book_srv_auth

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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type Request struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
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

func (m *Request) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Request) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Request) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
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

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "mu.micro.book.srv.auth.Error")
	proto.RegisterType((*Request)(nil), "mu.micro.book.srv.auth.Request")
	proto.RegisterType((*Response)(nil), "mu.micro.book.srv.auth.Response")
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5) }

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xcf, 0x4a, 0xf3, 0x50,
	0x10, 0xc5, 0xc9, 0xf7, 0x35, 0x4d, 0x1c, 0x17, 0xc2, 0xa0, 0x25, 0x14, 0xc4, 0x90, 0x55, 0x57,
	0x57, 0x68, 0x9e, 0x40, 0xd0, 0x85, 0x0b, 0x5d, 0xdc, 0x56, 0x5c, 0xa7, 0xc9, 0x80, 0x21, 0x4d,
	0xa7, 0xde, 0x3f, 0x7d, 0x36, 0x1f, 0x4f, 0xee, 0xf4, 0xea, 0xaa, 0xdd, 0xb9, 0x09, 0xf3, 0xe3,
	0x9c, 0x39, 0x93, 0xc3, 0x85, 0x9b, 0xbd, 0x61, 0xc7, 0xf7, 0x8d, 0x77, 0x1f, 0xf2, 0x51, 0xc2,
	0x38, 0x1b, 0xbd, 0x1a, 0xfb, 0xd6, 0xb0, 0xda, 0x30, 0x0f, 0xca, 0x9a, 0x83, 0x0a, 0x6a, 0x55,
	0x43, 0xfa, 0x64, 0x0c, 0x1b, 0x44, 0x98, 0xb4, 0xdc, 0x51, 0x91, 0x94, 0xc9, 0x22, 0xd5, 0x32,
	0xe3, 0x0c, 0xa6, 0x1d, 0xb9, 0xa6, 0xdf, 0x16, 0xff, 0xca, 0x64, 0x71, 0xa1, 0x23, 0x55, 0x2b,
	0xc8, 0x34, 0x7d, 0x7a, 0xb2, 0x2e, 0x58, 0xbc, 0x25, 0xf3, 0xdc, 0xc9, 0xe2, 0x44, 0x47, 0xc2,
	0x39, 0xe4, 0x61, 0x7a, 0x6d, 0x46, 0x8a, 0xcb, 0xbf, 0x8c, 0xd7, 0x90, 0x3a, 0x1e, 0x68, 0x57,
	0xfc, 0x17, 0xe1, 0x08, 0x15, 0x43, 0xae, 0xc9, 0xee, 0x79, 0x67, 0x09, 0x0b, 0xc8, 0xac, 0x6f,
	0x5b, 0xb2, 0x56, 0x62, 0x73, 0xfd, 0x83, 0x58, 0x43, 0x4a, 0xe1, 0x7f, 0x25, 0xf4, 0x72, 0x79,
	0xab, 0x4e, 0xf7, 0x52, 0x52, 0x4a, 0x1f, 0xbd, 0xa7, 0x0f, 0x2e, 0xbf, 0x12, 0xc8, 0x56, 0x64,
	0x0e, 0x7d, 0x4b, 0xb8, 0x86, 0xab, 0x97, 0x66, 0xa0, 0x07, 0x39, 0xb2, 0x0e, 0x32, 0xde, 0x9d,
	0x8b, 0x8e, 0xd5, 0xe7, 0xe5, 0x79, 0x43, 0xac, 0xf1, 0x0e, 0xf8, 0x48, 0xdb, 0x37, 0x4b, 0xe6,
	0x6f, 0x83, 0x37, 0x53, 0x79, 0xd4, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x32, 0xc1, 0xfc, 0x32,
	0xed, 0x01, 0x00, 0x00,
}
