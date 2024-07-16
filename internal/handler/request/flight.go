package request

import "time"

type CreateFlight struct {
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	Departure    string    `json:"departure"`
	Destination  string    `json:"destination"`
	Rank         string    `json:"rank"`
	Price        int64     `json:"price"`
	TotalTickets int       `json:"total_tickets"`
}
