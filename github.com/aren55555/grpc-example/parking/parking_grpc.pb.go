// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: parking.proto

package parking

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
	ParkingService_ParkCar_FullMethodName    = "/parking.ParkingService/ParkCar"
	ParkingService_CollectCar_FullMethodName = "/parking.ParkingService/CollectCar"
	ParkingService_GetStatus_FullMethodName  = "/parking.ParkingService/GetStatus"
)

// ParkingServiceClient is the client API for ParkingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Models a parking service that allows a car to be parked and later collected
type ParkingServiceClient interface {
	// Use to park the car in the parking lot
	ParkCar(ctx context.Context, in *ParkCarRequest, opts ...grpc.CallOption) (*ParkCarResponse, error)
	// Used to pickup a car from the parking lot
	CollectCar(ctx context.Context, in *CollectCarRequest, opts ...grpc.CallOption) (*CollectCarResponse, error)
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
}

type parkingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParkingServiceClient(cc grpc.ClientConnInterface) ParkingServiceClient {
	return &parkingServiceClient{cc}
}

func (c *parkingServiceClient) ParkCar(ctx context.Context, in *ParkCarRequest, opts ...grpc.CallOption) (*ParkCarResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ParkCarResponse)
	err := c.cc.Invoke(ctx, ParkingService_ParkCar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parkingServiceClient) CollectCar(ctx context.Context, in *CollectCarRequest, opts ...grpc.CallOption) (*CollectCarResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CollectCarResponse)
	err := c.cc.Invoke(ctx, ParkingService_CollectCar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parkingServiceClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, ParkingService_GetStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParkingServiceServer is the server API for ParkingService service.
// All implementations must embed UnimplementedParkingServiceServer
// for forward compatibility.
//
// Models a parking service that allows a car to be parked and later collected
type ParkingServiceServer interface {
	// Use to park the car in the parking lot
	ParkCar(context.Context, *ParkCarRequest) (*ParkCarResponse, error)
	// Used to pickup a car from the parking lot
	CollectCar(context.Context, *CollectCarRequest) (*CollectCarResponse, error)
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	mustEmbedUnimplementedParkingServiceServer()
}

// UnimplementedParkingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedParkingServiceServer struct{}

func (UnimplementedParkingServiceServer) ParkCar(context.Context, *ParkCarRequest) (*ParkCarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParkCar not implemented")
}
func (UnimplementedParkingServiceServer) CollectCar(context.Context, *CollectCarRequest) (*CollectCarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectCar not implemented")
}
func (UnimplementedParkingServiceServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedParkingServiceServer) mustEmbedUnimplementedParkingServiceServer() {}
func (UnimplementedParkingServiceServer) testEmbeddedByValue()                        {}

// UnsafeParkingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParkingServiceServer will
// result in compilation errors.
type UnsafeParkingServiceServer interface {
	mustEmbedUnimplementedParkingServiceServer()
}

func RegisterParkingServiceServer(s grpc.ServiceRegistrar, srv ParkingServiceServer) {
	// If the following call pancis, it indicates UnimplementedParkingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ParkingService_ServiceDesc, srv)
}

func _ParkingService_ParkCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParkCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkingServiceServer).ParkCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParkingService_ParkCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkingServiceServer).ParkCar(ctx, req.(*ParkCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParkingService_CollectCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkingServiceServer).CollectCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParkingService_CollectCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkingServiceServer).CollectCar(ctx, req.(*CollectCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParkingService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkingServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParkingService_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkingServiceServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ParkingService_ServiceDesc is the grpc.ServiceDesc for ParkingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ParkingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "parking.ParkingService",
	HandlerType: (*ParkingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParkCar",
			Handler:    _ParkingService_ParkCar_Handler,
		},
		{
			MethodName: "CollectCar",
			Handler:    _ParkingService_CollectCar_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _ParkingService_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parking.proto",
}
