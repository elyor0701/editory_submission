package models

type CreateUserDraftReq struct {
	JournalId    string `json:"journal_id,omitempty"`
	Type         string `json:"type,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	GroupId      string `json:"group_id,omitempty"`
	Manuscript   string `json:"manuscript,omitempty"`
	CoverLetter  string `json:"cover_letter,omitempty"`
	Supplemental string `json:"supplemental,omitempty"`
}
