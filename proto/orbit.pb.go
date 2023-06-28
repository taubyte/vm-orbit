// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: orbit.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Type int32

const (
	Type_unknown Type = 0
	Type_i32     Type = 127
	Type_i64     Type = 126
	Type_f32     Type = 125
	Type_f64     Type = 124
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0:   "unknown",
		127: "i32",
		126: "i64",
		125: "f32",
		124: "f64",
	}
	Type_value = map[string]int32{
		"unknown": 0,
		"i32":     127,
		"i64":     126,
		"f32":     125,
		"f64":     124,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_orbit_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_orbit_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_orbit_proto_rawDescGZIP(), []int{0}
}

type IOError int32

const (
	IOError_none         IOError = 0
	IOError_shortWrite   IOError = 16
	IOError_invalidWrite IOError = 17
	IOError_shortBuffer  IOError = 18
	IOError_eof          IOError = 19
	IOError_noProgress   IOError = 20
)

// Enum value maps for IOError.
var (
	IOError_name = map[int32]string{
		0:  "none",
		16: "shortWrite",
		17: "invalidWrite",
		18: "shortBuffer",
		19: "eof",
		20: "noProgress",
	}
	IOError_value = map[string]int32{
		"none":         0,
		"shortWrite":   16,
		"invalidWrite": 17,
		"shortBuffer":  18,
		"eof":          19,
		"noProgress":   20,
	}
)

func (x IOError) Enum() *IOError {
	p := new(IOError)
	*p = x
	return p
}

func (x IOError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IOError) Descriptor() protoreflect.EnumDescriptor {
	return file_orbit_proto_enumTypes[1].Descriptor()
}

func (IOError) Type() protoreflect.EnumType {
	return &file_orbit_proto_enumTypes[1]
}

func (x IOError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IOError.Descriptor instead.
func (IOError) EnumDescriptor() ([]byte, []int) {
	return file_orbit_proto_rawDescGZIP(), []int{1}
}

type FunctionDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Args []Type `protobuf:"varint,2,rep,packed,name=args,proto3,enum=proto.Type" json:"args,omitempty"`
	Rets []Type `protobuf:"varint,3,rep,packed,name=rets,proto3,enum=proto.Type" json:"rets,omitempty"`
}

func (x *FunctionDefinition) Reset() {
	*x = FunctionDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionDefinition) ProtoMessage() {}

func (x *FunctionDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_orbit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionDefinition.ProtoReflect.Descriptor instead.
func (*FunctionDefinition) Descriptor() ([]byte, []int) {
	return file_orbit_proto_rawDescGZIP(), []int{0}
}

func (x *FunctionDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FunctionDefinition) GetArgs() []Type {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *FunctionDefinition) GetRets() []Type {
	if x != nil {
		return x.Rets
	}
	return nil
}

type FunctionDefinitions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Functions []*FunctionDefinition `protobuf:"bytes,1,rep,name=functions,proto3" json:"functions,omitempty"`
}

func (x *FunctionDefinitions) Reset() {
	*x = FunctionDefinitions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionDefinitions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionDefinitions) ProtoMessage() {}

func (x *FunctionDefinitions) ProtoReflect() protoreflect.Message {
	mi := &file_orbit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionDefinitions.ProtoReflect.Descriptor instead.
func (*FunctionDefinitions) Descriptor() ([]byte, []int) {
	return file_orbit_proto_rawDescGZIP(), []int{1}
}

func (x *FunctionDefinitions) GetFunctions() []*FunctionDefinition {
	if x != nil {
		return x.Functions
	}
	return nil
}

// comes from main
type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Broker   uint32   `protobuf:"varint,1,opt,name=broker,proto3" json:"broker,omitempty"`
	Function string   `protobuf:"bytes,4,opt,name=function,proto3" json:"function,omitempty"`
	Inputs   []uint64 `protobuf:"varint,16,rep,packed,name=inputs,proto3" json:"inputs,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orbit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_orbit_proto_rawDescGZIP(), []int{2}
}

func (x *CallRequest) GetBroker() uint32 {
	if x != nil {
		return x.Broker
	}
	return 0
}

func (x *CallRequest) GetFunction() string {
	if x != nil {
		return x.Function
	}
	return ""
}

func (x *CallRequest) GetInputs() []uint64 {
	if x != nil {
		return x.Inputs
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Broker uint32 `protobuf:"varint,1,opt,name=broker,proto3" json:"broker,omitempty"`
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_orbit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_orbit_proto_rawDescGZIP(), []int{3}
}

func (x *Empty) GetBroker() uint32 {
	if x != nil {
		return x.Broker
	}
	return 0
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_orbit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_orbit_proto_rawDescGZIP(), []int{4}
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// comes from plugin
type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// uint32 broker =1;
	Offset uint32 `protobuf:"varint,16,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   uint32 `protobuf:"varint,17,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orbit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_orbit_proto_rawDescGZIP(), []int{5}
}

