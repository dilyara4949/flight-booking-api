package handler

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

type TicketService interface {
	Get(ctx context.Context, ticketID, userID uuid.UUID) (domain.Ticket, error)
}

func GetTicketHandler(service TicketService) gin.HandlerFunc {
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

		ticket, err := service.Get(c, ticketID, userID)
		if err != nil {
			c.JSON(http.StatusNotFound, response.Error{Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, ticket)
	}
}
