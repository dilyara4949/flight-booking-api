package handler

import (
	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/middleware"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	"github.com/dilyara4949/flight-booking-api/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAPI(cfg config.Config, database *gorm.DB) *gin.Engine {
	userRepo := repository.NewUserRepository(database)
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)

	flightRepo := repository.NewFlightRepository(database)
	flightService := service.NewFlightService(flightRepo)

	ticketRepo := repository.NewTicketRepository(database)
	ticketService := service.NewTicketService(ticketRepo)

	router := gin.Default()

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
			users := v1.Group("/users")
			{
				admin := users.Use(middleware.JWTAuth(cfg.JWTTokenSecret), middleware.AccessCheck(AdminRole))
				{
					admin.GET("/", GetUsersHandler(userService))
				}
				private := users.Use(middleware.JWTAuth(cfg.JWTTokenSecret))
				{
					users.PUT("/:userId", UpdateUserHandler(userService))
					private.DELETE("/:userId", DeleteUserHandler(userService))
					private.GET("/:userId", GetUserHandler(userService))
				}
			}

			flights := v1.Group("/flights")
			{
				private := flights.Use(middleware.JWTAuth(cfg.JWTTokenSecret))
				{
					private.GET("/", GetFlights(flightService))
					private.GET("/:flightId", GetFlightHandler(flightService))
				}
				admin := flights.Use(middleware.JWTAuth(cfg.JWTTokenSecret), middleware.AccessCheck(AdminRole))
				{
					admin.POST("/", CreateFlightHandler(flightService))
					admin.PUT("/:flightId", UpdateFlightHandler(flightService))
					admin.DELETE("/:flightId", DeleteFlightHandler(flightService))
				}
			}

			tickets := v1.Group("/users/:userId/tickets")
			{
				private := tickets.Use(middleware.JWTAuth(cfg.JWTTokenSecret))
				{
					private.GET("/", GetTickets(ticketService))
					private.GET(":ticketId", GetTicketHandler(ticketService))
					private.PUT("/:ticketId", UpdateTicketHandler(ticketService))
					private.DELETE("/:ticketId", DeleteTicketHandler(ticketService))
				}
				user := tickets.Use(middleware.JWTAuth(cfg.JWTTokenSecret), middleware.AccessCheck("user"))
				{
					user.POST("/", BookTicketHandler(ticketService, flightService))
				}
			}
		}
	}

	return router
}
