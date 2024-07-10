package postgres

import (
	"context"
	"fmt"
	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(ctx context.Context, cfg config.Postgres) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DB)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.MaxConnections)

	err = sqlDB.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.Role{}, &domain.User{}, &domain.Rank{}, &domain.Flight{}, &domain.Ticket{})
	if err != nil {
		return nil, fmt.Errorf("auto migration failed: %w", err)
	}

	err = insertInitialData(db)
	if err != nil {
		return nil, fmt.Errorf("inserting initial data failed: %w", err)
	}

	return db, nil
}

func insertInitialData(db *gorm.DB) error {
	var count int64
	db.Model(&domain.Role{}).Count(&count)
	if count == 0 {
		roles := []domain.Role{
			{ID: uuid.New(), Name: "user"},
			{ID: uuid.New(), Name: "admin"},
		}
		if err := db.Create(&roles).Error; err != nil {
			return err
		}
	}

	db.Model(&domain.Rank{}).Count(&count)
	if count == 0 {
		ranks := []domain.Rank{
			{ID: uuid.New(), Name: "economy"},
			{ID: uuid.New(), Name: "business"},
			{ID: uuid.New(), Name: "deluxe"},
		}
		if err := db.Create(&ranks).Error; err != nil {
			return err
		}
	}

	return nil
}
