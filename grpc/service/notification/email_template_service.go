package notification

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/notification_service"
	"editory_submission/grpc/client"
	"editory_submission/pkg/logger"
	"editory_submission/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type emailTmpService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedEmailTmpServiceServer
}

func NewEmailTmpService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *emailTmpService {
	return &emailTmpService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *emailTmpService) CreateEmailTmp(ctx context.Context, req *pb.CreateEmailTmpReq) (res *pb.CreateEmailTmpRes, err error) {
	s.log.Info("---CreateEmailTmp--->", logger.Any("req", req))

	res, err = s.strg.Notification().EmailTemplate().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateEmailTmp--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *emailTmpService) GetEmailTmp(ctx context.Context, req *pb.GetEmailTmpReq) (res *pb.GetEmailTmpRes, err error) {
	s.log.Info("---GetEmailTmp--->", logger.Any("req", req))

	res, err = s.strg.Notification().EmailTemplate().Get(ctx, req)
	if err != nil {
		s.log.Error("!!!GetEmailTmp--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *emailTmpService) GetEmailTmpList(ctx context.Context, req *pb.GetEmailTmpListReq) (res *pb.GetEmailTmpListRes, err error) {
	s.log.Info("---GetEmailTmpList--->", logger.Any("req", req))

	res, err = s.strg.Notification().EmailTemplate().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetEmailTmpList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *emailTmpService) UpdateEmailTmp(ctx context.Context, req *pb.UpdateEmailTmpReq) (res *pb.UpdateEmailTmpRes, err error) {
	s.log.Info("---UpdateEmailTmp--->", logger.Any("req", req))

	res, err = s.strg.Notification().EmailTemplate().Update(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateEmailTmp--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *emailTmpService) DeleteEmailTmp(ctx context.Context, req *pb.DeleteEmailTmpReq) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteEmailTmp--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Notification().EmailTemplate().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteEmailTmp--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
