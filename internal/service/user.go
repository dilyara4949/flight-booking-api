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
		ID:                   userid,
		Email:                signup.Email,
		Role:                 userRole,
		Password:             string(encryptedPassword),
		RequirePasswordReset: false,
	}

	err = service.repo.Create(ctx, &user)
	return user, err
}

func (service *User) Get(ctx context.Context, id uuid.UUID) (domain.User, error) {
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

func (service *User) UpdateUser(ctx context.Context, req request.UpdateUser, userID uuid.UUID) (domain.User, error) {
	user, err := service.Get(ctx, userID)
	if err != nil {
		return domain.User{}, errors.ErrUserNotFound
	}

	if req.Phone != "" {
		user.Phone = req.Phone
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Role != "" {
		user.Role = req.Role
	}

	user, err = service.repo.Update(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (service *User) ResetPassword(ctx context.Context, req request.ResetPassword, requirePasswordReset bool) error {
	user, err := service.ValidateUser(ctx, request.Signin{
		Email:    req.Email,
		Password: req.OldPassword,
	})

	if err != nil {
		return err
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.NewPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return fmt.Errorf("generate password error: %w", err)
	}

	err = service.repo.UpdatePassword(ctx, user.ID, string(encryptedPassword), requirePasswordReset)
	return err
}
