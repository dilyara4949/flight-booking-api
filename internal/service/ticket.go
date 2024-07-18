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

func (service *Ticket) Update(ctx context.Context, ticketID, userID uuid.UUID, req request.UpdateTicket) (domain.Ticket, error) {
	// check if user has the ticket, if not return not found
	ticket := domain.Ticket{
		ID:     ticketID,
		UserID: userID,
		Price:  req.Price,
	}
	return service.repo.Update(ctx, ticket)
}
