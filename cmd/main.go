package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/database/postgres"
	"github.com/dilyara4949/flight-booking-api/internal/route"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func runServer() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("error getting config: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()

	database, err := postgres.Connect(ctx, cfg.Postgres)
	if err != nil {
		return fmt.Errorf("database connection failed: %w", err)
	}
	defer database.Close()

	ginRouter := gin.Default()
	route.NewAPI(cfg, database, ginRouter)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Address, cfg.RestPort),
		Handler: ginRouter,
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return httpServer.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return httpServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("exit reason: %w", err)
	}

	return nil
}

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("Shutting down server...")
		os.Exit(0)
	}()

	if err := runServer(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
