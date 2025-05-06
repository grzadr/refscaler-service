// cmd/frontend/main.go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/grzadr/refscaler-service/internal"
	"github.com/grzadr/refscaler-service/internal/handlers"
)

var Version = ""

func main() {
	// Initialize the HTML template engine
	viewsPath := internal.DefaultEnv("VIEWS_PATH", "/assets/views")
	staticPath := internal.DefaultEnv("STATIC_PATH", "/assets/static")

	log.Printf("Using views path: %s", viewsPath)
	log.Printf("Using static path: %s", staticPath)

	engine := html.New(viewsPath, ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Add logger middleware for better debugging
	app.Use(logger.New())

	// Serve static files
	app.Static("/static", staticPath)

	// Get backend URL from environment or use default
	backendURL := internal.DefaultEnv("BACKEND_URL", "http://localhost:3000")
	log.Printf("Using backend URL: %s", backendURL)

	// Create a new handler with the backend URL
	handler := handlers.NewHandler(backendURL)

	// Homepage route
	app.Get("/", handler.Index)

	// API routes for HTMX interactions
	app.Post("/form/scale", handler.Scale)

	port := internal.DefaultEnv("PORT", "8080")
	log.Printf("Starting frontend on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
