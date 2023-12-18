package content

import (
	"editory_submission/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type contentRepo struct {
	db             *pgxpool.Pool
	journal        storage.JournalRepoI
	article        storage.ArticleRepoI
	edition        storage.EditionRepoI
	countryAndCity storage.CountryAndCityRepoI
	university     storage.UniversityRepoI
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

func (s *contentRepo) University() storage.UniversityRepoI {
	if s.university == nil {
		s.university = NewUniversityRepo(s.db)
	}

	return s.university
}

func (s *contentRepo) Edition() storage.EditionRepoI {
	if s.edition == nil {
		s.edition = NewEditionRepo(s.db)
	}

	return s.edition
}
