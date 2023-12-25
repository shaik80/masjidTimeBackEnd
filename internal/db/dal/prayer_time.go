package dal

import (
	"github.com/google/uuid"
	model "github.com/shaik80/SalahTimingsBackend/internal/models"
	"gorm.io/gorm"
)

// PrayerTimeDAL represents the Data Access Layer for PrayerTime model
type PrayerTimeDAL struct {
	db *gorm.DB
}

// NewPrayerTimeDAL creates a new instance of PrayerTimeDAL
func NewPrayerTimeDAL(db *gorm.DB) *PrayerTimeDAL {
	return &PrayerTimeDAL{db: db}
}

// CreatePrayerTime creates a new prayer time in the database
func (d *PrayerTimeDAL) CreatePrayerTime(prayerTime *model.PrayerTime) error {
	return d.db.Create(prayerTime).Error
}

// GetPrayerTimeByID retrieves a prayer time by its ID from the database
func (d *PrayerTimeDAL) GetPrayerTimeByID(id uuid.UUID) (*model.PrayerTime, error) {
	var prayerTime model.PrayerTime
	err := d.db.First(&prayerTime, id).Error
	return &prayerTime, err
}

// UpdatePrayerTime updates the details of a prayer time in the database
func (d *PrayerTimeDAL) UpdatePrayerTime(prayerTime *model.PrayerTime) error {
	return d.db.Save(prayerTime).Error
}

// DeletePrayerTime deletes a prayer time from the database
func (d *PrayerTimeDAL) DeletePrayerTime(id uuid.UUID) error {
	return d.db.Delete(&model.PrayerTime{}, id).Error
}
