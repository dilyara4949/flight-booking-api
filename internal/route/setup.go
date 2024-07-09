package route

import (
	"database/sql"

	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/gin-gonic/gin"
)

func Setup(cfg config.Config, db *sql.DB, gin *gin.Engine) {
	publicRouter := gin.Group("/api")

	NewSignupRoute(cfg, db, publicRouter)
}
