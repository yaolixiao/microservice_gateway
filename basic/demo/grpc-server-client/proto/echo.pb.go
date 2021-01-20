// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: echo.proto

package proto

import (
	context "context"
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

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{0}
}

func (x *EchoRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{1}
}

func (x *EchoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_echo_proto protoreflect.FileDescriptor

var file_echo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x63,
	0x68, 0x6f, 0x22, 0x27, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x45,
	0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x8b, 0x02, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x34,
	0x0a, 0x09, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x11, 0x2e, 0x65, 0x63,
	0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x45, 0x68, 0x63, 0x6f, 0x12, 0x11, 0x2e, 0x65, 0x63,
	0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x45, 0x68, 0x63, 0x6f, 0x12, 0x11, 0x2e,
	0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x49, 0x0a, 0x1a, 0x42, 0x69, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x65, 0x63, 0x68, 0x6f, 0x12, 0x11, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63,
	0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x65, 0x63, 0x68, 0x6f,
	0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_echo_proto_rawDescOnce sync.Once
	file_echo_proto_rawDescData = file_echo_proto_rawDesc
)

func file_echo_proto_rawDescGZIP() []byte {
	file_echo_proto_rawDescOnce.Do(func() {
		file_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_echo_proto_rawDescData)
	})
	return file_echo_proto_rawDescData
}

var file_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_echo_proto_goTypes = []interface{}{
	(*EchoRequest)(nil),  // 0: echo.EchoRequest
	(*EchoResponse)(nil), // 1: echo.EchoResponse
}
var file_echo_proto_depIdxs = []int32{
	0, // 0: echo.Echo.UnaryEcho:input_type -> echo.EchoRequest
	0, // 1: echo.Echo.ServerStreamingEhco:input_type -> echo.EchoRequest
	0, // 2: echo.Echo.ClientStreamingEhco:input_type -> echo.EchoRequest
	0, // 3: echo.Echo.BidirectionalStreamingecho:input_type -> echo.EchoRequest
	1, // 4: echo.Echo.UnaryEcho:output_type -> echo.EchoResponse
	1, // 5: echo.Echo.ServerStreamingEhco:output_type -> echo.EchoResponse
	1, // 6: echo.Echo.ClientStreamingEhco:output_type -> echo.EchoResponse
	1, // 7: echo.Echo.BidirectionalStreamingecho:output_type -> echo.EchoResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_echo_proto_init() }
func file_echo_proto_init() {
	if File_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
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
		file_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResponse); i {
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
			RawDescriptor: file_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_echo_proto_goTypes,
		DependencyIndexes: file_echo_proto_depIdxs,
		MessageInfos:      file_echo_proto_msgTypes,
	}.Build()
	File_echo_proto = out.File
	file_echo_proto_rawDesc = nil
	file_echo_proto_goTypes = nil
	file_echo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	UnaryEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	ServerStreamingEhco(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (Echo_ServerStreamingEhcoClient, error)
	ClientStreamingEhco(ctx context.Context, opts ...grpc.CallOption) (Echo_ClientStreamingEhcoClient, error)
	BidirectionalStreamingecho(ctx context.Context, opts ...grpc.CallOption) (Echo_BidirectionalStreamingechoClient, error)
}

type echoClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoClient(cc grpc.ClientConnInterface) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) UnaryEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/echo.Echo/UnaryEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) ServerStreamingEhco(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (Echo_ServerStreamingEhcoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Echo_serviceDesc.Streams[0], "/echo.Echo/ServerStreamingEhco", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServerStreamingEhcoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Echo_ServerStreamingEhcoClient interface {
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoServerStreamingEhcoClient struct {
	grpc.ClientStream
}

func (x *echoServerStreamingEhcoClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoClient) ClientStreamingEhco(ctx context.Context, opts ...grpc.CallOption) (Echo_ClientStreamingEhcoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Echo_serviceDesc.Streams[1], "/echo.Echo/ClientStreamingEhco", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoClientStreamingEhcoClient{stream}
	return x, nil
}

type Echo_ClientStreamingEhcoClient interface {
	Send(*EchoRequest) error
	CloseAndRecv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoClientStreamingEhcoClient struct {
	grpc.ClientStream
}

func (x *echoClientStreamingEhcoClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoClientStreamingEhcoClient) CloseAndRecv() (*EchoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoClient) BidirectionalStreamingecho(ctx context.Context, opts ...grpc.CallOption) (Echo_BidirectionalStreamingechoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Echo_serviceDesc.Streams[2], "/echo.Echo/BidirectionalStreamingecho", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoBidirectionalStreamingechoClient{stream}
	return x, nil
}

type Echo_BidirectionalStreamingechoClient interface {
	Send(*EchoRequest) error
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoBidirectionalStreamingechoClient struct {
	grpc.ClientStream
}

func (x *echoBidirectionalStreamingechoClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoBidirectionalStreamingechoClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	UnaryEcho(context.Context, *EchoRequest) (*EchoResponse, error)
	ServerStreamingEhco(*EchoRequest, Echo_ServerStreamingEhcoServer) error
	ClientStreamingEhco(Echo_ClientStreamingEhcoServer) error
	BidirectionalStreamingecho(Echo_BidirectionalStreamingechoServer) error
}

// UnimplementedEchoServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct {
}

func (*UnimplementedEchoServer) UnaryEcho(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryEcho not implemented")
}
func (*UnimplementedEchoServer) ServerStreamingEhco(*EchoRequest, Echo_ServerStreamingEhcoServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamingEhco not implemented")
}
func (*UnimplementedEchoServer) ClientStreamingEhco(Echo_ClientStreamingEhcoServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamingEhco not implemented")
}
func (*UnimplementedEchoServer) BidirectionalStreamingecho(Echo_BidirectionalStreamingechoServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreamingecho not implemented")
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_UnaryEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).UnaryEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/UnaryEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).UnaryEcho(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_ServerStreamingEhco_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EchoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EchoServer).ServerStreamingEhco(m, &echoServerStreamingEhcoServer{stream})
}

type Echo_ServerStreamingEhcoServer interface {
	Send(*EchoResponse) error
	grpc.ServerStream
}

type echoServerStreamingEhcoServer struct {
	grpc.ServerStream
}

func (x *echoServerStreamingEhcoServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Echo_ClientStreamingEhco_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServer).ClientStreamingEhco(&echoClientStreamingEhcoServer{stream})
}

type Echo_ClientStreamingEhcoServer interface {
	SendAndClose(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoClientStreamingEhcoServer struct {
	grpc.ServerStream
}

func (x *echoClientStreamingEhcoServer) SendAndClose(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoClientStreamingEhcoServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Echo_BidirectionalStreamingecho_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServer).BidirectionalStreamingecho(&echoBidirectionalStreamingechoServer{stream})
}

type Echo_BidirectionalStreamingechoServer interface {
	Send(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoBidirectionalStreamingechoServer struct {
	grpc.ServerStream
}

func (x *echoBidirectionalStreamingechoServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoBidirectionalStreamingechoServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryEcho",
			Handler:    _Echo_UnaryEcho_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamingEhco",
			Handler:       _Echo_ServerStreamingEhco_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamingEhco",
			Handler:       _Echo_ClientStreamingEhco_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStreamingecho",
			Handler:       _Echo_BidirectionalStreamingecho_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "echo.proto",
}
