package config

import (
	"time"
)

const (
	DatabaseQueryTimeLayout                 = `'YYYY-MM-DD"T"HH24:MI:SS"."MS"Z"TZ'`
	DatabaseTimeLayout        string        = time.RFC3339
	AccessTokenExpiresInTime  time.Duration = 1 * 24 * 60 * time.Minute
	RefreshTokenExpiresInTime time.Duration = 30 * 24 * 60 * time.Minute
)

const (
	// EDITOR user types
	EDITOR      = `EDITOR`
	SUPERADMIN  = `SUPERADMIN`
	PROOFREADER = `PROOFREADER`
	REVIEWER    = `REVIEWER`
	AUTHOR      = `AUTHOR`
)

const (
	// ADMIN platfrom types
	ADMIN = `ADMIN`
	USER  = `USER`
)

const (
	// DEFAULT_PASSWORD user password
	DEFAULT_PASSWORD = `default`
)
