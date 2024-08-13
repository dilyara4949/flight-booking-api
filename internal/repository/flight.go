package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

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

func (repo *FlightRepository) Update(ctx context.Context, flight domain.Flight) (domain.Flight, error) {
	if err := repo.db.WithContext(ctx).Save(&flight).Error; err != nil {
		return domain.Flight{}, fmt.Errorf("update flight error: %w", err)
	}
	return flight, nil
}

func (repo *FlightRepository) GetFlights(ctx context.Context, page, pageSize int, available bool) ([]domain.Flight, error) {
	flights := make([]domain.Flight, 0)

	offset := (page - 1) * pageSize

	query := repo.db.WithContext(ctx).Limit(pageSize).Offset(offset)

	if available {
		addDateFilter(query)
	}

	if err := query.Find(&flights).Error; err != nil {
		return nil, fmt.Errorf("get all flights error: %w", err)
	}
	return flights, nil
}

func (repo *FlightRepository) Create(ctx context.Context, flight domain.Flight) (domain.Flight, error) {
	if err := repo.db.WithContext(ctx).Create(&flight).Error; err != nil {
		return domain.Flight{}, fmt.Errorf("create flight error: %w", err)
	}
	return flight, nil
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
