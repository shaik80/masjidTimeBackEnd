package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Masjid represents the Masjid model
type Masjid struct {
	gorm.Model
	ID          uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string       `json:"name"`
	Lat         string       `json:"lat"`
	Long        string       `json:"long"`
	Email       string       `json:"email"`
	Pass        string       `json:"pass"`
	PhoneNo     string       `json:"phone_no"`
	Country     string       `json:"country"`
	CityVillage string       `json:"city_village"`
	Address     string       `json:"address"`
	Status      string       `json:"status"`
	Prayers     []PrayerTime `json:"Prayers"`
}
