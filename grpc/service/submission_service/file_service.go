package submission_service

import (
	"context"
	"editory_submission/genproto/submission_service"
	"editory_submission/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *articleService) AddFiles(ctx context.Context, req *submission_service.AddFilesReq) (res *submission_service.AddFilesRes, err error) {
	s.log.Info("---AddFiles--->", logger.Any("req", req))

	res, err = s.strg.Submission().File().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!AddFiles--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *articleService) GetFiles(ctx context.Context, req *submission_service.GetFilesReq) (res *submission_service.GetFilesRes, err error) {
	s.log.Info("---GetFiles--->", logger.Any("req", req))

	res, err = s.strg.Submission().File().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetFiles--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *articleService) DeleteFiles(ctx context.Context, req *submission_service.DeleteFilesReq) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteFiles--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Submission().File().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteFiles--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
