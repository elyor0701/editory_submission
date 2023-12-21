package util

import (
	"database/sql"
	"github.com/jackc/pgconn"
	"strings"
)

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func NewNullInt32(i int32) sql.NullInt32 {
	if i == 0 {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: i,
		Valid: true,
	}
}

func IsErrDuplicateKey(err error) bool {
	pgErr, ok := err.(*pgconn.PgError)
	if ok && pgErr.Code == "23505" {
		return true
	}
	return false
}

func IsErrNoRows(err error) bool {
	return strings.Contains(err.Error(), "no rows in result set")
}
