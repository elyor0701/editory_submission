package auth_service

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/auth_service"
	"editory_submission/grpc/client"
	"editory_submission/pkg/logger"
	"editory_submission/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type keywordService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedKeywordServiceServer
}

func NewKeywordService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *keywordService {
	return &keywordService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *keywordService) CreateKeyword(ctx context.Context, req *pb.CreateKeywordReq) (res *pb.CreateKeywordRes, err error) {
	s.log.Info("---CreateKeyword--->", logger.Any("req", req))

	res, err = s.strg.Auth().Keyword().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateKeyword--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *keywordService) GetKeyword(ctx context.Context, req *pb.GetKeywordReq) (res *pb.GetKeywordRes, err error) {
	s.log.Info("---GetKeyword--->", logger.Any("req", req))

	res, err = s.strg.Auth().Keyword().Get(ctx, req)
	if err != nil {
		s.log.Error("!!!GetKeyword--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *keywordService) GetKeywordList(ctx context.Context, req *pb.GetKeywordListReq) (res *pb.GetKeywordListRes, err error) {
	s.log.Info("---GetKeywordList--->", logger.Any("req", req))

	res, err = s.strg.Auth().Keyword().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetKeywordList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *keywordService) UpdateKeyword(ctx context.Context, req *pb.UpdateKeywordReq) (res *pb.UpdateKeywordRes, err error) {
	s.log.Info("---UpdateKeyword--->", logger.Any("req", req))

	res, err = s.strg.Auth().Keyword().Update(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateKeyword--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *keywordService) DeleteKeyword(ctx context.Context, req *pb.DeleteKeywordReq) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteKeyword--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Auth().Keyword().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteKeyword--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
