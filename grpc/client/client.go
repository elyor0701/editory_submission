package client

import (
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	UserService() auth_service.UserServiceClient
}

type grpcClients struct {
	userService auth_service.UserServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connAuthService, err := grpc.Dial(
		cfg.AuthServiceHost+cfg.AuthGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		userService: auth_service.NewUserServiceClient(connAuthService),
	}, nil
}

func (g *grpcClients) UserService() auth_service.UserServiceClient {
	return g.userService
}
