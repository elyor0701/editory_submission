package submission_service

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/submission_service"
	"editory_submission/grpc/client"
	"editory_submission/pkg/logger"
	"editory_submission/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type checkerService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedCheckerServiceServer
}

func NewCheckerService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *checkerService {
	return &checkerService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *checkerService) CreateArticleChecker(ctx context.Context, req *pb.CreateArticleCheckerReq) (res *pb.CreateArticleCheckerRes, err error) {
	s.log.Info("---CreateChecker--->", logger.Any("req", req))

	res, err = s.strg.Submission().Reviewer().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateChecker--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *checkerService) GetArticleChecker(ctx context.Context, req *pb.GetArticleCheckerReq) (res *pb.GetArticleCheckerRes, err error) {
	s.log.Info("---GetChecker--->", logger.Any("req", req))

	res, err = s.strg.Submission().Reviewer().Get(ctx, req)
	if err != nil {
		s.log.Error("!!!GetChecker--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *checkerService) GetArticleCheckerList(ctx context.Context, req *pb.GetArticleCheckerListReq) (res *pb.GetArticleCheckerListRes, err error) {
	s.log.Info("---GetCheckerList--->", logger.Any("req", req))

	res, err = s.strg.Submission().Reviewer().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetCheckerList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *checkerService) UpdateArticleChecker(ctx context.Context, req *pb.UpdateArticleCheckerReq) (res *pb.UpdateArticleCheckerRes, err error) {
	s.log.Info("---UpdateChecker--->", logger.Any("req", req))

	rowsAffected, err := s.strg.Submission().Reviewer().Update(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateChecker--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return &pb.UpdateArticleCheckerRes{}, nil
}

func (s *checkerService) DeleteArticleChecker(ctx context.Context, req *pb.DeleteArticleCheckerReq) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteChecker--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Submission().Reviewer().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteChecker--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
