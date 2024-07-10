package route

import (
	"database/sql"
	"github.com/dilyara4949/flight-booking-api/internal/handler"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	"github.com/dilyara4949/flight-booking-api/internal/service"

	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/gin-gonic/gin"
)

func NewAPI(cfg config.Config, db *sql.DB, gin *gin.Engine) {

	publicRouter := gin.Group("/api")

	repo := repository.NewUserRepository(db)
	authService := service.NewAuthService(repo)

	publicRouter.POST("signup", handler.SignupHandler(authService, cfg))

}
