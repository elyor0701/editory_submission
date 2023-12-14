package grpc

import (
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
	"editory_submission/genproto/content_service"
	"editory_submission/grpc/client"
	auth "editory_submission/grpc/service/auth_service"
	content "editory_submission/grpc/service/content_service"
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

	// content
	content_service.RegisterContentServiceServer(grpcServer, content.NewContentService(cfg, log, strg, svcs))
	content_service.RegisterUniversityServiceServer(grpcServer, content.NewUniversityService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)
	return
}
