package handler

import (
	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewAPI(cfg config.Config, authService AuthService, userService UserService, flightService FlightService) *gin.Engine {
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

			flights := v1.Group("/flights")
			{
				private := flights.Use(middleware.JWTAuth(cfg.JWTTokenSecret))
				{
					private.GET("/:flightId", GetFlightHandler(flightService))
				}
			}
		}
	}

	return router
}
