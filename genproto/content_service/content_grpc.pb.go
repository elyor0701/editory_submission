// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: content.proto

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

// ContentServiceClient is the client API for ContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentServiceClient interface {
	// Journal
	CreateJournal(ctx context.Context, in *CreateJournalReq, opts ...grpc.CallOption) (*Journal, error)
	GetJournal(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Journal, error)
	GetJournalList(ctx context.Context, in *GetList, opts ...grpc.CallOption) (*GetJournalListRes, error)
	UpdateJournal(ctx context.Context, in *Journal, opts ...grpc.CallOption) (*Journal, error)
	DeleteJournal(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Article
	CreateArticle(ctx context.Context, in *CreateArticleReq, opts ...grpc.CallOption) (*Article, error)
	GetArticle(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Article, error)
	GetArticleList(ctx context.Context, in *GetArticleListReq, opts ...grpc.CallOption) (*GetArticleListRes, error)
	UpdateArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error)
	DeleteArticle(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type contentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContentServiceClient(cc grpc.ClientConnInterface) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) CreateJournal(ctx context.Context, in *CreateJournalReq, opts ...grpc.CallOption) (*Journal, error) {
	out := new(Journal)
	err := c.cc.Invoke(ctx, "/content_service.ContentService/CreateJournal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetJournal(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Journal, error) {
	out := new(Journal)
	err := c.cc.Invoke(ctx, "/content_service.ContentService/GetJournal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetJournalList(ctx context.Context, in *GetList, opts ...grpc.CallOption) (*GetJournalListRes, error) {
	out := new(GetJournalListRes)
	err := c.cc.Invoke(ctx, "/content_service.ContentService/GetJournalList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) UpdateJournal(ctx context.Context, in *Journal, opts ...grpc.CallOption) (*Journal, error) {
	out := new(Journal)
	err := c.cc.Invoke(ctx, "/content_service.ContentService/UpdateJournal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DeleteJournal(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/content_service.ContentService/DeleteJournal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) CreateArticle(ctx context.Context, in *CreateArticleReq, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/content_service.ContentService/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetArticle(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/content_service.ContentService/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetArticleList(ctx context.Context, in *GetArticleListReq, opts ...grpc.CallOption) (*GetArticleListRes, error) {
	out := new(GetArticleListRes)
	err := c.cc.Invoke(ctx, "/content_service.ContentService/GetArticleList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) UpdateArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/content_service.ContentService/UpdateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DeleteArticle(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/content_service.ContentService/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServiceServer is the server API for ContentService service.
// All implementations must embed UnimplementedContentServiceServer
// for forward compatibility
type ContentServiceServer interface {
	// Journal
	CreateJournal(context.Context, *CreateJournalReq) (*Journal, error)
	GetJournal(context.Context, *PrimaryKey) (*Journal, error)
	GetJournalList(context.Context, *GetList) (*GetJournalListRes, error)
	UpdateJournal(context.Context, *Journal) (*Journal, error)
	DeleteJournal(context.Context, *PrimaryKey) (*emptypb.Empty, error)
	// Article
	CreateArticle(context.Context, *CreateArticleReq) (*Article, error)
	GetArticle(context.Context, *PrimaryKey) (*Article, error)
	GetArticleList(context.Context, *GetArticleListReq) (*GetArticleListRes, error)
	UpdateArticle(context.Context, *Article) (*Article, error)
	DeleteArticle(context.Context, *PrimaryKey) (*emptypb.Empty, error)
	mustEmbedUnimplementedContentServiceServer()
}

// UnimplementedContentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContentServiceServer struct {
}

func (UnimplementedContentServiceServer) CreateJournal(context.Context, *CreateJournalReq) (*Journal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJournal not implemented")
}
func (UnimplementedContentServiceServer) GetJournal(context.Context, *PrimaryKey) (*Journal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJournal not implemented")
}
func (UnimplementedContentServiceServer) GetJournalList(context.Context, *GetList) (*GetJournalListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJournalList not implemented")
}
func (UnimplementedContentServiceServer) UpdateJournal(context.Context, *Journal) (*Journal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJournal not implemented")
}
func (UnimplementedContentServiceServer) DeleteJournal(context.Context, *PrimaryKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJournal not implemented")
}
func (UnimplementedContentServiceServer) CreateArticle(context.Context, *CreateArticleReq) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedContentServiceServer) GetArticle(context.Context, *PrimaryKey) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedContentServiceServer) GetArticleList(context.Context, *GetArticleListReq) (*GetArticleListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleList not implemented")
}
func (UnimplementedContentServiceServer) UpdateArticle(context.Context, *Article) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedContentServiceServer) DeleteArticle(context.Context, *PrimaryKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedContentServiceServer) mustEmbedUnimplementedContentServiceServer() {}

// UnsafeContentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServiceServer will
// result in compilation errors.
type UnsafeContentServiceServer interface {
	mustEmbedUnimplementedContentServiceServer()
}

func RegisterContentServiceServer(s grpc.ServiceRegistrar, srv ContentServiceServer) {
	s.RegisterService(&ContentService_ServiceDesc, srv)
}

func _ContentService_CreateJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJournalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).CreateJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.ContentService/CreateJournal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).CreateJournal(ctx, req.(*CreateJournalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.ContentService/GetJournal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetJournal(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetJournalList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetJournalList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.ContentService/GetJournalList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetJournalList(ctx, req.(*GetList))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_UpdateJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Journal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).UpdateJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.ContentService/UpdateJournal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).UpdateJournal(ctx, req.(*Journal))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_DeleteJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeleteJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.ContentService/DeleteJournal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeleteJournal(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.ContentService/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).CreateArticle(ctx, req.(*CreateArticleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.ContentService/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetArticle(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetArticleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetArticleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.ContentService/GetArticleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetArticleList(ctx, req.(*GetArticleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Article)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.ContentService/UpdateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).UpdateArticle(ctx, req.(*Article))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.ContentService/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeleteArticle(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// ContentService_ServiceDesc is the grpc.ServiceDesc for ContentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content_service.ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJournal",
			Handler:    _ContentService_CreateJournal_Handler,
		},
		{
			MethodName: "GetJournal",
			Handler:    _ContentService_GetJournal_Handler,
		},
		{
			MethodName: "GetJournalList",
			Handler:    _ContentService_GetJournalList_Handler,
		},
		{
			MethodName: "UpdateJournal",
			Handler:    _ContentService_UpdateJournal_Handler,
		},
		{
			MethodName: "DeleteJournal",
			Handler:    _ContentService_DeleteJournal_Handler,
		},
		{
			MethodName: "CreateArticle",
			Handler:    _ContentService_CreateArticle_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _ContentService_GetArticle_Handler,
		},
		{
			MethodName: "GetArticleList",
			Handler:    _ContentService_GetArticleList_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _ContentService_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _ContentService_DeleteArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content.proto",
}
