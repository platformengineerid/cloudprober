// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: github.com/cloudprober/cloudprober/probes/external/proto/config.proto

package proto

import (
	proto "github.com/cloudprober/cloudprober/metrics/payload/proto"
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

// External probes support two mode: ONCE and SERVER. In ONCE mode, external
// command is re-executed for each probe run, while in SERVER mode, command
// is run in server mode, re-executed only if not running already.
type ProbeConf_Mode int32

const (
	ProbeConf_ONCE   ProbeConf_Mode = 0
	ProbeConf_SERVER ProbeConf_Mode = 1
)

// Enum value maps for ProbeConf_Mode.
var (
	ProbeConf_Mode_name = map[int32]string{
		0: "ONCE",
		1: "SERVER",
	}
	ProbeConf_Mode_value = map[string]int32{
		"ONCE":   0,
		"SERVER": 1,
	}
)

func (x ProbeConf_Mode) Enum() *ProbeConf_Mode {
	p := new(ProbeConf_Mode)
	*p = x
	return p
}

func (x ProbeConf_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProbeConf_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_enumTypes[0].Descriptor()
}

func (ProbeConf_Mode) Type() protoreflect.EnumType {
	return &file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_enumTypes[0]
}

func (x ProbeConf_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProbeConf_Mode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProbeConf_Mode(num)
	return nil
}

// Deprecated: Use ProbeConf_Mode.Descriptor instead.
func (ProbeConf_Mode) EnumDescriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_rawDescGZIP(), []int{0, 0}
}

type ProbeConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode *ProbeConf_Mode `protobuf:"varint,1,opt,name=mode,enum=cloudprober.probes.external.ProbeConf_Mode,def=0" json:"mode,omitempty"`
	// Command.  For ONCE probes, arguments are processed for the following field
	// substitutions:
	// @probe@    Name of the probe
	// @target@   Hostname of the target
	// @address@  IP address of the target
	//
	// For example, for target ig-us-central1-a, /tools/recreate_vm -vm @target@
	// will get converted to: /tools/recreate_vm -vm ig-us-central1-a
	Command *string             `protobuf:"bytes,2,req,name=command" json:"command,omitempty"`
	Options []*ProbeConf_Option `protobuf:"bytes,3,rep,name=options" json:"options,omitempty"`
	// Export output as metrics, where output is the output returned by the
	// external probe process, over stdout for ONCE probes, and through ProbeReply
	// for SERVER probes. Cloudprober expects variables to be in the following
	// format in the output:
	// var1 value1 (for example: total_errors 589)
	OutputAsMetrics      *bool                       `protobuf:"varint,4,opt,name=output_as_metrics,json=outputAsMetrics,def=1" json:"output_as_metrics,omitempty"`
	OutputMetricsOptions *proto.OutputMetricsOptions `protobuf:"bytes,5,opt,name=output_metrics_options,json=outputMetricsOptions" json:"output_metrics_options,omitempty"`
}

// Default values for ProbeConf fields.
const (
	Default_ProbeConf_Mode            = ProbeConf_ONCE
	Default_ProbeConf_OutputAsMetrics = bool(true)
)

func (x *ProbeConf) Reset() {
	*x = ProbeConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeConf) ProtoMessage() {}

func (x *ProbeConf) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeConf.ProtoReflect.Descriptor instead.
func (*ProbeConf) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *ProbeConf) GetMode() ProbeConf_Mode {
	if x != nil && x.Mode != nil {
		return *x.Mode
	}
	return Default_ProbeConf_Mode
}

func (x *ProbeConf) GetCommand() string {
	if x != nil && x.Command != nil {
		return *x.Command
	}
	return ""
}

func (x *ProbeConf) GetOptions() []*ProbeConf_Option {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *ProbeConf) GetOutputAsMetrics() bool {
	if x != nil && x.OutputAsMetrics != nil {
		return *x.OutputAsMetrics
	}
	return Default_ProbeConf_OutputAsMetrics
}

func (x *ProbeConf) GetOutputMetricsOptions() *proto.OutputMetricsOptions {
	if x != nil {
		return x.OutputMetricsOptions
	}
	return nil
}

// Options for the SERVER mode probe requests. These options are passed on to
// the external probe server as part of the ProbeRequest. Values are
// substituted similar to command arguments for the ONCE mode probes.
type ProbeConf_Option struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (x *ProbeConf_Option) Reset() {
	*x = ProbeConf_Option{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeConf_Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeConf_Option) ProtoMessage() {}

func (x *ProbeConf_Option) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeConf_Option.ProtoReflect.Descriptor instead.
func (*ProbeConf_Option) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProbeConf_Option) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ProbeConf_Option) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

var File_github_com_cloudprober_cloudprober_probes_external_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x03, 0x0a, 0x09,
	0x50, 0x72, 0x6f, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x45, 0x0a, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x3a, 0x04, 0x4f, 0x4e, 0x43, 0x45, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x47, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x30, 0x0a, 0x11, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x61, 0x73,
	0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x04,
	0x74, 0x72, 0x75, 0x65, 0x52, 0x0f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x41, 0x73, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x67, 0x0a, 0x16, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x72, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x14, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x32,
	0x0a, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x1c, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x4e,
	0x43, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x01,
	0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_rawDescData = file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_goTypes = []interface{}{
	(ProbeConf_Mode)(0),                // 0: cloudprober.probes.external.ProbeConf.Mode
	(*ProbeConf)(nil),                  // 1: cloudprober.probes.external.ProbeConf
	(*ProbeConf_Option)(nil),           // 2: cloudprober.probes.external.ProbeConf.Option
	(*proto.OutputMetricsOptions)(nil), // 3: cloudprober.metrics.payload.OutputMetricsOptions
}
var file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_depIdxs = []int32{
	0, // 0: cloudprober.probes.external.ProbeConf.mode:type_name -> cloudprober.probes.external.ProbeConf.Mode
	2, // 1: cloudprober.probes.external.ProbeConf.options:type_name -> cloudprober.probes.external.ProbeConf.Option
	3, // 2: cloudprober.probes.external.ProbeConf.output_metrics_options:type_name -> cloudprober.metrics.payload.OutputMetricsOptions
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_init() }
func file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_probes_external_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeConf); i {
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
		file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeConf_Option); i {
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
			RawDescriptor: file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_depIdxs,
		EnumInfos:         file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_enumTypes,
		MessageInfos:      file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_probes_external_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_probes_external_proto_config_proto_depIdxs = nil
}
