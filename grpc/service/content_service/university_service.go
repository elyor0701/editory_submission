package content_service

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/content_service"
	"editory_submission/grpc/client"
	"editory_submission/pkg/logger"
	"editory_submission/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type universityService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedUniversityServiceServer
}

func NewUniversityService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *universityService {
	return &universityService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *universityService) CreateUniversity(ctx context.Context, req *pb.CreateUniversityReq) (res *pb.CreateUniversityRes, err error) {
	s.log.Info("---CreateUniversity--->", logger.Any("req", req))

	res, err = s.strg.Content().University().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateUniversity--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *universityService) GetUniversity(ctx context.Context, req *pb.GetUniversityReq) (res *pb.GetUniversityRes, err error) {
	s.log.Info("---GetUniversity--->", logger.Any("req", req))

	res, err = s.strg.Content().University().Get(ctx, req)
	if err != nil {
		s.log.Error("!!!GetUniversity--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *universityService) GetUniversityList(ctx context.Context, req *pb.GetUniversityListReq) (res *pb.GetUniversityListRes, err error) {
	s.log.Info("---GetUniversityList--->", logger.Any("req", req))

	res, err = s.strg.Content().University().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetUniversityList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *universityService) UpdateUniversity(ctx context.Context, req *pb.UpdateUniversityReq) (res *pb.UpdateUniversityRes, err error) {
	s.log.Info("---UpdateUniversity--->", logger.Any("req", req))

	res, err = s.strg.Content().University().Update(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateUniversity--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *universityService) DeleteUniversity(ctx context.Context, req *pb.DeleteUniversityReq) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteUniversity--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Content().University().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteUniversity--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
