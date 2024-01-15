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

type articleService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedArticleServiceServer
}

func NewArticleService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *articleService {
	return &articleService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *articleService) CreateArticle(ctx context.Context, req *pb.CreateArticleReq) (res *pb.CreateArticleRes, err error) {
	s.log.Info("---CreateArticle--->", logger.Any("req", req))

	req.Step = "EDITOR"
	req.EditorStatus = "NEW"

	res, err = s.strg.Submission().Article().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateArticle--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *articleService) GetArticle(ctx context.Context, req *pb.GetArticleReq) (res *pb.GetArticleRes, err error) {
	s.log.Info("---GetArticle--->", logger.Any("req", req))

	res, err = s.strg.Submission().Article().Get(ctx, req)
	if err != nil {
		s.log.Error("!!!GetArticle--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *articleService) GetArticleList(ctx context.Context, req *pb.GetArticleListReq) (res *pb.GetArticleListRes, err error) {
	s.log.Info("---GetArticleList--->", logger.Any("req", req))

	res, err = s.strg.Submission().Article().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetArticleList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *articleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleReq) (res *pb.UpdateArticleRes, err error) {
	s.log.Info("---UpdateArticle--->", logger.Any("req", req))

	rowsAffected, err := s.strg.Submission().Article().Update(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateArticle--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return &pb.UpdateArticleRes{}, nil
}

func (s *articleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleReq) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteArticle--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Submission().Article().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteArticle--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
