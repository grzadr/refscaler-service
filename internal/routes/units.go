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

	units_group := app.Group("/units")

	// Register routes
	units_group.Get("/", unitHandler.GetAllUnits)
	units_group.Get("/:name", unitHandler.GetUnitByName)
}
