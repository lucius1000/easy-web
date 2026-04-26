package config

import (
	"fmt"

	"github.com/user/go-gin-gorm-starter/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto Migrate
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