func (x *ReadRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ReadRequest) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

// comes from plugin
type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// uint32 broker =1;
	Offset uint32 `protobuf:"varint,16,opt,name=offset,proto3" json:"offset,omitempty"`
	Data   []byte `protobuf:"bytes,17,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbit_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orbit_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_orbit_proto_rawDescGZIP(), []int{6}
}

func (x *WriteRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *WriteRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ReadReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []byte  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error IOError `protobuf:"varint,2,opt,name=error,proto3,enum=proto.IOError" json:"error,omitempty"`
}

func (x *ReadReturn) Reset() {
	*x = ReadReturn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbit_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadReturn) ProtoMessage() {}

func (x *ReadReturn) ProtoReflect() protoreflect.Message {
	mi := &file_orbit_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadReturn.ProtoReflect.Descriptor instead.
func (*ReadReturn) Descriptor() ([]byte, []int) {
	return file_orbit_proto_rawDescGZIP(), []int{7}
}

func (x *ReadReturn) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ReadReturn) GetError() IOError {
	if x != nil {
		return x.Error
	}
	return IOError_none
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    *uint64 `protobuf:"varint,2,opt,name=code,proto3,oneof" json:"code,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbit_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_orbit_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_orbit_proto_rawDescGZIP(), []int{8}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetCode() uint64 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

type WriteReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Written uint32  `protobuf:"varint,1,opt,name=written,proto3" json:"written,omitempty"`
	Error   IOError `protobuf:"varint,2,opt,name=error,proto3,enum=proto.IOError" json:"error,omitempty"`
}

func (x *WriteReturn) Reset() {
	*x = WriteReturn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbit_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteReturn) ProtoMessage() {}

func (x *WriteReturn) ProtoReflect() protoreflect.Message {
	mi := &file_orbit_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteReturn.ProtoReflect.Descriptor instead.
func (*WriteReturn) Descriptor() ([]byte, []int) {
	return file_orbit_proto_rawDescGZIP(), []int{9}
}

func (x *WriteReturn) GetWritten() uint32 {
	if x != nil {
		return x.Written
	}
	return 0
}

func (x *WriteReturn) GetError() IOError {
	if x != nil {
		return x.Error
	}
	return IOError_none
}

type CallReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rets  []uint64 `protobuf:"varint,1,rep,packed,name=rets,proto3" json:"rets,omitempty"`
	Error *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CallReturn) Reset() {
	*x = CallReturn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbit_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallReturn) ProtoMessage() {}

func (x *CallReturn) ProtoReflect() protoreflect.Message {
	mi := &file_orbit_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallReturn.ProtoReflect.Descriptor instead.
func (*CallReturn) Descriptor() ([]byte, []int) {
	return file_orbit_proto_rawDescGZIP(), []int{10}
}

func (x *CallReturn) GetRets() []uint64 {
	if x != nil {
		return x.Rets
	}
	return nil
}

func (x *CallReturn) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_orbit_proto protoreflect.FileDescriptor

var file_orbit_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x12, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12,
	0x1f, 0x0a, 0x04, 0x72, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x72, 0x65, 0x74, 0x73,
	0x22, 0x4e, 0x0a, 0x13, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x59, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x10, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x22, 0x1f, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x1e, 0x0a, 0x08,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x0b,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x3a, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x46, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x4f, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x43, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x4d, 0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x4f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x44, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x04, 0x72, 0x65, 0x74,
	0x73, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x37, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x69, 0x33,
	0x32, 0x10, 0x7f, 0x12, 0x07, 0x0a, 0x03, 0x69, 0x36, 0x34, 0x10, 0x7e, 0x12, 0x07, 0x0a, 0x03,
	0x66, 0x33, 0x32, 0x10, 0x7d, 0x12, 0x07, 0x0a, 0x03, 0x66, 0x36, 0x34, 0x10, 0x7c, 0x2a, 0x5f,
	0x0a, 0x07, 0x49, 0x4f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x6e, 0x6f, 0x6e,
	0x65, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x10, 0x10, 0x12, 0x10, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x10, 0x11, 0x12, 0x0f, 0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x10, 0x12, 0x12, 0x07, 0x0a, 0x03, 0x65, 0x6f, 0x66, 0x10, 0x13, 0x12,
	0x0e, 0x0a, 0x0a, 0x6e, 0x6f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x14, 0x32,
	0x93, 0x01, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x04, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x33, 0x0a, 0x07, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x12, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x32, 0x75, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x33, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x61, 0x64, 0x12, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x12, 0x36, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x08, 0x5a, 0x06,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orbit_proto_rawDescOnce sync.Once
	file_orbit_proto_rawDescData = file_orbit_proto_rawDesc
)

