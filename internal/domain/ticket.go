package domain

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	FlightID  uuid.UUID `gorm:"type:uuid"                            json:"flight_id"`
	UserID    uuid.UUID `gorm:"type:uuid"                            json:"user_id"`
	Rank      string    `json:"rank"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `gorm:"autoCreateTime"                       json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"                       json:"updated_at"`
}

func (Ticket) TableName() string {
	return "tickets"
}
