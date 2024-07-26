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

func (service *Flight) Get(ctx context.Context, id uuid.UUID, available bool) (*domain.Flight, error) {
	return service.repo.Get(ctx, id, available)
}

func (service *Flight) Delete(ctx context.Context, id uuid.UUID) error {
	return service.repo.Delete(ctx, id)
}
