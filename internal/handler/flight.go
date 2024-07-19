package handler

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/google/uuid"
)

type FlightService interface {
	Get(ctx context.Context, id uuid.UUID, available bool) (*domain.Flight, error)
}
