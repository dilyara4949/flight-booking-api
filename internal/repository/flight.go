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

func (repo *FlightRepository) Get(ctx context.Context, id uuid.UUID) (*domain.Flight, error) {
	flight := domain.Flight{}

	if err := repo.db.WithContext(ctx).First(&flight, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrFlightNotFound
		}
		return nil, fmt.Errorf("get flight error: %w", err)
	}
	return &flight, nil
}
