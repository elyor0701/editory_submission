package content

import (
	"editory_submission/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type contentRepo struct {
	db             *pgxpool.Pool
	journal        storage.JournalRepoI
	article        storage.ArticleRepoI
	countryAndCity storage.CountryAndCityRepoI
}

func NewContentRepo(db *pgxpool.Pool) storage.ContentRepoI {
	return &contentRepo{
		db: db,
	}
}

func (s *contentRepo) Journal() storage.JournalRepoI {
	if s.journal == nil {
		s.journal = NewJournalRepo(s.db)
	}

	return s.journal
}

func (s *contentRepo) Article() storage.ArticleRepoI {
	if s.article == nil {
		s.article = NewArticleRepo(s.db)
	}

	return s.article
}

func (s *contentRepo) CountryAndCity() storage.CountryAndCityRepoI {
	if s.countryAndCity == nil {
		s.countryAndCity = NewCountryAndCityRepo(s.db)
	}

	return s.countryAndCity
}
