// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: authority-api/v1/authority.proto

package v1

import (
	v1 "authority-api/api/service/authority-rpc/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_authority_api_v1_authority_proto protoreflect.FileDescriptor

var file_authority_api_v1_authority_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x28, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfb, 0x04, 0x0a,
	0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x75, 0x0a, 0x08, 0x4d, 0x65,
	0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a,
	0x22, 0x1f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2d, 0x72, 0x70, 0x63,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x7d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12,
	0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71,
	0x1a, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2d, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x7d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x1f,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x1a,
	0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x1a, 0x21, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2d, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x7a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x1f, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x1a, 0x20,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x2a, 0x21, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x2d, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x7d, 0x0a, 0x0a, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x2d, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x6e, 0x75, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x27, 0x5a, 0x25, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_authority_api_v1_authority_proto_goTypes = []interface{}{
	(*v1.MenuListReq)(nil),    // 0: authority.rpc.v1.MenuListReq
	(*v1.CreateMenuReq)(nil),  // 1: authority.rpc.v1.CreateMenuReq
	(*v1.UpdateMenuReq)(nil),  // 2: authority.rpc.v1.UpdateMenuReq
	(*v1.DeleteMenuReq)(nil),  // 3: authority.rpc.v1.DeleteMenuReq
	(*v1.DetailMenuReq)(nil),  // 4: authority.rpc.v1.DetailMenuReq
	(*v1.MenuListResp)(nil),   // 5: authority.rpc.v1.MenuListResp
	(*v1.CreateMenuResp)(nil), // 6: authority.rpc.v1.CreateMenuResp
	(*v1.UpdateMenuResp)(nil), // 7: authority.rpc.v1.UpdateMenuResp
	(*v1.DeleteMenuResp)(nil), // 8: authority.rpc.v1.DeleteMenuResp
	(*v1.DetailMenuResp)(nil), // 9: authority.rpc.v1.DetailMenuResp
}
var file_authority_api_v1_authority_proto_depIdxs = []int32{
	0, // 0: authority.api.v1.Authority.MenuList:input_type -> authority.rpc.v1.MenuListReq
	1, // 1: authority.api.v1.Authority.CreateMenu:input_type -> authority.rpc.v1.CreateMenuReq
	2, // 2: authority.api.v1.Authority.UpdateMenu:input_type -> authority.rpc.v1.UpdateMenuReq
	3, // 3: authority.api.v1.Authority.DeleteMenu:input_type -> authority.rpc.v1.DeleteMenuReq
	4, // 4: authority.api.v1.Authority.DetailMenu:input_type -> authority.rpc.v1.DetailMenuReq
	5, // 5: authority.api.v1.Authority.MenuList:output_type -> authority.rpc.v1.MenuListResp
	6, // 6: authority.api.v1.Authority.CreateMenu:output_type -> authority.rpc.v1.CreateMenuResp
	7, // 7: authority.api.v1.Authority.UpdateMenu:output_type -> authority.rpc.v1.UpdateMenuResp
	8, // 8: authority.api.v1.Authority.DeleteMenu:output_type -> authority.rpc.v1.DeleteMenuResp
	9, // 9: authority.api.v1.Authority.DetailMenu:output_type -> authority.rpc.v1.DetailMenuResp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authority_api_v1_authority_proto_init() }
func file_authority_api_v1_authority_proto_init() {
	if File_authority_api_v1_authority_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authority_api_v1_authority_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authority_api_v1_authority_proto_goTypes,
		DependencyIndexes: file_authority_api_v1_authority_proto_depIdxs,
	}.Build()
	File_authority_api_v1_authority_proto = out.File
	file_authority_api_v1_authority_proto_rawDesc = nil
	file_authority_api_v1_authority_proto_goTypes = nil
	file_authority_api_v1_authority_proto_depIdxs = nil
}
