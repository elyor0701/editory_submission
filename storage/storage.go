package storage

import (
	"context"
	pb "editory_submission/genproto/auth_service"
	cs_pb "editory_submission/genproto/content_service"
	"editory_submission/storage/postgres/models"
)

type StorageI interface {
	CloseDB()
	Auth() AuthRepoI
	Content() ContentRepoI
}

type AuthRepoI interface {
	User() UserRepoI
	Session() SessionRepoI
	Role() RoleRepoI
}

type ContentRepoI interface {
	Journal() JournalRepoI
	Article() ArticleRepoI
	Edition() EditionRepoI
	CountryAndCity() CountryAndCityRepoI
	University() UniversityRepoI
	Subject() SubjectRepoI
}

type UserRepoI interface {
	Create(ctx context.Context, req *pb.User) (res *pb.User, err error)
	Get(ctx context.Context, req *pb.GetUserReq) (res *pb.User, err error)
	GetList(ctx context.Context, req *pb.GetUserListReq) (res *pb.GetUserListRes, err error)
	GetListWithRole(ctx context.Context, req *pb.GetUserListByRoleReq) (res *pb.GetUserListByRoleRes, err error) // @TODO check performance
	Update(ctx context.Context, req *pb.User) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *pb.DeleteUserReq) (rowsAffected int64, err error)
	GetByEmail(ctx context.Context, req *pb.GetUserReq) (res *pb.User, err error)
	CreateEmailVerification(ctx context.Context, req *models.CreateEmailVerificationReq) (res *models.CreateEmailVerificationRes, err error)
	GetEmailVerificationList(ctx context.Context, req *models.GetEmailVerificationListReq) (res *models.GetEmailVerificationListRes, err error)
	DeleteEmailVerification(ctx context.Context, req *models.DeleteEmailVerificationReq) (rowsAffected int64, err error)
	UpdateEmailVerification(ctx context.Context, req *models.UpdateEmailVerificationReq) (res *models.UpdateEmailVerificationRes, err error)
	UpdateUserEmailVerificationStatus(ctx context.Context, req *models.UpdateUserEmailVerificationStatusReq) (rowsAffected int64, err error)
}

type SessionRepoI interface {
	Create(ctx context.Context, in *pb.CreateSessionReq) (pKey *pb.Session, err error)
	GetList(ctx context.Context, in *pb.SessionGetList) (res *pb.GetSessionListRes, err error)
	GetByPK(ctx context.Context, in *pb.SessionPrimaryKey) (res *pb.Session, err error)
	Update(ctx context.Context, in *pb.UpdateSessionReq) (res *pb.Session, err error)
	Delete(ctx context.Context, in *pb.SessionPrimaryKey) (rowsAffected int64, err error)
	DeleteExpiredUserSessions(ctx context.Context, userID string) (rowsAffected int64, err error)
	GetSessionListByUserID(ctx context.Context, userID string) (res *pb.GetSessionListRes, err error)
}

type JournalRepoI interface {
	Create(ctx context.Context, in *cs_pb.CreateJournalReq) (*cs_pb.Journal, error)
	Get(ctx context.Context, in *cs_pb.PrimaryKey) (*cs_pb.Journal, error)
	GetList(ctx context.Context, in *cs_pb.GetList) (*cs_pb.GetJournalListRes, error)
	Update(ctx context.Context, in *cs_pb.Journal) (*cs_pb.Journal, error)
	Delete(ctx context.Context, in *cs_pb.PrimaryKey) (rowsAffected int64, err error)
	UpsertJournalData(ctx context.Context, in *cs_pb.JournalData) (*cs_pb.JournalData, error)
	GetJournalData(ctx context.Context, in *cs_pb.PrimaryKey) ([]*cs_pb.JournalData, error)
	UpsertSubject(ctx context.Context, in *models.UpsertJournalSubjectReq) (*models.UpsertJournalSubjectRes, error)
	GetSubject(ctx context.Context, in *cs_pb.PrimaryKey) ([]*cs_pb.Subject, error)
	DeleteSubject(ctx context.Context, in *cs_pb.PrimaryKey) (rowsAffected int64, err error)
}

type ArticleRepoI interface {
	Create(ctx context.Context, in *cs_pb.CreateArticleReq) (*cs_pb.Article, error)
	Get(ctx context.Context, in *cs_pb.PrimaryKey) (*cs_pb.Article, error)
	GetList(ctx context.Context, in *cs_pb.GetArticleListReq) (*cs_pb.GetArticleListRes, error)
	Update(ctx context.Context, in *cs_pb.Article) (*cs_pb.Article, error)
	Delete(ctx context.Context, in *cs_pb.PrimaryKey) (rowsAffected int64, err error)
}

type EditionRepoI interface {
	Create(ctx context.Context, in *cs_pb.CreateEditionReq) (*cs_pb.Edition, error)
	Get(ctx context.Context, in *cs_pb.PrimaryKey) (*cs_pb.Edition, error)
	GetList(ctx context.Context, in *cs_pb.GetEditionListReq) (*cs_pb.GetEditionListRes, error)
	Update(ctx context.Context, in *cs_pb.Edition) (*cs_pb.Edition, error)
	Delete(ctx context.Context, in *cs_pb.PrimaryKey) (rowsAffected int64, err error)
}

type CountryAndCityRepoI interface {
	GetCountyList(ctx context.Context, in *cs_pb.GetCountryListReq) (*cs_pb.GetCountryListRes, error)
	GetCityList(ctx context.Context, in *cs_pb.GetCityListReq) (*cs_pb.GetCityListRes, error)
}

type UniversityRepoI interface {
	Create(ctx context.Context, in *cs_pb.CreateUniversityReq) (*cs_pb.CreateUniversityRes, error)
	Get(ctx context.Context, in *cs_pb.GetUniversityReq) (*cs_pb.GetUniversityRes, error)
	GetList(ctx context.Context, in *cs_pb.GetUniversityListReq) (*cs_pb.GetUniversityListRes, error)
	Update(ctx context.Context, in *cs_pb.UpdateUniversityReq) (*cs_pb.UpdateUniversityRes, error)
	Delete(ctx context.Context, in *cs_pb.DeleteUniversityReq) (rowsAffected int64, err error)
}

type SubjectRepoI interface {
	Create(ctx context.Context, in *cs_pb.CreateSubjectReq) (*cs_pb.CreateSubjectRes, error)
	Get(ctx context.Context, in *cs_pb.GetSubjectReq) (*cs_pb.GetSubjectRes, error)
	GetList(ctx context.Context, in *cs_pb.GetSubjectListReq) (*cs_pb.GetSubjectListRes, error)
	Update(ctx context.Context, in *cs_pb.UpdateSubjectReq) (*cs_pb.UpdateSubjectRes, error)
	Delete(ctx context.Context, in *cs_pb.DeleteSubjectReq) (rowsAffected int64, err error)
}

type RoleRepoI interface {
	Create(ctx context.Context, req *pb.Role) (res *pb.Role, err error)
	Get(ctx context.Context, req *pb.GetRoleReq) (res *pb.Role, err error)
	GetList(ctx context.Context, req *pb.GetRoleListReq) (res *pb.GetRoleListRes, err error)
	Update(ctx context.Context, req *pb.Role) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *pb.DeleteRoleReq) (rowsAffected int64, err error)
}
