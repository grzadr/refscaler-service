package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grzadr/refscaler-service/internal/handlers"
)

func SetupScalerRoutes(app *fiber.App) {
	scaler_group := app.Group("/scale")

	scaler_group.Post("/", handlers.PostScaledEnlistment)
}
