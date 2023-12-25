package dal

import (
	"fmt"

	"github.com/google/uuid"
	model "github.com/shaik80/SalahTimingsBackend/internal/models"
	"gorm.io/gorm"
)

// MasjidDAL represents the Data Access Layer for Masjid model
type MasjidDAL struct {
	db *gorm.DB
}

// NewMasjidDAL creates a new instance of MasjidDAL
func NewMasjidDAL(db *gorm.DB) *MasjidDAL {
	return &MasjidDAL{
		db: db,
	}

}

// CreateMasjid creates a new masjid in the database
func (d *MasjidDAL) CreateMasjid(masjid *model.Masjid) error {
	return d.db.Create(&masjid).Error
}

// GetMasjidByID retrieves a masjid by ID from the database
func (d *MasjidDAL) GetMasjidByID(id uuid.UUID) (*model.Masjid, error) {
	var masjid model.Masjid
	err := d.db.First(&masjid, id).Error
	return &masjid, err
}

// GetMasjidByID retrieves a masjid by ID from the database
func (d *MasjidDAL) GetMasjidByEmail(email string) (*model.Masjid, error) {
	var masjid model.Masjid
	fmt.Println(email)
	err := d.db.Where("email = ?", email).Find(&masjid).Error
	return &masjid, err
}

// UpdateMasjid updates the details of a masjid in the database
func (d *MasjidDAL) UpdateMasjid(masjid model.Masjid) error {
	fmt.Println(masjid)
	return d.db.Where("id = ?", masjid.ID).Updates(&masjid).Error
}

// DeleteMasjid deletes a masjid from the database
func (d *MasjidDAL) DeleteMasjid(id uuid.UUID) error {
	return d.db.Delete(&model.Masjid{}, id).Error
}

// GetMasjidWithPrayers retrieves a masjid along with associated prayers from the database
func (d *MasjidDAL) GetMasjidWithPrayers(id uuid.UUID) (*model.Masjid, error) {
	var masjid model.Masjid
	err := d.db.Preload("Prayers").First(&masjid, id).Error
	return &masjid, err
}
