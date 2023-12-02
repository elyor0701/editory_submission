package postgres

import (
	"context"
	"editory_submission/config"
	"editory_submission/storage"
	auth "editory_submission/storage/postgres/auth"
	content "editory_submission/storage/postgres/content"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db      *pgxpool.Pool
	auth    storage.AuthRepoI
	content storage.ContentRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	parseConfig, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))

	fmt.Printf("postgres://%s:%s@%s:%d/%s?sslmode=disable\n",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)
	if err != nil {
		return nil, err
	}

	parseConfig.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, parseConfig)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Auth() storage.AuthRepoI {
	if s.auth == nil {
		s.auth = auth.NewAuthRepo(s.db)
	}

	return s.auth
}

func (s *Store) Content() storage.ContentRepoI {
	if s.content == nil {
		s.content = content.NewContentRepo(s.db)
	}

	return s.content
}
