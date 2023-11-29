package service_test

import (
	"context"
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
	"editory_submission/grpc"
	"editory_submission/grpc/client"
	"editory_submission/pkg/logger"
	"editory_submission/storage/postgres"
	"fmt"
	"math/rand"
	"net"
	"testing"
	"time"
)

func TestRegisterCompany(t *testing.T) {
	conf := config.Load()
	log := logger.NewLogger(conf.ServiceName, logger.LevelDebug)

	pgStore, err := postgres.NewPostgres(context.Background(), conf)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer pgStore.CloseDB()

	svcs, err := client.NewGrpcClients(conf)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	grpcServer := grpc.SetUpServer(conf, log, pgStore, svcs)
	go func() {
		lis, err := net.Listen("tcp", conf.AuthGRPCPort)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		log.Info("GRPC: Server being started...", logger.String("port", conf.AuthGRPCPort))

		if err := grpcServer.Serve(lis); err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()

	n := rand.Intn(100)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	_, err = svcs.CompanyService().Register(ctx, &auth_service.RegisterCompanyRequest{
		Name: "Ucode",
		UserInfo: &auth_service.RegisterCompanyRequest_RegisterUserInfo{
			Phone:    fmt.Sprintf("+99890123456%d%d", rand.Intn(10), rand.Intn(10)),
			Email:    fmt.Sprintf("john_doe%d@gmail.com", n),
			Login:    fmt.Sprintf("john_login%d", n),
			Password: fmt.Sprintf("john_password%d", n),
		},
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

}
