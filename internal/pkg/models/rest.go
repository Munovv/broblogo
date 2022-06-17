package models

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

type CreatePostInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

type UpdatePostInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	UserId      string `json:"userId"`
}

type AuthServiceRequest struct {
	Token string `json:"token"`
}

type AuthServiceResponse struct {
	UserId string `json:"user_id"`
}

type CreatePartnerInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

type CreatePartnerOutput struct {
	Id string `json:"id"`
}

type DeletePartnerInput struct {
	Id string `json:"id"`
}

type ComposeInput struct {
	Images []string `json:"images"`
}
