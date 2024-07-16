package request

import (
	"time"
)

type UpdateFlight struct {
	StartDate    time.Time `gorm:"not null"                             json:"start_date"`
	EndDate      time.Time `gorm:"not null"                             json:"end_date"`
	Departure    string    `gorm:"not null"                             json:"departure"`
	Destination  string    `gorm:"not null"                             json:"destination"`
	Rank         string    `gorm:"type:uuid"                            json:"rank"`
	Price        int64     `gorm:"not null"                             json:"price"`
	TotalTickets int       `gorm:"not null"                             json:"total_tickets"`
}
