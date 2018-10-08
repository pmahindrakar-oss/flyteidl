// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/hive.proto

package plugins // import "github.com/lyft/flyteidl/gen/pb-go/plugins"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Defines a query to execute on a hive cluster.
type HiveQuery struct {
	// The query string.
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HiveQuery) Reset()         { *m = HiveQuery{} }
func (m *HiveQuery) String() string { return proto.CompactTextString(m) }
func (*HiveQuery) ProtoMessage()    {}
func (*HiveQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_hive_8a0071a286958895, []int{0}
}
func (m *HiveQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiveQuery.Unmarshal(m, b)
}
func (m *HiveQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiveQuery.Marshal(b, m, deterministic)
}
func (dst *HiveQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiveQuery.Merge(dst, src)
}
func (m *HiveQuery) XXX_Size() int {
	return xxx_messageInfo_HiveQuery.Size(m)
}
func (m *HiveQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_HiveQuery.DiscardUnknown(m)
}

var xxx_messageInfo_HiveQuery proto.InternalMessageInfo

func (m *HiveQuery) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

// Defines a collection of hive queries.
type HiveQueryCollection struct {
	Queries              []*HiveQuery `protobuf:"bytes,2,rep,name=queries,proto3" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HiveQueryCollection) Reset()         { *m = HiveQueryCollection{} }
func (m *HiveQueryCollection) String() string { return proto.CompactTextString(m) }
func (*HiveQueryCollection) ProtoMessage()    {}
func (*HiveQueryCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_hive_8a0071a286958895, []int{1}
}
func (m *HiveQueryCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiveQueryCollection.Unmarshal(m, b)
}
func (m *HiveQueryCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiveQueryCollection.Marshal(b, m, deterministic)
}
func (dst *HiveQueryCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiveQueryCollection.Merge(dst, src)
}
func (m *HiveQueryCollection) XXX_Size() int {
	return xxx_messageInfo_HiveQueryCollection.Size(m)
}
func (m *HiveQueryCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_HiveQueryCollection.DiscardUnknown(m)
}

var xxx_messageInfo_HiveQueryCollection proto.InternalMessageInfo

func (m *HiveQueryCollection) GetQueries() []*HiveQuery {
	if m != nil {
		return m.Queries
	}
	return nil
}

func init() {
	proto.RegisterType((*HiveQuery)(nil), "flyteidl.plugins.HiveQuery")
	proto.RegisterType((*HiveQueryCollection)(nil), "flyteidl.plugins.HiveQueryCollection")
}

func init() { proto.RegisterFile("flyteidl/plugins/hive.proto", fileDescriptor_hive_8a0071a286958895) }

var fileDescriptor_hive_8a0071a286958895 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xcb, 0xa9, 0x2c,
	0x49, 0xcd, 0x4c, 0xc9, 0xd1, 0x2f, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x2b, 0xd6, 0xcf, 0xc8, 0x2c,
	0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0x49, 0xea, 0x41, 0x25, 0xa5, 0x24,
	0xe1, 0xca, 0x93, 0xf3, 0x8b, 0x52, 0xf5, 0x4b, 0x12, 0x8b, 0xb3, 0x8b, 0x21, 0x8a, 0x95, 0x14,
	0xb9, 0x38, 0x3d, 0x32, 0xcb, 0x52, 0x03, 0x4b, 0x53, 0x8b, 0x2a, 0x85, 0x44, 0xb8, 0x58, 0x0b,
	0x41, 0x0c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0xc9, 0x87, 0x4b, 0x18, 0xae,
	0xc4, 0x39, 0x3f, 0x27, 0x27, 0x35, 0xb9, 0x24, 0x33, 0x3f, 0x4f, 0xc8, 0x94, 0x8b, 0x1d, 0x24,
	0x9f, 0x99, 0x5a, 0x2c, 0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x6d, 0x24, 0xad, 0x87, 0x6e, 0xb1, 0x1e,
	0x5c, 0x5f, 0x10, 0x4c, 0xad, 0x93, 0x4e, 0x94, 0x56, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
	0x72, 0x7e, 0xae, 0x7e, 0x4e, 0x65, 0x5a, 0x89, 0x3e, 0xdc, 0x75, 0xe9, 0xa9, 0x79, 0xfa, 0x05,
	0x49, 0xba, 0xe9, 0xf9, 0x30, 0x6f, 0x25, 0xb1, 0x81, 0x5d, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x6f, 0xb7, 0x64, 0xdd, 0xf1, 0x00, 0x00, 0x00,
}