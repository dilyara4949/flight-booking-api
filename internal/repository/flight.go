package repository

import (
	"context"
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


func (repo *FlightRepository) Delete(ctx context.Context, id uuid.UUID) error {
	res := repo.db.WithContext(ctx).Delete(&domain.Flight{}, id)
	if res.Error != nil {
		return fmt.Errorf("delete flight error: %w", res.Error)
	}

	if res.RowsAffected == 0 {
		return errs.ErrFlightNotFound
	}
	return nil
}
