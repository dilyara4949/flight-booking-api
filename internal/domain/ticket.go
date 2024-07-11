package domain

import (
	"github.com/google/uuid"
	"time"
)

type Ticket struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	FlightID  uuid.UUID `json:"flight_id" gorm:"type:uuid"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid"`
	Rank      string    `json:"rank"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Ticket) TableName() string {
	return "tickets"
}
