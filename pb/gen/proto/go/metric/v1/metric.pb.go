// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: metric/v1/metric.proto

package metricv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Counter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Counter) Reset() {
	*x = Counter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metric_v1_metric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Counter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Counter) ProtoMessage() {}

func (x *Counter) ProtoReflect() protoreflect.Message {
	mi := &file_metric_v1_metric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Counter.ProtoReflect.Descriptor instead.
func (*Counter) Descriptor() ([]byte, []int) {
	return file_metric_v1_metric_proto_rawDescGZIP(), []int{0}
}

func (x *Counter) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Gauge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Gauge) Reset() {
	*x = Gauge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metric_v1_metric_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gauge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gauge) ProtoMessage() {}

func (x *Gauge) ProtoReflect() protoreflect.Message {
	mi := &file_metric_v1_metric_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gauge.ProtoReflect.Descriptor instead.
func (*Gauge) Descriptor() ([]byte, []int) {
	return file_metric_v1_metric_proto_rawDescGZIP(), []int{1}
}

func (x *Gauge) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are assignable to Value:
	//
	//	*Metric_Counter
	//	*Metric_Gauge
	Value isMetric_Value `protobuf_oneof:"value"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metric_v1_metric_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_metric_v1_metric_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_metric_v1_metric_proto_rawDescGZIP(), []int{2}
}

func (x *Metric) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (m *Metric) GetValue() isMetric_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Metric) GetCounter() *Counter {
	if x, ok := x.GetValue().(*Metric_Counter); ok {
		return x.Counter
	}
	return nil
}

func (x *Metric) GetGauge() *Gauge {
	if x, ok := x.GetValue().(*Metric_Gauge); ok {
		return x.Gauge
	}
	return nil
}

type isMetric_Value interface {
	isMetric_Value()
}

type Metric_Counter struct {
	Counter *Counter `protobuf:"bytes,2,opt,name=counter,proto3,oneof"`
}

type Metric_Gauge struct {
	Gauge *Gauge `protobuf:"bytes,3,opt,name=gauge,proto3,oneof"`
}

func (*Metric_Counter) isMetric_Value() {}

func (*Metric_Gauge) isMetric_Value() {}

var File_metric_v1_metric_proto protoreflect.FileDescriptor

var file_metric_v1_metric_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x47, 0x61, 0x75, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xa5, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12,
	0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x48, 0x00,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x05, 0x67, 0x61, 0x75,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x75, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x67, 0x61,
	0x75, 0x67, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0xa5, 0x01, 0x0a,
	0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x42, 0x0b,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x69, 0x72,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x74, 0x72, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x2f, 0x70,
	0x62, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x15, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metric_v1_metric_proto_rawDescOnce sync.Once
	file_metric_v1_metric_proto_rawDescData = file_metric_v1_metric_proto_rawDesc
)

func file_metric_v1_metric_proto_rawDescGZIP() []byte {
	file_metric_v1_metric_proto_rawDescOnce.Do(func() {
		file_metric_v1_metric_proto_rawDescData = protoimpl.X.CompressGZIP(file_metric_v1_metric_proto_rawDescData)
	})
	return file_metric_v1_metric_proto_rawDescData
}

var file_metric_v1_metric_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_metric_v1_metric_proto_goTypes = []interface{}{
	(*Counter)(nil),               // 0: metric.v1.Counter
	(*Gauge)(nil),                 // 1: metric.v1.Gauge
	(*Metric)(nil),                // 2: metric.v1.Metric
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_metric_v1_metric_proto_depIdxs = []int32{
	3, // 0: metric.v1.Metric.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: metric.v1.Metric.counter:type_name -> metric.v1.Counter
	1, // 2: metric.v1.Metric.gauge:type_name -> metric.v1.Gauge
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_metric_v1_metric_proto_init() }
func file_metric_v1_metric_proto_init() {
	if File_metric_v1_metric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metric_v1_metric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Counter); i {
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
		file_metric_v1_metric_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gauge); i {
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
		file_metric_v1_metric_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
	file_metric_v1_metric_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Metric_Counter)(nil),
		(*Metric_Gauge)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_metric_v1_metric_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_metric_v1_metric_proto_goTypes,
		DependencyIndexes: file_metric_v1_metric_proto_depIdxs,
		MessageInfos:      file_metric_v1_metric_proto_msgTypes,
	}.Build()
	File_metric_v1_metric_proto = out.File
	file_metric_v1_metric_proto_rawDesc = nil
	file_metric_v1_metric_proto_goTypes = nil
	file_metric_v1_metric_proto_depIdxs = nil
}
