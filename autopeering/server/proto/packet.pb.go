// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/proto/packet.proto

package proto

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

type Packet struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_c39379c1dd924f72, []int{0}
}

func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Packet) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Packet) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Packet)(nil), "proto.Packet")
}

func init() { proto.RegisterFile("server/proto/packet.proto", fileDescriptor_c39379c1dd924f72) }

var fileDescriptor_c39379c1dd924f72 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0xd1,
	0x03, 0x73, 0x84, 0x58, 0xc1, 0x94, 0x52, 0x24, 0x17, 0x5b, 0x00, 0x58, 0x58, 0x48, 0x96, 0x8b,
	0xab, 0xa0, 0x34, 0x29, 0x27, 0x33, 0x39, 0x3e, 0x3b, 0xb5, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x27, 0x88, 0x13, 0x22, 0xe2, 0x9d, 0x5a, 0x29, 0x24, 0xc3, 0xc5, 0x59, 0x9c, 0x99, 0x9e, 0x97,
	0x58, 0x52, 0x5a, 0x94, 0x2a, 0xc1, 0x04, 0x91, 0x85, 0x0b, 0x08, 0x09, 0x71, 0xb1, 0xa4, 0x24,
	0x96, 0x24, 0x4a, 0x30, 0x83, 0x25, 0xc0, 0x6c, 0x27, 0x93, 0x28, 0xa3, 0xf4, 0xcc, 0x92, 0x8c,
	0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xcc, 0xfc, 0x92, 0xc4, 0x9c, 0xd4, 0x94, 0x74, 0xa0,
	0x6b, 0x12, 0x4b, 0x4b, 0xf2, 0x0b, 0x52, 0x53, 0x8b, 0x32, 0xf3, 0xd2, 0x75, 0x8b, 0x33, 0x73,
	0xf5, 0x91, 0x1d, 0x99, 0xc4, 0x06, 0xa6, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x50, 0x6a,
	0x7d, 0xb7, 0xbb, 0x00, 0x00, 0x00,
}
