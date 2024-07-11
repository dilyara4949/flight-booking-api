package domain

import "github.com/google/uuid"

type Role struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name" gorm:"unique;not null"`
}

func (Role) TableName() string {
	return "roles"
}
