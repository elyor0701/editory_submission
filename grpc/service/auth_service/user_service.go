package auth_service

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/auth_service"
	"editory_submission/grpc/client"
	"editory_submission/pkg/security"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"fmt"
	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedUserServiceServer
}

func NewUserService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *userService {
	return &userService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *userService) CreateUser(ctx context.Context, req *pb.User) (res *pb.User, err error) {
	s.log.Info("---CreateUser--->", logger.Any("req", req))

	if ok := util.IsValidEmail(req.Email); !ok {
		err = fmt.Errorf("email is not valid")
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, err
	}

	if len(req.Password) < 6 {
		err := fmt.Errorf("password must not be less than 6 characters")
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, err
	}

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	req.Password = hashedPassword

	res, err = s.strg.Auth().User().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// TODO add author role for every user

	return res, nil
}

func (s *userService) GetUser(ctx context.Context, req *pb.GetUserReq) (res *pb.User, err error) {
	s.log.Info("---GetUser--->", logger.Any("req", req))

	res, err = s.strg.Auth().User().Get(ctx, req)
	if err != nil {
		s.log.Error("!!!GetUser--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *userService) GetUserList(ctx context.Context, req *pb.GetUserListReq) (res *pb.GetUserListRes, err error) {
	s.log.Info("---GetUserList--->", logger.Any("req", req))

	res, err = s.strg.Auth().User().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetUserList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *userService) UpdateUser(ctx context.Context, req *pb.User) (res *pb.User, err error) {
	s.log.Info("---UpdateUser--->", logger.Any("req", req))

	rowsAffected, err := s.strg.Auth().User().Update(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return req, nil
}

func (s *userService) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteUser--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Auth().User().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteUser--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
