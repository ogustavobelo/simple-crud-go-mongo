package models

type CreateUserPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	ImageURL string `json:"image_url"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
	Token   string `json:"token"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
