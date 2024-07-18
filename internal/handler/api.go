package handler

import (
	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/gin-gonic/gin"
)

func NewAPI(cfg config.Config, authService AuthService, userService UserService, ticketService TicketService) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.POST("/signup", SignupHandler(authService, userService, cfg))
			}

			tickets := v1.Group("/tickets")
			{
				tickets.GET(":ticketId", GetTicketHandler(ticketService))
			}
		}
	}

	return router
}
