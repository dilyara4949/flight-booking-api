package service

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	"github.com/google/uuid"
)

type Flight struct {
	repo repository.FlightRepository
}

func NewFlightService(repo repository.FlightRepository) *Flight {
	return &Flight{repo: repo}
}

func (service *Flight) Create(ctx context.Context, req request.CreateFlight) (domain.Flight, error) {
	flight := domain.Flight{
		ID:           uuid.New(),
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
		Departure:    req.Departure,
		Destination:  req.Destination,
		Rank:         req.Rank,
		Price:        req.Price,
		TotalTickets: req.TotalTickets,
	}

	return service.repo.Create(ctx, flight)

}

func (service *Flight) Get(ctx context.Context, id uuid.UUID, available bool) (*domain.Flight, error) {
	return service.repo.Get(ctx, id, available)
}

func (service *Flight) Delete(ctx context.Context, id uuid.UUID) error {
	return service.repo.Delete(ctx, id)
}
