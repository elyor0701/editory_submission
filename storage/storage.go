package storage

import (
	"context"
	pb "editory_submission/genproto/auth_service"
	cs_pb "editory_submission/genproto/content_service"
)

type StorageI interface {
	CloseDB()
	Auth() AuthRepoI
	Content() ContentRepoI
}

type AuthRepoI interface {
	User() UserRepoI
}

type ContentRepoI interface {
	Journal() JournalRepoI
	Article() ArticleRepoI
}

type UserRepoI interface {
	Create(ctx context.Context, req *pb.User) (res *pb.User, err error)
	Get(ctx context.Context, req *pb.GetUserReq) (res *pb.User, err error)
	GetList(ctx context.Context, req *pb.GetUserListReq) (res *pb.GetUserListRes, err error)
	Update(ctx context.Context, req *pb.User) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *pb.DeleteUserReq) (rowsAffected int64, err error)
}

type JournalRepoI interface {
	Create(ctx context.Context, in *cs_pb.CreateJournalReq) (*cs_pb.Journal, error)
	Get(ctx context.Context, in *cs_pb.PrimaryKey) (*cs_pb.Journal, error)
	GetList(ctx context.Context, in *cs_pb.GetList) (*cs_pb.GetJournalListRes, error)
	Update(ctx context.Context, in *cs_pb.Journal) (*cs_pb.Journal, error)
	Delete(ctx context.Context, in *cs_pb.PrimaryKey) (rowsAffected int64, err error)
}

type ArticleRepoI interface {
	Create(ctx context.Context, in *cs_pb.CreateArticleReq) (*cs_pb.Article, error)
	Get(ctx context.Context, in *cs_pb.PrimaryKey) (*cs_pb.Article, error)
	GetList(ctx context.Context, in *cs_pb.GetList) (*cs_pb.GetArticleListRes, error)
	Update(ctx context.Context, in *cs_pb.Article) (*cs_pb.Article, error)
	Delete(ctx context.Context, in *cs_pb.PrimaryKey) (rowsAffected int64, err error)
}
