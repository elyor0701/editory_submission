package models

type UpsertJournalSubjectReq struct {
	JournalId string
	SubjectId string
}

type UpsertJournalSubjectRes struct {
	Id        string
	JournalId string
	SubjectId string
}

type GetSubjectRes struct {
	Id        string
	JournalId string
	SubjectId string
	Title     string
}
