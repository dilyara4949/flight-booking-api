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

type TicketService interface {
	BookTicket(ctx context.Context, req request.BookTicket, userID uuid.UUID) (domain.Ticket, error)
}

func BookTicketHandler(service TicketService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.BookTicket

		if err := c.ShouldBind(&req); err != nil {
			slog.Error("error at binding request body", "error", err.Error())

			c.JSON(http.StatusBadRequest, response.Error{Error: "error at binding request body"})

			return
		}

		if req.FlightID == uuid.Nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: response.ErrEmptyRequestFields.Error()})
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

		ticket, err := service.BookTicket(c, req, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.Error{Error: err.Error()})
		}

		c.JSON(http.StatusOK, ticket)
	}
}
