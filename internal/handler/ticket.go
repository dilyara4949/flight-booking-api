package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/dilyara4949/flight-booking-api/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TicketService interface {
	Update(ctx context.Context, ticketID, userID uuid.UUID, req request.UpdateTicket) (domain.Ticket, error)
}

func UpdateTicketHandler(service TicketService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.UpdateTicket

		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "error at binding request body"})
			return
		}

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

		ticket, err := service.Update(c, ticketID, userID, req)
		if err != nil {
			c.JSON(http.StatusNotFound, response.Error{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, ticket)
	}
}
