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

type RegistrationEmailReq struct {
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
	RedirectUrl     string `json:"redirect_url,omitempty"`
}

type RegistrationEmailRes struct {
	Email         string `json:"email,omitempty"`
	MessageStatus bool   `json:"message_status"`
	UserId        string `json:"user_id,omitempty"`
}

type RegisterDetailReq struct {
	Id         string `json:"id,omitempty"`
	Username   string `json:"username,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Phone      string `json:"phone,omitempty"`
	ExtraPhone string `json:"extra_phone,omitempty"`
	CountryId  string `json:"country_id,omitempty"`
	CityId     string `json:"city_id,omitempty"`
	ProfSphere string `json:"prof_sphere,omitempty"`
	Degree     string `json:"degree,omitempty"`
	Address    string `json:"address,omitempty"`
	PostCode   string `json:"post_code,omitempty"`
	Gender     string `json:"gender,omitempty"`
	IsReviewer bool   `json:"is_reviewer,omitempty"`
}

type RegisterDetailRes struct {
	Id         string `json:"id,omitempty"`
	Username   string `json:"username,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Phone      string `json:"phone,omitempty"`
	ExtraPhone string `json:"extra_phone,omitempty"`
	CountryId  string `json:"country_id,omitempty"`
	CityId     string `json:"city_id,omitempty"`
	ProfSphere string `json:"prof_sphere,omitempty"`
	Degree     string `json:"degree,omitempty"`
	Address    string `json:"address,omitempty"`
	PostCode   string `json:"post_code,omitempty"`
	Gender     string `json:"gender,omitempty"`
	IsReviewer bool   `json:"is_reviewer,omitempty"`
}
