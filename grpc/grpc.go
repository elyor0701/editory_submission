package grpc

import (
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
	"editory_submission/genproto/content_service"
	"editory_submission/genproto/notification_service"
	"editory_submission/grpc/client"
	auth "editory_submission/grpc/service/auth_service"
	content "editory_submission/grpc/service/content_service"
	"editory_submission/grpc/service/notification"
	"editory_submission/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	// auth
	auth_service.RegisterUserServiceServer(grpcServer, auth.NewUserService(cfg, log, strg, svcs))
	auth_service.RegisterSessionServiceServer(grpcServer, auth.NewSessionService(cfg, log, strg, svcs))
	auth_service.RegisterRoleServiceServer(grpcServer, auth.NewRoleService(cfg, log, strg, svcs))
	auth_service.RegisterKeywordServiceServer(grpcServer, auth.NewKeywordService(cfg, log, strg, svcs))

	// content
	content_service.RegisterContentServiceServer(grpcServer, content.NewContentService(cfg, log, strg, svcs))
	content_service.RegisterUniversityServiceServer(grpcServer, content.NewUniversityService(cfg, log, strg, svcs))
	content_service.RegisterSubjectServiceServer(grpcServer, content.NewSubjectService(cfg, log, strg, svcs))

	// notification
	notification_service.RegisterEmailTmpServiceServer(grpcServer, notification.NewEmailTmpService(cfg, log, strg, svcs))
	notification_service.RegisterNotificationServiceServer(grpcServer, notification.NewNotificationService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)
	return
}
