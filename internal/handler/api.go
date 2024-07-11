package handler

import (
	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/gin-gonic/gin"
)

func NewAPI(cfg config.Config, authService AuthService) *gin.Engine {

	ginRouter := gin.Default()
	publicRouter := ginRouter.Group("/api")

	publicRouter.POST("signup", SignupHandler(authService, cfg))

	return ginRouter
}
