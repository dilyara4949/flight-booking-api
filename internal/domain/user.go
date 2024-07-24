package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Email     string    `gorm:"unique;not null"                      json:"email"`
	Phone     string    `gorm:"phone"                                json:"phone"`
	Password  string    `gorm:"not null"                             json:"password"`
	Role      string    `gorm:"role"                                 json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime"                       json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"                       json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
