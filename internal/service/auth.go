package service

import (
	"context"
	"fmt"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

type Auth struct {
	repo repository.UserRepository
}
type AuthService interface {
	CreateUser(ctx context.Context, user *domain.User, password string) error
	GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	CreateAccessToken(ctx context.Context, user domain.User, secret string, expiry int) (accessToken string, err error)
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &Auth{
		repo: userRepo,
	}
}

func (service *Auth) CreateUser(ctx context.Context, user *domain.User, password string) error {
	user.ID = uuid.New()

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return fmt.Errorf("generate password error: %v", err)
	}

	user.Password = string(encryptedPassword)

	return service.repo.Create(ctx, user)
}

func (service *Auth) GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return service.repo.Get(ctx, id)
}

func (service *Auth) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return service.repo.Delete(ctx, id)
}

func (service *Auth) CreateAccessToken(ctx context.Context, user domain.User, jwtSecret string, expiry int) (accessToken string, err error) {
	expirationTime := time.Now().Add(time.Duration(expiry) * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err = token.SignedString([]byte(jwtSecret))
	if err != nil {
		err2 := service.DeleteUser(ctx, user.ID)
		if err2 != nil {
			slog.Error("delete user failed: %v", err2)
		}
		return "", err
	}

	return accessToken, nil
}
