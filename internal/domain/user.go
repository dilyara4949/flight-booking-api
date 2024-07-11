package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password" gorm:"not null"`
	RoleID    uuid.UUID `json:"role_id" gorm:"type:uuid"`
	Role      Role      `json:"role" gorm:"foreignKey:RoleID;references:ID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}
