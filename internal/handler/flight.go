package handler

import (
	"context"
	"fmt"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

type FlightService interface {
	Create(ctx context.Context, flight request.CreateFlight) (domain.Flight, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func CreateFlightHandler(service FlightService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.CreateFlight

		err := c.ShouldBind(&req)
		if err != nil {
			slog.Error("error at binding request body", "error", err.Error())

			c.JSON(http.StatusBadRequest, response.Error{Error: "error at binding request body"})

			return
		}

		if err = req.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, response.Error{
				Error: fmt.Sprintf("request fields cannot be empty: %v", err.Error()),
			})
			return
		}

		flight, err := service.Create(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.Error{Error: err.Error()})

			return
		}
		c.JSON(http.StatusOK, flight)
	}
}

func DeleteFlightHandler(service FlightService) gin.HandlerFunc {
	return func(c *gin.Context) {
		flightID, err := uuid.Parse(c.Param("flightId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "id format is not correct"})

			return
		}

		err = service.Delete(c, flightID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.Error{Error: err.Error()})

			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}

func ValidateFlightRequest(req request.CreateFlight) bool {
	if req.StartDate.IsZero() || req.EndDate.IsZero() ||
		req.Departure == "" || req.Destination == "" ||
		req.Rank == "" || req.TotalTickets == 0 || req.Price == 0 {
		return false
	}
	return true
}
