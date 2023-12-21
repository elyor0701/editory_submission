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

const (
	REGISTRATION          = `REGISTRATION`
	RESET_PASSWORD        = `RESET_PASSWORD`
	ACCOUNT_DEACTIVATION  = `ACCOUNT_DEACTIVATION`
	NEW_ARTICLE_TO_REVIEW = `NEW_ARTICLE_TO_REVIEW`
	NEW_JOURNAL_USER      = `NEW_JOURNAL_USER`
)

const (
	EMAIL_STATUS_NEW     = `NEW`
	EMAIL_STATUS_PENDING = `PENDING`
	EMAIL_STATUS_SENT    = `SENT`
	EMAIL_STATUS_FAILED  = `FAILED`
)
