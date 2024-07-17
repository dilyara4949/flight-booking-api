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

func (service *Flight) Update(ctx context.Context, req request.UpdateFlight, id uuid.UUID) (domain.Flight, error) {
	flight := domain.Flight{
		ID:           id,
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
		Departure:    req.Departure,
		Destination:  req.Destination,
		Rank:         req.Rank,
		Price:        req.Price,
		TotalTickets: req.TotalTickets,
	}
	return service.repo.Update(ctx, flight)
}
