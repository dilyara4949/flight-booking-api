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

func (service *Flight) Get(ctx context.Context, id uuid.UUID) (*domain.Flight, error) {
	return service.repo.Get(ctx, id)
}

func (service *Flight) Create(ctx context.Context, flight domain.Flight) (domain.Flight, error) {
	return flight, service.repo.Create(ctx, flight)
}

func (service *Flight) Update(ctx context.Context, flight domain.Flight) error {
	return service.repo.Update(ctx, flight)
}

func (service *Flight) Delete(ctx context.Context, id uuid.UUID) error {
	return service.repo.Delete(ctx, id)
}

func (service *Flight) GetAll(ctx context.Context, page, pageSize int) ([]domain.Flight, error) {
	return service.repo.GetAll(ctx, page, pageSize)
}
