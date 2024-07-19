package repository

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/google/uuid"
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

func (repo *TicketRepository) CheckAvailability(ctx context.Context, flightID uuid.UUID) (bool, error) {
	var count int64
	if err := repo.db.WithContext(ctx).
		Model(&domain.Ticket{}).
		Where("flight_id = ?", flightID).
		Count(&count).Error; err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}
