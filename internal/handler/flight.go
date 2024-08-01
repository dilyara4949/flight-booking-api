package handler

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FlightService interface {
	GetFlights(ctx context.Context, page, pageSize int, available bool) ([]domain.Flight, error)
	Get(ctx context.Context, id uuid.UUID, available bool) (*domain.Flight, error)
	Create(ctx context.Context, flight request.CreateFlight) (domain.Flight, error)
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
