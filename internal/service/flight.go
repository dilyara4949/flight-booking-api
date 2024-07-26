package service

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
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
