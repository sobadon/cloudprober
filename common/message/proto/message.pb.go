// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: github.com/cloudprober/cloudprober/common/message/proto/message.proto

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

type DataNode_Type int32

const (
	DataNode_UNKNOWN DataNode_Type = 0
	DataNode_CLIENT  DataNode_Type = 1
	DataNode_SERVER  DataNode_Type = 2
)

// Enum value maps for DataNode_Type.
var (
	DataNode_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "CLIENT",
		2: "SERVER",
	}
	DataNode_Type_value = map[string]int32{
		"UNKNOWN": 0,
		"CLIENT":  1,
		"SERVER":  2,
	}
)

func (x DataNode_Type) Enum() *DataNode_Type {
	p := new(DataNode_Type)
	*p = x
	return p
}

func (x DataNode_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataNode_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_enumTypes[0].Descriptor()
}

func (DataNode_Type) Type() protoreflect.EnumType {
	return &file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_enumTypes[0]
}

func (x DataNode_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DataNode_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DataNode_Type(num)
	return nil
}

// Deprecated: Use DataNode_Type.Descriptor instead.
func (DataNode_Type) EnumDescriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_rawDescGZIP(), []int{1, 0}
}

// Constants defines constants with default values.
type Constants struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Magic *uint64 `protobuf:"varint,1,opt,name=magic,def=257787339638762" json:"magic,omitempty"`
}

// Default values for Constants fields.
const (
	Default_Constants_Magic = uint64(257787339638762)
)

func (x *Constants) Reset() {
	*x = Constants{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Constants) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Constants) ProtoMessage() {}

func (x *Constants) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Constants.ProtoReflect.Descriptor instead.
func (*Constants) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_rawDescGZIP(), []int{0}
}

func (x *Constants) GetMagic() uint64 {
	if x != nil && x.Magic != nil {
		return *x.Magic
	}
	return Default_Constants_Magic
}

// Datanode is something that see's a message AND can modify it.
type DataNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *DataNode_Type `protobuf:"varint,1,opt,name=type,enum=message.DataNode_Type,def=1" json:"type,omitempty"`
	Name *string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Port *string        `protobuf:"bytes,4,opt,name=port" json:"port,omitempty"`
	// 8 bytes of timestamp in pcap-friendly network byte order.
	TimestampUsec []byte `protobuf:"bytes,3,opt,name=timestamp_usec,json=timestampUsec" json:"timestamp_usec,omitempty"`
}

// Default values for DataNode fields.
const (
	Default_DataNode_Type = DataNode_CLIENT
)

func (x *DataNode) Reset() {
	*x = DataNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataNode) ProtoMessage() {}

func (x *DataNode) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataNode.ProtoReflect.Descriptor instead.
func (*DataNode) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_rawDescGZIP(), []int{1}
}

func (x *DataNode) GetType() DataNode_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Default_DataNode_Type
}

func (x *DataNode) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *DataNode) GetPort() string {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return ""
}

func (x *DataNode) GetTimestampUsec() []byte {
	if x != nil {
		return x.TimestampUsec
	}
	return nil
}

// Msg is a message sent over the network.
// magic, seq, src and dst are required fields.
type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Magic *uint64 `protobuf:"fixed64,1,opt,name=magic" json:"magic,omitempty"` // required.
	// 8 bytes of sequence in pcap-friendly network byte order.
	Seq []byte `protobuf:"bytes,2,opt,name=seq" json:"seq,omitempty"` // required.
	// Datanodes seen by this message.
	Src   *DataNode   `protobuf:"bytes,3,opt,name=src" json:"src,omitempty"`     // required.
	Dst   *DataNode   `protobuf:"bytes,4,opt,name=dst" json:"dst,omitempty"`     // required.
	Nodes []*DataNode `protobuf:"bytes,5,rep,name=nodes" json:"nodes,omitempty"` // Intermediate nodes.
	// Optional payload.
	Payload []byte `protobuf:"bytes,99,opt,name=payload" json:"payload,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_rawDescGZIP(), []int{2}
}

func (x *Msg) GetMagic() uint64 {
	if x != nil && x.Magic != nil {
		return *x.Magic
	}
	return 0
}

func (x *Msg) GetSeq() []byte {
	if x != nil {
		return x.Seq
	}
	return nil
}

func (x *Msg) GetSrc() *DataNode {
	if x != nil {
		return x.Src
	}
	return nil
}

func (x *Msg) GetDst() *DataNode {
	if x != nil {
		return x.Dst
	}
	return nil
}

func (x *Msg) GetNodes() []*DataNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *Msg) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_github_com_cloudprober_cloudprober_common_message_proto_message_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x32, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a,
	0x05, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x3a, 0x0f, 0x32, 0x35,
	0x37, 0x37, 0x38, 0x37, 0x33, 0x33, 0x39, 0x36, 0x33, 0x38, 0x37, 0x36, 0x32, 0x52, 0x05, 0x6d,
	0x61, 0x67, 0x69, 0x63, 0x22, 0xba, 0x01, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f,
	0x64, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x06, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x55, 0x73, 0x65, 0x63, 0x22, 0x2b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x49,
	0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10,
	0x02, 0x22, 0xba, 0x01, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x67,
	0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x05, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x65,
	0x71, 0x12, 0x23, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x23, 0x0a, 0x03, 0x64, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x03, 0x64, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x63, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x39,
	0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_rawDescData = file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_goTypes = []interface{}{
	(DataNode_Type)(0), // 0: message.DataNode.Type
	(*Constants)(nil),  // 1: message.Constants
	(*DataNode)(nil),   // 2: message.DataNode
	(*Msg)(nil),        // 3: message.Msg
}
var file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_depIdxs = []int32{
	0, // 0: message.DataNode.type:type_name -> message.DataNode.Type
	2, // 1: message.Msg.src:type_name -> message.DataNode
	2, // 2: message.Msg.dst:type_name -> message.DataNode
	2, // 3: message.Msg.nodes:type_name -> message.DataNode
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_init() }
func file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_init() {
	if File_github_com_cloudprober_cloudprober_common_message_proto_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Constants); i {
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
		file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataNode); i {
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
		file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
			RawDescriptor: file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_depIdxs,
		EnumInfos:         file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_enumTypes,
		MessageInfos:      file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_common_message_proto_message_proto = out.File
	file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_common_message_proto_message_proto_depIdxs = nil
}
