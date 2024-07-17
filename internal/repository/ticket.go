package repository

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (repo *TicketRepository) BookTicket(ctx context.Context, ticket domain.Ticket) (domain.Ticket, error) {
	if err := repo.db.WithContext(ctx).Create(&ticket).Error; err != nil {
		return domain.Ticket{}, err
	}
	return ticket, nil
}
