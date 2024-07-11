package domain

import "github.com/google/uuid"

type Rank struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name" gorm:"unique;not null"`
}

func (Rank) TableName() string {
	return "ranks"
}