func file_orbit_proto_rawDescGZIP() []byte {
	file_orbit_proto_rawDescOnce.Do(func() {
		file_orbit_proto_rawDescData = protoimpl.X.CompressGZIP(file_orbit_proto_rawDescData)
	})
	return file_orbit_proto_rawDescData
}

var file_orbit_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_orbit_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_orbit_proto_goTypes = []interface{}{
	(Type)(0),                   // 0: proto.Type
	(IOError)(0),                // 1: proto.IOError
	(*FunctionDefinition)(nil),  // 2: proto.FunctionDefinition
	(*FunctionDefinitions)(nil), // 3: proto.FunctionDefinitions
	(*CallRequest)(nil),         // 4: proto.CallRequest
	(*Empty)(nil),               // 5: proto.Empty
	(*Metadata)(nil),            // 6: proto.Metadata
	(*ReadRequest)(nil),         // 7: proto.ReadRequest
	(*WriteRequest)(nil),        // 8: proto.WriteRequest
	(*ReadReturn)(nil),          // 9: proto.ReadReturn
	(*Error)(nil),               // 10: proto.Error
	(*WriteReturn)(nil),         // 11: proto.WriteReturn
	(*CallReturn)(nil),          // 12: proto.CallReturn
}
var file_orbit_proto_depIdxs = []int32{
	0,  // 0: proto.FunctionDefinition.args:type_name -> proto.Type
	0,  // 1: proto.FunctionDefinition.rets:type_name -> proto.Type
	2,  // 2: proto.FunctionDefinitions.functions:type_name -> proto.FunctionDefinition
	1,  // 3: proto.ReadReturn.error:type_name -> proto.IOError
	1,  // 4: proto.WriteReturn.error:type_name -> proto.IOError
	10, // 5: proto.CallReturn.error:type_name -> proto.Error
	5,  // 6: proto.Plugin.Meta:input_type -> proto.Empty
	5,  // 7: proto.Plugin.Symbols:input_type -> proto.Empty
	4,  // 8: proto.Plugin.Call:input_type -> proto.CallRequest
	7,  // 9: proto.Module.MemoryRead:input_type -> proto.ReadRequest
	8,  // 10: proto.Module.MemoryWrite:input_type -> proto.WriteRequest
	6,  // 11: proto.Plugin.Meta:output_type -> proto.Metadata
	3,  // 12: proto.Plugin.Symbols:output_type -> proto.FunctionDefinitions
	12, // 13: proto.Plugin.Call:output_type -> proto.CallReturn
	9,  // 14: proto.Module.MemoryRead:output_type -> proto.ReadReturn
	11, // 15: proto.Module.MemoryWrite:output_type -> proto.WriteReturn
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_orbit_proto_init() }
func file_orbit_proto_init() {
	if File_orbit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orbit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionDefinition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orbit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionDefinitions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orbit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orbit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orbit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orbit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orbit_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orbit_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadReturn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orbit_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orbit_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteReturn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orbit_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallReturn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_orbit_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orbit_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_orbit_proto_goTypes,
		DependencyIndexes: file_orbit_proto_depIdxs,
		EnumInfos:         file_orbit_proto_enumTypes,
		MessageInfos:      file_orbit_proto_msgTypes,
	}.Build()
	File_orbit_proto = out.File
	file_orbit_proto_rawDesc = nil
	file_orbit_proto_goTypes = nil
	file_orbit_proto_depIdxs = nil
}