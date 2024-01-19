package submission_service

import (
	"context"
	"editory_submission/genproto/submission_service"
	"editory_submission/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *articleService) AddCoAuthor(ctx context.Context, req *submission_service.AddCoAuthorReq) (res *submission_service.AddCoAuthorRes, err error) {
	s.log.Info("---AddCoAuthor--->", logger.Any("req", req))

	res, err = s.strg.Submission().CoAuthor().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!AddCoAuthor--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *articleService) GetCoAuthors(ctx context.Context, req *submission_service.GetCoAuthorsReq) (res *submission_service.GetCoAuthorsRes, err error) {
	s.log.Info("---GetCoAuthors--->", logger.Any("req", req))

	res, err = s.strg.Submission().CoAuthor().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetCoAuthors--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *articleService) DeleteCoAuthor(ctx context.Context, req *submission_service.DeleteCoAuthorReq) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteCoAuthor--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Submission().CoAuthor().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteCoAuthor--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
