// Code generated by protoc-gen-go.
// source: kairosdb.proto
// DO NOT EDIT!

/*
Package kairosdb is a generated protocol buffer package.

It is generated from these files:
	kairosdb.proto
	types.proto

It has these top-level messages:
	Datapoint
	Sampling
	Aggregator
	GroupBy
	QueryMetric
	Result
	Query
	QueryMetricsRequest
	QueryMetricsResponse
	StringList
*/
package kairosdb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Represents a kairosdb datapoint
// https://github.com/kairosdb/kairosdb-client/blob/master/src/main/java/org/kairosdb/client/builder/DataPoint.java
type Datapoint struct {
	Timestamp int64   `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Value     float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
}

func (m *Datapoint) Reset()                    { *m = Datapoint{} }
func (m *Datapoint) String() string            { return proto.CompactTextString(m) }
func (*Datapoint) ProtoMessage()               {}
func (*Datapoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Represents a sampling in the QueryMetrics request body
type Sampling struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	Unit  string `protobuf:"bytes,2,opt,name=unit" json:"unit,omitempty"`
}

func (m *Sampling) Reset()                    { *m = Sampling{} }
func (m *Sampling) String() string            { return proto.CompactTextString(m) }
func (*Sampling) ProtoMessage()               {}
func (*Sampling) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Represents a kairosdb aggregator
// https://kairosdb.github.io/docs/build/html/restapi/Aggregators.html
type Aggregator struct {
	Name          string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	AlignSampling bool      `protobuf:"varint,2,opt,name=align_sampling,json=alignSampling" json:"align_sampling,omitempty"`
	Sampling      *Sampling `protobuf:"bytes,3,opt,name=sampling" json:"sampling,omitempty"`
}

func (m *Aggregator) Reset()                    { *m = Aggregator{} }
func (m *Aggregator) String() string            { return proto.CompactTextString(m) }
func (*Aggregator) ProtoMessage()               {}
func (*Aggregator) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Aggregator) GetSampling() *Sampling {
	if m != nil {
		return m.Sampling
	}
	return nil
}

// Represents the group_property in the request
// I've decided that it's not bad
type GroupBy struct {
	Name  string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tags  []string          `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty"`
	Group map[string]string `protobuf:"bytes,3,rep,name=group" json:"group,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GroupBy) Reset()                    { *m = GroupBy{} }
func (m *GroupBy) String() string            { return proto.CompactTextString(m) }
func (*GroupBy) ProtoMessage()               {}
func (*GroupBy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GroupBy) GetGroup() map[string]string {
	if m != nil {
		return m.Group
	}
	return nil
}

// Reprents a kairosdb QueryMetric
// https://kairosdb.github.io/docs/build/html/restapi/QueryMetrics.html#id3
type QueryMetric struct {
	Name        string                 `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tags        map[string]*StringList `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	GroupBy     []*GroupBy             `protobuf:"bytes,3,rep,name=group_by,json=groupBy" json:"group_by,omitempty"`
	Aggregators []*Aggregator          `protobuf:"bytes,4,rep,name=aggregators" json:"aggregators,omitempty"`
	Limit       int64                  `protobuf:"varint,5,opt,name=limit" json:"limit,omitempty"`
}

func (m *QueryMetric) Reset()                    { *m = QueryMetric{} }
func (m *QueryMetric) String() string            { return proto.CompactTextString(m) }
func (*QueryMetric) ProtoMessage()               {}
func (*QueryMetric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *QueryMetric) GetTags() map[string]*StringList {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *QueryMetric) GetGroupBy() []*GroupBy {
	if m != nil {
		return m.GroupBy
	}
	return nil
}

func (m *QueryMetric) GetAggregators() []*Aggregator {
	if m != nil {
		return m.Aggregators
	}
	return nil
}

type Result struct {
	Name    string                 `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	GroupBy []*GroupBy             `protobuf:"bytes,2,rep,name=group_by,json=groupBy" json:"group_by,omitempty"`
	Tags    map[string]*StringList `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Values  []*Datapoint           `protobuf:"bytes,4,rep,name=values" json:"values,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Result) GetGroupBy() []*GroupBy {
	if m != nil {
		return m.GroupBy
	}
	return nil
}

