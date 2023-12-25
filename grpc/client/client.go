package client

import (
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
	"editory_submission/genproto/content_service"
	"editory_submission/genproto/notification_service"
	"editory_submission/genproto/submission_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	UserService() auth_service.UserServiceClient
	SessionService() auth_service.SessionServiceClient
	RoleService() auth_service.RoleServiceClient
	KeywordService() auth_service.KeywordServiceClient
	ContentService() content_service.ContentServiceClient
	UniversityService() content_service.UniversityServiceClient
	SubjectService() content_service.SubjectServiceClient
	NotificationService() notification_service.NotificationServiceClient
	EmailTmpService() notification_service.EmailTmpServiceClient
	ArticleService() submission_service.ArticleServiceClient
	ReviewerService() submission_service.ReviewerServiceClient
}

type grpcClients struct {
	// auth
	userService    auth_service.UserServiceClient
	sessionService auth_service.SessionServiceClient
	roleService    auth_service.RoleServiceClient
	keywordService auth_service.KeywordServiceClient

	// content
	contentService    content_service.ContentServiceClient
	universityService content_service.UniversityServiceClient
	subjectService    content_service.SubjectServiceClient

	// notification
	notificationService notification_service.NotificationServiceClient
	emailTmpService     notification_service.EmailTmpServiceClient

	// submission
	articleService  submission_service.ArticleServiceClient
	reviewerService submission_service.ReviewerServiceClient
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
		userService:         auth_service.NewUserServiceClient(connAuthService),
		sessionService:      auth_service.NewSessionServiceClient(connAuthService),
		roleService:         auth_service.NewRoleServiceClient(connAuthService),
		keywordService:      auth_service.NewKeywordServiceClient(connAuthService),
		contentService:      content_service.NewContentServiceClient(connAuthService),
		universityService:   content_service.NewUniversityServiceClient(connAuthService),
		subjectService:      content_service.NewSubjectServiceClient(connAuthService),
		emailTmpService:     notification_service.NewEmailTmpServiceClient(connAuthService),
		notificationService: notification_service.NewNotificationServiceClient(connAuthService),
		articleService:      submission_service.NewArticleServiceClient(connAuthService),
		reviewerService:     submission_service.NewReviewerServiceClient(connAuthService),
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

func (g *grpcClients) KeywordService() auth_service.KeywordServiceClient {
	return g.keywordService
}

func (g *grpcClients) NotificationService() notification_service.NotificationServiceClient {
	return g.notificationService
}

func (g *grpcClients) EmailTmpService() notification_service.EmailTmpServiceClient {
	return g.emailTmpService
}

func (g *grpcClients) ArticleService() submission_service.ArticleServiceClient {
	return g.articleService
}
func (g *grpcClients) ReviewerService() submission_service.ReviewerServiceClient {
	return g.reviewerService
}
