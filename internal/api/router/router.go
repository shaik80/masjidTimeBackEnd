package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shaik80/SalahTimingsBackend/internal/api/handler"
)

// Setup initializes and returns the Fiber app with defined routes
func Setup(app *fiber.App) *fiber.App {
	// Initialize Handler
	masjidHandler := handler.NewMasjidHandler()
	prayerTimeHandler := handler.NewPrayerTimeHandler()

	// Masjid routes
	app.Post("/masjids", masjidHandler.CreateMasjid)
	// app.Post("/masjids/login", masjidHandler.MasjidLogin)
	app.Get("/masjids/:id", masjidHandler.GetMasjid)
	app.Put("/masjids/:id", masjidHandler.UpdateMasjid)
	app.Delete("/masjids/:id", masjidHandler.DeleteMasjid)

	// Prayer Time routes
	app.Post("/prayer-times", prayerTimeHandler.CreatePrayerTime)
	app.Get("/prayer-times/:id", prayerTimeHandler.GetPrayerTime)
	app.Put("/prayer-times/:id", prayerTimeHandler.UpdatePrayerTime)
	app.Delete("/prayer-times/:id", prayerTimeHandler.DeletePrayerTime)

	// Masjid with Array of Prayers route
	app.Get("/masjids/:id/prayers", masjidHandler.GetMasjidWithPrayers)

	return app
}

func SecurityMiddleware(c *fiber.Ctx) error {
	// Define a list of secured paths
	securedPaths := []string{"/masjids"}

	// Check if the requested path is in the list of secured paths
	for _, path := range securedPaths {
		if path == c.Path() {
			// Handle secured API logic here
			return c.SendString("This is a secured API.")
		}
	}

	// If the path is not in the list of secured paths, continue to the next middleware
	return c.Next()
}
