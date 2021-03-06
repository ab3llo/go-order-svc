package db

import (
	"fmt"
	"log"
	"time"

	"github.com/ab3llo/go-order-svc/pkg/config"
	"github.com/ab3llo/go-order-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConnection struct {
	DB *gorm.DB
}

func Connect(cfg *config.Config) DatabaseConnection {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/London",
		cfg.DatabaseHost,
		cfg.DatabaseUsername,
		cfg.DatabasePassword,
		cfg.DatabaseName,
		cfg.DatabasePort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(5)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(10)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	db.AutoMigrate(&models.Order{})

	return DatabaseConnection{db}
}
