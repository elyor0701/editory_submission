package content_service

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/content_service"
	"editory_submission/grpc/client"
	"editory_submission/pkg/logger"
	"editory_submission/storage"
	"editory_submission/storage/postgres/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type contentService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedContentServiceServer
}

func NewContentService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *contentService {
	return &contentService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *contentService) CreateJournal(ctx context.Context, req *pb.CreateJournalReq) (res *pb.Journal, err error) {
	s.log.Info("---CreateJournal--->", logger.Any("req", req))

	res, err = s.strg.Content().Journal().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateJournal--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	for _, val := range req.GetJournalData() {
		_, err = s.strg.Content().Journal().UpsertJournalData(ctx, &pb.JournalData{
			JournalId: res.GetId(),
			Text:      val.GetText(),
			Type:      val.GetType(),
		})
		if err != nil {
			s.log.Error("!!!CreateJournalData--->", logger.Error(err))
			continue
		}
	}

	for _, val := range req.GetSubjects() {
		_, err = s.strg.Content().Journal().UpsertSubject(ctx, &models.UpsertJournalSubjectReq{
			JournalId: res.GetId(),
			SubjectId: val.GetId(),
		})
		if err != nil {
			s.log.Error("!!!CreateSubject--->", logger.Error(err))
			continue
		}
	}

	return res, nil
}

