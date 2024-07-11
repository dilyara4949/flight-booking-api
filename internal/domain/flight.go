package domain

import (
	"github.com/google/uuid"
	"time"
)

type Flight struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	Departure    string    `json:"departure"`
	Destination  string    `json:"destination"`
	Rank         string    `json:"rank" gorm:"type:uuid"`
	Price        int64     `json:"price"`
	TotalTickets int       `json:"total_tickets"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Flight) TableName() string {
	return "flights"
}
