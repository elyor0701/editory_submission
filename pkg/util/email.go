package util

import "regexp"

const (
	emailShortString = `^[\w-\.]+@([\w-]+\.)+[\w-]{2,}$`
	emailString      = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

// IsValidEmail ...
func IsValidEmail(email string) bool {
	r := regexp.MustCompile(emailString)
	return r.MatchString(email)
}
