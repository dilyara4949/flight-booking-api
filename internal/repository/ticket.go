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

func (repo *TicketRepository) Delete(ctx context.Context, ticketID, userID uuid.UUID) error {
	res := repo.db.WithContext(ctx).Delete(domain.Ticket{ID: ticketID, UserID: userID})
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errs.ErrTicketNotFound
	}
	return nil
}
