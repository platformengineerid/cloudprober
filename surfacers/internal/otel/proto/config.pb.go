// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.5
// source: github.com/cloudprober/cloudprober/surfacers/internal/otel/proto/config.proto

package proto

import (
	proto "github.com/cloudprober/cloudprober/internal/tlsconfig/proto"
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

type Compression int32

const (
	Compression_NONE Compression = 0
	Compression_GZIP Compression = 1
)

// Enum value maps for Compression.
var (
	Compression_name = map[int32]string{
		0: "NONE",
		1: "GZIP",
	}
	Compression_value = map[string]int32{
		"NONE": 0,
		"GZIP": 1,
	}
)

func (x Compression) Enum() *Compression {
	p := new(Compression)
	*p = x
	return p
}

func (x Compression) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Compression) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_enumTypes[0].Descriptor()
}

func (Compression) Type() protoreflect.EnumType {
	return &file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_enumTypes[0]
}

func (x Compression) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Compression) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Compression(num)
	return nil
}

// Deprecated: Use Compression.Descriptor instead.
func (Compression) EnumDescriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_rawDescGZIP(), []int{0}
}

type HTTPExporter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If no URL is provided, OpenTelemetry SDK will use the environment variable
	// OTEL_EXPORTER_OTLP_METRICS_ENDPOINT or OTEL_EXPORTER_OTLP_ENDPOINT in that
	// preference order.
	EndpointUrl *string          `protobuf:"bytes,1,opt,name=endpoint_url,json=endpointUrl" json:"endpoint_url,omitempty"`
	TlsConfig   *proto.TLSConfig `protobuf:"bytes,2,opt,name=tls_config,json=tlsConfig" json:"tls_config,omitempty"`
	// HTTP request headers. These can also be set using environment variables.
	HttpHeader map[string]string `protobuf:"bytes,3,rep,name=http_header,json=httpHeader" json:"http_header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Compression algorithm to use for HTTP requests.
	Compression *Compression `protobuf:"varint,4,opt,name=compression,enum=cloudprober.surfacer.otel.Compression" json:"compression,omitempty"`
}

func (x *HTTPExporter) Reset() {
	*x = HTTPExporter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPExporter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPExporter) ProtoMessage() {}

func (x *HTTPExporter) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPExporter.ProtoReflect.Descriptor instead.
func (*HTTPExporter) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *HTTPExporter) GetEndpointUrl() string {
	if x != nil && x.EndpointUrl != nil {
		return *x.EndpointUrl
	}
	return ""
}

func (x *HTTPExporter) GetTlsConfig() *proto.TLSConfig {
	if x != nil {
		return x.TlsConfig
	}
	return nil
}

func (x *HTTPExporter) GetHttpHeader() map[string]string {
	if x != nil {
		return x.HttpHeader
	}
	return nil
}

func (x *HTTPExporter) GetCompression() Compression {
	if x != nil && x.Compression != nil {
		return *x.Compression
	}
	return Compression_NONE
}

type GRPCExporter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If no URL is provided, OpenTelemetry SDK will use the environment variable
	// OTEL_EXPORTER_OTLP_METRICS_ENDPOINT or OTEL_EXPORTER_OTLP_ENDPOINT in that
	// preference order.
	Endpoint  *string          `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	TlsConfig *proto.TLSConfig `protobuf:"bytes,2,opt,name=tls_config,json=tlsConfig" json:"tls_config,omitempty"`
	// HTTP request headers. These can also be set using environment variables.
	HttpHeader map[string]string `protobuf:"bytes,3,rep,name=http_header,json=httpHeader" json:"http_header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Compression algorithm to use for gRPC requests.
	Compression *Compression `protobuf:"varint,4,opt,name=compression,enum=cloudprober.surfacer.otel.Compression" json:"compression,omitempty"`
	// Whether to use insecure gRPC connection.
	Insecure *bool `protobuf:"varint,5,opt,name=insecure" json:"insecure,omitempty"`
}

func (x *GRPCExporter) Reset() {
	*x = GRPCExporter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPCExporter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCExporter) ProtoMessage() {}

func (x *GRPCExporter) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCExporter.ProtoReflect.Descriptor instead.
func (*GRPCExporter) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_rawDescGZIP(), []int{1}
}

func (x *GRPCExporter) GetEndpoint() string {
	if x != nil && x.Endpoint != nil {
		return *x.Endpoint
	}
	return ""
}

func (x *GRPCExporter) GetTlsConfig() *proto.TLSConfig {
	if x != nil {
		return x.TlsConfig
	}
	return nil
}

func (x *GRPCExporter) GetHttpHeader() map[string]string {
	if x != nil {
		return x.HttpHeader
	}
	return nil
}

func (x *GRPCExporter) GetCompression() Compression {
	if x != nil && x.Compression != nil {
		return *x.Compression
	}
	return Compression_NONE
}

func (x *GRPCExporter) GetInsecure() bool {
	if x != nil && x.Insecure != nil {
		return *x.Insecure
	}
	return false
}

type SurfacerConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Exporter:
	//
	//	*SurfacerConf_OtlpHttpExporter
	//	*SurfacerConf_OtlpGrpcExporter
	Exporter isSurfacerConf_Exporter `protobuf_oneof:"exporter"`
	// How often metrics will be exported. Note that metrics are accumulated
	// internally and exported at this interval. Increasing this value will
	// increase the memory usage.
	ExportIntervalSec *int32 `protobuf:"varint,3,opt,name=export_interval_sec,json=exportIntervalSec,def=10" json:"export_interval_sec,omitempty"`
	// Prefix to use for metrics. Defaults to "cloudprober_".
	MetricsPrefix *string `protobuf:"bytes,4,opt,name=metrics_prefix,json=metricsPrefix,def=cloudprober_" json:"metrics_prefix,omitempty"`
}

// Default values for SurfacerConf fields.
const (
	Default_SurfacerConf_ExportIntervalSec = int32(10)
	Default_SurfacerConf_MetricsPrefix     = string("cloudprober_")
)

func (x *SurfacerConf) Reset() {
	*x = SurfacerConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurfacerConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurfacerConf) ProtoMessage() {}

func (x *SurfacerConf) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurfacerConf.ProtoReflect.Descriptor instead.
func (*SurfacerConf) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_rawDescGZIP(), []int{2}
}

func (m *SurfacerConf) GetExporter() isSurfacerConf_Exporter {
	if m != nil {
		return m.Exporter
	}
	return nil
}

func (x *SurfacerConf) GetOtlpHttpExporter() *HTTPExporter {
	if x, ok := x.GetExporter().(*SurfacerConf_OtlpHttpExporter); ok {
		return x.OtlpHttpExporter
	}
	return nil
}

func (x *SurfacerConf) GetOtlpGrpcExporter() *GRPCExporter {
	if x, ok := x.GetExporter().(*SurfacerConf_OtlpGrpcExporter); ok {
		return x.OtlpGrpcExporter
	}
	return nil
}

func (x *SurfacerConf) GetExportIntervalSec() int32 {
	if x != nil && x.ExportIntervalSec != nil {
		return *x.ExportIntervalSec
	}
	return Default_SurfacerConf_ExportIntervalSec
}

func (x *SurfacerConf) GetMetricsPrefix() string {
	if x != nil && x.MetricsPrefix != nil {
		return *x.MetricsPrefix
	}
	return Default_SurfacerConf_MetricsPrefix
}

type isSurfacerConf_Exporter interface {
	isSurfacerConf_Exporter()
}

type SurfacerConf_OtlpHttpExporter struct {
	// OTLP HTTP exporter.
	OtlpHttpExporter *HTTPExporter `protobuf:"bytes,1,opt,name=otlp_http_exporter,json=otlpHttpExporter,oneof"`
}

type SurfacerConf_OtlpGrpcExporter struct {
	// OTLP gRPC exporter.
	OtlpGrpcExporter *GRPCExporter `protobuf:"bytes,2,opt,name=otlp_grpc_exporter,json=otlpGrpcExporter,oneof"`
}

func (*SurfacerConf_OtlpHttpExporter) isSurfacerConf_Exporter() {}

func (*SurfacerConf_OtlpGrpcExporter) isSurfacerConf_Exporter() {}

var File_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x72, 0x73, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6f, 0x74, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x73, 0x75, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x72, 0x2e, 0x6f, 0x74, 0x65, 0x6c, 0x1a, 0x48, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62,
	0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x6c, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x02, 0x0a, 0x0c, 0x48, 0x54, 0x54, 0x50, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x3f, 0x0a, 0x0a, 0x74, 0x6c, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x6c, 0x73, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09,
	0x74, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x58, 0x0a, 0x0b, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x73, 0x75, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x72, 0x2e, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x72, 0x2e,
	0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x3d, 0x0a,
	0x0f, 0x48, 0x74, 0x74, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xea, 0x02, 0x0a,
	0x0c, 0x47, 0x52, 0x50, 0x43, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x0a, 0x74, 0x6c, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x6c, 0x73, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x09, 0x74, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x58, 0x0a, 0x0b, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x73, 0x75,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x72, 0x2e, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x47, 0x52, 0x50, 0x43,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x72,
	0x2e, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x1a, 0x3d, 0x0a, 0x0f, 0x48, 0x74,
	0x74, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb5, 0x02, 0x0a, 0x0c, 0x53, 0x75,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x57, 0x0a, 0x12, 0x6f, 0x74,
	0x6c, 0x70, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x72, 0x2e, 0x6f, 0x74,
	0x65, 0x6c, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x48,
	0x00, 0x52, 0x10, 0x6f, 0x74, 0x6c, 0x70, 0x48, 0x74, 0x74, 0x70, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x12, 0x57, 0x0a, 0x12, 0x6f, 0x74, 0x6c, 0x70, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x73, 0x75,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x72, 0x2e, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x47, 0x52, 0x50, 0x43,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x10, 0x6f, 0x74, 0x6c, 0x70,
	0x47, 0x72, 0x70, 0x63, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x13,
	0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f,
	0x73, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x31, 0x30, 0x52, 0x11, 0x65,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63,
	0x12, 0x33, 0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x0c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x72, 0x5f, 0x52, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x42, 0x0a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x2a, 0x21, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x5a,
	0x49, 0x50, 0x10, 0x01, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x75, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x72, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6f, 0x74,
	0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_rawDescData = file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_goTypes = []interface{}{
	(Compression)(0),        // 0: cloudprober.surfacer.otel.Compression
	(*HTTPExporter)(nil),    // 1: cloudprober.surfacer.otel.HTTPExporter
	(*GRPCExporter)(nil),    // 2: cloudprober.surfacer.otel.GRPCExporter
	(*SurfacerConf)(nil),    // 3: cloudprober.surfacer.otel.SurfacerConf
	nil,                     // 4: cloudprober.surfacer.otel.HTTPExporter.HttpHeaderEntry
	nil,                     // 5: cloudprober.surfacer.otel.GRPCExporter.HttpHeaderEntry
	(*proto.TLSConfig)(nil), // 6: cloudprober.tlsconfig.TLSConfig
}
var file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_depIdxs = []int32{
	6, // 0: cloudprober.surfacer.otel.HTTPExporter.tls_config:type_name -> cloudprober.tlsconfig.TLSConfig
	4, // 1: cloudprober.surfacer.otel.HTTPExporter.http_header:type_name -> cloudprober.surfacer.otel.HTTPExporter.HttpHeaderEntry
	0, // 2: cloudprober.surfacer.otel.HTTPExporter.compression:type_name -> cloudprober.surfacer.otel.Compression
	6, // 3: cloudprober.surfacer.otel.GRPCExporter.tls_config:type_name -> cloudprober.tlsconfig.TLSConfig
	5, // 4: cloudprober.surfacer.otel.GRPCExporter.http_header:type_name -> cloudprober.surfacer.otel.GRPCExporter.HttpHeaderEntry
	0, // 5: cloudprober.surfacer.otel.GRPCExporter.compression:type_name -> cloudprober.surfacer.otel.Compression
	1, // 6: cloudprober.surfacer.otel.SurfacerConf.otlp_http_exporter:type_name -> cloudprober.surfacer.otel.HTTPExporter
	2, // 7: cloudprober.surfacer.otel.SurfacerConf.otlp_grpc_exporter:type_name -> cloudprober.surfacer.otel.GRPCExporter
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() {
	file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_init()
}
func file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPExporter); i {
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
		file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRPCExporter); i {
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
		file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurfacerConf); i {
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
	file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*SurfacerConf_OtlpHttpExporter)(nil),
		(*SurfacerConf_OtlpGrpcExporter)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_depIdxs,
		EnumInfos:         file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_enumTypes,
		MessageInfos:      file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_surfacers_internal_otel_proto_config_proto_depIdxs = nil
}
