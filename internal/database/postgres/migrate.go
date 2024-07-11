package postgres

import (
	"fmt"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&domain.Role{}, &domain.User{}, &domain.Rank{}, &domain.Flight{}, &domain.Ticket{})
	if err != nil {
		return fmt.Errorf("auto migration failed: %w", err)
	}

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
		if err = db.Create(&ranks).Error; err != nil {
			return err
		}
	}

	return nil
}
