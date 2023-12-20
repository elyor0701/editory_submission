package models

type LoginReq struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
