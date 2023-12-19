package auth

import (
	"editory_submission/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type authRepo struct {
	db      *pgxpool.Pool
	user    storage.UserRepoI
	session storage.SessionRepoI
	role    storage.RoleRepoI
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

func (s *authRepo) Role() storage.RoleRepoI {
	if s.role == nil {
		s.role = NewRoleRepo(s.db)
	}

	return s.role
}
