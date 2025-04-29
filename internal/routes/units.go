package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grzadr/refscaler-service/internal/handlers"
	"github.com/grzadr/refscaler-service/internal/services"
)

// SetupUnitsRoutes configures all application routes
func SetupUnitsRoutes(app *fiber.App) {
	// Initialize services
	unitService := services.NewUnitService()

	// Initialize handlers
	unitHandler := handlers.NewUnitHandler(unitService)

	// Register routes
	app.Get("/units", unitHandler.GetAllUnits)
	app.Get("/units/:name", unitHandler.GetUnitByName)
}
