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

func (repo *TicketRepository) BookTicket(ctx context.Context, ticket domain.Ticket) (domain.Ticket, error) {
	if err := repo.db.WithContext(ctx).Create(&ticket).Error; err != nil {
		return domain.Ticket{}, err
	}
	return ticket, nil
}

func (repo *TicketRepository) CheckAvailability(ctx context.Context, flightID uuid.UUID, totalTickets int) (bool, error) {
	var count int64
	if err := repo.db.WithContext(ctx).
		Model(&domain.Ticket{}).
		Where("flight_id = ?", flightID).
		Count(&count).Error; err != nil {
		return false, err
	}

	if int64(totalTickets)-count > 0 {
		return true, nil
	}
	return false, nil
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

func (repo *TicketRepository) GetTickets(ctx context.Context, userID uuid.UUID, page, pageSize int) ([]domain.Ticket, error) {
	tickets := make([]domain.Ticket, 0)

	offset := (page - 1) * pageSize

	if err := repo.db.WithContext(ctx).Limit(pageSize).Offset(offset).Find(tickets, "user_id = ?", userID).Error; err != nil {
		return nil, errs.ErrTicketNotFound
	}
	return tickets, nil
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
