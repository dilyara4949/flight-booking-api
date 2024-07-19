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

func (service *Flight) GetAll(ctx context.Context, page, pageSize int, hasSeats bool) ([]domain.Flight, error) {
	return service.repo.GetAll(ctx, page, pageSize, hasSeats)
}
