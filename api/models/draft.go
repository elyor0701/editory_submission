package models

type CreateUserDraftReq struct {
	JournalId    string        `json:"journal_id,omitempty"`
	Type         string        `json:"type,omitempty"`
	Title        string        `json:"title,omitempty"`
	Description  string        `json:"description,omitempty"`
	GroupId      string        `json:"group_id,omitempty"`
	Conflict     bool          `json:"conflict,omitempty"`
	Availability string        `json:"availability,omitempty"`
	Funding      string        `json:"funding,omitempty"`
	Status       string        `json:"status,omitempty"`
	DraftStep    string        `json:"draft_step,omitempty"`
	Step         string        `json:"step,omitempty"`
	Files        []*AddFileReq `json:"files,omitempty"`
}

type AddFileReq struct {
	Url  string `json:"url,omitempty"`
	Type string `json:"type,omitempty"`
}

type File struct {
	Id        string `json:"id,omitempty"`
	Url       string `json:"url,omitempty"`
	Type      string `json:"type,omitempty"`
	ArticleId string `json:"article_id,omitempty"`
}
