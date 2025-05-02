package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/grzadr/refscaler-service/internal/routes"
	"github.com/grzadr/refscaler-service/internal/services"

	_ "github.com/grzadr/refscaler-service/cmd/service/docs"
)

var Version = ""

// @title RefScaler
// @version 1.0.0
// @description This is a service for refscaler app
// @contact.name Adrian Grzemski
// @contact.email adrian.grzemski@gmail.com
// @license.name MIT
// @host localhost:3000
// @BasePath /
func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupUnitsRoutes(app)
	routes.SetupScalerRoutes(app)
	routes.SetupGeneralRoutes(app)

	services.SetupService(Version)

	app.Get("/swagger/*", swagger.HandlerDefault)

	// Start server
	log.Println("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
