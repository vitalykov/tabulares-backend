package model

type RegisterResponse struct {
	UserID string `json:"user_id"`
}

type LoginResponse struct {
	UserID string `json:"user_id"`
}

type GetByIDResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type FindByUsernameResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type DeleteResponse struct {
	UserID string `json:"user_id"`
}
