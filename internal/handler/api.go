package handler

import (
	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/middleware"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	"github.com/dilyara4949/flight-booking-api/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAPI(cfg config.Config, database *gorm.DB, authService AuthService, userService UserService, flightService FlightService) *gin.Engine {
	router := gin.Default()

	ticketRepo := repository.NewTicketRepository(database)
	ticketService := service.NewTicketService(ticketRepo)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.POST("/signup", SignupHandler(authService, userService, cfg))
				auth.POST("/signin", SigninHandler(authService, userService, cfg))
				auth.POST("/reset-password", ResetPasswordHandler(userService))
			}
			users := v1.Group("/users").Use(middleware.JWTAuth(cfg.JWTTokenSecret))
			{
				users.DELETE("/:userId", DeleteUserHandler(userService))
			}

			flights := v1.Group("/flights")
			{
				admin := flights.Use(middleware.JWTAuth(cfg.JWTTokenSecret), middleware.AccessCheck("admin"))
				{
					admin.DELETE("/:flightId", DeleteFlightHandler(flightService))
				}
			}

			tickets := v1.Group("/users/tickets")
			{
				private := tickets.Use(middleware.JWTAuth(cfg.JWTTokenSecret))
				{
					private.POST("/", BookTicketHandler(ticketService, flightService))
				}
			}
		}
	}

	return router
}
