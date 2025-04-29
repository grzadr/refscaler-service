package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/grzadr/refscaler-service/internal/routes"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupGeneralRoutes(app)
	routes.SetupUnitsRoutes(app)

	// Start server
	log.Println("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
