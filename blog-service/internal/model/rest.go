package model

type CreatePostInput struct {
	Title       string
	Description string
	Content     string
}

type UpdatePostInput struct {
	Title       string
	Description string
	Content     string
	UserId      string
}

type AuthServiceRequest struct {
	Token string `json:"token"`
}

type AuthServiceResponse struct {
	UserId string `json:"user_id"`
}
