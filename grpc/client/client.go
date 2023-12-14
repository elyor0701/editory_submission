package client

import (
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
	"editory_submission/genproto/content_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	UserService() auth_service.UserServiceClient
	SessionService() auth_service.SessionServiceClient
	ContentService() content_service.ContentServiceClient
	UniversityService() content_service.UniversityServiceClient
}

type grpcClients struct {
	// auth
	userService    auth_service.UserServiceClient
	sessionService auth_service.SessionServiceClient

	// content
	contentService    content_service.ContentServiceClient
	universityService content_service.UniversityServiceClient
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
		userService:       auth_service.NewUserServiceClient(connAuthService),
		sessionService:    auth_service.NewSessionServiceClient(connAuthService),
		contentService:    content_service.NewContentServiceClient(connAuthService),
		universityService: content_service.NewUniversityServiceClient(connAuthService),
	}, nil
}

func (g *grpcClients) UserService() auth_service.UserServiceClient {
	return g.userService
}

func (g *grpcClients) SessionService() auth_service.SessionServiceClient {
	return g.sessionService
}

func (g *grpcClients) ContentService() content_service.ContentServiceClient {
	return g.contentService
}

func (g *grpcClients) UniversityService() content_service.UniversityServiceClient {
	return g.universityService
}
