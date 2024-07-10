package main

import (
	"fmt"
	"log/slog"

	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/database/postgres"
	"github.com/dilyara4949/flight-booking-api/internal/route"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := slog.Default()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Error("error at getting config:", err.Error())
	}

	database, err := postgres.ConnectPostgres(cfg.PostgresCfg)
	if err != nil {
		logger.Error("database connection failed:", err.Error())
	}

	defer database.Close()

	ginRouter := gin.Default()
	route.Setup(cfg, database, ginRouter)

	err = ginRouter.Run(fmt.Sprintf("%s:%s", cfg.Address, cfg.RestPort))
	if err != nil {
		logger.Error("server failed to start:", err.Error())
	}
}
