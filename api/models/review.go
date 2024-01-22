package models

type CreateArticleReviewerReq struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}

type UpdateUserReviewReq struct {
	Id           string
	Status       string
	Comment      string
	FileComments []*FileComment
}
