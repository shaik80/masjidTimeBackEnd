package server

import (
	"github.com/gofiber/fiber/v2"
	api_log "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/shaik80/SalahTimingsBackend/config"
	"github.com/shaik80/SalahTimingsBackend/internal/api/router"
	"github.com/shaik80/SalahTimingsBackend/utils/logger"
)

func New() *fiber.App {

	app := fiber.New()
	logger.Logger.Printf("Server is running on %s:%d", config.GetConfig().Server.Address, config.GetConfig().Server.Port)
	// Define your routes here
	app.Use(api_log.New())

	return router.Setup(app)
}
