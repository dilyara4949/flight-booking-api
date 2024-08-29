package main

import (
	"context"
	"errors"
	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/database/postgres"
	"github.com/dilyara4949/flight-booking-api/internal/database/redis"
	"github.com/dilyara4949/flight-booking-api/internal/handler"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("error getting config:", "error", err.Error())
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := postgres.Connect(ctx, cfg.Postgres)
	if err != nil {
		slog.Error("database connection failed:", "error", err.Error())
		return
	}

	cache, err := redis.Connect(ctx, cfg.Redis)
	if err != nil {
		slog.Error("redis connection failed", "error", err.Error())
		return
	}

	apiHandler := handler.NewAPI(cfg, database, cache)

	httpServer := &http.Server{
		Addr:              net.JoinHostPort(cfg.Address, cfg.RestPort),
		Handler:           apiHandler,
		ReadHeaderTimeout: cfg.HeaderTimeout,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		slog.Info("Shutting down server...")
		cancel()

		if err = httpServer.Shutdown(ctx); err != nil {
			slog.Error("Server Shutdown Failed", "error", err)
		}
	}()

	if err = httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("listen:", "error", err)
		return
	}
}
