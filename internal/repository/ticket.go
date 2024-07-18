package repository

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	errs "github.com/dilyara4949/flight-booking-api/internal/repository/errors"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return TicketRepository{db: db}
}

func (repo *TicketRepository) Get(ctx context.Context, ticket domain.Ticket) (domain.Ticket, error) {
	if err := repo.db.WithContext(ctx).First(&ticket).Error; err != nil {
		return domain.Ticket{}, errs.ErrTickerNotFound
	}
	return ticket, nil
}
