package handler

import (
	"log/slog"
	"net/http"

	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/dilyara4949/flight-booking-api/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type TicketService interface {
	Delete(ctx context.Context, ticketID, userID uuid.UUID) error
}

func DeleteTicketHandler(service TicketService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ticketID, err := uuid.Parse(c.Param("ticketId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "id format is not correct"})
			return
		}

		userID, err := uuid.Parse(c.GetString(middleware.UserIDKey))
		if err != nil {
			slog.Error("user id format is not correct at jwt", "error", err.Error())
			c.JSON(http.StatusBadRequest, response.Error{Error: "user token set incorrectly"})
			return
		}

		err = service.Delete(c, ticketID, userID)
		if err != nil {
			c.JSON(http.StatusNotFound, response.Error{Error: err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}
