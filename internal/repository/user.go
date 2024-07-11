package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

var (
	ErrUserNotFound = errors.New("user not found")
)

func (repo *UserRepository) Create(ctx context.Context, user *domain.User) error {
	if err := repo.db.WithContext(ctx).Create(&user).Error; err != nil {
		return fmt.Errorf("create user error: %v", err)
	}

	return nil
}

func (repo *UserRepository) Get(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user := domain.User{}

	if err := repo.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("get user error: %v", err)
	}

	return &user, nil
}

func (repo *UserRepository) Update(ctx context.Context, user domain.User) error {
	if err := repo.db.WithContext(ctx).Save(&user).Error; err != nil {
		return fmt.Errorf("update user error: %v", err)
	}
	return nil
}

func (repo *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := repo.db.WithContext(ctx).Delete(&domain.User{}, id).Error; err != nil {
		return fmt.Errorf("delete user error: %v", err)
	}
	return nil
}

func (repo *UserRepository) GetAll(ctx context.Context, page, pageSize int) ([]domain.User, error) {
	var users []domain.User

	offset := (page - 1) * pageSize

	if err := repo.db.WithContext(ctx).Limit(pageSize).Offset(offset).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("get all users error: %v", err)
	}

	return users, nil
}
