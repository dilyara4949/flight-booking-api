package domain

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name" gorm:"unique;not null"`
}

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

type Rank struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name" gorm:"unique;not null"`
}

type Flight struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	Departure    string    `json:"departure"`
	Destination  string    `json:"destination"`
	RankID       uuid.UUID `json:"rank_id" gorm:"type:uuid"`
	Rank         Rank      `json:"rank" gorm:"foreignKey:RankID;references:ID"`
	Price        int64     `json:"price"`
	TotalTickets int       `json:"total_tickets"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

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

func (User) TableName() string {
	return "users"
}

func (Role) TableName() string {
	return "roles"
}

func (Rank) TableName() string {
	return "ranks"
}

func (Flight) TableName() string {
	return "flights"
}

func (Ticket) TableName() string {
	return "tickets"
}
