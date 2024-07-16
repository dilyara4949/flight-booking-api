package handler

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type FlightService interface {
	Update(ctx context.Context, flight request.UpdateFlight, id uuid.UUID) (domain.Flight, error)
}

func UpdateFlightHandler(service FlightService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.UpdateFlight

		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "error at binding request body"})
			return
		}

		if req.StartDate.IsZero() || req.EndDate.IsZero() ||
			req.Departure == "" || req.Destination == "" ||
			req.Rank == "" || req.TotalTickets == 0 || req.Price == 0 {
			c.JSON(http.StatusBadRequest, response.Error{Error: "request fields cannot be empty"})

			return
		}
		flightID, err := uuid.Parse(c.Param("flightId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "id format is not correct"})
			return
		}

		flight, err := service.Update(c, req, flightID)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, flight)
	}
}
