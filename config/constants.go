package config

import (
	"time"
)

const (
	DatabaseQueryTimeLayout = `'YYYY-MM-DD"T"HH24:MI:SS"."MS"Z"TZ'`
	// DatabaseTimeLayout
	DatabaseTimeLayout string = time.RFC3339
	// AccessTokenExpiresInTime ... 1 * 24 *
	AccessTokenExpiresInTime time.Duration = 1 * 24 * 60 * time.Minute
	// RefreshTokenExpiresInTime ... 30 * 24 * 60
	RefreshTokenExpiresInTime time.Duration = 30 * 24 * 60 * time.Minute
)
