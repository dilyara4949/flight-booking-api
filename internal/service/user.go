package service

import (
	"context"
	"fmt"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	"github.com/dilyara4949/flight-booking-api/internal/repository/errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const userRole = "user"

type User struct {
	repo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *User {
	return &User{
		repo: userRepo,
	}
}

func (service *User) CreateUser(ctx context.Context, signup request.Signup, password string) (domain.User, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("generate password error: %w", err)
	}

	userid, err := uuid.NewV7()
	if err != nil {
		return domain.User{}, fmt.Errorf("generate uuid error: %w", err)
	}

	user := domain.User{
		ID:       userid,
		Email:    signup.Email,
		Role:     userRole,
		Password: string(encryptedPassword),
	}

	err = service.repo.Create(ctx, &user)

	return user, err
}

func (service *User) GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return service.repo.Get(ctx, id)
}

func (service *User) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return service.repo.GetByEmail(ctx, email)
}

func (service *User) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return service.repo.Delete(ctx, id)
}

func (service *User) ValidateUser(ctx context.Context, signin request.Signin) (domain.User, error) {
	user, err := service.repo.GetByEmail(ctx, signin.Email)
	if err != nil {
		return domain.User{}, errors.ErrUserNotFound
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signin.Password)) != nil {
		return domain.User{}, errors.ErrInvalidEmailPassword
	}

	return *user, nil
}
