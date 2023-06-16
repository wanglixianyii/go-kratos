// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: conf/registry.proto

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

// 注册发现中心
type Registry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       string               `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Consul     *Registry_Consul     `protobuf:"bytes,2,opt,name=consul,proto3" json:"consul,omitempty"`         // Consul
	Etcd       *Registry_Etcd       `protobuf:"bytes,3,opt,name=etcd,proto3" json:"etcd,omitempty"`             // Etcd
	Kubernetes *Registry_Kubernetes `protobuf:"bytes,4,opt,name=kubernetes,proto3" json:"kubernetes,omitempty"` // Kubernetes
}

func (x *Registry) Reset() {
	*x = Registry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_conf_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_conf_registry_proto_rawDescGZIP(), []int{0}
}

func (x *Registry) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Registry) GetConsul() *Registry_Consul {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *Registry) GetEtcd() *Registry_Etcd {
	if x != nil {
		return x.Etcd
	}
	return nil
}

func (x *Registry) GetKubernetes() *Registry_Kubernetes {
	if x != nil {
		return x.Kubernetes
	}
	return nil
}

// Consul
type Registry_Consul struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scheme      string `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`                               // 网络样式
	Address     string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`                             // 服务端地址
	HealthCheck bool   `protobuf:"varint,3,opt,name=health_check,json=healthCheck,proto3" json:"health_check,omitempty"` // 健康检查
}

func (x *Registry_Consul) Reset() {
	*x = Registry_Consul{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry_Consul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_Consul) ProtoMessage() {}

func (x *Registry_Consul) ProtoReflect() protoreflect.Message {
	mi := &file_conf_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry_Consul.ProtoReflect.Descriptor instead.
func (*Registry_Consul) Descriptor() ([]byte, []int) {
	return file_conf_registry_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Registry_Consul) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *Registry_Consul) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Registry_Consul) GetHealthCheck() bool {
	if x != nil {
		return x.HealthCheck
	}
	return false
}

// Etcd
type Registry_Etcd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *Registry_Etcd) Reset() {
	*x = Registry_Etcd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry_Etcd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_Etcd) ProtoMessage() {}

func (x *Registry_Etcd) ProtoReflect() protoreflect.Message {
	mi := &file_conf_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry_Etcd.ProtoReflect.Descriptor instead.
func (*Registry_Etcd) Descriptor() ([]byte, []int) {
	return file_conf_registry_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Registry_Etcd) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

// Kubernetes
type Registry_Kubernetes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Registry_Kubernetes) Reset() {
	*x = Registry_Kubernetes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry_Kubernetes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_Kubernetes) ProtoMessage() {}

func (x *Registry_Kubernetes) ProtoReflect() protoreflect.Message {
	mi := &file_conf_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry_Kubernetes.ProtoReflect.Descriptor instead.
func (*Registry_Kubernetes) Descriptor() ([]byte, []int) {
	return file_conf_registry_proto_rawDescGZIP(), []int{0, 2}
}

var File_conf_registry_proto protoreflect.FileDescriptor

var file_conf_registry_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x22, 0xe2, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12,
	0x31, 0x0a, 0x04, 0x65, 0x74, 0x63, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x52, 0x04, 0x65, 0x74,
	0x63, 0x64, 0x12, 0x43, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x1a, 0x5d, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x24, 0x0a, 0x04, 0x45, 0x74, 0x63, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x0c, 0x0a, 0x0a,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x42, 0x1e, 0x5a, 0x1c, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_conf_registry_proto_rawDescOnce sync.Once
	file_conf_registry_proto_rawDescData = file_conf_registry_proto_rawDesc
)

func file_conf_registry_proto_rawDescGZIP() []byte {
	file_conf_registry_proto_rawDescOnce.Do(func() {
		file_conf_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_registry_proto_rawDescData)
	})
	return file_conf_registry_proto_rawDescData
}

var file_conf_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_conf_registry_proto_goTypes = []interface{}{
	(*Registry)(nil),            // 0: authority.conf.Registry
	(*Registry_Consul)(nil),     // 1: authority.conf.Registry.Consul
	(*Registry_Etcd)(nil),       // 2: authority.conf.Registry.Etcd
	(*Registry_Kubernetes)(nil), // 3: authority.conf.Registry.Kubernetes
}
var file_conf_registry_proto_depIdxs = []int32{
	1, // 0: authority.conf.Registry.consul:type_name -> authority.conf.Registry.Consul
	2, // 1: authority.conf.Registry.etcd:type_name -> authority.conf.Registry.Etcd
	3, // 2: authority.conf.Registry.kubernetes:type_name -> authority.conf.Registry.Kubernetes
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_conf_registry_proto_init() }
func file_conf_registry_proto_init() {
	if File_conf_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry); i {
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
		file_conf_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry_Consul); i {
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
		file_conf_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry_Etcd); i {
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
		file_conf_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry_Kubernetes); i {
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
			RawDescriptor: file_conf_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_registry_proto_goTypes,
		DependencyIndexes: file_conf_registry_proto_depIdxs,
		MessageInfos:      file_conf_registry_proto_msgTypes,
	}.Build()
	File_conf_registry_proto = out.File
	file_conf_registry_proto_rawDesc = nil
	file_conf_registry_proto_goTypes = nil
	file_conf_registry_proto_depIdxs = nil
}