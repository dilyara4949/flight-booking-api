package domain

import (
	"github.com/google/uuid"
	"time"
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

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Email     string    `gorm:"unique;not null"                      json:"email"`
	Phone     string    `gorm:"-"                                    json:"phone"`
	Password  string    `gorm:"not null"                             json:"password"`
	Role      string    `gorm:"-"                                    json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime"                       json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"                       json:"updated_at"`
}
