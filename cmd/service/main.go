package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"github.com/grzadr/refscaler-service/internal/routes"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	log.Println("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
