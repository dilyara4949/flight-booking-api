package main

import (
	"context"
	"fmt"
	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/database/postgres"
	"github.com/dilyara4949/flight-booking-api/internal/route"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func runServer() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("error getting config: %w", err)
	}

	ctx := context.Background()

	database, err := postgres.Connect(ctx, cfg.Postgres)
	if err != nil {
		return fmt.Errorf("database connection failed: %w", err)
	}

	defer func() {
		sqlDB, err := database.DB()
		if err != nil {
			log.Printf("Error getting sql.DB from gorm.DB: %v", err)
			return
		}
		sqlDB.Close()
	}()

	ginRouter := gin.Default()
	route.NewAPI(cfg, database, ginRouter)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Address, cfg.RestPort),
		Handler: ginRouter,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		slog.Info("Shutting down server...")
		if err := httpServer.Shutdown(context.Background()); err != nil {
			slog.Error("Server Shutdown Failed:", err.Error())
		}
	}()

	if err = httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("listen: %w", err)
	}

	return nil
}

func main() {
	if err := runServer(); err != nil {
		slog.Error("server failed:", err.Error())
	}
}
