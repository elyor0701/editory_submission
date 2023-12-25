package util

import (
	"database/sql"
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

//func IsErrDuplicateKey(err error) bool {
//	pgErr, ok := err.(*pgconn.PgError)
//	fmt.Println(pgErr.Code)
//	if ok && pgErr.Code == "23505" {
//		return true
//	}
//	return false
//}

func IsErrDuplicateKey(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}

func IsErrNoRows(err error) bool {
	return strings.Contains(err.Error(), "no rows in result set")
}
