package main

import (
	"context"
	"editory_submission/api"
	"editory_submission/api/handlers"
	"editory_submission/config"
	"editory_submission/grpc"
	"editory_submission/grpc/client"
	"editory_submission/storage/postgres"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/saidamir98/udevs_pkg/logger"
	"net"

	"github.com/gin-gonic/gin"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func main() {
	cfg := config.Load()

	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer func() {
		if err := logger.Cleanup(log); err != nil {
			log.Panic("logger.Cleanup", logger.Error(err))
		}
	}()

	m, err := migrate.New("file:///home/euler/Documents/projects/editory_submission/migrations/postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDatabase,
		),
	)

	if err != nil {
		log.Panic("migrate.Postgres", logger.Error(err))
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Panic("migrate.Postgres", logger.Error(err))
	}

	pgStore, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("postgres.NewPostgres", logger.Error(err))
	}
	defer pgStore.CloseDB()

	svcs, err := client.NewGrpcClients(cfg)
	if err != nil {
		log.Panic("client.NewGrpcClients", logger.Error(err))
	}

	grpcServer := grpc.SetUpServer(cfg, log, pgStore, svcs)
	go func() {
		lis, err := net.Listen("tcp", cfg.AuthGRPCPort)
		if err != nil {
			log.Panic("net.Listen", logger.Error(err))
		}

		log.Info("GRPC: Server being started...", logger.String("port", cfg.AuthGRPCPort))

		if err := grpcServer.Serve(lis); err != nil {
			log.Panic("grpcServer.Serve", logger.Error(err))
		}
	}()
	h := handlers.NewHandler(cfg, log, svcs)

	r := api.SetUpRouter(h, cfg)

	if err = r.Run(cfg.HTTPPort); err != nil {
		log.Panic("router.Run", logger.Error(err))
	}
}
