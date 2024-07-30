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

func (service *Ticket) GetAll(ctx context.Context, userID uuid.UUID, page, pageSize int) ([]domain.Ticket, error) {
	return service.repo.GetAll(ctx, userID, page, pageSize)
}
