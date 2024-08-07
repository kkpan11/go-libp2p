// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: pb/holepunch.proto

package pb

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

type HolePunch_Type int32

const (
	HolePunch_CONNECT HolePunch_Type = 100
	HolePunch_SYNC    HolePunch_Type = 300
)

// Enum value maps for HolePunch_Type.
var (
	HolePunch_Type_name = map[int32]string{
		100: "CONNECT",
		300: "SYNC",
	}
	HolePunch_Type_value = map[string]int32{
		"CONNECT": 100,
		"SYNC":    300,
	}
)

func (x HolePunch_Type) Enum() *HolePunch_Type {
	p := new(HolePunch_Type)
	*p = x
	return p
}

func (x HolePunch_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HolePunch_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_holepunch_proto_enumTypes[0].Descriptor()
}

func (HolePunch_Type) Type() protoreflect.EnumType {
	return &file_pb_holepunch_proto_enumTypes[0]
}

func (x HolePunch_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *HolePunch_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = HolePunch_Type(num)
	return nil
}

// Deprecated: Use HolePunch_Type.Descriptor instead.
func (HolePunch_Type) EnumDescriptor() ([]byte, []int) {
	return file_pb_holepunch_proto_rawDescGZIP(), []int{0, 0}
}

// spec: https://github.com/libp2p/specs/blob/master/relay/DCUtR.md
type HolePunch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     *HolePunch_Type `protobuf:"varint,1,req,name=type,enum=holepunch.pb.HolePunch_Type" json:"type,omitempty"`
	ObsAddrs [][]byte        `protobuf:"bytes,2,rep,name=ObsAddrs" json:"ObsAddrs,omitempty"`
}

func (x *HolePunch) Reset() {
	*x = HolePunch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_holepunch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HolePunch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HolePunch) ProtoMessage() {}

func (x *HolePunch) ProtoReflect() protoreflect.Message {
	mi := &file_pb_holepunch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HolePunch.ProtoReflect.Descriptor instead.
func (*HolePunch) Descriptor() ([]byte, []int) {
	return file_pb_holepunch_proto_rawDescGZIP(), []int{0}
}

func (x *HolePunch) GetType() HolePunch_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return HolePunch_CONNECT
}

func (x *HolePunch) GetObsAddrs() [][]byte {
	if x != nil {
		return x.ObsAddrs
	}
	return nil
}

var File_pb_holepunch_proto protoreflect.FileDescriptor

var file_pb_holepunch_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x68, 0x6f, 0x6c, 0x65, 0x70, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x68, 0x6f, 0x6c, 0x65, 0x70, 0x75, 0x6e, 0x63, 0x68, 0x2e,
	0x70, 0x62, 0x22, 0x79, 0x0a, 0x09, 0x48, 0x6f, 0x6c, 0x65, 0x50, 0x75, 0x6e, 0x63, 0x68, 0x12,
	0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x68, 0x6f, 0x6c, 0x65, 0x70, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x6c,
	0x65, 0x50, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x62, 0x73, 0x41, 0x64, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x08, 0x4f, 0x62, 0x73, 0x41, 0x64, 0x64, 0x72, 0x73, 0x22, 0x1e, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x10, 0x64, 0x12, 0x09, 0x0a, 0x04, 0x53, 0x59, 0x4e, 0x43, 0x10, 0xac, 0x02,
}

var (
	file_pb_holepunch_proto_rawDescOnce sync.Once
	file_pb_holepunch_proto_rawDescData = file_pb_holepunch_proto_rawDesc
)

func file_pb_holepunch_proto_rawDescGZIP() []byte {
	file_pb_holepunch_proto_rawDescOnce.Do(func() {
		file_pb_holepunch_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_holepunch_proto_rawDescData)
	})
	return file_pb_holepunch_proto_rawDescData
}

var file_pb_holepunch_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_holepunch_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_holepunch_proto_goTypes = []any{
	(HolePunch_Type)(0), // 0: holepunch.pb.HolePunch.Type
	(*HolePunch)(nil),   // 1: holepunch.pb.HolePunch
}
var file_pb_holepunch_proto_depIdxs = []int32{
	0, // 0: holepunch.pb.HolePunch.type:type_name -> holepunch.pb.HolePunch.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_holepunch_proto_init() }
func file_pb_holepunch_proto_init() {
	if File_pb_holepunch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_holepunch_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HolePunch); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_holepunch_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_holepunch_proto_goTypes,
		DependencyIndexes: file_pb_holepunch_proto_depIdxs,
		EnumInfos:         file_pb_holepunch_proto_enumTypes,
		MessageInfos:      file_pb_holepunch_proto_msgTypes,
	}.Build()
	File_pb_holepunch_proto = out.File
	file_pb_holepunch_proto_rawDesc = nil
	file_pb_holepunch_proto_goTypes = nil
	file_pb_holepunch_proto_depIdxs = nil
}
