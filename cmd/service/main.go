package main

import (
	"log"
	"os"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/grzadr/refscaler-service/cmd/service/docs"
	"github.com/grzadr/refscaler-service/internal/routes"
	"github.com/grzadr/refscaler-service/internal/services"
)

var Version = ""

// @title RefScaler
// @version 1.0.0
// @description This is a service for refscaler app
// @contact.name Adrian Grzemski
// @contact.email adrian.grzemski@gmail.com
// @license.name MIT
// @BasePath /
func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupUnitsRoutes(app)
	routes.SetupScalerRoutes(app)
	routes.SetupGeneralRoutes(app)

	services.SetupService(Version)

	hostname, found := os.LookupEnv("SWAGGER_HOST")

	if !found {
		hostname = "localhost:3000"
	}

	//app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{
	    DeepLinking: true,
	    // Use a relative URL, not an absolute one
	    URL: fmt.Sprintf("%s/swagger/doc.json", hostname),
	    // Optional improvements
	    DocExpansion: "list",
	}))

	// Start server
	log.Println("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
