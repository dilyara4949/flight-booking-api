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

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("error at getting config:", err.Error())
	}

	database, err := postgres.ConnectPostgres(cfg.PostgresCfg)
	if err != nil {
		slog.Error("database connection failed:", err.Error())
	}

	defer database.Close()

	ginRouter := gin.Default()
	route.Setup(cfg, database, ginRouter)

	err = ginRouter.Run(fmt.Sprintf("%s:%s", cfg.Address, cfg.RestPort))
	if err != nil {
		slog.Error("server failed to start:", err.Error())
	}
}
