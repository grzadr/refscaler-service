package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/grzadr/refscaler-service/internal/routes"
	"github.com/grzadr/refscaler-service/internal/services"
)

var Version = ""

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupUnitsRoutes(app)
	routes.SetupScalerRoutes(app)
	routes.SetupGeneralRoutes(app)

	services.SetupService(Version)

	// app.Get("/swagger/*", swagger.HandlerDefault)

	// Start server
	log.Println("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
