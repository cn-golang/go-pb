// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: data.proto

package data

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
	Data_AddBindingDoctor_FullMethodName = "/data.Data/AddBindingDoctor"
	Data_DelBindingDoctor_FullMethodName = "/data.Data/DelBindingDoctor"
)

// DataClient is the client API for Data service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataClient interface {
	AddBindingDoctor(ctx context.Context, in *AddBindingDoctorReq, opts ...grpc.CallOption) (*AddBindingDoctorRes, error)
	DelBindingDoctor(ctx context.Context, in *AddBindingDoctorReq, opts ...grpc.CallOption) (*AddBindingDoctorRes, error)
}

type dataClient struct {
	cc grpc.ClientConnInterface
}

func NewDataClient(cc grpc.ClientConnInterface) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) AddBindingDoctor(ctx context.Context, in *AddBindingDoctorReq, opts ...grpc.CallOption) (*AddBindingDoctorRes, error) {
	out := new(AddBindingDoctorRes)
	err := c.cc.Invoke(ctx, Data_AddBindingDoctor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) DelBindingDoctor(ctx context.Context, in *AddBindingDoctorReq, opts ...grpc.CallOption) (*AddBindingDoctorRes, error) {
	out := new(AddBindingDoctorRes)
	err := c.cc.Invoke(ctx, Data_DelBindingDoctor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServer is the server API for Data service.
// All implementations must embed UnimplementedDataServer
// for forward compatibility
type DataServer interface {
	AddBindingDoctor(context.Context, *AddBindingDoctorReq) (*AddBindingDoctorRes, error)
	DelBindingDoctor(context.Context, *AddBindingDoctorReq) (*AddBindingDoctorRes, error)
	mustEmbedUnimplementedDataServer()
}

// UnimplementedDataServer must be embedded to have forward compatible implementations.
type UnimplementedDataServer struct {
}

func (UnimplementedDataServer) AddBindingDoctor(context.Context, *AddBindingDoctorReq) (*AddBindingDoctorRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBindingDoctor not implemented")
}
func (UnimplementedDataServer) DelBindingDoctor(context.Context, *AddBindingDoctorReq) (*AddBindingDoctorRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelBindingDoctor not implemented")
}
func (UnimplementedDataServer) mustEmbedUnimplementedDataServer() {}

// UnsafeDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServer will
// result in compilation errors.
type UnsafeDataServer interface {
	mustEmbedUnimplementedDataServer()
}

func RegisterDataServer(s grpc.ServiceRegistrar, srv DataServer) {
	s.RegisterService(&Data_ServiceDesc, srv)
}

func _Data_AddBindingDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBindingDoctorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).AddBindingDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Data_AddBindingDoctor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).AddBindingDoctor(ctx, req.(*AddBindingDoctorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_DelBindingDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBindingDoctorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).DelBindingDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Data_DelBindingDoctor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).DelBindingDoctor(ctx, req.(*AddBindingDoctorReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Data_ServiceDesc is the grpc.ServiceDesc for Data service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Data_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "data.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBindingDoctor",
			Handler:    _Data_AddBindingDoctor_Handler,
		},
		{
			MethodName: "DelBindingDoctor",
			Handler:    _Data_DelBindingDoctor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}
