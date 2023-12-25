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

type reviewerService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedReviewerServiceServer
}

func NewReviewerService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *reviewerService {
	return &reviewerService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *reviewerService) CreateArticleReviewer(ctx context.Context, req *pb.CreateArticleReviewerReq) (res *pb.CreateArticleReviewerRes, err error) {
	s.log.Info("---CreateReviewer--->", logger.Any("req", req))

	res, err = s.strg.Submission().Reviewer().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateReviewer--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *reviewerService) GetArticleReviewer(ctx context.Context, req *pb.GetArticleReviewerReq) (res *pb.GetArticleReviewerRes, err error) {
	s.log.Info("---GetReviewer--->", logger.Any("req", req))

	res, err = s.strg.Submission().Reviewer().Get(ctx, req)
	if err != nil {
		s.log.Error("!!!GetReviewer--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *reviewerService) GetArticleReviewerList(ctx context.Context, req *pb.GetArticleReviewerListReq) (res *pb.GetArticleReviewerListRes, err error) {
	s.log.Info("---GetReviewerList--->", logger.Any("req", req))

	res, err = s.strg.Submission().Reviewer().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetReviewerList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *reviewerService) UpdateArticleReviewer(ctx context.Context, req *pb.UpdateArticleReviewerReq) (res *pb.UpdateArticleReviewerRes, err error) {
	s.log.Info("---UpdateReviewer--->", logger.Any("req", req))

	rowsAffected, err := s.strg.Submission().Reviewer().Update(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateReviewer--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return &pb.UpdateArticleReviewerRes{}, nil
}

func (s *reviewerService) DeleteArticleReviewer(ctx context.Context, req *pb.DeleteArticleReviewerReq) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteReviewer--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Submission().Reviewer().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteReviewer--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
