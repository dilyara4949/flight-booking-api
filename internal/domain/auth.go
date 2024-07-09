package domain

import "context"

type Signup struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	AccessToken string
	User
}

type AuthService interface {
	CreateUser(ctx context.Context, user *User, password string) error
	GetUser(ctx context.Context, id string) (*User, error)
	CreateAccessToken(userID string, secret string, expiry int) (accessToken string, err error)
}
