package domain

import (
	"context"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(ctx context.Context, user User, password string) (*User, error)
	Get(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context, page, pageSize int) ([]User, error)
}
