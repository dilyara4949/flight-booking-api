package service

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

type Auth struct {
	repo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *Auth {
	return &Auth{
		repo: userRepo,
	}
}

func (service *Auth) CreateAccessToken(ctx context.Context, user domain.User, jwtSecret string, expiry int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(expiry) * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (service *Auth) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return service.repo.Delete(ctx, id)
}
