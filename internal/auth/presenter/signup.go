package presenter

type SignUpInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Limit    int    `json:"limit"`
}

type SignUpResponse struct {
	Username string `json:"username"`
	Limit    int    `json:"limit"`
}