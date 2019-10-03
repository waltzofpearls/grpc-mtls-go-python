// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/metrics.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type QueryMetricsRequest struct {
	MetricName           string               `protobuf:"bytes,1,opt,name=metricName,proto3" json:"metricName,omitempty"`
	Labels               map[string]string    `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Start                *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QueryMetricsRequest) Reset()         { *m = QueryMetricsRequest{} }
func (m *QueryMetricsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryMetricsRequest) ProtoMessage()    {}
func (*QueryMetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82f4d3bdff1c5591, []int{0}
}

func (m *QueryMetricsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryMetricsRequest.Unmarshal(m, b)
}
func (m *QueryMetricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryMetricsRequest.Marshal(b, m, deterministic)
}
func (m *QueryMetricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMetricsRequest.Merge(m, src)
}
func (m *QueryMetricsRequest) XXX_Size() int {
	return xxx_messageInfo_QueryMetricsRequest.Size(m)
}
func (m *QueryMetricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMetricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMetricsRequest proto.InternalMessageInfo

func (m *QueryMetricsRequest) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

func (m *QueryMetricsRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *QueryMetricsRequest) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *QueryMetricsRequest) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

type QueryMetricsResponse struct {
	Metrics              []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QueryMetricsResponse) Reset()         { *m = QueryMetricsResponse{} }
func (m *QueryMetricsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryMetricsResponse) ProtoMessage()    {}
func (*QueryMetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82f4d3bdff1c5591, []int{1}
}

func (m *QueryMetricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryMetricsResponse.Unmarshal(m, b)
}
func (m *QueryMetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryMetricsResponse.Marshal(b, m, deterministic)
}
func (m *QueryMetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMetricsResponse.Merge(m, src)
}
func (m *QueryMetricsResponse) XXX_Size() int {
	return xxx_messageInfo_QueryMetricsResponse.Size(m)
}
func (m *QueryMetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMetricsResponse proto.InternalMessageInfo

func (m *QueryMetricsResponse) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type Metric struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string    `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Values               []*Metric_SamplePair `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_82f4d3bdff1c5591, []int{2}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metric) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Metric) GetValues() []*Metric_SamplePair {
	if m != nil {
		return m.Values
	}
	return nil
}

type Metric_SamplePair struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Value                float64              `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Metric_SamplePair) Reset()         { *m = Metric_SamplePair{} }
func (m *Metric_SamplePair) String() string { return proto.CompactTextString(m) }
func (*Metric_SamplePair) ProtoMessage()    {}
func (*Metric_SamplePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_82f4d3bdff1c5591, []int{2, 1}
}

func (m *Metric_SamplePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric_SamplePair.Unmarshal(m, b)
}
func (m *Metric_SamplePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric_SamplePair.Marshal(b, m, deterministic)
}
func (m *Metric_SamplePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric_SamplePair.Merge(m, src)
}
func (m *Metric_SamplePair) XXX_Size() int {
	return xxx_messageInfo_Metric_SamplePair.Size(m)
}
func (m *Metric_SamplePair) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric_SamplePair.DiscardUnknown(m)
}

var xxx_messageInfo_Metric_SamplePair proto.InternalMessageInfo

func (m *Metric_SamplePair) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Metric_SamplePair) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryMetricsRequest)(nil), "api.QueryMetricsRequest")
	proto.RegisterMapType((map[string]string)(nil), "api.QueryMetricsRequest.LabelsEntry")
	proto.RegisterType((*QueryMetricsResponse)(nil), "api.QueryMetricsResponse")
	proto.RegisterType((*Metric)(nil), "api.Metric")
	proto.RegisterMapType((map[string]string)(nil), "api.Metric.LabelsEntry")
	proto.RegisterType((*Metric_SamplePair)(nil), "api.Metric.SamplePair")
}

func init() { proto.RegisterFile("api/metrics.proto", fileDescriptor_82f4d3bdff1c5591) }

