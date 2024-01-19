package models

type CreateUserDraftReq struct {
	JournalId    string            `json:"journal_id,omitempty"`
	Type         string            `json:"type,omitempty"`
	Title        string            `json:"title,omitempty"`
	Description  string            `json:"description,omitempty"`
	GroupId      string            `json:"group_id,omitempty"`
	Conflict     bool              `json:"conflict,omitempty"`
	Availability string            `json:"availability,omitempty"`
	Funding      string            `json:"funding,omitempty"`
	Status       string            `json:"status,omitempty"`
	DraftStep    string            `json:"draft_step,omitempty"`
	Files        []*AddFileReq     `json:"files,omitempty"`
	Coauthors    []*AddCoAuthorReq `json:"coauthors"`
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

type UpdateJournalDraftReq struct {
	Status        string         `json:"status,omitempty"`
	EditorId      string         `json:"editor_id,omitempty"`
	Id            string         `json:"id,omitempty"`
	CheckerStatus string         `json:"checker_status,omitempty"`
	Comment       string         `json:"comment"`
	FileComments  []*FileComment `json:"file_comment,omitempty"`
}

type FileComment struct {
	Id      string `json:"id"`
	Type    string `json:"type,omitempty"`
	FileId  string `json:"file_id,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type UpdateUserDraftReq struct {
	Id           string        `json:"id"`
	JournalId    string        `json:"journal_id,omitempty"`
	Type         string        `json:"type,omitempty"`
	Title        string        `json:"title,omitempty"`
	Description  string        `json:"description,omitempty"`
	Conflict     bool          `json:"conflict,omitempty"`
	Availability string        `json:"availability,omitempty"`
	Funding      string        `json:"funding,omitempty"`
	Status       string        `json:"status,omitempty"`
	DraftStep    string        `json:"draft_step,omitempty"`
	Files        []*AddFileReq `json:"files,omitempty"`
}

type AddCoAuthorReq struct {
	Email        string `json:"email,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	UniversityId string `json:"university_id,omitempty"`
	CountryId    string `json:"country_id,omitempty"`
}
