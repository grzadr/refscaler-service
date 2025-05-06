package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	docs "github.com/grzadr/refscaler-service/cmd/service/docs"
	"github.com/grzadr/refscaler-service/internal/routes"
	"github.com/grzadr/refscaler-service/internal/services"
)

var Version = ""

func DefaultEnv(key, value string) string {
	env, found := os.LookupEnv(key)

	if !found {
		return value
	}

	return env
}

func updateSwaggerDocs() {
	docs.SwaggerInfo.Version = Version
	docs.SwaggerInfo.Host = DefaultEnv("API_URL_BASE", "localhost:3000")
	docs.SwaggerInfo.BasePath = DefaultEnv("API_URL_PREFIX", "/")
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

	log.Println("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
