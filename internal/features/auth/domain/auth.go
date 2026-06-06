package domain

import user_domain "authTest/internal/features/user/domain"

type AuthResponse struct {
	Token string           `json:"token"`
	User  user_domain.User `json:"user"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
