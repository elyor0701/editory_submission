// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: article_service.proto

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

const (
	ArticleService_CreateArticle_FullMethodName  = "/submission_service.ArticleService/CreateArticle"
	ArticleService_GetArticle_FullMethodName     = "/submission_service.ArticleService/GetArticle"
	ArticleService_GetArticleList_FullMethodName = "/submission_service.ArticleService/GetArticleList"
	ArticleService_UpdateArticle_FullMethodName  = "/submission_service.ArticleService/UpdateArticle"
	ArticleService_DeleteArticle_FullMethodName  = "/submission_service.ArticleService/DeleteArticle"
	ArticleService_AddFiles_FullMethodName       = "/submission_service.ArticleService/AddFiles"
	ArticleService_GetFiles_FullMethodName       = "/submission_service.ArticleService/GetFiles"
	ArticleService_DeleteFiles_FullMethodName    = "/submission_service.ArticleService/DeleteFiles"
)

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleServiceClient interface {
	// Article
	CreateArticle(ctx context.Context, in *CreateArticleReq, opts ...grpc.CallOption) (*CreateArticleRes, error)
	GetArticle(ctx context.Context, in *GetArticleReq, opts ...grpc.CallOption) (*GetArticleRes, error)
	GetArticleList(ctx context.Context, in *GetArticleListReq, opts ...grpc.CallOption) (*GetArticleListRes, error)
	UpdateArticle(ctx context.Context, in *UpdateArticleReq, opts ...grpc.CallOption) (*UpdateArticleRes, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// File
	AddFiles(ctx context.Context, in *AddFilesReq, opts ...grpc.CallOption) (*AddFilesRes, error)
	GetFiles(ctx context.Context, in *GetFilesReq, opts ...grpc.CallOption) (*GetFilesRes, error)
	DeleteFiles(ctx context.Context, in *DeleteFilesReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type articleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleServiceClient(cc grpc.ClientConnInterface) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) CreateArticle(ctx context.Context, in *CreateArticleReq, opts ...grpc.CallOption) (*CreateArticleRes, error) {
	out := new(CreateArticleRes)
	err := c.cc.Invoke(ctx, ArticleService_CreateArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticle(ctx context.Context, in *GetArticleReq, opts ...grpc.CallOption) (*GetArticleRes, error) {
	out := new(GetArticleRes)
	err := c.cc.Invoke(ctx, ArticleService_GetArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticleList(ctx context.Context, in *GetArticleListReq, opts ...grpc.CallOption) (*GetArticleListRes, error) {
	out := new(GetArticleListRes)
	err := c.cc.Invoke(ctx, ArticleService_GetArticleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) UpdateArticle(ctx context.Context, in *UpdateArticleReq, opts ...grpc.CallOption) (*UpdateArticleRes, error) {
	out := new(UpdateArticleRes)
	err := c.cc.Invoke(ctx, ArticleService_UpdateArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) DeleteArticle(ctx context.Context, in *DeleteArticleReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ArticleService_DeleteArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) AddFiles(ctx context.Context, in *AddFilesReq, opts ...grpc.CallOption) (*AddFilesRes, error) {
	out := new(AddFilesRes)
	err := c.cc.Invoke(ctx, ArticleService_AddFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetFiles(ctx context.Context, in *GetFilesReq, opts ...grpc.CallOption) (*GetFilesRes, error) {
	out := new(GetFilesRes)
	err := c.cc.Invoke(ctx, ArticleService_GetFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) DeleteFiles(ctx context.Context, in *DeleteFilesReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ArticleService_DeleteFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
// All implementations must embed UnimplementedArticleServiceServer
// for forward compatibility
type ArticleServiceServer interface {
	// Article
	CreateArticle(context.Context, *CreateArticleReq) (*CreateArticleRes, error)
	GetArticle(context.Context, *GetArticleReq) (*GetArticleRes, error)
	GetArticleList(context.Context, *GetArticleListReq) (*GetArticleListRes, error)
	UpdateArticle(context.Context, *UpdateArticleReq) (*UpdateArticleRes, error)
	DeleteArticle(context.Context, *DeleteArticleReq) (*emptypb.Empty, error)
	// File
	AddFiles(context.Context, *AddFilesReq) (*AddFilesRes, error)
	GetFiles(context.Context, *GetFilesReq) (*GetFilesRes, error)
	DeleteFiles(context.Context, *DeleteFilesReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedArticleServiceServer()
}

// UnimplementedArticleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (UnimplementedArticleServiceServer) CreateArticle(context.Context, *CreateArticleReq) (*CreateArticleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedArticleServiceServer) GetArticle(context.Context, *GetArticleReq) (*GetArticleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedArticleServiceServer) GetArticleList(context.Context, *GetArticleListReq) (*GetArticleListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleList not implemented")
}
func (UnimplementedArticleServiceServer) UpdateArticle(context.Context, *UpdateArticleReq) (*UpdateArticleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedArticleServiceServer) DeleteArticle(context.Context, *DeleteArticleReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedArticleServiceServer) AddFiles(context.Context, *AddFilesReq) (*AddFilesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFiles not implemented")
}
func (UnimplementedArticleServiceServer) GetFiles(context.Context, *GetFilesReq) (*GetFilesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiles not implemented")
}
func (UnimplementedArticleServiceServer) DeleteFiles(context.Context, *DeleteFilesReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFiles not implemented")
}
func (UnimplementedArticleServiceServer) mustEmbedUnimplementedArticleServiceServer() {}

// UnsafeArticleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServiceServer will
// result in compilation errors.
type UnsafeArticleServiceServer interface {
	mustEmbedUnimplementedArticleServiceServer()
}

func RegisterArticleServiceServer(s grpc.ServiceRegistrar, srv ArticleServiceServer) {
	s.RegisterService(&ArticleService_ServiceDesc, srv)
}

func _ArticleService_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_CreateArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).CreateArticle(ctx, req.(*CreateArticleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_GetArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticle(ctx, req.(*GetArticleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_GetArticleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticleList(ctx, req.(*GetArticleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_UpdateArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).UpdateArticle(ctx, req.(*UpdateArticleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_DeleteArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).DeleteArticle(ctx, req.(*DeleteArticleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_AddFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFilesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).AddFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_AddFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).AddFiles(ctx, req.(*AddFilesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_GetFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetFiles(ctx, req.(*GetFilesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_DeleteFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFilesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).DeleteFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_DeleteFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).DeleteFiles(ctx, req.(*DeleteFilesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ArticleService_ServiceDesc is the grpc.ServiceDesc for ArticleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArticleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "submission_service.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArticle",
			Handler:    _ArticleService_CreateArticle_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _ArticleService_GetArticle_Handler,
		},
		{
			MethodName: "GetArticleList",
			Handler:    _ArticleService_GetArticleList_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _ArticleService_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _ArticleService_DeleteArticle_Handler,
		},
		{
			MethodName: "AddFiles",
			Handler:    _ArticleService_AddFiles_Handler,
		},
		{
			MethodName: "GetFiles",
			Handler:    _ArticleService_GetFiles_Handler,
		},
		{
			MethodName: "DeleteFiles",
			Handler:    _ArticleService_DeleteFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article_service.proto",
}
