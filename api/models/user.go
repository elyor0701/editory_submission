package models

type SendVerificationMessageReq struct {
	UserId      string `json:"user_id"`
	RedirectUrl string `json:"redirect_url"`
}

type EmailVerificationReq struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
