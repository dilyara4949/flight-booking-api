package domain

import (
	"time"

	"github.com/google/uuid"
)

type Flight struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	StartDate    time.Time `gorm:"not null"                             json:"start_date"`
	EndDate      time.Time `gorm:"not null"                             json:"end_date"`
	Departure    string    `gorm:"not null"                             json:"departure"`
	Destination  string    `gorm:"not null"                             json:"destination"`
	Rank         string    `gorm:"type:uuid"                            json:"rank"`
	Price        int64     `gorm:"not null"                             json:"price"`
	TotalTickets int       `gorm:"not null"                             json:"total_tickets"`
	CreatedAt    time.Time `gorm:"autoCreateTime"                       json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"                       json:"updated_at"`
}

func (Flight) TableName() string {
	return "flights"
}
