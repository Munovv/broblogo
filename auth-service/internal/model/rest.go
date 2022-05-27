package model

type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInOutput struct {
	Token string `json:"token"`
}

type SignUpInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpOutput struct {
	Message string `json:"message"`
}

type VerifyInput struct {
	Token string `json:"token"`
}

type VerifyOutput struct {
	UserId string `json:"user_id"`
}
