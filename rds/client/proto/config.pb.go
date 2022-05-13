// Configuration proto for RDS targets.
// Example:
// {
//   request {
//     resource_uri: "gcp://gce_instances/google.com:bbmc-stackdriver/*"
//   }
// }

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: github.com/cloudprober/cloudprober/rds/client/proto/config.proto

package proto

import (
	proto1 "github.com/cloudprober/cloudprober/common/oauth/proto"
	proto2 "github.com/cloudprober/cloudprober/common/tlsconfig/proto"
	proto "github.com/cloudprober/cloudprober/rds/proto"
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

// ClientConf represents resource discovery service (RDS) based targets.
// Next tag: 6
type ClientConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerOptions *ClientConf_ServerOptions   `protobuf:"bytes,1,opt,name=server_options,json=serverOptions" json:"server_options,omitempty"`
	Request       *proto.ListResourcesRequest `protobuf:"bytes,2,req,name=request" json:"request,omitempty"`
	// How often targets should be evaluated. Any number less than or equal to 0
	// will result in no target caching (targets will be reevaluated on demand).
	// Note that individual target types may have their own caches implemented
	// (specifically GCE instances/forwarding rules). This does not impact those
	// caches.
	ReEvalSec *int32 `protobuf:"varint,3,opt,name=re_eval_sec,json=reEvalSec,def=30" json:"re_eval_sec,omitempty"`
}

// Default values for ClientConf fields.
const (
	Default_ClientConf_ReEvalSec = int32(30)
)

func (x *ClientConf) Reset() {
	*x = ClientConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConf) ProtoMessage() {}

func (x *ClientConf) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConf.ProtoReflect.Descriptor instead.
func (*ClientConf) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *ClientConf) GetServerOptions() *ClientConf_ServerOptions {
	if x != nil {
		return x.ServerOptions
	}
	return nil
}

func (x *ClientConf) GetRequest() *proto.ListResourcesRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *ClientConf) GetReEvalSec() int32 {
	if x != nil && x.ReEvalSec != nil {
		return *x.ReEvalSec
	}
	return Default_ClientConf_ReEvalSec
}

type ClientConf_ServerOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerAddress *string `protobuf:"bytes,1,opt,name=server_address,json=serverAddress" json:"server_address,omitempty"`
	// Optional oauth config for authentication.
	OauthConfig *proto1.Config `protobuf:"bytes,2,opt,name=oauth_config,json=oauthConfig" json:"oauth_config,omitempty"`
	// TLS config, it can be used to:
	// - Specify a CA cert for server cert verification:
	//     tls_config {
	//       ca_cert_file: "...."
	//     }
	//
	// - Specify client's TLS cert and key:
	//     tls_config {
	//       tls_cert_file: "..."
	//       tls_key_file: "..."
	//     }
	TlsConfig *proto2.TLSConfig `protobuf:"bytes,3,opt,name=tls_config,json=tlsConfig" json:"tls_config,omitempty"`
}

func (x *ClientConf_ServerOptions) Reset() {
	*x = ClientConf_ServerOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConf_ServerOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConf_ServerOptions) ProtoMessage() {}

func (x *ClientConf_ServerOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConf_ServerOptions.ProtoReflect.Descriptor instead.
func (*ClientConf_ServerOptions) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ClientConf_ServerOptions) GetServerAddress() string {
	if x != nil && x.ServerAddress != nil {
		return *x.ServerAddress
	}
	return ""
}

func (x *ClientConf_ServerOptions) GetOauthConfig() *proto1.Config {
	if x != nil {
		return x.OauthConfig
	}
	return nil
}

func (x *ClientConf_ServerOptions) GetTlsConfig() *proto2.TLSConfig {
	if x != nil {
		return x.TlsConfig
	}
	return nil
}

var File_github_com_cloudprober_cloudprober_rds_client_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x64, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e,
	0x72, 0x64, 0x73, 0x1a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x74, 0x6c, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x72, 0x2f, 0x72, 0x64, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x02, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x50, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64, 0x73,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0b, 0x72, 0x65, 0x5f,
	0x65, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02,
	0x33, 0x30, 0x52, 0x09, 0x72, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x1a, 0xb5, 0x01,
	0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x3f, 0x0a, 0x0a, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x6c, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x74, 0x6c, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x64, 0x73, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_rawDescData = file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_goTypes = []interface{}{
	(*ClientConf)(nil),                 // 0: cloudprober.rds.ClientConf
	(*ClientConf_ServerOptions)(nil),   // 1: cloudprober.rds.ClientConf.ServerOptions
	(*proto.ListResourcesRequest)(nil), // 2: cloudprober.rds.ListResourcesRequest
	(*proto1.Config)(nil),              // 3: cloudprober.oauth.Config
	(*proto2.TLSConfig)(nil),           // 4: cloudprober.tlsconfig.TLSConfig
}
var file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_depIdxs = []int32{
	1, // 0: cloudprober.rds.ClientConf.server_options:type_name -> cloudprober.rds.ClientConf.ServerOptions
	2, // 1: cloudprober.rds.ClientConf.request:type_name -> cloudprober.rds.ListResourcesRequest
	3, // 2: cloudprober.rds.ClientConf.ServerOptions.oauth_config:type_name -> cloudprober.oauth.Config
	4, // 3: cloudprober.rds.ClientConf.ServerOptions.tls_config:type_name -> cloudprober.tlsconfig.TLSConfig
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_init() }
func file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_rds_client_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConf); i {
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
		file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConf_ServerOptions); i {
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
			RawDescriptor: file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_depIdxs,
		MessageInfos:      file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_rds_client_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_rds_client_proto_config_proto_depIdxs = nil
}
