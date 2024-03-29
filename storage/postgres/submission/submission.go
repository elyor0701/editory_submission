package submission

import (
	"editory_submission/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type submissionRepo struct {
	db       *pgxpool.Pool
	reviewer storage.ReviewerRepoI
	article  storage.ArticleRepoI
	file     storage.FileRepoI
	coAuthor storage.CoAuthorRepoI
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

func (s submissionRepo) File() storage.FileRepoI {
	if s.file == nil {
		s.file = NewFileRepo(s.db)
	}

	return s.file
}

func (s submissionRepo) CoAuthor() storage.CoAuthorRepoI {
	if s.coAuthor == nil {
		s.coAuthor = NewCoAuthorRepo(s.db)
	}

	return s.coAuthor
}
