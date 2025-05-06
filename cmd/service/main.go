package main

import (
	// "fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	docs "github.com/grzadr/refscaler-service/cmd/service/docs"
	"github.com/grzadr/refscaler-service/internal/routes"
	"github.com/grzadr/refscaler-service/internal/services"
)

var Version = ""

func updateSwaggerDocs() {
	apiUrlBase, found := os.LookupEnv("API_URL_BASE")

	if !found {
		apiUrlBase = "localhost:3000"
	}

	apiUrlPrefix, found := os.LookupEnv("API_URL_PREFIX")

	if !found {
		apiUrlBase = "/"
	}

	docs.SwaggerInfo.Version = Version
	docs.SwaggerInfo.Host = apiUrlBase
	docs.SwaggerInfo.BasePath = apiUrlPrefix
}

// @title RefScaler
// @description This is a service for refscaler app
// @contact.name Adrian Grzemski
// @contact.email adrian.grzemski@gmail.com
// @license.name MIT
func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupUnitsRoutes(app)
	routes.SetupScalerRoutes(app)
	routes.SetupGeneralRoutes(app)

	services.SetupService(Version)

	updateSwaggerDocs()

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{
		DeepLinking: true,
		// Use a relative URL, not an absolute one
		URL: "./swagger/doc.json",
		// Optional improvements
		DocExpansion: "list",
	}))

	// Start server
	log.Println("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
