// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: hiddifyrpc/hiddify.proto

package hiddifyrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Hello_SayHello_FullMethodName       = "/hiddifyrpc.Hello/SayHello"
	Hello_SayHelloStream_FullMethodName = "/hiddifyrpc.Hello/SayHelloStream"
)

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	SayHelloStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[HelloRequest, HelloResponse], error)
}

type helloClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloClient(cc grpc.ClientConnInterface) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, Hello_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) SayHelloStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[HelloRequest, HelloResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Hello_ServiceDesc.Streams[0], Hello_SayHelloStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HelloRequest, HelloResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Hello_SayHelloStreamClient = grpc.BidiStreamingClient[HelloRequest, HelloResponse]

// HelloServer is the server API for Hello service.
// All implementations must embed UnimplementedHelloServer
// for forward compatibility.
type HelloServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	SayHelloStream(grpc.BidiStreamingServer[HelloRequest, HelloResponse]) error
	mustEmbedUnimplementedHelloServer()
}

// UnimplementedHelloServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHelloServer struct{}

func (UnimplementedHelloServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloServer) SayHelloStream(grpc.BidiStreamingServer[HelloRequest, HelloResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloStream not implemented")
}
func (UnimplementedHelloServer) mustEmbedUnimplementedHelloServer() {}
func (UnimplementedHelloServer) testEmbeddedByValue()               {}

// UnsafeHelloServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServer will
// result in compilation errors.
type UnsafeHelloServer interface {
	mustEmbedUnimplementedHelloServer()
}

func RegisterHelloServer(s grpc.ServiceRegistrar, srv HelloServer) {
	// If the following call pancis, it indicates UnimplementedHelloServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Hello_ServiceDesc, srv)
}

func _Hello_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hello_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_SayHelloStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServer).SayHelloStream(&grpc.GenericServerStream[HelloRequest, HelloResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Hello_SayHelloStreamServer = grpc.BidiStreamingServer[HelloRequest, HelloResponse]

// Hello_ServiceDesc is the grpc.ServiceDesc for Hello service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hello_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hiddifyrpc.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Hello_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloStream",
			Handler:       _Hello_SayHelloStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hiddifyrpc/hiddify.proto",
}

const (
	Core_Start_FullMethodName                 = "/hiddifyrpc.Core/Start"
	Core_CoreInfoListener_FullMethodName      = "/hiddifyrpc.Core/CoreInfoListener"
	Core_OutboundsInfo_FullMethodName         = "/hiddifyrpc.Core/OutboundsInfo"
	Core_MainOutboundsInfo_FullMethodName     = "/hiddifyrpc.Core/MainOutboundsInfo"
	Core_GetSystemInfo_FullMethodName         = "/hiddifyrpc.Core/GetSystemInfo"
	Core_Setup_FullMethodName                 = "/hiddifyrpc.Core/Setup"
	Core_StartService_FullMethodName          = "/hiddifyrpc.Core/StartService"
	Core_Stop_FullMethodName                  = "/hiddifyrpc.Core/Stop"
	Core_Restart_FullMethodName               = "/hiddifyrpc.Core/Restart"
	Core_SelectOutbound_FullMethodName        = "/hiddifyrpc.Core/SelectOutbound"
	Core_UrlTest_FullMethodName               = "/hiddifyrpc.Core/UrlTest"
	Core_GetSystemProxyStatus_FullMethodName  = "/hiddifyrpc.Core/GetSystemProxyStatus"
	Core_SetSystemProxyEnabled_FullMethodName = "/hiddifyrpc.Core/SetSystemProxyEnabled"
	Core_LogListener_FullMethodName           = "/hiddifyrpc.Core/LogListener"
)

// CoreClient is the client API for Core service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*CoreInfoResponse, error)
	CoreInfoListener(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StopRequest, CoreInfoResponse], error)
	OutboundsInfo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StopRequest, OutboundGroupList], error)
	MainOutboundsInfo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StopRequest, OutboundGroupList], error)
	GetSystemInfo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StopRequest, SystemInfo], error)
	Setup(ctx context.Context, in *SetupRequest, opts ...grpc.CallOption) (*Response, error)
	StartService(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*CoreInfoResponse, error)
	Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CoreInfoResponse, error)
	Restart(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*CoreInfoResponse, error)
	SelectOutbound(ctx context.Context, in *SelectOutboundRequest, opts ...grpc.CallOption) (*Response, error)
	UrlTest(ctx context.Context, in *UrlTestRequest, opts ...grpc.CallOption) (*Response, error)
	GetSystemProxyStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SystemProxyStatus, error)
	SetSystemProxyEnabled(ctx context.Context, in *SetSystemProxyEnabledRequest, opts ...grpc.CallOption) (*Response, error)
	LogListener(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StopRequest, LogMessage], error)
}

type coreClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreClient(cc grpc.ClientConnInterface) CoreClient {
	return &coreClient{cc}
}

func (c *coreClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*CoreInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CoreInfoResponse)
	err := c.cc.Invoke(ctx, Core_Start_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) CoreInfoListener(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StopRequest, CoreInfoResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Core_ServiceDesc.Streams[0], Core_CoreInfoListener_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StopRequest, CoreInfoResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Core_CoreInfoListenerClient = grpc.BidiStreamingClient[StopRequest, CoreInfoResponse]

func (c *coreClient) OutboundsInfo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StopRequest, OutboundGroupList], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Core_ServiceDesc.Streams[1], Core_OutboundsInfo_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StopRequest, OutboundGroupList]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Core_OutboundsInfoClient = grpc.BidiStreamingClient[StopRequest, OutboundGroupList]

func (c *coreClient) MainOutboundsInfo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StopRequest, OutboundGroupList], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Core_ServiceDesc.Streams[2], Core_MainOutboundsInfo_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StopRequest, OutboundGroupList]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Core_MainOutboundsInfoClient = grpc.BidiStreamingClient[StopRequest, OutboundGroupList]

func (c *coreClient) GetSystemInfo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StopRequest, SystemInfo], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Core_ServiceDesc.Streams[3], Core_GetSystemInfo_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StopRequest, SystemInfo]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Core_GetSystemInfoClient = grpc.BidiStreamingClient[StopRequest, SystemInfo]

func (c *coreClient) Setup(ctx context.Context, in *SetupRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Core_Setup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) StartService(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*CoreInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CoreInfoResponse)
	err := c.cc.Invoke(ctx, Core_StartService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CoreInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CoreInfoResponse)
	err := c.cc.Invoke(ctx, Core_Stop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) Restart(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*CoreInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CoreInfoResponse)
	err := c.cc.Invoke(ctx, Core_Restart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) SelectOutbound(ctx context.Context, in *SelectOutboundRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Core_SelectOutbound_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) UrlTest(ctx context.Context, in *UrlTestRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Core_UrlTest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetSystemProxyStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SystemProxyStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemProxyStatus)
	err := c.cc.Invoke(ctx, Core_GetSystemProxyStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) SetSystemProxyEnabled(ctx context.Context, in *SetSystemProxyEnabledRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Core_SetSystemProxyEnabled_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) LogListener(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StopRequest, LogMessage], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Core_ServiceDesc.Streams[4], Core_LogListener_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StopRequest, LogMessage]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Core_LogListenerClient = grpc.BidiStreamingClient[StopRequest, LogMessage]

// CoreServer is the server API for Core service.
// All implementations must embed UnimplementedCoreServer
// for forward compatibility.
type CoreServer interface {
	Start(context.Context, *StartRequest) (*CoreInfoResponse, error)
	CoreInfoListener(grpc.BidiStreamingServer[StopRequest, CoreInfoResponse]) error
	OutboundsInfo(grpc.BidiStreamingServer[StopRequest, OutboundGroupList]) error
	MainOutboundsInfo(grpc.BidiStreamingServer[StopRequest, OutboundGroupList]) error
	GetSystemInfo(grpc.BidiStreamingServer[StopRequest, SystemInfo]) error
	Setup(context.Context, *SetupRequest) (*Response, error)
	StartService(context.Context, *StartRequest) (*CoreInfoResponse, error)
	Stop(context.Context, *Empty) (*CoreInfoResponse, error)
	Restart(context.Context, *StartRequest) (*CoreInfoResponse, error)
	SelectOutbound(context.Context, *SelectOutboundRequest) (*Response, error)
	UrlTest(context.Context, *UrlTestRequest) (*Response, error)
	GetSystemProxyStatus(context.Context, *Empty) (*SystemProxyStatus, error)
	SetSystemProxyEnabled(context.Context, *SetSystemProxyEnabledRequest) (*Response, error)
	LogListener(grpc.BidiStreamingServer[StopRequest, LogMessage]) error
	mustEmbedUnimplementedCoreServer()
}

// UnimplementedCoreServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCoreServer struct{}

