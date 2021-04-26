// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: play.proto

package api

import (
	context "context"
	"fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PlayReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
	Nums    int64  `protobuf:"varint,2,opt,name=Nums,json=nums,proto3" json:"Nums,omitempty"`
}

func (x *PlayReq) Reset() {
	*x = PlayReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_play_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayReq) ProtoMessage() {}

func (x *PlayReq) ProtoReflect() protoreflect.Message {
	mi := &file_play_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayReq.ProtoReflect.Descriptor instead.
func (*PlayReq) Descriptor() ([]byte, []int) {
	return file_play_proto_rawDescGZIP(), []int{0}
}

func (x *PlayReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PlayReq) GetNums() int64 {
	if x != nil {
		return x.Nums
	}
	return 0
}

type PlayRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64 `protobuf:"varint,1,opt,name=Code,json=code,proto3" json:"Code,omitempty"`
}

func (x *PlayRes) Reset() {
	*x = PlayRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_play_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayRes) ProtoMessage() {}

func (x *PlayRes) ProtoReflect() protoreflect.Message {
	mi := &file_play_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayRes.ProtoReflect.Descriptor instead.
func (*PlayRes) Descriptor() ([]byte, []int) {
	return file_play_proto_rawDescGZIP(), []int{1}
}

func (x *PlayRes) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

type StopReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
}

func (x *StopReq) Reset() {
	*x = StopReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_play_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopReq) ProtoMessage() {}

func (x *StopReq) ProtoReflect() protoreflect.Message {
	mi := &file_play_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopReq.ProtoReflect.Descriptor instead.
func (*StopReq) Descriptor() ([]byte, []int) {
	return file_play_proto_rawDescGZIP(), []int{2}
}

func (x *StopReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type StopRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=Status,json=status,proto3" json:"Status,omitempty"`
}

func (x *StopRes) Reset() {
	*x = StopRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_play_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRes) ProtoMessage() {}

func (x *StopRes) ProtoReflect() protoreflect.Message {
	mi := &file_play_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRes.ProtoReflect.Descriptor instead.
func (*StopRes) Descriptor() ([]byte, []int) {
	return file_play_proto_rawDescGZIP(), []int{3}
}

func (x *StopRes) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_play_proto protoreflect.FileDescriptor

var file_play_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70,
	0x69, 0x22, 0x37, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x75, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x1d, 0x0a, 0x07, 0x50, 0x6c,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x19, 0x0a, 0x07, 0x53, 0x74, 0x6f,
	0x70, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x55, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x53,
	0x76, 0x63, 0x12, 0x24, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x79, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70,
	0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x0c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x22, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_play_proto_rawDescOnce sync.Once
	file_play_proto_rawDescData = file_play_proto_rawDesc
)

func file_play_proto_rawDescGZIP() []byte {
	file_play_proto_rawDescOnce.Do(func() {
		file_play_proto_rawDescData = protoimpl.X.CompressGZIP(file_play_proto_rawDescData)
	})
	return file_play_proto_rawDescData
}

var file_play_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_play_proto_goTypes = []interface{}{
	(*PlayReq)(nil), // 0: api.PlayReq
	(*PlayRes)(nil), // 1: api.PlayRes
	(*StopReq)(nil), // 2: api.StopReq
	(*StopRes)(nil), // 3: api.StopRes
}
var file_play_proto_depIdxs = []int32{
	0, // 0: api.PlaySvc.Play:input_type -> api.PlayReq
	2, // 1: api.PlaySvc.Stop:input_type -> api.StopReq
	1, // 2: api.PlaySvc.Play:output_type -> api.PlayRes
	3, // 3: api.PlaySvc.Stop:output_type -> api.StopRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_play_proto_init() }
func file_play_proto_init() {
	if File_play_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_play_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayReq); i {
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
		file_play_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayRes); i {
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
		file_play_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopReq); i {
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
		file_play_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopRes); i {
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
			RawDescriptor: file_play_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_play_proto_goTypes,
		DependencyIndexes: file_play_proto_depIdxs,
		MessageInfos:      file_play_proto_msgTypes,
	}.Build()
	File_play_proto = out.File
	file_play_proto_rawDesc = nil
	file_play_proto_goTypes = nil
	file_play_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PlaySvcClient is the client API for PlaySvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlaySvcClient interface {
	Play(ctx context.Context, in *PlayReq, opts ...grpc.CallOption) (*PlayRes, error)
	Stop(ctx context.Context, in *StopReq, opts ...grpc.CallOption) (*StopRes, error)
}

type playSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaySvcClient(cc grpc.ClientConnInterface) PlaySvcClient {
	return &playSvcClient{cc}
}

func (c *playSvcClient) Play(ctx context.Context, in *PlayReq, opts ...grpc.CallOption) (*PlayRes, error) {
	out := new(PlayRes)
	err := c.cc.Invoke(ctx, "/api.PlaySvc/Play", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playSvcClient) Stop(ctx context.Context, in *StopReq, opts ...grpc.CallOption) (*StopRes, error) {
	out := new(StopRes)
	err := c.cc.Invoke(ctx, "/api.PlaySvc/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaySvcServer is the server API for PlaySvc service.
type PlaySvcServer interface {
	Play(context.Context, *PlayReq) (*PlayRes, error)
	Stop(context.Context, *StopReq) (*StopRes, error)
}

// UnimplementedPlaySvcServer can be embedded to have forward compatible implementations.
type UnimplementedPlaySvcServer struct {
}

func (*UnimplementedPlaySvcServer) Play(context.Context, *PlayReq) (*PlayRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Play not implemented")
}
func (*UnimplementedPlaySvcServer) Stop(context.Context, *StopReq) (*StopRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}

func RegisterPlaySvcServer(s *grpc.Server, srv PlaySvcServer) {
	s.RegisterService(&_PlaySvc_serviceDesc, srv)
}

func _PlaySvc_Play_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		fmt.Println("play-->", in)
		return srv.(PlaySvcServer).Play(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PlaySvc/Play",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaySvcServer).Play(ctx, req.(*PlayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaySvc_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaySvcServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PlaySvc/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaySvcServer).Stop(ctx, req.(*StopReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlaySvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.PlaySvc",
	HandlerType: (*PlaySvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Play",
			Handler:    _PlaySvc_Play_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _PlaySvc_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "play.proto",
}
