// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GreetingServiceClient is the client API for GreetingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetingServiceClient interface {
	GetGreeting(ctx context.Context, in *GetGreetingRequest, opts ...grpc.CallOption) (*GetGreetingResponse, error)
}

type greetingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetingServiceClient(cc grpc.ClientConnInterface) GreetingServiceClient {
	return &greetingServiceClient{cc}
}

func (c *greetingServiceClient) GetGreeting(ctx context.Context, in *GetGreetingRequest, opts ...grpc.CallOption) (*GetGreetingResponse, error) {
	out := new(GetGreetingResponse)
	err := c.cc.Invoke(ctx, "/simple.GreetingService/GetGreeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetingServiceServer is the server API for GreetingService service.
// All implementations should embed UnimplementedGreetingServiceServer
// for forward compatibility
type GreetingServiceServer interface {
	GetGreeting(context.Context, *GetGreetingRequest) (*GetGreetingResponse, error)
}

// UnimplementedGreetingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGreetingServiceServer struct {
}

func (UnimplementedGreetingServiceServer) GetGreeting(context.Context, *GetGreetingRequest) (*GetGreetingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGreeting not implemented")
}

// UnsafeGreetingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetingServiceServer will
// result in compilation errors.
type UnsafeGreetingServiceServer interface {
	mustEmbedUnimplementedGreetingServiceServer()
}

func RegisterGreetingServiceServer(s grpc.ServiceRegistrar, srv GreetingServiceServer) {
	s.RegisterService(&_GreetingService_serviceDesc, srv)
}

func _GreetingService_GetGreeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGreetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServiceServer).GetGreeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simple.GreetingService/GetGreeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServiceServer).GetGreeting(ctx, req.(*GetGreetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "simple.GreetingService",
	HandlerType: (*GreetingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGreeting",
			Handler:    _GreetingService_GetGreeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simple.proto",
}