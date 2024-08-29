package request

import "github.com/google/uuid"

type BookTicket struct {
	FlightID uuid.UUID `json:"flight_id"`
}

type UpdateTicket struct {
	Price int64 `json:"price"`
}
