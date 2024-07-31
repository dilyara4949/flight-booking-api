package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/dilyara4949/flight-booking-api/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TicketService interface {
	GetTickets(ctx context.Context, userID uuid.UUID, page, pageSize int) ([]domain.Ticket, error)
}

func GetAllTickets(service TicketService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := uuid.Parse(c.GetString(middleware.UserIDKey))
		if err != nil {
			slog.Error("user id format is not correct at jwt", "error", err.Error())
			c.JSON(http.StatusBadRequest, response.Error{Error: "user token set incorrectly"})
			return
		}

		page, pageSize := GetPageInfo(c)

		tickets, err := service.GetTickets(c, userID, page, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.Error{Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, tickets)
	}
}
