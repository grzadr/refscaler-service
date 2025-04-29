package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grzadr/refscaler-service/internal/handlers"
)

// SetupUnitsRoutes configures all application routes
func SetupGeneralRoutes(app *fiber.App) {
	app.Get("/version", handlers.GetVersion)
	app.Get("/health", handlers.GetCurrentStatus)
}
