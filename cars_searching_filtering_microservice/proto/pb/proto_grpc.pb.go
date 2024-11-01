// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto.proto

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

const (
	CarsManagement_CreateCarDoc_FullMethodName = "/CarsManagement/CreateCarDoc"
)

// CarsManagementClient is the client API for CarsManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarsManagementClient interface {
	CreateCarDoc(ctx context.Context, in *CreateCarDocReq, opts ...grpc.CallOption) (*CreateCarDocRes, error)
}

type carsManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewCarsManagementClient(cc grpc.ClientConnInterface) CarsManagementClient {
	return &carsManagementClient{cc}
}

func (c *carsManagementClient) CreateCarDoc(ctx context.Context, in *CreateCarDocReq, opts ...grpc.CallOption) (*CreateCarDocRes, error) {
	out := new(CreateCarDocRes)
	err := c.cc.Invoke(ctx, CarsManagement_CreateCarDoc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarsManagementServer is the server API for CarsManagement service.
// All implementations must embed UnimplementedCarsManagementServer
// for forward compatibility
type CarsManagementServer interface {
	CreateCarDoc(context.Context, *CreateCarDocReq) (*CreateCarDocRes, error)
	//mustEmbedUnimplementedCarsManagementServer()
}

// UnimplementedCarsManagementServer must be embedded to have forward compatible implementations.
type UnimplementedCarsManagementServer struct {
}

func (UnimplementedCarsManagementServer) CreateCarDoc(context.Context, *CreateCarDocReq) (*CreateCarDocRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCarDoc not implemented")
}
func (UnimplementedCarsManagementServer) mustEmbedUnimplementedCarsManagementServer() {}

// UnsafeCarsManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarsManagementServer will
// result in compilation errors.
type UnsafeCarsManagementServer interface {
	mustEmbedUnimplementedCarsManagementServer()
}

func RegisterCarsManagementServer(s grpc.ServiceRegistrar, srv CarsManagementServer) {
	s.RegisterService(&CarsManagement_ServiceDesc, srv)
}

func _CarsManagement_CreateCarDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCarDocReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarsManagementServer).CreateCarDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CarsManagement_CreateCarDoc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarsManagementServer).CreateCarDoc(ctx, req.(*CreateCarDocReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CarsManagement_ServiceDesc is the grpc.ServiceDesc for CarsManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CarsManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CarsManagement",
	HandlerType: (*CarsManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCarDoc",
			Handler:    _CarsManagement_CreateCarDoc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}
