package service

import (
	"context"
	"time"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type authService struct {
	repo domain.UserRepository
}

func NewAuthService(userRepo domain.UserRepository) domain.AuthService {
	return &authService{
		repo: userRepo,
	}
}

func (service *authService) CreateUser(ctx context.Context, user *domain.User, password string) error {
	return service.repo.Create(ctx, user, password)
}
func (service *authService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	return service.repo.Get(ctx, id)
}
func (service *authService) CreateAccessToken(userID string, jwtSecret string, expiry int) (accessToken string, err error) {
	expirationTime := time.Now().Add(time.Duration(expiry) * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