func (m *Result) GetTags() map[string]*StringList {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Result) GetValues() []*Datapoint {
	if m != nil {
		return m.Values
	}
	return nil
}

// Represents a single set of results for a given Query
type Query struct {
	Results []*Result `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Query) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

// Reprents a QueryMetrics request to kairosdb
// https://kairosdb.github.io/docs/build/html/restapi/QueryMetrics.html#query-properties
type QueryMetricsRequest struct {
	Metrics       []*QueryMetric `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
	CacheTime     int64          `protobuf:"varint,2,opt,name=cache_time,json=cacheTime" json:"cache_time,omitempty"`
	StartAbsolute int64          `protobuf:"varint,3,opt,name=start_absolute,json=startAbsolute" json:"start_absolute,omitempty"`
	EndAbsolute   int64          `protobuf:"varint,4,opt,name=end_absolute,json=endAbsolute" json:"end_absolute,omitempty"`
}

func (m *QueryMetricsRequest) Reset()                    { *m = QueryMetricsRequest{} }
func (m *QueryMetricsRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryMetricsRequest) ProtoMessage()               {}
func (*QueryMetricsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *QueryMetricsRequest) GetMetrics() []*QueryMetric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

// Represents response to kairosdb QueryMetrics request
type QueryMetricsResponse struct {
	Queries []*Query `protobuf:"bytes,1,rep,name=queries" json:"queries,omitempty"`
	Errors  []string `protobuf:"bytes,5,rep,name=errors" json:"errors,omitempty"`
}

func (m *QueryMetricsResponse) Reset()                    { *m = QueryMetricsResponse{} }
func (m *QueryMetricsResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryMetricsResponse) ProtoMessage()               {}
func (*QueryMetricsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *QueryMetricsResponse) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

