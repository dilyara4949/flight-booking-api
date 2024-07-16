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
	GetAll(ctx context.Context, page, pageSize int) ([]domain.Flight, error)
}

const (
	pageDefault     = 1
	pageSizeDefault = 30
)

func GetAllFlights(service FlightService) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = pageDefault
		}

		pageSize, err := strconv.Atoi(c.Query("page_size"))
		if err != nil {
			pageSize = pageSizeDefault
		}

		flights, err := service.GetAll(c, page, pageSize)
		if err != nil {
			slog.Error("error at getting flights", "error", err.Error())
			c.JSON(http.StatusInternalServerError, response.Error{Error: "error at getting flights"})

			return
		}
		c.JSON(http.StatusOK, flights)
	}
}
