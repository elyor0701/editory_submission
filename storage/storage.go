package storage

import (
	"context"
	pb "editory_submission/genproto/auth_service"
	cs_pb "editory_submission/genproto/content_service"
	"editory_submission/genproto/notification_service"
	"editory_submission/genproto/submission_service"
	"editory_submission/storage/postgres/models"
)

type StorageI interface {
	CloseDB()
	Auth() AuthRepoI
	Content() ContentRepoI
	Notification() NotificationRepoI
	Submission() SubmissionRepoI
}

type AuthRepoI interface {
	User() UserRepoI
	Session() SessionRepoI
	Role() RoleRepoI
	Keyword() KeywordRepoI
}

type ContentRepoI interface {
	Journal() JournalRepoI
	JournalAuthor() JournalAuthorRepoI
	Edition() EditionRepoI
	CountryAndCity() CountryAndCityRepoI
	University() UniversityRepoI
	Subject() SubjectRepoI
	Article() ContentArticleRepoI
}

type NotificationRepoI interface {
	Notification() NotifyRepoI
	EmailTemplate() EmailTemplateRepoI
}

type SubmissionRepoI interface {
	Article() ArticleRepoI
	File() FileRepoI
	CoAuthor() CoAuthorRepoI
	Reviewer() ReviewerRepoI
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

type EditionRepoI interface {
	Create(ctx context.Context, in *cs_pb.CreateEditionReq) (*cs_pb.Edition, error)
	Get(ctx context.Context, in *cs_pb.PrimaryKey) (*cs_pb.Edition, error)
	GetList(ctx context.Context, in *cs_pb.GetEditionListReq) (*cs_pb.GetEditionListRes, error)
	Update(ctx context.Context, in *cs_pb.Edition) (*cs_pb.Edition, error)
	Delete(ctx context.Context, in *cs_pb.PrimaryKey) (rowsAffected int64, err error)
}

type ContentArticleRepoI interface {
	Create(ctx context.Context, in *cs_pb.CreateArticleReq) (*cs_pb.CreateArticleRes, error)
	Get(ctx context.Context, in *cs_pb.GetArticleReq) (*cs_pb.GetArticleRes, error)
	GetList(ctx context.Context, in *cs_pb.GetArticleListReq) (*cs_pb.GetArticleListRes, error)
	Update(ctx context.Context, in *cs_pb.UpdateArticleReq) (*cs_pb.UpdateArticleRes, error)
	Delete(ctx context.Context, in *cs_pb.DeleteArticleReq) (rowsAffected int64, err error)
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

type NotifyRepoI interface {
	Create(ctx context.Context, in *notification_service.CreateNotificationReq) (*notification_service.CreateNotificationRes, error)
	Get(ctx context.Context, in *notification_service.GetNotificationReq) (*notification_service.GetNotificationRes, error)
	GetList(ctx context.Context, in *notification_service.GetNotificationListReq) (*notification_service.GetNotificationListRes, error)
	Update(ctx context.Context, in *notification_service.UpdateNotificationReq) (*notification_service.UpdateNotificationRes, error)
	Delete(ctx context.Context, in *notification_service.DeleteNotificationReq) (rowsAffected int64, err error)
}

type EmailTemplateRepoI interface {
	Create(ctx context.Context, in *notification_service.CreateEmailTmpReq) (*notification_service.CreateEmailTmpRes, error)
	Get(ctx context.Context, in *notification_service.GetEmailTmpReq) (*notification_service.GetEmailTmpRes, error)
	GetList(ctx context.Context, in *notification_service.GetEmailTmpListReq) (*notification_service.GetEmailTmpListRes, error)
	Update(ctx context.Context, in *notification_service.UpdateEmailTmpReq) (*notification_service.UpdateEmailTmpRes, error)
	Delete(ctx context.Context, in *notification_service.DeleteEmailTmpReq) (rowsAffected int64, err error)
}

type KeywordRepoI interface {
	Create(ctx context.Context, in *pb.CreateKeywordReq) (*pb.CreateKeywordRes, error)
	Get(ctx context.Context, in *pb.GetKeywordReq) (*pb.GetKeywordRes, error)
	GetList(ctx context.Context, in *pb.GetKeywordListReq) (*pb.GetKeywordListRes, error)
	Update(ctx context.Context, in *pb.UpdateKeywordReq) (*pb.UpdateKeywordRes, error)
	Delete(ctx context.Context, in *pb.DeleteKeywordReq) (rowsAffected int64, err error)
}

type ArticleRepoI interface {
	Create(ctx context.Context, in *submission_service.CreateArticleReq) (*submission_service.CreateArticleRes, error)
	Get(ctx context.Context, in *submission_service.GetArticleReq) (*submission_service.GetArticleRes, error)
	GetList(ctx context.Context, in *submission_service.GetArticleListReq) (*submission_service.GetArticleListRes, error)
	Update(ctx context.Context, in *submission_service.UpdateArticleReq) (rowsAffected int64, err error)
	Delete(ctx context.Context, in *submission_service.DeleteArticleReq) (rowsAffected int64, err error)
}

type FileRepoI interface {
	Create(ctx context.Context, in *submission_service.AddFilesReq) (*submission_service.AddFilesRes, error)
	//Get(ctx context.Context, in *cs_pb.GetSubjectReq) (*cs_pb.GetSubjectRes, error)
	GetList(ctx context.Context, in *submission_service.GetFilesReq) (*submission_service.GetFilesRes, error)
	//Update(ctx context.Context, in *cs_pb.UpdateSubjectReq) (*cs_pb.UpdateSubjectRes, error)
	Delete(ctx context.Context, in *submission_service.DeleteFilesReq) (rowsAffected int64, err error)
}

type CoAuthorRepoI interface {
	Create(ctx context.Context, in *submission_service.AddCoAuthorReq) (*submission_service.AddCoAuthorRes, error)
	//Get(ctx context.Context, in *cs_pb.GetSubjectReq) (*cs_pb.GetSubjectRes, error)
	GetList(ctx context.Context, in *submission_service.GetCoAuthorsReq) (*submission_service.GetCoAuthorsRes, error)
	//Update(ctx context.Context, in *cs_pb.UpdateSubjectReq) (*cs_pb.UpdateSubjectRes, error)
	Delete(ctx context.Context, in *submission_service.DeleteCoAuthorReq) (rowsAffected int64, err error)
}

type ReviewerRepoI interface {
	Create(ctx context.Context, in *submission_service.CreateArticleCheckerReq) (*submission_service.CreateArticleCheckerRes, error)
	Get(ctx context.Context, in *submission_service.GetArticleCheckerReq) (*submission_service.GetArticleCheckerRes, error)
	GetList(ctx context.Context, in *submission_service.GetArticleCheckerListReq) (*submission_service.GetArticleCheckerListRes, error)
	Update(ctx context.Context, in *submission_service.UpdateArticleCheckerReq) (rowsAffected int64, err error)
	Delete(ctx context.Context, in *submission_service.DeleteArticleCheckerReq) (rowsAffected int64, err error)
}

type JournalAuthorRepoI interface {
	Create(ctx context.Context, in *cs_pb.CreateJournalAuthorReq) (*cs_pb.CreateJournalAuthorRes, error)
	Get(ctx context.Context, in *cs_pb.GetJournalAuthorReq) (*cs_pb.GetJournalAuthorRes, error)
	GetList(ctx context.Context, in *cs_pb.GetJournalAuthorListReq) (*cs_pb.GetJournalAuthorListRes, error)
	Update(ctx context.Context, in *cs_pb.UpdateJournalAuthorReq) (rowsAffected int64, err error)
	Delete(ctx context.Context, in *cs_pb.DeleteJournalAuthorReq) (rowsAffected int64, err error)
}
