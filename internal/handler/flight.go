package handler

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type FlightService interface {
	GetFlights(ctx context.Context, page, pageSize int, available bool) ([]domain.Flight, error)
}

const (
	pageDefault      = 1
	pageSizeDefault  = 30
	availableDefault = false
)

func GetFlights(service FlightService) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page <= 0 {
			page = pageDefault
		}

		pageSize, err := strconv.Atoi(c.Query("page_size"))
		if err != nil || pageSize <= 0 {
			pageSize = pageSizeDefault
		}

		available, err := strconv.ParseBool(c.Query("available"))
		if err != nil {
			available = availableDefault
		}

		flights, err := service.GetFlights(c, page, pageSize, available)
		if err != nil {
			slog.Error("error at getting flights", "error", err.Error())
			c.JSON(http.StatusInternalServerError, response.Error{Error: "error at getting flights"})

			return
		}
		c.JSON(http.StatusOK, flights)
	}
}
