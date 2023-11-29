package auth_service

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/auth_service"
	"editory_submission/grpc/client"
	"editory_submission/storage"
	"github.com/saidamir98/udevs_pkg/logger"
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

func (s *userService) CreateUser(context.Context, *pb.User) (*pb.User, error) {
	return nil, nil
}

func (s *userService) GetUser(context.Context, *pb.GetUserReq) (*pb.User, error) {
	return nil, nil
}

func (s *userService) GetUserList(context.Context, *pb.GetUserListReq) (*pb.GetUserListReq, error) {
	return nil, nil
}

func (s *userService) UpdateUser(context.Context, *pb.User) (*pb.User, error) {
	return nil, nil
}

func (s *userService) DeleteUser(context.Context, *pb.DeleteUserReq) (*emptypb.Empty, error) {
	return nil, nil
}
