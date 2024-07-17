package handler

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type FlightService interface {
	Delete(ctx context.Context, id uuid.UUID) error
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
