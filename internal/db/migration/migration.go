package migration

import (
	model "github.com/shaik80/SalahTimingsBackend/internal/models"
	"gorm.io/gorm"
)

// Migrate runs the database migrations
func Migrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	db.AutoMigrate(
		&model.Masjid{},
		&model.PrayerTime{},
	)
	// Define foreign key relationship
	db.Migrator().CreateConstraint(&model.PrayerTime{}, "masjids")
	db.Migrator().CreateConstraint(&model.PrayerTime{}, "fk_masjids_prayers")
	return nil
}
