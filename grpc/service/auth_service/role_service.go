package auth_service

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/auth_service"
	"editory_submission/grpc/client"
	"editory_submission/pkg/logger"
	"editory_submission/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type roleService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedRoleServiceServer
}

func NewRoleService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *roleService {
	return &roleService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *roleService) CreateRole(ctx context.Context, req *pb.Role) (res *pb.Role, err error) {
	s.log.Info("---CreateRole--->", logger.Any("req", req))

	res, err = s.strg.Auth().Role().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateRole--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *roleService) GetRole(ctx context.Context, req *pb.GetRoleReq) (res *pb.Role, err error) {
	s.log.Info("---GetRole--->", logger.Any("req", req))

	res, err = s.strg.Auth().Role().Get(ctx, req)
	if err != nil {
		s.log.Error("!!!GetRole--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *roleService) GetRoleList(ctx context.Context, req *pb.GetRoleListReq) (res *pb.GetRoleListRes, err error) {
	s.log.Info("---GetRoleList--->", logger.Any("req", req))

	res, err = s.strg.Auth().Role().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetRoleList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *roleService) UpdateRole(ctx context.Context, req *pb.Role) (res *pb.Role, err error) {
	s.log.Info("---UpdateRole--->", logger.Any("req", req))

	rowsAffected, err := s.strg.Auth().Role().Update(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateRole--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}

func (s *roleService) DeleteRole(ctx context.Context, req *pb.DeleteRoleReq) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteRole--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Auth().Role().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteRole--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
