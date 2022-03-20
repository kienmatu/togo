package presenter

type LogInResponse struct {
	Token string `json:"token"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}