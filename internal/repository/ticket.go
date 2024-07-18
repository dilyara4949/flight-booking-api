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

func (repo *TicketRepository) Get(ctx context.Context, id, userID uuid.UUID) (domain.Ticket, error) {
	ticket := domain.Ticket{
		ID:     id,
		UserID: userID,
	}

	if err := repo.db.WithContext(ctx).First(&ticket, "id = ?", id).Error; err != nil {
		return domain.Ticket{}, errs.ErrTickerNotFound
	}
	return ticket, nil
}
