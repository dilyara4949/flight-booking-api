package domain

import "context"

type Signup struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	accessToken string
	User
}

type SignupService interface {
	Create(ctx context.Context, user User, password string) (*User, error)
	Get(ctx context.Context, id string) (*User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
}
