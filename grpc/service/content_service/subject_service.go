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

type subjectService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedSubjectServiceServer
}

func NewSubjectService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *subjectService {
	return &subjectService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *subjectService) CreateSubject(ctx context.Context, req *pb.CreateSubjectReq) (res *pb.CreateSubjectRes, err error) {
	s.log.Info("---CreateSubject--->", logger.Any("req", req))

	res, err = s.strg.Content().Subject().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateSubject--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *subjectService) GetSubject(ctx context.Context, req *pb.GetSubjectReq) (res *pb.GetSubjectRes, err error) {
	s.log.Info("---GetSubject--->", logger.Any("req", req))

	res, err = s.strg.Content().Subject().Get(ctx, req)
	if err != nil {
		s.log.Error("!!!GetSubject--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *subjectService) GetSubjectList(ctx context.Context, req *pb.GetSubjectListReq) (res *pb.GetSubjectListRes, err error) {
	s.log.Info("---GetSubjectList--->", logger.Any("req", req))

	res, err = s.strg.Content().Subject().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetSubjectList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *subjectService) UpdateSubject(ctx context.Context, req *pb.UpdateSubjectReq) (res *pb.UpdateSubjectRes, err error) {
	s.log.Info("---UpdateSubject--->", logger.Any("req", req))

	res, err = s.strg.Content().Subject().Update(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateSubject--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *subjectService) DeleteSubject(ctx context.Context, req *pb.DeleteSubjectReq) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteSubject--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Content().Subject().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteSubject--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
