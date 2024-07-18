package handler

import (
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"log/slog"
	"net/http"
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

		userIDStr := c.GetString("user_id")
		if userIDStr == "" {
			c.JSON(http.StatusInternalServerError, response.Error{Error: "user token set incorrectly"})
			return
		}

		userID, err := uuid.Parse(userIDStr)
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
