package models

type CreateUserPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	ImageURL string `json:"image_url"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
	Token   string `json:"token"`
}
type UpdateUserResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}
