package service

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	"github.com/google/uuid"
)

type Ticket struct {
	repo repository.TicketRepository
}

func NewTicketService(repo repository.TicketRepository) *Ticket {
	return &Ticket{repo: repo}
}

func (service *Ticket) BookTicket(ctx context.Context, req request.BookTicket, userID uuid.UUID) (domain.Ticket, error) {
	ticket := domain.Ticket{
		ID:       uuid.New(),
		FlightID: req.FlightID,
		UserID:   userID,
		//////
	}

	ticket, err := service.repo.BookTicket(ctx, ticket)
	if err != nil {
		return domain.Ticket{}, err
	}

	return ticket, nil
}
