// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: submission_service.proto

package submission_service

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

// ReviewerServiceClient is the client API for ReviewerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewerServiceClient interface {
	CreateArticleReviewer(ctx context.Context, in *CreateArticleReviewerReq, opts ...grpc.CallOption) (*CreateArticleReviewerRes, error)
	GetArticleReviewer(ctx context.Context, in *GetArticleReviewerReq, opts ...grpc.CallOption) (*GetArticleReviewerRes, error)
	GetArticleReviewerList(ctx context.Context, in *GetArticleReviewerListReq, opts ...grpc.CallOption) (*GetArticleReviewerListRes, error)
	UpdateArticleReviewer(ctx context.Context, in *UpdateArticleReviewerReq, opts ...grpc.CallOption) (*UpdateArticleReviewerRes, error)
	DeleteArticleReviewer(ctx context.Context, in *DeleteArticleReviewerReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type reviewerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewerServiceClient(cc grpc.ClientConnInterface) ReviewerServiceClient {
	return &reviewerServiceClient{cc}
}

func (c *reviewerServiceClient) CreateArticleReviewer(ctx context.Context, in *CreateArticleReviewerReq, opts ...grpc.CallOption) (*CreateArticleReviewerRes, error) {
	out := new(CreateArticleReviewerRes)
	err := c.cc.Invoke(ctx, "/submission_service.ReviewerService/CreateArticleReviewer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewerServiceClient) GetArticleReviewer(ctx context.Context, in *GetArticleReviewerReq, opts ...grpc.CallOption) (*GetArticleReviewerRes, error) {
	out := new(GetArticleReviewerRes)
	err := c.cc.Invoke(ctx, "/submission_service.ReviewerService/GetArticleReviewer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewerServiceClient) GetArticleReviewerList(ctx context.Context, in *GetArticleReviewerListReq, opts ...grpc.CallOption) (*GetArticleReviewerListRes, error) {
	out := new(GetArticleReviewerListRes)
	err := c.cc.Invoke(ctx, "/submission_service.ReviewerService/GetArticleReviewerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewerServiceClient) UpdateArticleReviewer(ctx context.Context, in *UpdateArticleReviewerReq, opts ...grpc.CallOption) (*UpdateArticleReviewerRes, error) {
	out := new(UpdateArticleReviewerRes)
	err := c.cc.Invoke(ctx, "/submission_service.ReviewerService/UpdateArticleReviewer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewerServiceClient) DeleteArticleReviewer(ctx context.Context, in *DeleteArticleReviewerReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/submission_service.ReviewerService/DeleteArticleReviewer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewerServiceServer is the server API for ReviewerService service.
// All implementations must embed UnimplementedReviewerServiceServer
// for forward compatibility
type ReviewerServiceServer interface {
	CreateArticleReviewer(context.Context, *CreateArticleReviewerReq) (*CreateArticleReviewerRes, error)
	GetArticleReviewer(context.Context, *GetArticleReviewerReq) (*GetArticleReviewerRes, error)
	GetArticleReviewerList(context.Context, *GetArticleReviewerListReq) (*GetArticleReviewerListRes, error)
	UpdateArticleReviewer(context.Context, *UpdateArticleReviewerReq) (*UpdateArticleReviewerRes, error)
	DeleteArticleReviewer(context.Context, *DeleteArticleReviewerReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedReviewerServiceServer()
}

// UnimplementedReviewerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReviewerServiceServer struct {
}

func (UnimplementedReviewerServiceServer) CreateArticleReviewer(context.Context, *CreateArticleReviewerReq) (*CreateArticleReviewerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticleReviewer not implemented")
}
func (UnimplementedReviewerServiceServer) GetArticleReviewer(context.Context, *GetArticleReviewerReq) (*GetArticleReviewerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleReviewer not implemented")
}
func (UnimplementedReviewerServiceServer) GetArticleReviewerList(context.Context, *GetArticleReviewerListReq) (*GetArticleReviewerListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleReviewerList not implemented")
}
func (UnimplementedReviewerServiceServer) UpdateArticleReviewer(context.Context, *UpdateArticleReviewerReq) (*UpdateArticleReviewerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticleReviewer not implemented")
}
func (UnimplementedReviewerServiceServer) DeleteArticleReviewer(context.Context, *DeleteArticleReviewerReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticleReviewer not implemented")
}
func (UnimplementedReviewerServiceServer) mustEmbedUnimplementedReviewerServiceServer() {}

// UnsafeReviewerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewerServiceServer will
// result in compilation errors.
type UnsafeReviewerServiceServer interface {
	mustEmbedUnimplementedReviewerServiceServer()
}

func RegisterReviewerServiceServer(s grpc.ServiceRegistrar, srv ReviewerServiceServer) {
	s.RegisterService(&ReviewerService_ServiceDesc, srv)
}

func _ReviewerService_CreateArticleReviewer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleReviewerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewerServiceServer).CreateArticleReviewer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/submission_service.ReviewerService/CreateArticleReviewer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewerServiceServer).CreateArticleReviewer(ctx, req.(*CreateArticleReviewerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewerService_GetArticleReviewer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleReviewerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewerServiceServer).GetArticleReviewer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/submission_service.ReviewerService/GetArticleReviewer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewerServiceServer).GetArticleReviewer(ctx, req.(*GetArticleReviewerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewerService_GetArticleReviewerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleReviewerListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewerServiceServer).GetArticleReviewerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/submission_service.ReviewerService/GetArticleReviewerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewerServiceServer).GetArticleReviewerList(ctx, req.(*GetArticleReviewerListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewerService_UpdateArticleReviewer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleReviewerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewerServiceServer).UpdateArticleReviewer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/submission_service.ReviewerService/UpdateArticleReviewer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewerServiceServer).UpdateArticleReviewer(ctx, req.(*UpdateArticleReviewerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewerService_DeleteArticleReviewer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleReviewerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewerServiceServer).DeleteArticleReviewer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/submission_service.ReviewerService/DeleteArticleReviewer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewerServiceServer).DeleteArticleReviewer(ctx, req.(*DeleteArticleReviewerReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ReviewerService_ServiceDesc is the grpc.ServiceDesc for ReviewerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReviewerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "submission_service.ReviewerService",
	HandlerType: (*ReviewerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArticleReviewer",
			Handler:    _ReviewerService_CreateArticleReviewer_Handler,
		},
		{
			MethodName: "GetArticleReviewer",
			Handler:    _ReviewerService_GetArticleReviewer_Handler,
		},
		{
			MethodName: "GetArticleReviewerList",
			Handler:    _ReviewerService_GetArticleReviewerList_Handler,
		},
		{
			MethodName: "UpdateArticleReviewer",
			Handler:    _ReviewerService_UpdateArticleReviewer_Handler,
		},
		{
			MethodName: "DeleteArticleReviewer",
			Handler:    _ReviewerService_DeleteArticleReviewer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "submission_service.proto",
}
