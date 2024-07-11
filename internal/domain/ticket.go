package domain

import (
	"github.com/google/uuid"
	"time"
)

type Ticket struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	FlightID  uuid.UUID `json:"flight_id" gorm:"type:uuid"`
	Flight    Flight    `json:"flight" gorm:"foreignKey:FlightID;references:ID"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid"`
	User      User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
	RankID    uuid.UUID `json:"rank_id" gorm:"type:uuid"`
	Rank      Rank      `json:"rank" gorm:"foreignKey:RankID;references:ID"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Ticket) TableName() string {
	return "tickets"
}