func (s *contentService) GetJournal(ctx context.Context, req *pb.PrimaryKey) (res *pb.Journal, err error) {
	s.log.Info("---GetJournal--->", logger.Any("req", req))

	res, err = s.strg.Content().Journal().Get(ctx, req)
	if err != nil {
		s.log.Error("!!!GetJournal--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	journalData, err := s.strg.Content().Journal().GetJournalData(ctx, &pb.PrimaryKey{
		Id: req.GetId(),
	})
	if err != nil {
		s.log.Error("!!!GetJournalData--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	res.JournalData = journalData

	subjects, err := s.strg.Content().Journal().GetSubject(ctx, &pb.PrimaryKey{
		Id: req.GetId(),
	})
	if err != nil {
		s.log.Error("!!!GetJournalSubject--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	res.Subjects = subjects

	return res, nil
}

func (s *contentService) GetJournalList(ctx context.Context, req *pb.GetList) (res *pb.GetJournalListRes, err error) {
	s.log.Info("---GetJournalList--->", logger.Any("req", req))

	res, err = s.strg.Content().Journal().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetJournalList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *contentService) UpdateJournal(ctx context.Context, req *pb.Journal) (res *pb.Journal, err error) {
	s.log.Info("---UpdateJournal--->", logger.Any("req", req))

	res, err = s.strg.Content().Journal().Update(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateJournal--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	for _, val := range req.GetJournalData() {
		_, err = s.strg.Content().Journal().UpsertJournalData(ctx, &pb.JournalData{
			JournalId: res.GetId(),
			Text:      val.GetText(),
			Type:      val.GetType(),
		})
		if err != nil {
			s.log.Error("!!!CreateJournalData--->", logger.Error(err))
			continue
		}
	}

	if len(req.GetSubjects()) > 0 {
		_, err = s.strg.Content().Journal().DeleteSubject(ctx, &pb.PrimaryKey{
			Id: res.GetId(),
		})
		if err != nil {
			s.log.Error("!!!DeleteSubject--->", logger.Error(err))
		}

		for _, val := range req.GetSubjects() {
			_, err = s.strg.Content().Journal().UpsertSubject(ctx, &models.UpsertJournalSubjectReq{
				JournalId: res.GetId(),
				SubjectId: val.GetId(),
			})
			if err != nil {
				s.log.Error("!!!CreateSubject--->", logger.Error(err))
				continue
			}
		}
	}

	return res, nil
}

func (s *contentService) DeleteJournal(ctx context.Context, req *pb.PrimaryKey) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteJournal--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Content().Journal().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteJournal--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}

//func (s *contentService) CreateArticle(ctx context.Context, req *pb.CreateArticleReq) (res *pb.Article, err error) {
//	s.log.Info("---CreateArticle--->", logger.Any("req", req))
//
//	res, err = s.strg.Content().Article().Create(ctx, req)
//	if err != nil {
//		s.log.Error("!!!CreateArticle--->", logger.Error(err))
//		return nil, status.Error(codes.InvalidArgument, err.Error())
//	}
//
//	return res, nil
//}

//func (s *contentService) GetArticle(ctx context.Context, req *pb.PrimaryKey) (res *pb.Article, err error) {
//	s.log.Info("---GetArticle--->", logger.Any("req", req))
//
//	res, err = s.strg.Content().Article().Get(ctx, req)
//	if err != nil {
//		s.log.Error("!!!GetArticle--->", logger.Error(err))
//		return nil, status.Error(codes.NotFound, err.Error())
//	}
//
//	return res, nil
//}

//func (s *contentService) GetArticleList(ctx context.Context, req *pb.GetArticleListReq) (res *pb.GetArticleListRes, err error) {
//	s.log.Info("---GetArticleList--->", logger.Any("req", req))
//
//	res, err = s.strg.Content().Article().GetList(ctx, req)
//	if err != nil {
//		s.log.Error("!!!GetArticleList--->", logger.Error(err))
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return res, nil
//}

//func (s *contentService) UpdateArticle(ctx context.Context, req *pb.Article) (res *pb.Article, err error) {
//	s.log.Info("---UpdateArticle--->", logger.Any("req", req))
//
//	res, err = s.strg.Content().Article().Update(ctx, req)
//	if err != nil {
//		s.log.Error("!!!UpdateArticle--->", logger.Error(err))
//		return nil, status.Error(codes.InvalidArgument, err.Error())
//	}
//
//	return res, nil
//}

//func (s *contentService) DeleteArticle(ctx context.Context, req *pb.PrimaryKey) (res *emptypb.Empty, err error) {
//	s.log.Info("---DeleteArticle--->", logger.Any("req", req))
//
//	res = &emptypb.Empty{}
//
//	rowsAffected, err := s.strg.Content().Article().Delete(ctx, req)
//	if err != nil {
//		s.log.Error("!!!DeleteArticle--->", logger.Error(err))
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	if rowsAffected <= 0 {
//		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
//	}
//
//	return res, nil
//}

func (s *contentService) GetCountryList(ctx context.Context, req *pb.GetCountryListReq) (res *pb.GetCountryListRes, err error) {
	s.log.Info("---GetCountryList--->", logger.Any("req", req))

	res, err = s.strg.Content().CountryAndCity().GetCountyList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetCountryList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *contentService) GetCityList(ctx context.Context, req *pb.GetCityListReq) (res *pb.GetCityListRes, err error) {
	s.log.Info("---GetCityList--->", logger.Any("req", req))

	res, err = s.strg.Content().CountryAndCity().GetCityList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetCityList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *contentService) CreateEdition(ctx context.Context, req *pb.CreateEditionReq) (res *pb.Edition, err error) {
	s.log.Info("---CreateEdition--->", logger.Any("req", req))

	res, err = s.strg.Content().Edition().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateEdition--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *contentService) GetEdition(ctx context.Context, req *pb.PrimaryKey) (res *pb.Edition, err error) {
	s.log.Info("---GetEdition--->", logger.Any("req", req))

	res, err = s.strg.Content().Edition().Get(ctx, req)
	if err != nil {
		s.log.Error("!!!GetEdition--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *contentService) GetEditionList(ctx context.Context, req *pb.GetEditionListReq) (res *pb.GetEditionListRes, err error) {
	s.log.Info("---GetEditionList--->", logger.Any("req", req))

	res, err = s.strg.Content().Edition().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetEditionList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *contentService) UpdateEdition(ctx context.Context, req *pb.Edition) (res *pb.Edition, err error) {
	s.log.Info("---UpdateEdition--->", logger.Any("req", req))

	res, err = s.strg.Content().Edition().Update(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateEdition--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *contentService) DeleteEdition(ctx context.Context, req *pb.PrimaryKey) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteEdition--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Content().Edition().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteEdition--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
