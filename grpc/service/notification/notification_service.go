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

type notificationService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedNotificationServiceServer
}

func NewNotificationService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *notificationService {
	return &notificationService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *notificationService) CreateNotification(ctx context.Context, req *pb.CreateNotificationReq) (res *pb.CreateNotificationRes, err error) {
	s.log.Info("---CreateNotification--->", logger.Any("req", req))

	res, err = s.strg.Notification().Notification().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateNotification--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *notificationService) GetNotification(ctx context.Context, req *pb.GetNotificationReq) (res *pb.GetNotificationRes, err error) {
	s.log.Info("---GetNotification--->", logger.Any("req", req))

	res, err = s.strg.Notification().Notification().Get(ctx, req)
	if err != nil {
		s.log.Error("!!!GetNotification--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *notificationService) GetNotificationList(ctx context.Context, req *pb.GetNotificationListReq) (res *pb.GetNotificationListRes, err error) {
	s.log.Info("---GetNotificationList--->", logger.Any("req", req))

	res, err = s.strg.Notification().Notification().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetNotificationList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *notificationService) UpdateNotification(ctx context.Context, req *pb.UpdateNotificationReq) (res *pb.UpdateNotificationRes, err error) {
	s.log.Info("---UpdateNotification--->", logger.Any("req", req))

	res, err = s.strg.Notification().Notification().Update(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateNotification--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *notificationService) DeleteNotification(ctx context.Context, req *pb.DeleteNotificationReq) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteNotification--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Notification().Notification().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteNotification--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
