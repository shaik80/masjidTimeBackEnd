package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	db "github.com/shaik80/SalahTimingsBackend/internal/db"
	"github.com/shaik80/SalahTimingsBackend/internal/db/dal"
	model "github.com/shaik80/SalahTimingsBackend/internal/models"
)

// MasjidHandler handles HTTP requests related to Masjid
type MasjidHandler struct {
	masjidDAL *dal.MasjidDAL
}

// NewMasjidHandler creates a new instance of MasjidHandler
func NewMasjidHandler() *MasjidHandler {
	return &MasjidHandler{
		masjidDAL: dal.NewMasjidDAL(db.GetDB()), // Assuming `db.DB` is your GORM database instance
	}
}

// CreateMasjid handles the creation of a new masjid
func (c *MasjidHandler) CreateMasjid(ctx *fiber.Ctx) error {
	// Parse request body
	var newMasjid model.Masjid
	newMasjid.Status = "inactive"
	if err := ctx.BodyParser(&newMasjid); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	isDataFound, _, err := c.GetMasjidByEmail(ctx, newMasjid.Email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email already in use"})
	}

	if !isDataFound {
		// Create masjid in the database
		if err := c.masjidDAL.CreateMasjid(&newMasjid); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create masjid"})
		}
	}

	// Return the created masjid
	return ctx.Status(fiber.StatusCreated).JSON(newMasjid)
}

func (c *MasjidHandler) GetMasjidByEmail(ctx *fiber.Ctx, email string) (bool, *model.Masjid, error) {
	// get data from DB
	data, err := c.masjidDAL.GetMasjidByEmail(email)
	if err != nil {
		return false, &model.Masjid{}, ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch from db"})
	}
	if data.Email == "" {
		return false, &model.Masjid{}, nil
	}
	return true, data, nil
}

// GetMasjid handles retrieving details of a specific masjid
func (c *MasjidHandler) GetMasjid(ctx *fiber.Ctx) error {
	// Get masjid ID from URL parameter
	id := ctx.Params("id")

	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid masjid ID"})
	}

	// Check if the provided masjidID is a valid UUID
	masjidID, err := uuid.Parse(id)
	if err != nil {
		// Return an error response if the masjidID is not a valid UUID
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid masjidID: %s", id),
		})
	}

	// Retrieve masjid details from the database
	masjid, err := c.masjidDAL.GetMasjidByID(masjidID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Masjid not found"})
	}

	// Return masjid details
	return ctx.JSON(masjid)
}

// UpdateMasjid handles updating details of a specific masjid
func (c *MasjidHandler) UpdateMasjid(ctx *fiber.Ctx) error {
	// Get masjid ID from URL parameter
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid masjid ID"})
	}

	// Check if the provided masjidID is a valid UUID
	masjidID, err := uuid.Parse(id)
	if err != nil {
		// Return an error response if the masjidID is not a valid UUID
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid masjidID: %s", id),
		})
	}

	// Retrieve masjid details from the database
	masjid, err := c.masjidDAL.GetMasjidByID(masjidID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Masjid not found"})
	}

	// Parse request body for updated masjid details
	var updatedMasjid model.Masjid
	if err := ctx.BodyParser(&updatedMasjid); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Update masjid details
	masjid.ID = masjidID
	masjid.Name = updatedMasjid.Name
	masjid.Lat = updatedMasjid.Lat
	masjid.Long = updatedMasjid.Long
	masjid.Pass = updatedMasjid.Pass
	masjid.PhoneNo = updatedMasjid.PhoneNo
	masjid.Country = updatedMasjid.Country
	masjid.CityVillage = updatedMasjid.CityVillage
	masjid.Address = updatedMasjid.Address
	masjid.Status = updatedMasjid.Status
	masjid.Prayers = updatedMasjid.Prayers

	if masjid.Email != "" {
		// Save the updated masjid to the database
		if err := c.masjidDAL.UpdateMasjid(*masjid); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update masjid"})
		}
	} else {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "there is no masjid available"})
	}
	// Return the updated masjid
	return ctx.JSON(masjid)
}

// DeleteMasjid handles deleting a specific masjid
func (c *MasjidHandler) DeleteMasjid(ctx *fiber.Ctx) error {
	// Get masjid ID from URL parameter
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid masjid ID"})
	}

	// check masjidId is valid or not
	masjidID, err := uuid.Parse(id)
	if err != nil {
		// Return an error response if the masjidID is not a valid UUID
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid masjidID: %s", id),
		})
	}

	// Delete masjid from the database
	if err := c.masjidDAL.DeleteMasjid(masjidID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete masjid"})
	}

	// Return success message
	return ctx.JSON(fiber.Map{"message": "Masjid deleted successfully"})
}

// GetMasjidWithPrayers handles retrieving a specific masjid along with an array of associated prayers
func (c *MasjidHandler) GetMasjidWithPrayers(ctx *fiber.Ctx) error {
	// Get masjid ID from URL parameter
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid masjid ID"})
	}

	// check masjidId is valid or not
	masjidID, err := uuid.Parse(id)
	if err != nil {
		// Return an error response if the masjidID is not a valid UUID
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid masjidID: %s", id),
		})
	}

	// Retrieve masjid details from the database, including associated prayers
	masjid, err := c.masjidDAL.GetMasjidWithPrayers(masjidID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Masjid not found"})
	}

	// Return masjid details along with associated prayers
	return ctx.JSON(masjid)
}

func (c *MasjidHandler) MasjidLogin(ctx *fiber.Ctx) error {
	// Parse request body for updated masjid details
	var masjid model.Masjid
	if err := ctx.BodyParser(&masjid); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	_, data, err := c.GetMasjidByEmail(ctx, masjid.Email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email already in use"})
	}

	// Replace this with your actual authentication logic
	if (masjid.Email != "admin" && masjid.Email != data.Email) || (masjid.Pass != "password" && masjid.Pass != data.Pass) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"username": masjid.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as a response.
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{"token": tokenString})
}
