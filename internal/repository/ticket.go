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

func (repo *TicketRepository) GetAll(ctx context.Context, userID uuid.UUID) ([]domain.Ticket, error) {
	tickets := make([]domain.Ticket, 0)

	if err := repo.db.WithContext(ctx).Find(tickets, "user_id = ?", userID).Error; err != nil {
		return nil, errs.ErrTicketNotFound
	}
	return tickets, nil
}
