package submission

import (
	"editory_submission/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type submissionRepo struct {
	db       *pgxpool.Pool
	reviewer storage.ReviewerRepoI
	article  storage.ArticleRepoI
}

func NewSubmissionRepo(db *pgxpool.Pool) storage.SubmissionRepoI {
	return &submissionRepo{
		db: db,
	}
}

func (s submissionRepo) Article() storage.ArticleRepoI {
	if s.article == nil {
		s.article = NewArticleRepo(s.db)
	}

	return s.article
}

func (s submissionRepo) Reviewer() storage.ReviewerRepoI {
	if s.reviewer == nil {
		s.reviewer = NewReviewerRepo(s.db)
	}

	return s.reviewer
}
