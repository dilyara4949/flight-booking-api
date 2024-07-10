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
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type authService struct {
	repo           repository.UserRepository
	contextTimeout time.Duration
}
type AuthService interface {
	CreateUser(ctx context.Context, user *domain.User, password string) error
	GetUser(ctx context.Context, id string) (*domain.User, error)
	DeleteUser(ctx context.Context, id string) error
	CreateAccessToken(ctx context.Context, user domain.User, secret string, expiry int) (accessToken string, err error)
}

func NewAuthService(userRepo repository.UserRepository, contextTimeout time.Duration) AuthService {
	return &authService{
		repo:           userRepo,
		contextTimeout: contextTimeout,
	}
}

func (service *authService) CreateUser(ctx context.Context, user *domain.User, password string) error {
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeout)
	defer cancel()

	user.ID = uuid.New().String()

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return fmt.Errorf("generate password error: %v", err)
	}

	return service.repo.Create(ctx, user, string(encryptedPassword))
}
func (service *authService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeout)
	defer cancel()

	return service.repo.Get(ctx, id)
}

func (service *authService) DeleteUser(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeout)
	defer cancel()

	return service.repo.Delete(ctx, id)
}

func (service *authService) CreateAccessToken(ctx context.Context, user domain.User, jwtSecret string, expiry int) (accessToken string, err error) {
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
