package handler

import (
	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/gin-gonic/gin"
)

func NewAPI(cfg config.Config, authService AuthService) *gin.Engine {

	router := gin.Default()
	publicRouter := router.Group("/api")

	publicRouter.POST("signup", SignupHandler(authService, cfg))

	return router
}
