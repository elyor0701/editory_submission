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
	RoleService() auth_service.RoleServiceClient
	ContentService() content_service.ContentServiceClient
	UniversityService() content_service.UniversityServiceClient
	SubjectService() content_service.SubjectServiceClient
}

type grpcClients struct {
	// auth
	userService    auth_service.UserServiceClient
	sessionService auth_service.SessionServiceClient
	roleService    auth_service.RoleServiceClient

	// content
	contentService    content_service.ContentServiceClient
	universityService content_service.UniversityServiceClient
	subjectService    content_service.SubjectServiceClient
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
		roleService:       auth_service.NewRoleServiceClient(connAuthService),
		contentService:    content_service.NewContentServiceClient(connAuthService),
		universityService: content_service.NewUniversityServiceClient(connAuthService),
		subjectService:    content_service.NewSubjectServiceClient(connAuthService),
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

func (g *grpcClients) SubjectService() content_service.SubjectServiceClient {
	return g.subjectService
}

func (g *grpcClients) RoleService() auth_service.RoleServiceClient {
	return g.roleService
}
