package service

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	"github.com/google/uuid"
)

type Flight struct {
	repo repository.FlightRepository
}

func NewFlightService(repo repository.FlightRepository) *Flight {
	return &Flight{repo: repo}
}

func (service *Flight) GetFlights(ctx context.Context, page, pageSize int, available bool) ([]domain.Flight, error) {
	return service.repo.GetFlights(ctx, page, pageSize, available)
}

func (service *Flight) Delete(ctx context.Context, id uuid.UUID) error {
	return service.repo.Delete(ctx, id)
}
