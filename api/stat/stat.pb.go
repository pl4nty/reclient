// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: api/stat/stat.proto

package stat

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

type Outlier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandId string `protobuf:"bytes,1,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	Value     int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Outlier) Reset() {
	*x = Outlier{}
	mi := &file_api_stat_stat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Outlier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outlier) ProtoMessage() {}

func (x *Outlier) ProtoReflect() protoreflect.Message {
	mi := &file_api_stat_stat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outlier.ProtoReflect.Descriptor instead.
func (*Outlier) Descriptor() ([]byte, []int) {
	return file_api_stat_stat_proto_rawDescGZIP(), []int{0}
}

func (x *Outlier) GetCommandId() string {
	if x != nil {
		return x.CommandId
	}
	return ""
}

func (x *Outlier) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Stat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count         int64         `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	CountsByValue []*Stat_Value `protobuf:"bytes,3,rep,name=counts_by_value,json=countsByValue,proto3" json:"counts_by_value,omitempty"`
	Outliers      []*Outlier    `protobuf:"bytes,4,rep,name=outliers,proto3" json:"outliers,omitempty"`
	Median        int64         `protobuf:"varint,5,opt,name=median,proto3" json:"median,omitempty"`
	Percentile75  int64         `protobuf:"varint,6,opt,name=percentile75,proto3" json:"percentile75,omitempty"`
	Percentile85  int64         `protobuf:"varint,7,opt,name=percentile85,proto3" json:"percentile85,omitempty"`
	Percentile95  int64         `protobuf:"varint,8,opt,name=percentile95,proto3" json:"percentile95,omitempty"`
	Average       float64       `protobuf:"fixed64,9,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *Stat) Reset() {
	*x = Stat{}
	mi := &file_api_stat_stat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Stat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stat) ProtoMessage() {}

func (x *Stat) ProtoReflect() protoreflect.Message {
	mi := &file_api_stat_stat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stat.ProtoReflect.Descriptor instead.
func (*Stat) Descriptor() ([]byte, []int) {
	return file_api_stat_stat_proto_rawDescGZIP(), []int{1}
}

func (x *Stat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Stat) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Stat) GetCountsByValue() []*Stat_Value {
	if x != nil {
		return x.CountsByValue
	}
	return nil
}

func (x *Stat) GetOutliers() []*Outlier {
	if x != nil {
		return x.Outliers
	}
	return nil
}

func (x *Stat) GetMedian() int64 {
	if x != nil {
		return x.Median
	}
	return 0
}

func (x *Stat) GetPercentile75() int64 {
	if x != nil {
		return x.Percentile75
	}
	return 0
}

func (x *Stat) GetPercentile85() int64 {
	if x != nil {
		return x.Percentile85
	}
	return 0
}

func (x *Stat) GetPercentile95() int64 {
	if x != nil {
		return x.Percentile95
	}
	return 0
}

func (x *Stat) GetAverage() float64 {
	if x != nil {
		return x.Average
	}
	return 0
}

type Stat_Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Stat_Value) Reset() {
	*x = Stat_Value{}
	mi := &file_api_stat_stat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Stat_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stat_Value) ProtoMessage() {}

func (x *Stat_Value) ProtoReflect() protoreflect.Message {
	mi := &file_api_stat_stat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stat_Value.ProtoReflect.Descriptor instead.
func (*Stat_Value) Descriptor() ([]byte, []int) {
	return file_api_stat_stat_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Stat_Value) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Stat_Value) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_api_stat_stat_proto protoreflect.FileDescriptor

var file_api_stat_stat_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x3e, 0x0a, 0x07,
	0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xe8, 0x02, 0x0a,
	0x04, 0x53, 0x74, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x39, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x42, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x6f, 0x75,
	0x74, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x08, 0x6f, 0x75,
	0x74, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x12, 0x22,
	0x0a, 0x0c, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x37, 0x35, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65,
	0x37, 0x35, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65,
	0x38, 0x35, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x69, 0x6c, 0x65, 0x38, 0x35, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x69, 0x6c, 0x65, 0x39, 0x35, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x39, 0x35, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x1a, 0x31, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2f, 0x72, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_stat_stat_proto_rawDescOnce sync.Once
	file_api_stat_stat_proto_rawDescData = file_api_stat_stat_proto_rawDesc
)

func file_api_stat_stat_proto_rawDescGZIP() []byte {
	file_api_stat_stat_proto_rawDescOnce.Do(func() {
		file_api_stat_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_stat_stat_proto_rawDescData)
	})
	return file_api_stat_stat_proto_rawDescData
}

var file_api_stat_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_stat_stat_proto_goTypes = []any{
	(*Outlier)(nil),    // 0: stats.Outlier
	(*Stat)(nil),       // 1: stats.Stat
	(*Stat_Value)(nil), // 2: stats.Stat.Value
}
var file_api_stat_stat_proto_depIdxs = []int32{
	2, // 0: stats.Stat.counts_by_value:type_name -> stats.Stat.Value
	0, // 1: stats.Stat.outliers:type_name -> stats.Outlier
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_stat_stat_proto_init() }
func file_api_stat_stat_proto_init() {
	if File_api_stat_stat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_stat_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_stat_stat_proto_goTypes,
		DependencyIndexes: file_api_stat_stat_proto_depIdxs,
		MessageInfos:      file_api_stat_stat_proto_msgTypes,
	}.Build()
	File_api_stat_stat_proto = out.File
	file_api_stat_stat_proto_rawDesc = nil
	file_api_stat_stat_proto_goTypes = nil
	file_api_stat_stat_proto_depIdxs = nil
}
