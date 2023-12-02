package models

type CreateSessionReq struct {
	UserId    string
	RoleId    string
	Ip        string
	Data      string
	ExpiresAt string
}
