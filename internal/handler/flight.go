package handler

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"strconv"
)

type FlightService interface {
	GetFlights(ctx context.Context, page, pageSize int, available bool) ([]domain.Flight, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

const (
	availableDefault = false
)

func GetFlights(service FlightService) gin.HandlerFunc {
	return func(c *gin.Context) {
		available, err := strconv.ParseBool(c.Query("available"))
		if err != nil {
			available = availableDefault
		}

		page, pageSize := GetPageInfo(c)

		flights, err := service.GetFlights(c, page, pageSize, available)
		if err != nil {
			slog.Error("error at getting flights", "error", err.Error())
			c.JSON(http.StatusInternalServerError, response.Error{Error: "error at getting flights"})

			return
		}
		c.JSON(http.StatusOK, flights)
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
