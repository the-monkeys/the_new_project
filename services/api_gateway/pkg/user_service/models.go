package user_service

type ProfileRequestBody struct {
	Id int64 `json:"id"`
}

type UpdateProfile struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	MobileNo    string `json:"mobile,omitempty"`
	About       string `json:"about,omitempty"`
	Instagram   string `json:"instagram,omitempty"`
	Twitter     string `json:"twitter,omitempty"`
	Email       string `json:"email,omitempty"`
}
