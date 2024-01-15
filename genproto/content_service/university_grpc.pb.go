// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: university.proto

package content_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UniversityService_CreateUniversity_FullMethodName  = "/content_service.UniversityService/CreateUniversity"
	UniversityService_GetUniversity_FullMethodName     = "/content_service.UniversityService/GetUniversity"
	UniversityService_GetUniversityList_FullMethodName = "/content_service.UniversityService/GetUniversityList"
	UniversityService_UpdateUniversity_FullMethodName  = "/content_service.UniversityService/UpdateUniversity"
	UniversityService_DeleteUniversity_FullMethodName  = "/content_service.UniversityService/DeleteUniversity"
)

// UniversityServiceClient is the client API for UniversityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UniversityServiceClient interface {
	CreateUniversity(ctx context.Context, in *CreateUniversityReq, opts ...grpc.CallOption) (*CreateUniversityRes, error)
	GetUniversity(ctx context.Context, in *GetUniversityReq, opts ...grpc.CallOption) (*GetUniversityRes, error)
	GetUniversityList(ctx context.Context, in *GetUniversityListReq, opts ...grpc.CallOption) (*GetUniversityListRes, error)
	UpdateUniversity(ctx context.Context, in *UpdateUniversityReq, opts ...grpc.CallOption) (*UpdateUniversityRes, error)
	DeleteUniversity(ctx context.Context, in *DeleteUniversityReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type universityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUniversityServiceClient(cc grpc.ClientConnInterface) UniversityServiceClient {
	return &universityServiceClient{cc}
}

func (c *universityServiceClient) CreateUniversity(ctx context.Context, in *CreateUniversityReq, opts ...grpc.CallOption) (*CreateUniversityRes, error) {
	out := new(CreateUniversityRes)
	err := c.cc.Invoke(ctx, UniversityService_CreateUniversity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universityServiceClient) GetUniversity(ctx context.Context, in *GetUniversityReq, opts ...grpc.CallOption) (*GetUniversityRes, error) {
	out := new(GetUniversityRes)
	err := c.cc.Invoke(ctx, UniversityService_GetUniversity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universityServiceClient) GetUniversityList(ctx context.Context, in *GetUniversityListReq, opts ...grpc.CallOption) (*GetUniversityListRes, error) {
	out := new(GetUniversityListRes)
	err := c.cc.Invoke(ctx, UniversityService_GetUniversityList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universityServiceClient) UpdateUniversity(ctx context.Context, in *UpdateUniversityReq, opts ...grpc.CallOption) (*UpdateUniversityRes, error) {
	out := new(UpdateUniversityRes)
	err := c.cc.Invoke(ctx, UniversityService_UpdateUniversity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universityServiceClient) DeleteUniversity(ctx context.Context, in *DeleteUniversityReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UniversityService_DeleteUniversity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UniversityServiceServer is the server API for UniversityService service.
// All implementations must embed UnimplementedUniversityServiceServer
// for forward compatibility
type UniversityServiceServer interface {
	CreateUniversity(context.Context, *CreateUniversityReq) (*CreateUniversityRes, error)
	GetUniversity(context.Context, *GetUniversityReq) (*GetUniversityRes, error)
	GetUniversityList(context.Context, *GetUniversityListReq) (*GetUniversityListRes, error)
	UpdateUniversity(context.Context, *UpdateUniversityReq) (*UpdateUniversityRes, error)
	DeleteUniversity(context.Context, *DeleteUniversityReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedUniversityServiceServer()
}

// UnimplementedUniversityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUniversityServiceServer struct {
}

func (UnimplementedUniversityServiceServer) CreateUniversity(context.Context, *CreateUniversityReq) (*CreateUniversityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUniversity not implemented")
}
func (UnimplementedUniversityServiceServer) GetUniversity(context.Context, *GetUniversityReq) (*GetUniversityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUniversity not implemented")
}
func (UnimplementedUniversityServiceServer) GetUniversityList(context.Context, *GetUniversityListReq) (*GetUniversityListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUniversityList not implemented")
}
func (UnimplementedUniversityServiceServer) UpdateUniversity(context.Context, *UpdateUniversityReq) (*UpdateUniversityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUniversity not implemented")
}
func (UnimplementedUniversityServiceServer) DeleteUniversity(context.Context, *DeleteUniversityReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUniversity not implemented")
}
func (UnimplementedUniversityServiceServer) mustEmbedUnimplementedUniversityServiceServer() {}

// UnsafeUniversityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UniversityServiceServer will
// result in compilation errors.
type UnsafeUniversityServiceServer interface {
	mustEmbedUnimplementedUniversityServiceServer()
}

func RegisterUniversityServiceServer(s grpc.ServiceRegistrar, srv UniversityServiceServer) {
	s.RegisterService(&UniversityService_ServiceDesc, srv)
}

func _UniversityService_CreateUniversity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUniversityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversityServiceServer).CreateUniversity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UniversityService_CreateUniversity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversityServiceServer).CreateUniversity(ctx, req.(*CreateUniversityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UniversityService_GetUniversity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUniversityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversityServiceServer).GetUniversity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UniversityService_GetUniversity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversityServiceServer).GetUniversity(ctx, req.(*GetUniversityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UniversityService_GetUniversityList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUniversityListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversityServiceServer).GetUniversityList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UniversityService_GetUniversityList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversityServiceServer).GetUniversityList(ctx, req.(*GetUniversityListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UniversityService_UpdateUniversity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUniversityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversityServiceServer).UpdateUniversity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UniversityService_UpdateUniversity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversityServiceServer).UpdateUniversity(ctx, req.(*UpdateUniversityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UniversityService_DeleteUniversity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUniversityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversityServiceServer).DeleteUniversity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UniversityService_DeleteUniversity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversityServiceServer).DeleteUniversity(ctx, req.(*DeleteUniversityReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UniversityService_ServiceDesc is the grpc.ServiceDesc for UniversityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UniversityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content_service.UniversityService",
	HandlerType: (*UniversityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUniversity",
			Handler:    _UniversityService_CreateUniversity_Handler,
		},
		{
			MethodName: "GetUniversity",
			Handler:    _UniversityService_GetUniversity_Handler,
		},
		{
			MethodName: "GetUniversityList",
			Handler:    _UniversityService_GetUniversityList_Handler,
		},
		{
			MethodName: "UpdateUniversity",
			Handler:    _UniversityService_UpdateUniversity_Handler,
		},
		{
			MethodName: "DeleteUniversity",
			Handler:    _UniversityService_DeleteUniversity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "university.proto",
}
