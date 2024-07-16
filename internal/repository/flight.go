package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	errs "github.com/dilyara4949/flight-booking-api/internal/repository/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FlightRepository struct {
	db *gorm.DB
}

func NewFlightRepository(db *gorm.DB) FlightRepository {
	return FlightRepository{db: db}
}

func (repo *FlightRepository) Create(ctx context.Context, flight domain.Flight) (domain.Flight, error) {
	if err := repo.db.WithContext(ctx).Create(&flight).Error; err != nil {
		return domain.Flight{}, fmt.Errorf("create flight error: %w", err)
	}
	return flight, nil
}

func (repo *FlightRepository) Get(ctx context.Context, id uuid.UUID) (*domain.Flight, error) {
	flight := domain.Flight{}

	if err := repo.db.WithContext(ctx).First(&flight, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUserNotFound
		}
		return nil, fmt.Errorf("get flight error: %w", err)
	}
	return &flight, nil
}

func (repo *FlightRepository) Update(ctx context.Context, flight domain.Flight) error {
	if err := repo.db.WithContext(ctx).Save(&flight).Error; err != nil {
		return fmt.Errorf("update flight error: %w", err)
	}
	return nil
}

func (repo *FlightRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := repo.db.WithContext(ctx).Delete(&domain.Flight{}, id).Error; err != nil {
		return fmt.Errorf("delete flight error: %w", err)
	}
	return nil
}

func (repo *FlightRepository) GetAll(ctx context.Context, page, pageSize int) ([]domain.Flight, error) {
	var flights []domain.Flight

	offset := (page - 1) * pageSize

	if err := repo.db.WithContext(ctx).Limit(pageSize).Offset(offset).Find(&flights).Error; err != nil {
		return nil, fmt.Errorf("get all flights error: %w", err)
	}
	return flights, nil
}
