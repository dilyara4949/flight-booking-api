package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TicketService interface {
	BookTicket(ctx context.Context, req request.BookTicket, userID uuid.UUID, flight domain.Flight) (domain.Ticket, error)
	CheckAvailability(ctx context.Context, flightID uuid.UUID, totalTickets int) (bool, error)
	Get(ctx context.Context, ticketID, userID uuid.UUID) (domain.Ticket, error)
	Delete(ctx context.Context, ticketID, userID uuid.UUID) error
	GetTickets(ctx context.Context, userID uuid.UUID, page, pageSize int) ([]domain.Ticket, error)
	Update(ctx context.Context, ticketID, userID uuid.UUID, req request.UpdateTicket) (domain.Ticket, error)
}

func BookTicketHandler(ticketService TicketService, flightService FlightService) gin.HandlerFunc {
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

		userID, err := uuid.Parse(c.GetString("user_id"))
		if err != nil {
			slog.Error("user id format is not correct at jwt", "error", err.Error())
			c.JSON(http.StatusBadRequest, response.Error{Error: "user token set incorrectly"})
			return
		}

		flight, err := flightService.Get(c, req.FlightID, true)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.Error{Error: err.Error()})
			return
		}

		available, err := ticketService.CheckAvailability(c, req.FlightID, flight.TotalTickets)
		if err != nil {
			slog.Error(err.Error(), "error")
			c.JSON(http.StatusInternalServerError, response.Error{Error: "ticket is not available"})
			return
		}

		if !available {
			c.JSON(http.StatusInternalServerError, response.Error{Error: "tickets aut of stock"})
			return
		}

		ticket, err := ticketService.BookTicket(c, req, userID, *flight)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.Error{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, ticket)
	}
}


func UpdateTicketHandler(service TicketService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !AccessCheck(*c, c.GetString(middleware.UserIDKey), userIDParamKey) {
			c.JSON(http.StatusForbidden, response.Error{Error: "access denied"})
			return
		}

		var req request.UpdateTicket

		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "error at binding request body"})
			return
		}

		ticketID, err := uuid.Parse(c.Param("ticketId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "ticket id format is not correct"})
			return
		}

		userID, err := uuid.Parse(c.Param(userIDParamKey))
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "user id is not correct"})
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

func GetTicketHandler(service TicketService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !AccessCheck(*c, c.GetString(middleware.UserIDKey), userIDParamKey) {
			c.JSON(http.StatusForbidden, response.Error{Error: "access denied"})
			return
		}

		ticketID, err := uuid.Parse(c.Param("ticketId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "ticket id format is not correct"})
			return
		}

		userID, err := uuid.Parse(c.Param(userIDParamKey))
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "user id is not correct"})
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

func DeleteTicketHandler(service TicketService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !AccessCheck(*c, c.GetString(middleware.UserIDKey), userIDParamKey) {
			c.JSON(http.StatusForbidden, response.Error{Error: "access denied"})
			return
		}

		ticketID, err := uuid.Parse(c.Param("ticketId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "ticket id format is not correct"})
			return
		}

		userID, err := uuid.Parse(c.Param(userIDParamKey))
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "user id is not correct"})
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
