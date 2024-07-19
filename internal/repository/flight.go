package repository

import (
	"context"
	"fmt"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"gorm.io/gorm"
	"time"
)

type FlightRepository struct {
	db *gorm.DB
}

func NewFlightRepository(db *gorm.DB) FlightRepository {
	return FlightRepository{db: db}
}

func (repo *FlightRepository) GetAll(ctx context.Context, page, pageSize int, available bool) ([]domain.Flight, error) {
	flights := make([]domain.Flight, 0)

	offset := (page - 1) * pageSize

	query := repo.db.WithContext(ctx).Limit(pageSize).Offset(offset)

	if available {
		now := time.Now()
		twoHoursLater := now.Add(2 * time.Hour)
		query = query.Where("start_date > ?", twoHoursLater)
	}

	if err := query.Find(&flights).Error; err != nil {
		return nil, fmt.Errorf("get all flights error: %w", err)
	}
	return flights, nil
}
