package service

import (
	"context"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	errs "github.com/dilyara4949/flight-booking-api/internal/repository/errors"
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

func (service *Ticket) Update(ctx context.Context, ticketID, userID uuid.UUID, req request.UpdateTicket) (domain.Ticket, error) {
	ticket, err := service.Get(ctx, ticketID, userID)
	if err != nil {
		return domain.Ticket{}, errs.ErrTicketNotFound
	}

	ticket.Price = req.Price

	return service.repo.Update(ctx, ticket)
}
