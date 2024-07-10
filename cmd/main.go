package main

import (
	"fmt"
	"log"

	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/database/postgres"
	"github.com/dilyara4949/flight-booking-api/internal/route"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error at getting config: %v", err)
	}
	fmt.Println(cfg.ContextTimeout)

	db, err := postgres.ConnectPostgres(cfg.PostgresCfg)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	defer db.Close()

	ginRouter := gin.Default()
	route.Setup(cfg, db, ginRouter)

	err = ginRouter.Run(fmt.Sprintf("%s:%s", cfg.Address, cfg.RestPort))
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
