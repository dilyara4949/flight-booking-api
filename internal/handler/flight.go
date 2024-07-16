package handler

import (
	"context"
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
	Get(ctx context.Context, id uuid.UUID) (*domain.Flight, error)
	Update(ctx context.Context, flight domain.Flight) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context, page, pageSize int) ([]domain.Flight, error)
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

		if req.StartDate.IsZero() || req.EndDate.IsZero() ||
			req.Departure == "" || req.Destination == "" ||
			req.Rank == "" || req.TotalTickets == 0 || req.Price == 0 {
			c.JSON(http.StatusBadRequest, response.Error{Error: "request fields cannot be empty"})

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
