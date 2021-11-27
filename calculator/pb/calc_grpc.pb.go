// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalcServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	DecomposePrimeNumber(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalcService_DecomposePrimeNumberClient, error)
	Average(ctx context.Context, opts ...grpc.CallOption) (CalcService_AverageClient, error)
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalcService_FindMaximumClient, error)
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type calcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcServiceClient(cc grpc.ClientConnInterface) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calc.CalcService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) DecomposePrimeNumber(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalcService_DecomposePrimeNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalcService_ServiceDesc.Streams[0], "/calc.CalcService/DecomposePrimeNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceDecomposePrimeNumberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalcService_DecomposePrimeNumberClient interface {
	Recv() (*PrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type calcServiceDecomposePrimeNumberClient struct {
	grpc.ClientStream
}

func (x *calcServiceDecomposePrimeNumberClient) Recv() (*PrimeNumberDecompositionResponse, error) {
	m := new(PrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (CalcService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalcService_ServiceDesc.Streams[1], "/calc.CalcService/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceAverageClient{stream}
	return x, nil
}

type CalcService_AverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calcServiceAverageClient struct {
	grpc.ClientStream
}

func (x *calcServiceAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcServiceAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalcService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalcService_ServiceDesc.Streams[2], "/calc.CalcService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceFindMaximumClient{stream}
	return x, nil
}

type CalcService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type calcServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calcServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/calc.CalcService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServiceServer is the server API for CalcService service.
// All implementations must embed UnimplementedCalcServiceServer
// for forward compatibility
type CalcServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	DecomposePrimeNumber(*PrimeNumberDecompositionRequest, CalcService_DecomposePrimeNumberServer) error
	Average(CalcService_AverageServer) error
	FindMaximum(CalcService_FindMaximumServer) error
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
	mustEmbedUnimplementedCalcServiceServer()
}

// UnimplementedCalcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalcServiceServer struct {
}

func (UnimplementedCalcServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalcServiceServer) DecomposePrimeNumber(*PrimeNumberDecompositionRequest, CalcService_DecomposePrimeNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method DecomposePrimeNumber not implemented")
}
func (UnimplementedCalcServiceServer) Average(CalcService_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (UnimplementedCalcServiceServer) FindMaximum(CalcService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}
func (UnimplementedCalcServiceServer) SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}
func (UnimplementedCalcServiceServer) mustEmbedUnimplementedCalcServiceServer() {}

// UnsafeCalcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalcServiceServer will
// result in compilation errors.
type UnsafeCalcServiceServer interface {
	mustEmbedUnimplementedCalcServiceServer()
}

func RegisterCalcServiceServer(s grpc.ServiceRegistrar, srv CalcServiceServer) {
	s.RegisterService(&CalcService_ServiceDesc, srv)
}

func _CalcService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalcService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_DecomposePrimeNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalcServiceServer).DecomposePrimeNumber(m, &calcServiceDecomposePrimeNumberServer{stream})
}

type CalcService_DecomposePrimeNumberServer interface {
	Send(*PrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type calcServiceDecomposePrimeNumberServer struct {
	grpc.ServerStream
}

func (x *calcServiceDecomposePrimeNumberServer) Send(m *PrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalcService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServiceServer).Average(&calcServiceAverageServer{stream})
}

type CalcService_AverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calcServiceAverageServer struct {
	grpc.ServerStream
}

func (x *calcServiceAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcServiceAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalcService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServiceServer).FindMaximum(&calcServiceFindMaximumServer{stream})
}

type CalcService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type calcServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calcServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalcService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalcService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalcService_ServiceDesc is the grpc.ServiceDesc for CalcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calc.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalcService_Sum_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _CalcService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DecomposePrimeNumber",
			Handler:       _CalcService_DecomposePrimeNumber_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _CalcService_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _CalcService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/pb/calc.proto",
}
