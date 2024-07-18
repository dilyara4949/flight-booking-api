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

func (repo *FlightRepository) GetAll(ctx context.Context, page, pageSize int) ([]domain.Flight, error) {
	flights := make([]domain.Flight, 0)

	offset := (page - 1) * pageSize

	if err := repo.db.WithContext(ctx).Limit(pageSize).Offset(offset).Find(&flights).Error; err != nil {
		return nil, fmt.Errorf("get all flights error: %w", err)
	}
	return flights, nil
}
