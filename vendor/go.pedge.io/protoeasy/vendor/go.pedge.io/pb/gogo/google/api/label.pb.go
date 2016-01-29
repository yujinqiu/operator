// Code generated by protoc-gen-gogo.
// source: google/api/label.proto
// DO NOT EDIT!

package google_api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Value types that can be used as label values.
type LabelDescriptor_ValueType int32

const (
	// A variable-length string. This is the default.
	LabelDescriptor_STRING LabelDescriptor_ValueType = 0
	// Boolean; true or false.
	LabelDescriptor_BOOL LabelDescriptor_ValueType = 1
	// A 64-bit signed integer.
	LabelDescriptor_INT64 LabelDescriptor_ValueType = 2
)

var LabelDescriptor_ValueType_name = map[int32]string{
	0: "STRING",
	1: "BOOL",
	2: "INT64",
}
var LabelDescriptor_ValueType_value = map[string]int32{
	"STRING": 0,
	"BOOL":   1,
	"INT64":  2,
}

func (x LabelDescriptor_ValueType) String() string {
	return proto.EnumName(LabelDescriptor_ValueType_name, int32(x))
}

// A description of a label.
type LabelDescriptor struct {
	// The label key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The type of data that can be assigned to the label.
	ValueType LabelDescriptor_ValueType `protobuf:"varint,2,opt,name=value_type,proto3,enum=google.api.LabelDescriptor_ValueType" json:"value_type,omitempty"`
	// A human-readable description for the label.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *LabelDescriptor) Reset()         { *m = LabelDescriptor{} }
func (m *LabelDescriptor) String() string { return proto.CompactTextString(m) }
func (*LabelDescriptor) ProtoMessage()    {}

func init() {
	proto.RegisterType((*LabelDescriptor)(nil), "google.api.LabelDescriptor")
	proto.RegisterEnum("google.api.LabelDescriptor_ValueType", LabelDescriptor_ValueType_name, LabelDescriptor_ValueType_value)
}