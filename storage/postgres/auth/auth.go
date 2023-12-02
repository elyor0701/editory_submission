package auth

import (
	"editory_submission/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type authRepo struct {
	db      *pgxpool.Pool
	user    storage.UserRepoI
	session storage.SessionRepoI
}

func NewAuthRepo(db *pgxpool.Pool) storage.AuthRepoI {
	return &authRepo{
		db: db,
	}
}

func (s *authRepo) User() storage.UserRepoI {
	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}

func (s *authRepo) Session() storage.SessionRepoI {
	if s.session == nil {
		s.session = NewSessionRepo(s.db)
	}

	return s.session
}
