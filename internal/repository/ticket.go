package repository

import (
	"context"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
	errs "github.com/dilyara4949/flight-booking-api/internal/repository/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return TicketRepository{db: db}
}

func (repo *TicketRepository) GetTickets(ctx context.Context, userID uuid.UUID, page, pageSize int) ([]domain.Ticket, error) {
	tickets := make([]domain.Ticket, 0)

	offset := (page - 1) * pageSize

	if err := repo.db.WithContext(ctx).Limit(pageSize).Offset(offset).Find(tickets, "user_id = ?", userID).Error; err != nil {
		return nil, errs.ErrTicketNotFound
	}
	return tickets, nil
}