func (UnimplementedCoreServer) Start(context.Context, *StartRequest) (*CoreInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedCoreServer) CoreInfoListener(grpc.BidiStreamingServer[StopRequest, CoreInfoResponse]) error {
	return status.Errorf(codes.Unimplemented, "method CoreInfoListener not implemented")
}
func (UnimplementedCoreServer) OutboundsInfo(grpc.BidiStreamingServer[StopRequest, OutboundGroupList]) error {
	return status.Errorf(codes.Unimplemented, "method OutboundsInfo not implemented")
}
func (UnimplementedCoreServer) MainOutboundsInfo(grpc.BidiStreamingServer[StopRequest, OutboundGroupList]) error {
	return status.Errorf(codes.Unimplemented, "method MainOutboundsInfo not implemented")
}
func (UnimplementedCoreServer) GetSystemInfo(grpc.BidiStreamingServer[StopRequest, SystemInfo]) error {
	return status.Errorf(codes.Unimplemented, "method GetSystemInfo not implemented")
}
func (UnimplementedCoreServer) Setup(context.Context, *SetupRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Setup not implemented")
}
func (UnimplementedCoreServer) StartService(context.Context, *StartRequest) (*CoreInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartService not implemented")
}
func (UnimplementedCoreServer) Stop(context.Context, *Empty) (*CoreInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedCoreServer) Restart(context.Context, *StartRequest) (*CoreInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedCoreServer) SelectOutbound(context.Context, *SelectOutboundRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectOutbound not implemented")
}
func (UnimplementedCoreServer) UrlTest(context.Context, *UrlTestRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UrlTest not implemented")
}
func (UnimplementedCoreServer) GetSystemProxyStatus(context.Context, *Empty) (*SystemProxyStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemProxyStatus not implemented")
}
func (UnimplementedCoreServer) SetSystemProxyEnabled(context.Context, *SetSystemProxyEnabledRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSystemProxyEnabled not implemented")
}
func (UnimplementedCoreServer) LogListener(grpc.BidiStreamingServer[StopRequest, LogMessage]) error {
	return status.Errorf(codes.Unimplemented, "method LogListener not implemented")
}
func (UnimplementedCoreServer) mustEmbedUnimplementedCoreServer() {}
func (UnimplementedCoreServer) testEmbeddedByValue()              {}

// UnsafeCoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreServer will
// result in compilation errors.
type UnsafeCoreServer interface {
	mustEmbedUnimplementedCoreServer()
}

func RegisterCoreServer(s grpc.ServiceRegistrar, srv CoreServer) {
	// If the following call pancis, it indicates UnimplementedCoreServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Core_ServiceDesc, srv)
}

func _Core_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_CoreInfoListener_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CoreServer).CoreInfoListener(&grpc.GenericServerStream[StopRequest, CoreInfoResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Core_CoreInfoListenerServer = grpc.BidiStreamingServer[StopRequest, CoreInfoResponse]

func _Core_OutboundsInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CoreServer).OutboundsInfo(&grpc.GenericServerStream[StopRequest, OutboundGroupList]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Core_OutboundsInfoServer = grpc.BidiStreamingServer[StopRequest, OutboundGroupList]

func _Core_MainOutboundsInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CoreServer).MainOutboundsInfo(&grpc.GenericServerStream[StopRequest, OutboundGroupList]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Core_MainOutboundsInfoServer = grpc.BidiStreamingServer[StopRequest, OutboundGroupList]

func _Core_GetSystemInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CoreServer).GetSystemInfo(&grpc.GenericServerStream[StopRequest, SystemInfo]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Core_GetSystemInfoServer = grpc.BidiStreamingServer[StopRequest, SystemInfo]

func _Core_Setup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).Setup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_Setup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).Setup(ctx, req.(*SetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_StartService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).StartService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_StartService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).StartService(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).Stop(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).Restart(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_SelectOutbound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectOutboundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).SelectOutbound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_SelectOutbound_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).SelectOutbound(ctx, req.(*SelectOutboundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_UrlTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).UrlTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_UrlTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).UrlTest(ctx, req.(*UrlTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetSystemProxyStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetSystemProxyStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_GetSystemProxyStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetSystemProxyStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_SetSystemProxyEnabled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSystemProxyEnabledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).SetSystemProxyEnabled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Core_SetSystemProxyEnabled_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).SetSystemProxyEnabled(ctx, req.(*SetSystemProxyEnabledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_LogListener_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CoreServer).LogListener(&grpc.GenericServerStream[StopRequest, LogMessage]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Core_LogListenerServer = grpc.BidiStreamingServer[StopRequest, LogMessage]

// Core_ServiceDesc is the grpc.ServiceDesc for Core service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Core_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hiddifyrpc.Core",
	HandlerType: (*CoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Core_Start_Handler,
		},
		{
			MethodName: "Setup",
			Handler:    _Core_Setup_Handler,
		},
		{
			MethodName: "StartService",
			Handler:    _Core_StartService_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Core_Stop_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _Core_Restart_Handler,
		},
		{
			MethodName: "SelectOutbound",
			Handler:    _Core_SelectOutbound_Handler,
		},
		{
			MethodName: "UrlTest",
			Handler:    _Core_UrlTest_Handler,
		},
		{
			MethodName: "GetSystemProxyStatus",
			Handler:    _Core_GetSystemProxyStatus_Handler,
		},
		{
			MethodName: "SetSystemProxyEnabled",
			Handler:    _Core_SetSystemProxyEnabled_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CoreInfoListener",
			Handler:       _Core_CoreInfoListener_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "OutboundsInfo",
			Handler:       _Core_OutboundsInfo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "MainOutboundsInfo",
			Handler:       _Core_MainOutboundsInfo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetSystemInfo",
			Handler:       _Core_GetSystemInfo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "LogListener",
			Handler:       _Core_LogListener_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hiddifyrpc/hiddify.proto",
}

const (
	TunnelService_Start_FullMethodName  = "/hiddifyrpc.TunnelService/Start"
	TunnelService_Stop_FullMethodName   = "/hiddifyrpc.TunnelService/Stop"
	TunnelService_Status_FullMethodName = "/hiddifyrpc.TunnelService/Status"
	TunnelService_Exit_FullMethodName   = "/hiddifyrpc.TunnelService/Exit"
)

// TunnelServiceClient is the client API for TunnelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TunnelServiceClient interface {
	Start(ctx context.Context, in *TunnelStartRequest, opts ...grpc.CallOption) (*TunnelResponse, error)
	Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TunnelResponse, error)
	Status(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TunnelResponse, error)
	Exit(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TunnelResponse, error)
}

type tunnelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTunnelServiceClient(cc grpc.ClientConnInterface) TunnelServiceClient {
	return &tunnelServiceClient{cc}
}

func (c *tunnelServiceClient) Start(ctx context.Context, in *TunnelStartRequest, opts ...grpc.CallOption) (*TunnelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TunnelResponse)
	err := c.cc.Invoke(ctx, TunnelService_Start_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelServiceClient) Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TunnelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TunnelResponse)
	err := c.cc.Invoke(ctx, TunnelService_Stop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelServiceClient) Status(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TunnelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TunnelResponse)
	err := c.cc.Invoke(ctx, TunnelService_Status_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelServiceClient) Exit(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TunnelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TunnelResponse)
	err := c.cc.Invoke(ctx, TunnelService_Exit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TunnelServiceServer is the server API for TunnelService service.
// All implementations must embed UnimplementedTunnelServiceServer
// for forward compatibility.
type TunnelServiceServer interface {
	Start(context.Context, *TunnelStartRequest) (*TunnelResponse, error)
	Stop(context.Context, *Empty) (*TunnelResponse, error)
	Status(context.Context, *Empty) (*TunnelResponse, error)
	Exit(context.Context, *Empty) (*TunnelResponse, error)
	mustEmbedUnimplementedTunnelServiceServer()
}

// UnimplementedTunnelServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTunnelServiceServer struct{}

func (UnimplementedTunnelServiceServer) Start(context.Context, *TunnelStartRequest) (*TunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedTunnelServiceServer) Stop(context.Context, *Empty) (*TunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedTunnelServiceServer) Status(context.Context, *Empty) (*TunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedTunnelServiceServer) Exit(context.Context, *Empty) (*TunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exit not implemented")
}
func (UnimplementedTunnelServiceServer) mustEmbedUnimplementedTunnelServiceServer() {}
func (UnimplementedTunnelServiceServer) testEmbeddedByValue()                       {}

// UnsafeTunnelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TunnelServiceServer will
// result in compilation errors.
type UnsafeTunnelServiceServer interface {
	mustEmbedUnimplementedTunnelServiceServer()
}

func RegisterTunnelServiceServer(s grpc.ServiceRegistrar, srv TunnelServiceServer) {
	// If the following call pancis, it indicates UnimplementedTunnelServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TunnelService_ServiceDesc, srv)
}

func _TunnelService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TunnelStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TunnelService_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServiceServer).Start(ctx, req.(*TunnelStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TunnelService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TunnelService_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServiceServer).Stop(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TunnelService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TunnelService_Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServiceServer).Status(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TunnelService_Exit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServiceServer).Exit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TunnelService_Exit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServiceServer).Exit(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TunnelService_ServiceDesc is the grpc.ServiceDesc for TunnelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TunnelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hiddifyrpc.TunnelService",
	HandlerType: (*TunnelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _TunnelService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _TunnelService_Stop_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _TunnelService_Status_Handler,
		},
		{
			MethodName: "Exit",
			Handler:    _TunnelService_Exit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hiddifyrpc/hiddify.proto",
}
