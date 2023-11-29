package client

import (
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	AuthService() auth_service.UserServiceClient
}

type grpcClients struct {
	authService auth_service.UserServiceClient
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
		authService: auth_service.NewUserServiceClient(connAuthService),
	}, nil
}

func (g *grpcClients) AuthService() auth_service.UserServiceClient {
	return g.authService
}