var fileDescriptor_82f4d3bdff1c5591 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x4f, 0xe3, 0x30,
	0x14, 0x54, 0x92, 0x36, 0xd5, 0xbe, 0x5c, 0x76, 0xdf, 0x56, 0xbb, 0x26, 0x07, 0xa8, 0x2a, 0x90,
	0x7a, 0x40, 0x0e, 0x2a, 0x17, 0x40, 0xe5, 0x88, 0xb8, 0x00, 0x02, 0xc3, 0x1f, 0x70, 0xe1, 0x51,
	0x45, 0xe4, 0xc3, 0xc4, 0x0e, 0x52, 0xff, 0x03, 0xe2, 0x37, 0xa3, 0xda, 0xa9, 0x9a, 0xa2, 0x02,
	0x07, 0x6e, 0xf6, 0xf3, 0xcc, 0x78, 0x66, 0xf4, 0xe0, 0x8f, 0x54, 0x69, 0x92, 0x93, 0xa9, 0xd2,
	0x7b, 0xcd, 0x55, 0x55, 0x9a, 0x12, 0x03, 0xa9, 0xd2, 0x78, 0x67, 0x56, 0x96, 0xb3, 0x8c, 0x12,
	0x3b, 0x9a, 0xd6, 0x8f, 0x89, 0x49, 0x73, 0xd2, 0x46, 0xe6, 0xca, 0xa1, 0x86, 0x6f, 0x3e, 0xfc,
	0xbd, 0xa9, 0xa9, 0x9a, 0x5f, 0x3a, 0xb2, 0xa0, 0xe7, 0x9a, 0xb4, 0xc1, 0x6d, 0x00, 0x27, 0x77,
	0x25, 0x73, 0x62, 0xde, 0xc0, 0x1b, 0xfd, 0x12, 0xad, 0x09, 0x4e, 0x20, 0xcc, 0xe4, 0x94, 0x32,
	0xcd, 0xfc, 0x41, 0x30, 0x8a, 0xc6, 0xbb, 0x5c, 0xaa, 0x94, 0x6f, 0x50, 0xe2, 0x17, 0x16, 0x76,
	0x56, 0x98, 0x6a, 0x2e, 0x1a, 0x0e, 0x1e, 0x40, 0x57, 0x1b, 0x59, 0x19, 0x16, 0x0c, 0xbc, 0x51,
	0x34, 0x8e, 0xb9, 0xb3, 0xc9, 0x97, 0x36, 0xf9, 0xdd, 0xd2, 0xa6, 0x70, 0x40, 0xdc, 0x87, 0x80,
	0x8a, 0x07, 0xd6, 0xf9, 0x16, 0xbf, 0x80, 0xc5, 0xc7, 0x10, 0xb5, 0xbe, 0xc5, 0xdf, 0x10, 0x3c,
	0xd1, 0xbc, 0x49, 0xb1, 0x38, 0x62, 0x1f, 0xba, 0x2f, 0x32, 0xab, 0x89, 0xf9, 0x76, 0xe6, 0x2e,
	0x27, 0xfe, 0x91, 0x37, 0x3c, 0x85, 0xfe, 0x7a, 0x0a, 0xad, 0xca, 0x42, 0x13, 0xee, 0x41, 0xaf,
	0xe9, 0x97, 0x79, 0x36, 0x71, 0x64, 0x13, 0x3b, 0x98, 0x58, 0xbe, 0x0d, 0x5f, 0x7d, 0x08, 0xdd,
	0x0c, 0x11, 0x3a, 0xc5, 0xaa, 0x3c, 0x7b, 0xc6, 0xe4, 0x43, 0x6d, 0xff, 0x5b, 0x22, 0x1b, 0x9b,
	0xe2, 0x10, 0x5a, 0x6f, 0x9a, 0x05, 0x96, 0xf0, 0xaf, 0x4d, 0xb8, 0x95, 0xb9, 0xca, 0xe8, 0x5a,
	0xa6, 0x95, 0x68, 0x50, 0x3f, 0x48, 0x1e, 0x0b, 0x80, 0x95, 0x20, 0x72, 0xe8, 0x2c, 0x76, 0xc5,
	0x52, 0xbf, 0x6e, 0xdc, 0xe2, 0xd6, 0x75, 0xbd, 0x46, 0x77, 0x7c, 0x0e, 0xbd, 0xa6, 0x48, 0x9c,
	0x40, 0xd7, 0x16, 0x8b, 0xec, 0xb3, 0x55, 0x89, 0xb7, 0x36, 0xbc, 0xb8, 0xfa, 0xa7, 0xa1, 0xfd,
	0xf8, 0xf0, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x93, 0x83, 0x5e, 0xcf, 0xe9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsClient is the client API for Metrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsClient interface {
	Query(ctx context.Context, in *QueryMetricsRequest, opts ...grpc.CallOption) (*QueryMetricsResponse, error)
}

type metricsClient struct {
	cc *grpc.ClientConn
}

func NewMetricsClient(cc *grpc.ClientConn) MetricsClient {
	return &metricsClient{cc}
}

func (c *metricsClient) Query(ctx context.Context, in *QueryMetricsRequest, opts ...grpc.CallOption) (*QueryMetricsResponse, error) {
	out := new(QueryMetricsResponse)
	err := c.cc.Invoke(ctx, "/api.Metrics/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServer is the server API for Metrics service.
type MetricsServer interface {
	Query(context.Context, *QueryMetricsRequest) (*QueryMetricsResponse, error)
}

// UnimplementedMetricsServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsServer struct {
}

func (*UnimplementedMetricsServer) Query(ctx context.Context, req *QueryMetricsRequest) (*QueryMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

func RegisterMetricsServer(s *grpc.Server, srv MetricsServer) {
	s.RegisterService(&_Metrics_serviceDesc, srv)
}

func _Metrics_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Metrics/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).Query(ctx, req.(*QueryMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metrics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Metrics",
	HandlerType: (*MetricsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Metrics_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/metrics.proto",
}
