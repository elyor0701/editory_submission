package models

type CreateArticleCheckReq struct {
	EditorId     string         `json:"editor_id,omitempty"`
	Comment      string         `json:"comment,omitempty"`
	Status       string         `json:"status"`
	FileComments []*FileComment `json:"file_comments,omitempty"`
}

type UpdateArticleCheckReq struct {
	Id           string
	Status       string
	Comment      string
	FileComments []*FileComment
}
