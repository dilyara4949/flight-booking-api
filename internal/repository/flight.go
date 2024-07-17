package repository

import (
	"context"
	"fmt"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
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
