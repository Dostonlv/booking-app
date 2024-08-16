package db

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"

	"github.com/Dostonlv/booking-app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvConfig, DBMigration func(db *gorm.DB) error) *gorm.DB {
	uri := fmt.Sprintf(`
host=%s user=%s password=%s dbname=%s sslmode=%s port=5432,`,
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBSSSLMode,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Info("Database connected")

	if err := DBMigration(db); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	return db
}
