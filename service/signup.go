package service

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
)

type signupService struct {
	repo domain.UserRepository
}

func NewSignupService(userRepo domain.UserRepository) domain.SignupService {
	return &signupService{
		repo: userRepo,
	}
}

func (service *signupService) Create(ctx context.Context, user domain.User, password string) (*domain.User, error) {
	return service.repo.Create(ctx, user, password)
}
func (service *signupService) Get(ctx context.Context, id string) (*domain.User, error) {
	return service.repo.Get(ctx, id)
}
func (service *signupService) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return "", nil
}
