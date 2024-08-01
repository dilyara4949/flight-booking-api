package request

import (
	"errors"
	"time"
)

type CreateFlight struct {
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	Departure    string    `json:"departure"`
	Destination  string    `json:"destination"`
	Rank         string    `json:"rank"`
	Price        int64     `json:"price"`
	TotalTickets int       `json:"total_tickets"`
}

func (f CreateFlight) Validate() error {
	var err error
	if f.StartDate.IsZero() {
		err = errors.Join(err, errors.New("start_date cannot be empty"))
	}

	if f.EndDate.IsZero() {
		err = errors.Join(err, errors.New("end_date cannot be empty"))
	}
	if f.Departure == "" {
		err = errors.Join(err, errors.New("departure cannot be empty"))
	}

	if f.Destination == "" {
		err = errors.Join(err, errors.New("destination cannot be empty"))
	}

	if f.Rank == "" {
		err = errors.Join(err, errors.New("rank cannot be empty"))
	}

	if f.TotalTickets == 0 {
		err = errors.Join(err, errors.New("total_tickets cannot be empty"))
	}

	if f.Price == 0 {
		err = errors.Join(err, errors.New("price cannot be empty"))
	}
	return err
}

type UpdateFlight struct {
	StartDate    time.Time `gorm:"not null"                             json:"start_date"`
	EndDate      time.Time `gorm:"not null"                             json:"end_date"`
	Departure    string    `gorm:"not null"                             json:"departure"`
	Destination  string    `gorm:"not null"                             json:"destination"`
	Rank         string    `gorm:"type:uuid"                            json:"rank"`
	Price        int64     `gorm:"not null"                             json:"price"`
	TotalTickets int       `gorm:"not null"                             json:"total_tickets"`
}
