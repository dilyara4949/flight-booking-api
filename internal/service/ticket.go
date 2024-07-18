package service

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	"github.com/google/uuid"
)

type Ticket struct {
	repo repository.TicketRepository
}

func NewTicketService(repo repository.TicketRepository) *Ticket {
	return &Ticket{repo: repo}
}

func (service *Ticket) Get(ctx context.Context, ticketID, userID uuid.UUID) (domain.Ticket, error) {
	ticket := domain.Ticket{
		ID:     ticketID,
		UserID: userID,
	}

	ticket, err := service.repo.Get(ctx, ticket)
	if err != nil {
		return domain.Ticket{}, err
	}
	return ticket, nil
}
