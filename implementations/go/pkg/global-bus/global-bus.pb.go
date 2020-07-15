// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: lib/global-bus.proto

package global_bus

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Used internally in the package for testing. DO NOT USE.
type MyValidTestEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *MyValidTestEvent) Reset() {
	*x = MyValidTestEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_global_bus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyValidTestEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyValidTestEvent) ProtoMessage() {}

func (x *MyValidTestEvent) ProtoReflect() protoreflect.Message {
	mi := &file_lib_global_bus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyValidTestEvent.ProtoReflect.Descriptor instead.
func (*MyValidTestEvent) Descriptor() ([]byte, []int) {
	return file_lib_global_bus_proto_rawDescGZIP(), []int{0}
}

func (x *MyValidTestEvent) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// Used internally in the package for testing. DO NOT USE.
type MyInvalidTestEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *MyInvalidTestEvent) Reset() {
	*x = MyInvalidTestEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_global_bus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyInvalidTestEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyInvalidTestEvent) ProtoMessage() {}

func (x *MyInvalidTestEvent) ProtoReflect() protoreflect.Message {
	mi := &file_lib_global_bus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyInvalidTestEvent.ProtoReflect.Descriptor instead.
func (*MyInvalidTestEvent) Descriptor() ([]byte, []int) {
	return file_lib_global_bus_proto_rawDescGZIP(), []int{1}
}

func (x *MyInvalidTestEvent) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// Describes the events that are send around
type EventPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The event path of this event. Used for distinguishing between the actual event types
	EventPath string `protobuf:"bytes,1,opt,name=eventPath,proto3" json:"eventPath,omitempty"`
	//  The actual event data, encoding as a protobuf structure
	EventData []byte `protobuf:"bytes,2,opt,name=eventData,proto3" json:"eventData,omitempty"`
}

func (x *EventPackage) Reset() {
	*x = EventPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_global_bus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventPackage) ProtoMessage() {}

func (x *EventPackage) ProtoReflect() protoreflect.Message {
	mi := &file_lib_global_bus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventPackage.ProtoReflect.Descriptor instead.
func (*EventPackage) Descriptor() ([]byte, []int) {
	return file_lib_global_bus_proto_rawDescGZIP(), []int{2}
}

func (x *EventPackage) GetEventPath() string {
	if x != nil {
		return x.EventPath
	}
	return ""
}

func (x *EventPackage) GetEventData() []byte {
	if x != nil {
		return x.EventData
	}
	return nil
}

var file_lib_global_bus_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         58691,
		Name:          "event_path",
		Tag:           "bytes,58691,opt,name=event_path",
		Filename:      "lib/global-bus.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	// optional string event_path = 58691;
	E_EventPath = &file_lib_global_bus_proto_extTypes[0]
)

var File_lib_global_bus_proto protoreflect.FileDescriptor

var file_lib_global_bus_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2d, 0x62, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x10, 0x4d, 0x79, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x54, 0x65, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x3a, 0x0d, 0x9a, 0xd4, 0x1c, 0x09, 0x74, 0x65, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22,
	0x28, 0x0a, 0x12, 0x4d, 0x79, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x65, 0x73, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x4a, 0x0a, 0x0c, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x3a, 0x40, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc3, 0xca, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6c, 0x65, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2d, 0x62, 0x75, 0x73, 0x2f, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2d, 0x62, 0x75, 0x73, 0x3b, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x5f, 0x62, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lib_global_bus_proto_rawDescOnce sync.Once
	file_lib_global_bus_proto_rawDescData = file_lib_global_bus_proto_rawDesc
)

func file_lib_global_bus_proto_rawDescGZIP() []byte {
	file_lib_global_bus_proto_rawDescOnce.Do(func() {
		file_lib_global_bus_proto_rawDescData = protoimpl.X.CompressGZIP(file_lib_global_bus_proto_rawDescData)
	})
	return file_lib_global_bus_proto_rawDescData
}

var file_lib_global_bus_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_lib_global_bus_proto_goTypes = []interface{}{
	(*MyValidTestEvent)(nil),          // 0: MyValidTestEvent
	(*MyInvalidTestEvent)(nil),        // 1: MyInvalidTestEvent
	(*EventPackage)(nil),              // 2: EventPackage
	(*descriptor.MessageOptions)(nil), // 3: google.protobuf.MessageOptions
}
var file_lib_global_bus_proto_depIdxs = []int32{
	3, // 0: event_path:extendee -> google.protobuf.MessageOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lib_global_bus_proto_init() }
func file_lib_global_bus_proto_init() {
	if File_lib_global_bus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lib_global_bus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyValidTestEvent); i {
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
		file_lib_global_bus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyInvalidTestEvent); i {
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
		file_lib_global_bus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventPackage); i {
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
			RawDescriptor: file_lib_global_bus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_lib_global_bus_proto_goTypes,
		DependencyIndexes: file_lib_global_bus_proto_depIdxs,
		MessageInfos:      file_lib_global_bus_proto_msgTypes,
		ExtensionInfos:    file_lib_global_bus_proto_extTypes,
	}.Build()
	File_lib_global_bus_proto = out.File
	file_lib_global_bus_proto_rawDesc = nil
	file_lib_global_bus_proto_goTypes = nil
	file_lib_global_bus_proto_depIdxs = nil
}
