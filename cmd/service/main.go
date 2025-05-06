package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	docs "github.com/grzadr/refscaler-service/cmd/service/docs"
	"github.com/grzadr/refscaler-service/internal"
	"github.com/grzadr/refscaler-service/internal/routes"
	"github.com/grzadr/refscaler-service/internal/services"
)

var Version = ""

func updateSwaggerDocs() {
	docs.SwaggerInfo.Version = Version
	docs.SwaggerInfo.Host = internal.DefaultEnv(
		"API_URL_BASE",
		"localhost:3000",
	)
	docs.SwaggerInfo.BasePath = internal.DefaultEnv("API_URL_PREFIX", "/")
}

// @title RefScaler
// @description This is a service for refscaler app
// @contact.name Adrian Grzemski
// @contact.email adrian.grzemski@gmail.com
// @license.name MIT
func main() {
	app := fiber.New()

	routes.SetupUnitsRoutes(app)
	routes.SetupScalerRoutes(app)
	routes.SetupGeneralRoutes(app)

	services.SetupService(Version)

	updateSwaggerDocs()

	app.Get("/swagger/*", swagger.HandlerDefault)

	port := internal.DefaultEnv("PORT", "3000")

	log.Printf("Starting frontend on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
