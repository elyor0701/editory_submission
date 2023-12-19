package models

type Role struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	RoleType  string `json:"role_type"`
	JournalId string `json:"journal_id"`
}
