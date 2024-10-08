package service

import (
	"context"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	errs "github.com/dilyara4949/flight-booking-api/internal/repository/errors"
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

func (service *Flight) Create(ctx context.Context, req request.Flight) (domain.Flight, error) {
	flight := domain.Flight{
		ID:           uuid.New(),
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
		Departure:    req.Departure,
		Destination:  req.Destination,
		Rank:         req.Rank,
		Price:        *req.Price,
		TotalTickets: *req.TotalTickets,
	}

	return service.repo.Create(ctx, flight)
}

func (service *Flight) Get(ctx context.Context, id uuid.UUID, available bool) (*domain.Flight, error) {
	return service.repo.Get(ctx, id, available)
}

func (service *Flight) Delete(ctx context.Context, id uuid.UUID) error {
	return service.repo.Delete(ctx, id)
}

func (service *Flight) Update(ctx context.Context, req request.Flight, id uuid.UUID) (domain.Flight, error) {
	flight, err := service.Get(ctx, id, false)
	if err != nil {
		return domain.Flight{}, errs.ErrFlightNotFound
	}

	if req.Departure != "" {
		flight.Departure = req.Departure
	}

	if req.Destination != "" {
		flight.Destination = req.Destination
	}

	if !req.StartDate.IsZero() {
		flight.StartDate = req.StartDate
	}

	if !req.EndDate.IsZero() {
		flight.EndDate = req.EndDate
	}

	if req.TotalTickets != nil {
		flight.TotalTickets = *req.TotalTickets
	}

	if req.Rank != "" {
		flight.Rank = req.Rank
	}

	if req.Price != nil {
		flight.Price = *req.Price
	}

	return service.repo.Update(ctx, *flight)
}
