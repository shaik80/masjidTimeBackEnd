package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	db "github.com/shaik80/SalahTimingsBackend/internal/db"
	"github.com/shaik80/SalahTimingsBackend/internal/db/dal"
	model "github.com/shaik80/SalahTimingsBackend/internal/models"
)

// PrayerTimeHandler handles HTTP requests related to PrayerTime
type PrayerTimeHandler struct {
	prayerTimeDAL *dal.PrayerTimeDAL
}

// NewPrayerTimeHandler creates a new instance of PrayerTimeController
func NewPrayerTimeHandler() *PrayerTimeHandler {
	return &PrayerTimeHandler{
		prayerTimeDAL: dal.NewPrayerTimeDAL(db.GetDB()), // Assuming `db.DB` is your GORM database instance
	}
}

// CreatePrayerTime handles the creation of a new prayer time
func (c *PrayerTimeHandler) CreatePrayerTime(ctx *fiber.Ctx) error {
	// Parse request body
	var newPrayerTime model.PrayerTime
	if err := ctx.BodyParser(&newPrayerTime); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Create prayer time in the database
	if err := c.prayerTimeDAL.CreatePrayerTime(&newPrayerTime); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create prayer time"})
	}

	// Return the created prayer time
	return ctx.Status(fiber.StatusCreated).JSON(newPrayerTime)
}

// GetPrayerTime handles retrieving details of a specific prayer time
func (c *PrayerTimeHandler) GetPrayerTime(ctx *fiber.Ctx) error {
	// Get prayer time ID from URL parameter
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid prayer time ID"})
	}

	// check masjidId is valid or not
	prayerTimeID, err := uuid.Parse(id)
	if err != nil {
		// Return an error response if the masjidID is not a valid UUID
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid masjidID: %s", id),
		})
	}

	// Retrieve prayer time details from the database
	prayerTime, err := c.prayerTimeDAL.GetPrayerTimeByID(prayerTimeID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Prayer time not found"})
	}

	// Return prayer time details
	return ctx.JSON(prayerTime)
}

// UpdatePrayerTime handles updating details of a specific prayer time
func (c *PrayerTimeHandler) UpdatePrayerTime(ctx *fiber.Ctx) error {
	// Get prayer time ID from URL parameter
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid prayer time ID"})
	}

	// check masjidId is valid or not
	prayerTimeID, err := uuid.Parse(id)
	if err != nil {
		// Return an error response if the masjidID is not a valid UUID
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid masjidID: %s", id),
		})
	}

	// Retrieve prayer time details from the database
	prayerTime, err := c.prayerTimeDAL.GetPrayerTimeByID(prayerTimeID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Prayer time not found"})
	}

	// Parse request body for updated prayer time details
	var updatedPrayerTime model.PrayerTime
	if err := ctx.BodyParser(&updatedPrayerTime); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Update prayer time details
	prayerTime.Name = updatedPrayerTime.Name
	prayerTime.Azan = updatedPrayerTime.Azan
	prayerTime.Iqama = updatedPrayerTime.Iqama
	// Update other fields as needed

	// Save the updated prayer time to the database
	if err := c.prayerTimeDAL.UpdatePrayerTime(prayerTime); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update prayer time"})
	}

	// Return the updated prayer time
	return ctx.JSON(prayerTime)
}

// DeletePrayerTime handles deleting a specific prayer time
func (c *PrayerTimeHandler) DeletePrayerTime(ctx *fiber.Ctx) error {
	// Get prayer time ID from URL parameter
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid prayer time ID"})
	}

	// check masjidId is valid or not
	prayerTimeID, err := uuid.Parse(id)
	if err != nil {
		// Return an error response if the masjidID is not a valid UUID
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid masjidID: %s", id),
		})
	}

	// Delete prayer time from the database
	if err := c.prayerTimeDAL.DeletePrayerTime(prayerTimeID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete prayer time"})
	}

	// Return success message
	return ctx.JSON(fiber.Map{"message": "Prayer time deleted successfully"})
}
