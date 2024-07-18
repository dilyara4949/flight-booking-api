package service

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	"github.com/google/uuid"
)

type Ticket struct {
	repo repository.TicketRepository
}

func NewTicketService(repo repository.TicketRepository) *Ticket {
	return &Ticket{repo: repo}
}

func (service *Ticket) Delete(ctx context.Context, ticketID, userID uuid.UUID) error {
	return service.repo.Delete(ctx, ticketID, userID)
}
