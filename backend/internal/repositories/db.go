package repositories

import (
	"backend/internal/models"
	"backend/pkg/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	cfg := config.GetDBConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Автомиграция (создание таблиц, если их нет)
	db.AutoMigrate(&models.User{}, &models.Room{}, &models.Track{})

	return db, nil
}
