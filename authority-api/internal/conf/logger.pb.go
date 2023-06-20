// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: conf/logger.proto

package conf

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

// 日志
type Logger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string      `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Zap  *Logger_Zap `protobuf:"bytes,2,opt,name=zap,proto3" json:"zap,omitempty"`
}

func (x *Logger) Reset() {
	*x = Logger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_logger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger) ProtoMessage() {}

func (x *Logger) ProtoReflect() protoreflect.Message {
	mi := &file_conf_logger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger.ProtoReflect.Descriptor instead.
func (*Logger) Descriptor() ([]byte, []int) {
	return file_conf_logger_proto_rawDescGZIP(), []int{0}
}

func (x *Logger) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Logger) GetZap() *Logger_Zap {
	if x != nil {
		return x.Zap
	}
	return nil
}

// Zap
type Logger_Zap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level      string `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	FilePath   string `protobuf:"bytes,2,opt,name=filePath,proto3" json:"filePath,omitempty"`
	FileName   string `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName,omitempty"`
	MaxSize    int32  `protobuf:"varint,4,opt,name=maxSize,proto3" json:"maxSize,omitempty"`
	MaxAge     int32  `protobuf:"varint,5,opt,name=maxAge,proto3" json:"maxAge,omitempty"`
	MaxBackups int32  `protobuf:"varint,6,opt,name=maxBackups,proto3" json:"maxBackups,omitempty"`
	LogStdout  bool   `protobuf:"varint,7,opt,name=logStdout,proto3" json:"logStdout,omitempty"`
}

func (x *Logger_Zap) Reset() {
	*x = Logger_Zap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_logger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logger_Zap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger_Zap) ProtoMessage() {}

func (x *Logger_Zap) ProtoReflect() protoreflect.Message {
	mi := &file_conf_logger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger_Zap.ProtoReflect.Descriptor instead.
func (*Logger_Zap) Descriptor() ([]byte, []int) {
	return file_conf_logger_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Logger_Zap) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Logger_Zap) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *Logger_Zap) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Logger_Zap) GetMaxSize() int32 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

func (x *Logger_Zap) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *Logger_Zap) GetMaxBackups() int32 {
	if x != nil {
		return x.MaxBackups
	}
	return 0
}

func (x *Logger_Zap) GetLogStdout() bool {
	if x != nil {
		return x.LogStdout
	}
	return false
}

var File_conf_logger_proto protoreflect.FileDescriptor

var file_conf_logger_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x22, 0x90, 0x02, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x7a, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x5a, 0x61, 0x70, 0x52, 0x03, 0x7a, 0x61, 0x70,
	0x1a, 0xc3, 0x01, 0x0a, 0x03, 0x5a, 0x61, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61,
	0x78, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x53,
	0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6c, 0x6f, 0x67,
	0x53, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x2d, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f,
	0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_logger_proto_rawDescOnce sync.Once
	file_conf_logger_proto_rawDescData = file_conf_logger_proto_rawDesc
)

func file_conf_logger_proto_rawDescGZIP() []byte {
	file_conf_logger_proto_rawDescOnce.Do(func() {
		file_conf_logger_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_logger_proto_rawDescData)
	})
	return file_conf_logger_proto_rawDescData
}

var file_conf_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_conf_logger_proto_goTypes = []interface{}{
	(*Logger)(nil),     // 0: authority.conf.Logger
	(*Logger_Zap)(nil), // 1: authority.conf.Logger.Zap
}
var file_conf_logger_proto_depIdxs = []int32{
	1, // 0: authority.conf.Logger.zap:type_name -> authority.conf.Logger.Zap
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_conf_logger_proto_init() }
func file_conf_logger_proto_init() {
	if File_conf_logger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_logger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logger); i {
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
		file_conf_logger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logger_Zap); i {
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
			RawDescriptor: file_conf_logger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_logger_proto_goTypes,
		DependencyIndexes: file_conf_logger_proto_depIdxs,
		MessageInfos:      file_conf_logger_proto_msgTypes,
	}.Build()
	File_conf_logger_proto = out.File
	file_conf_logger_proto_rawDesc = nil
	file_conf_logger_proto_goTypes = nil
	file_conf_logger_proto_depIdxs = nil
}
