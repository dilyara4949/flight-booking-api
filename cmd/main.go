package main

import (
	"context"
	"fmt"
	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/database/postgres"
	"github.com/dilyara4949/flight-booking-api/internal/handler"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	"github.com/dilyara4949/flight-booking-api/internal/service"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("error getting config: %w", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := postgres.Connect(ctx, cfg.Postgres)
	if err != nil {
		slog.Error("database connection failed: %w", err)
		return
	}

	repo := repository.NewUserRepository(database)
	authService := service.NewAuthService(repo)

	apiHandler := handler.NewAPI(cfg, authService)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Address, cfg.RestPort),
		Handler: apiHandler,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		slog.Info("Shutting down server...")
		if err := httpServer.Shutdown(ctx); err != nil {
			slog.Error("Server Shutdown Failed:", err.Error())
		}
	}()

	if err = httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("listen: %w", err)
		return
	}
}
