package models

type EmailVerification struct {
	Email     string
	Token     string
	Sent      bool
	ExpiresAt string
	CreatedAt string
}

type CreateEmailVerificationReq struct {
	Email     string
	Token     string
	ExpiresAt string
}

type CreateEmailVerificationRes struct {
	Email     string
	Token     string
	ExpiresAt string
}

type UpdateEmailVerificationReq struct {
	Email string
	Token string
	Sent  bool
}

type UpdateEmailVerificationRes struct {
	Email string
	Token string
	Sent  bool
}

type GetEmailVerificationListReq struct {
	Email string
}

type GetEmailVerificationListRes struct {
	Tokens []*EmailVerification
	Count  int
}

type DeleteEmailVerificationReq struct {
	Email string
}

type UpdateUserEmailVerificationStatusReq struct {
	Email              string
	VerificationStatus bool
}
