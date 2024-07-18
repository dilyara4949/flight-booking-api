package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	errs "github.com/dilyara4949/flight-booking-api/internal/repository/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (repo *UserRepository) Create(ctx context.Context, user *domain.User) error {
	if err := repo.db.WithContext(ctx).Create(&user).Error; err != nil {
		return fmt.Errorf("create user error: %w", err)
	}

	return nil
}

func (repo *UserRepository) Get(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var user domain.User

	if err := repo.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUserNotFound
		}

		return nil, fmt.Errorf("get user error: %w", err)
	}

	return &user, nil
}

func (repo *UserRepository) Update(ctx context.Context, user domain.User) (domain.User, error) {
	if err := repo.db.WithContext(ctx).Model(&user).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"email":      user.Email,
			"phone":      user.Phone,
			"updated_at": time.Now(),
		}).Error; err != nil {
		return domain.User{}, fmt.Errorf("update user error: %w", err)
	}

	return user, nil
}

func (repo *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := repo.db.WithContext(ctx).Delete(&domain.User{}, id).Error; err != nil {
		return fmt.Errorf("delete user error: %w", err)
	}
	return nil
}

func (repo *UserRepository) GetAll(ctx context.Context, page, pageSize int) ([]domain.User, error) {
	var users []domain.User

	offset := (page - 1) * pageSize

	if err := repo.db.WithContext(ctx).Limit(pageSize).Offset(offset).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("get all users error: %w", err)
	}

	return users, nil
}
