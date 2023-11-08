// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0
// source: proto/price/price.proto

package price

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

type PriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *PriceRequest) Reset() {
	*x = PriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_price_price_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceRequest) ProtoMessage() {}

func (x *PriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_price_price_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceRequest.ProtoReflect.Descriptor instead.
func (*PriceRequest) Descriptor() ([]byte, []int) {
	return file_proto_price_price_proto_rawDescGZIP(), []int{0}
}

func (x *PriceRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type PriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date  string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Price int64  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *PriceResponse) Reset() {
	*x = PriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_price_price_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceResponse) ProtoMessage() {}

func (x *PriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_price_price_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceResponse.ProtoReflect.Descriptor instead.
func (*PriceResponse) Descriptor() ([]byte, []int) {
	return file_proto_price_price_proto_rawDescGZIP(), []int{1}
}

func (x *PriceResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *PriceResponse) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_proto_price_price_proto protoreflect.FileDescriptor

var file_proto_price_price_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x39, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32,
	0x45, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x35, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_price_price_proto_rawDescOnce sync.Once
	file_proto_price_price_proto_rawDescData = file_proto_price_price_proto_rawDesc
)

func file_proto_price_price_proto_rawDescGZIP() []byte {
	file_proto_price_price_proto_rawDescOnce.Do(func() {
		file_proto_price_price_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_price_price_proto_rawDescData)
	})
	return file_proto_price_price_proto_rawDescData
}

var file_proto_price_price_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_price_price_proto_goTypes = []interface{}{
	(*PriceRequest)(nil),  // 0: price.PriceRequest
	(*PriceResponse)(nil), // 1: price.PriceResponse
}
var file_proto_price_price_proto_depIdxs = []int32{
	0, // 0: price.PriceService.GetPrice:input_type -> price.PriceRequest
	1, // 1: price.PriceService.GetPrice:output_type -> price.PriceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_price_price_proto_init() }
func file_proto_price_price_proto_init() {
	if File_proto_price_price_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_price_price_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceRequest); i {
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
		file_proto_price_price_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceResponse); i {
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
			RawDescriptor: file_proto_price_price_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_price_price_proto_goTypes,
		DependencyIndexes: file_proto_price_price_proto_depIdxs,
		MessageInfos:      file_proto_price_price_proto_msgTypes,
	}.Build()
	File_proto_price_price_proto = out.File
	file_proto_price_price_proto_rawDesc = nil
	file_proto_price_price_proto_goTypes = nil
	file_proto_price_price_proto_depIdxs = nil
}
