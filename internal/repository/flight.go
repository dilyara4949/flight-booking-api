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

type FlightRepository struct {
	db *gorm.DB
}

func NewFlightRepository(db *gorm.DB) FlightRepository {
	return FlightRepository{db: db}
}

func (repo *FlightRepository) Get(ctx context.Context, id uuid.UUID, available bool) (*domain.Flight, error) {
	flight := domain.Flight{}

	query := repo.db.WithContext(ctx)

	if available {
		addDateFilter(query)
	}

	if err := query.First(&flight, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrFlightNotFound
		}
		return nil, fmt.Errorf("get flight error: %w", err)
	}
	return &flight, nil
}

func addDateFilter(query *gorm.DB) {
	now := time.Now()
	twoHoursLater := now.Add(2 * time.Hour)
	query.Where("start_date > ?", twoHoursLater)
}
