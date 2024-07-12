package domain

import "github.com/google/uuid"

type Role struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name string    `gorm:"unique;not null"                      json:"name"`
}

func (Role) TableName() string {
	return "roles"
}
