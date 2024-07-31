package repository

import (
	"context"
	"fmt"
	"time"

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

func (repo *TicketRepository) Get(ctx context.Context, ticket domain.Ticket) (domain.Ticket, error) {
	if err := repo.db.WithContext(ctx).First(&ticket).Error; err != nil {
		return domain.Ticket{}, errs.ErrTicketNotFound
	}
	return ticket, nil
}

func (repo *TicketRepository) Update(ctx context.Context, ticket domain.Ticket) (domain.Ticket, error) {
	if err := repo.db.WithContext(ctx).Model(&ticket).
		Where("id = ? AND user_id = ?", ticket.ID, ticket.UserID).
		Updates(map[string]interface{}{
			"price":      ticket.Price,
			"updated_at": time.Now(),
		}).Error; err != nil {
		return ticket, fmt.Errorf("update ticket error: %w", err)
	}
	return ticket, nil
}
