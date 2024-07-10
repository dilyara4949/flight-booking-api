package route

import (
	"database/sql"

	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/controller"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	"github.com/dilyara4949/flight-booking-api/internal/service"
	"github.com/gin-gonic/gin"
)

func NewSignupRoute(cfg config.Config, db *sql.DB, group *gin.RouterGroup) {
	repo := repository.NewUserRepository(db)
	authService := service.NewAuthService(repo, cfg.ContextTimeout)
	authController := controller.NewAuthController(authService, cfg)

	group.POST("signup", authController.Signup)
}
