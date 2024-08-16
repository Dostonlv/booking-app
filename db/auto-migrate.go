package db

import (
	"github.com/Dostonlv/booking-app/models"
	"gorm.io/gorm"
)

func DBMigration(db *gorm.DB) error {
	db.Migrator().DropTable(&models.Event{})
	return db.AutoMigrate(&models.Event{})
}
