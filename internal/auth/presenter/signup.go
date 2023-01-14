package presenter

type SignUpInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}
