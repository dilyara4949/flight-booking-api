package handler

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

const availableDefault = false

type FlightService interface {
	Get(ctx context.Context, id uuid.UUID, available bool) (*domain.Flight, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func GetFlightHandler(service FlightService) gin.HandlerFunc {
	return func(c *gin.Context) {
		flightID, err := uuid.Parse(c.Param("flightId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "id format is not correct"})

			return
		}

		available, err := strconv.ParseBool(c.Query("available"))
		if err != nil {
			available = availableDefault
		}

		flight, err := service.Get(c, flightID, available)
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