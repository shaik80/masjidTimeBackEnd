package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PrayerTime represents the PrayerTime model
type PrayerTime struct {
	gorm.Model
	// ENUM('fajar', 'zuhar', 'ashar', 'magrib', 'isha', 'juma', 'eid')
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name     string    `json:"name"`
	Azan     string    `json:"azan"`
	Iqama    string    `json:"iqama"`
	MasjidID uuid.UUID `json:"masjid_id" gorm:"foreignKey:MasjidID;references:masjids(id)"`
}

