package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	errs "github.com/dilyara4949/flight-booking-api/internal/repository/errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

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
		Role:     signup.Role,
		Password: string(encryptedPassword),
	}

	err = service.repo.Create(ctx, &user)
	return user, err
}

func (service *User) GetUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return service.repo.Get(ctx, id)
}

func (service *User) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return service.repo.Delete(ctx, id)
}

func (service *User) UpdateUser(ctx context.Context, req request.UpdateUser, userID uuid.UUID) (domain.User, error) {
	user, err := service.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, errs.ErrUserNotFound
	}

	user.Email = req.Email
	user.Phone = req.Phone

	if req.Role != "" {
		user.Role = req.Role
	}

	user, err = service.repo.Update(ctx, user)
	if err != nil {
		return domain.User{}, errors.New("email already exists")
	}

	return user, nil
}
