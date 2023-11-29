package auth

import (
	"editory_submission/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type authRepo struct {
	db   *pgxpool.Pool
	user storage.UserRepoI
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
