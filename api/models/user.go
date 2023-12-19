package models

type SendVerificationMessageReq struct {
	UserId      string `json:"user_id"`
	RedirectUrl string `json:"redirect_url"`
}

type EmailVerificationReq struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type UpdateAdminUserReq struct {
	Gender          string `json:"gender,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	Email           string `json:"email,omitempty"`
	CurrentPassword string `json:"current_password,omitempty"`
	NewPassword     string `json:"new_password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
}

type UpdateAdminUserRes struct {
	Gender    string `json:"gender,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

type GetAdminUserRes struct {
	Gender    string `json:"gender,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}

type EmailVerificationRes struct {
	Status bool   `json:"status"`
	UserId string `json:"user_id"`
}
