package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/shaik80/SalahTimingsBackend/internal/api/handler"
)

// Setup initializes and returns the Fiber app with defined routes
func Setup(app *fiber.App) *fiber.App {
	// Initialize Handler
	masjidHandler := handler.NewMasjidHandler()
	prayerTimeHandler := handler.NewPrayerTimeHandler()

	// Prayer Time unsecured routes
	app.Get("/prayer-times/:id", prayerTimeHandler.GetPrayerTime)

	// Masjid unsecured routes
	app.Get("/masjids/:id/prayers", masjidHandler.GetMasjidWithPrayers)
	app.Post("/masjids/login", masjidHandler.MasjidLogin)
	app.Get("/masjids/:id", masjidHandler.GetMasjid)
	app.Post("/masjids", masjidHandler.CreateMasjid)



	app.Use(SecurityMiddleware)
	// Masjid secure routes
	app.Put("/masjids/:id", masjidHandler.UpdateMasjid)
	app.Delete("/masjids/:id", masjidHandler.DeleteMasjid)

	// Prayer secure routes
	app.Post("/prayer-times", prayerTimeHandler.CreatePrayerTime)
	app.Put("/prayer-times/:id", prayerTimeHandler.UpdatePrayerTime)
	app.Delete("/prayer-times/:id", prayerTimeHandler.DeletePrayerTime)



	return app
}

// SecurityMiddleware handles the security logic for secured APIs
func SecurityMiddleware(c *fiber.Ctx) error {
	// Get the token from the Authorization header
	authorizationHeader := c.Get("Authorization")
	if authorizationHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
	}

	// Extract the token from the Authorization header
	// The expected format is "Bearer <token>"
	tokenString := ""
	fmt.Sscanf(authorizationHeader, "Bearer %s", &tokenString)

	// Parse and validate the token
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Provide the key for validation
		return []byte("secret"), nil
	})
	if err != nil || !parsedToken.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid parsed or error token"})
	}

	// Access the username from the JWT payload
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid claims"})
	}
	username, ok := claims["username"].(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid username"})
	}

	// Set the username in the locals to make it accessible in the route handlers
	c.Locals("username", username)

	// Continue to the next middleware or route handler
	return c.Next()
}
