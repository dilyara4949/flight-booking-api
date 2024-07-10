package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dilyara4949/flight-booking-api/internal/config"
	_ "github.com/lib/pq"
	"time"
)

func ConnectPostgres(cfg config.Postgres) (*sql.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DB)
	database, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	database.SetMaxOpenConns(cfg.MaxConnections)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()

	err = database.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return database, nil
}
