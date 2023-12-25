package database

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/shaik80/SalahTimingsBackend/config"
	"github.com/shaik80/SalahTimingsBackend/internal/db/migration"
	"github.com/shaik80/SalahTimingsBackend/utils/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return dbInstance
}

// Initialize initializes the database connection
func Initialize() (*gorm.DB, error) {
	cfg := config.GetConfig().Database

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println(db, "database initialized")

	// Set connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	dbInstance = db

	// Run migrations
	if err := migration.Migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDB() error {
	if dbInstance != nil {
		db, err := dbInstance.DB()
		if err != nil {
			logger.Logger.Printf("Error getting the database instance: %v", err)
			return errors.New("Error getting the database instance:" + err.Error())
		}
		if err := db.Close(); err != nil {
			logger.Logger.Printf("Error closing the database connection: %v", err.Error())
			return errors.New("Error closing the database connection:" + err.Error())

		}
	}
	return nil
}

func GracefulShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-stop
		CloseDB()
		os.Exit(0)
	}()
}
