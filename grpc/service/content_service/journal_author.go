package content_service

import (
	"context"
	pb "editory_submission/genproto/content_service"
	"editory_submission/pkg/logger"
	"editory_submission/pkg/util"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *contentService) CreateJournalAuthor(ctx context.Context, req *pb.CreateJournalAuthorReq) (res *pb.CreateJournalAuthorRes, err error) {
	s.log.Info("---CreateJournalAuthor--->", logger.Any("req", req))

	if !util.IsValidEmail(req.Email) {
		err = fmt.Errorf("email is not valid")
		s.log.Error("!!!CreateJournalAuthor--->", logger.Error(err))
		return nil, err
	}

	res, err = s.strg.Content().JournalAuthor().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateJournalAuthor--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *contentService) GetJournalAuthor(ctx context.Context, req *pb.GetJournalAuthorReq) (res *pb.GetJournalAuthorRes, err error) {
	s.log.Info("---GetJournalAuthor--->", logger.Any("req", req))

	res, err = s.strg.Content().JournalAuthor().Get(ctx, req)
	if err != nil {
		s.log.Error("!!!GetJournalAuthor--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *contentService) GetJournalAuthorList(ctx context.Context, req *pb.GetJournalAuthorListReq) (res *pb.GetJournalAuthorListRes, err error) {
	s.log.Info("---GetJournalAuthorList--->", logger.Any("req", req))

	res, err = s.strg.Content().JournalAuthor().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetJournalAuthorList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *contentService) UpdateJournalAuthor(ctx context.Context, req *pb.UpdateJournalAuthorReq) (res *pb.UpdateJournalAuthorRes, err error) {
	s.log.Info("---UpdateJournalAuthor--->", logger.Any("req", req))

	if req.Email != "" && !util.IsValidEmail(req.Email) {
		err = fmt.Errorf("email is not valid")
		s.log.Error("!!!CreateJournalAuthor--->", logger.Error(err))
		return nil, err
	}

	rowsAffected, err := s.strg.Content().JournalAuthor().Update(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateJournalAuthor--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return &pb.UpdateJournalAuthorRes{}, nil
}

func (s *contentService) DeleteJournalAuthor(ctx context.Context, req *pb.DeleteJournalAuthorReq) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteJournalAuthor--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Content().JournalAuthor().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteJournalAuthor--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
