package models

type AdminJournalCreateReq struct {
	Title       string
	Description string
	Isbn        string
	Author      string
	Email       string
}

type AdminJournalCreateRes struct {
	Id          string
	Title       string
	Description string
	Isbn        string
	Author      string
	Email       string
}