func init() {
	proto.RegisterType((*Datapoint)(nil), "kairosdb.Datapoint")
	proto.RegisterType((*Sampling)(nil), "kairosdb.Sampling")
	proto.RegisterType((*Aggregator)(nil), "kairosdb.Aggregator")
	proto.RegisterType((*GroupBy)(nil), "kairosdb.GroupBy")
	proto.RegisterType((*QueryMetric)(nil), "kairosdb.QueryMetric")
	proto.RegisterType((*Result)(nil), "kairosdb.Result")
	proto.RegisterType((*Query)(nil), "kairosdb.Query")
	proto.RegisterType((*QueryMetricsRequest)(nil), "kairosdb.QueryMetricsRequest")
	proto.RegisterType((*QueryMetricsResponse)(nil), "kairosdb.QueryMetricsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for KairosDB service

type KairosDBClient interface {
	QueryMetrics(ctx context.Context, in *QueryMetricsRequest, opts ...grpc.CallOption) (*QueryMetricsResponse, error)
}

type kairosDBClient struct {
	cc *grpc.ClientConn
}

func NewKairosDBClient(cc *grpc.ClientConn) KairosDBClient {
	return &kairosDBClient{cc}
}

func (c *kairosDBClient) QueryMetrics(ctx context.Context, in *QueryMetricsRequest, opts ...grpc.CallOption) (*QueryMetricsResponse, error) {
	out := new(QueryMetricsResponse)
	err := grpc.Invoke(ctx, "/kairosdb.KairosDB/QueryMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KairosDB service

type KairosDBServer interface {
	QueryMetrics(context.Context, *QueryMetricsRequest) (*QueryMetricsResponse, error)
}

func RegisterKairosDBServer(s *grpc.Server, srv KairosDBServer) {
	s.RegisterService(&_KairosDB_serviceDesc, srv)
}

func _KairosDB_QueryMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KairosDBServer).QueryMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kairosdb.KairosDB/QueryMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KairosDBServer).QueryMetrics(ctx, req.(*QueryMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KairosDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kairosdb.KairosDB",
	HandlerType: (*KairosDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryMetrics",
			Handler:    _KairosDB_QueryMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("kairosdb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xcb, 0x6e, 0xd3, 0x4c,
	0x18, 0x95, 0xe3, 0x26, 0x71, 0x3e, 0xb7, 0xfd, 0xfb, 0x4f, 0x03, 0x8a, 0xa2, 0x96, 0x8b, 0x25,
	0x24, 0x6e, 0x0a, 0x52, 0x82, 0x50, 0xc5, 0x06, 0x35, 0x2a, 0x62, 0x01, 0x59, 0x30, 0xed, 0xa6,
	0x6c, 0xa2, 0x49, 0x3a, 0x32, 0x56, 0x1d, 0xdb, 0xcc, 0x8c, 0x41, 0x7e, 0x1a, 0xc4, 0x13, 0xf0,
	0x76, 0xac, 0x19, 0xcf, 0xc5, 0x76, 0x5a, 0x23, 0x36, 0xac, 0xec, 0x39, 0x73, 0xce, 0x7c, 0xe7,
	0x3b, 0x73, 0x81, 0xfd, 0x6b, 0x12, 0xb1, 0x94, 0x5f, 0xad, 0x26, 0x19, 0x4b, 0x45, 0x8a, 0x3c,
	0x3b, 0x1e, 0xfb, 0xa2, 0xc8, 0x28, 0xd7, 0x70, 0xf0, 0x06, 0x06, 0x67, 0x44, 0x90, 0x2c, 0x8d,
	0x12, 0x81, 0x8e, 0x60, 0x20, 0xa2, 0x0d, 0xe5, 0x82, 0x6c, 0xb2, 0x91, 0xf3, 0xc0, 0x79, 0xec,
	0xe2, 0x1a, 0x40, 0x43, 0xe8, 0x7e, 0x25, 0x71, 0x4e, 0x47, 0x1d, 0x39, 0xe3, 0x60, 0x3d, 0x08,
	0x5e, 0x82, 0x77, 0x2e, 0x67, 0xe3, 0x28, 0x09, 0x6b, 0x46, 0xa9, 0x1d, 0x18, 0x06, 0x42, 0xb0,
	0x93, 0x27, 0x91, 0x50, 0xb2, 0x01, 0x56, 0xff, 0xc1, 0x37, 0x80, 0xd3, 0x30, 0x64, 0x34, 0x24,
	0x22, 0x65, 0x25, 0x23, 0x21, 0x1b, 0x2b, 0x53, 0xff, 0xe8, 0x11, 0xec, 0x93, 0x38, 0x0a, 0x93,
	0x25, 0x37, 0xab, 0x2b, 0xbd, 0x87, 0xf7, 0x14, 0x5a, 0x95, 0x9c, 0x80, 0x57, 0x11, 0x5c, 0x49,
	0xf0, 0xa7, 0x68, 0x52, 0x75, 0x6e, 0x59, 0xb8, 0xe2, 0x04, 0xdf, 0x1d, 0xe8, 0xbf, 0x63, 0x69,
	0x9e, 0xcd, 0x8b, 0xd6, 0xb2, 0x12, 0x13, 0x24, 0xe4, 0xb2, 0x98, 0x5b, 0x62, 0xe5, 0x3f, 0x9a,
	0x42, 0x37, 0x2c, 0x25, 0xb2, 0x80, 0x2b, 0x0b, 0x1c, 0xd5, 0x05, 0xcc, 0x4a, 0xfa, 0xfb, 0x36,
	0x11, 0xac, 0xc0, 0x9a, 0x3a, 0x3e, 0x01, 0xa8, 0x41, 0x74, 0x00, 0xee, 0x35, 0x2d, 0x4c, 0xa1,
	0xf2, 0x77, 0x3b, 0x4c, 0x1b, 0xd5, 0xeb, 0xce, 0x89, 0x13, 0xfc, 0xe8, 0x80, 0xff, 0x31, 0xa7,
	0xac, 0x58, 0x50, 0xc1, 0xa2, 0x75, 0xab, 0xcb, 0x59, 0xc3, 0xa5, 0x3f, 0xbd, 0x5f, 0x1b, 0x6a,
	0x08, 0x27, 0x17, 0x92, 0xa1, 0x3d, 0xe9, 0x36, 0x9e, 0x83, 0xa7, 0xbc, 0x2d, 0x57, 0x85, 0xe9,
	0xe4, 0xff, 0x5b, 0x9d, 0xe0, 0x7e, 0x68, 0xc2, 0x79, 0x05, 0x3e, 0xa9, 0x76, 0x88, 0x8f, 0x76,
	0x94, 0x60, 0x58, 0x0b, 0xea, 0xed, 0xc3, 0x4d, 0x62, 0xd9, 0x58, 0x1c, 0x6d, 0xe4, 0x76, 0x77,
	0xd5, 0xf9, 0xd1, 0x83, 0xf1, 0x02, 0x06, 0x95, 0x9d, 0x96, 0x34, 0x9e, 0x36, 0xd3, 0xd8, 0x2a,
	0x73, 0x2e, 0x5b, 0x49, 0xc2, 0x0f, 0x11, 0x17, 0xcd, 0x8c, 0x7e, 0x39, 0xd0, 0xc3, 0x94, 0xe7,
	0xb1, 0x68, 0x8d, 0xa7, 0xd9, 0x69, 0xe7, 0xaf, 0x9d, 0x4e, 0x4c, 0x98, 0x3a, 0x93, 0x71, 0xcd,
	0xd4, 0x15, 0x6e, 0xe5, 0xf8, 0x0c, 0x7a, 0xca, 0x89, 0x0d, 0xe5, 0xb0, 0x56, 0x54, 0x57, 0x09,
	0x1b, 0xca, 0xbf, 0x6e, 0x7c, 0x06, 0x5d, 0xb5, 0xc5, 0x52, 0xd8, 0x67, 0xca, 0x1e, 0x97, 0xcb,
	0x95, 0x2e, 0x0e, 0x6e, 0xfa, 0xc6, 0x96, 0x10, 0xfc, 0x74, 0xe0, 0xb0, 0x71, 0x30, 0x38, 0xa6,
	0x5f, 0xa4, 0x35, 0x81, 0x5e, 0x40, 0x7f, 0xa3, 0x11, 0xb3, 0xc6, 0x9d, 0xd6, 0x83, 0x84, 0x2d,
	0x0b, 0x1d, 0x03, 0xac, 0xc9, 0xfa, 0x33, 0x5d, 0x96, 0x8f, 0x82, 0xb2, 0x2c, 0x1f, 0x08, 0x85,
	0x5c, 0x44, 0xfa, 0xca, 0xca, 0x97, 0x82, 0x89, 0x25, 0x59, 0xf1, 0x34, 0xce, 0x05, 0x55, 0x37,
	0xd2, 0xc5, 0x7b, 0x0a, 0x3d, 0x35, 0x20, 0x7a, 0x08, 0xbb, 0x34, 0xb9, 0xaa, 0x49, 0x3b, 0x8a,
	0xe4, 0x4b, 0xcc, 0x52, 0x82, 0x4b, 0x18, 0x6e, 0x1b, 0xe6, 0x59, 0x9a, 0x70, 0x8a, 0x9e, 0x40,
	0x5f, 0x5a, 0x67, 0x11, 0xb5, 0x8e, 0xff, 0xbb, 0xe1, 0x18, 0xdb, 0x79, 0x74, 0x17, 0x7a, 0x94,
	0xb1, 0xf2, 0xe8, 0x76, 0xd5, 0x55, 0x36, 0xa3, 0xe9, 0x25, 0x78, 0xef, 0x95, 0xe4, 0x6c, 0x8e,
	0x16, 0xb0, 0xdb, 0x2c, 0x83, 0x8e, 0x5b, 0xfb, 0xb7, 0x79, 0x8d, 0xef, 0xfd, 0x69, 0x5a, 0xbb,
	0x9b, 0xc3, 0xa7, 0xea, 0x91, 0x5d, 0xf5, 0xd4, 0xf3, 0x3a, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0xcd, 0xc0, 0xa0, 0x27, 0x87, 0x05, 0x00, 0x00,
}
