// internal/handlers/handler.go
package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/grzadr/refscaler-service/internal/models"
)

// Handler handles HTTP requests for the frontend
type Handler struct {
	backendURL string
}

// NewHandler creates a new handler with the backend URL
func NewHandler(backendURL string) *Handler {
	return &Handler{
		backendURL: backendURL,
	}
}

// Index renders the home page
func (h *Handler) Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "RefScaler",
	})
}

// Scale handles the scaling request and returns the results
func (h *Handler) Scale(c *fiber.Ctx) error {
	// Parse form data
	enlistment := c.FormValue("enlistment")
	scale := c.FormValue("scale")

	// Log the received values for debugging
	log.Printf(
		"Received scale request - Scale: %s, Enlistment: %s",
		scale,
		enlistment,
	)

	if enlistment == "" || scale == "" {
		log.Printf("Missing required fields")
		return c.Render("partials/results", fiber.Map{
			"Error": "Both enlistment and scale must be provided",
		})
	}

	// Create request to backend
	request := models.EnlistmentRequest{
		Enlistment: enlistment,
		Scale:      scale,
	}

	// Convert to JSON
	requestBody, err := json.Marshal(request)
	if err != nil {
		log.Printf("Error marshaling request: %v", err)
		return c.Render("partials/results", fiber.Map{
			"Error": "Failed to marshal request",
		})
	}

	// Send request to backend
	backendURL := fmt.Sprintf("%s/scale", h.backendURL)
	log.Printf("Sending request to backend: %s", backendURL)

	resp, err := http.Post(
		backendURL,
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		log.Printf("Error communicating with backend: %v", err)
		return c.Render("partials/results", fiber.Map{
			"Error": fmt.Sprintf("Failed to communicate with backend: %v", err),
		})
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("Error closing response body: %v", closeErr)
		}
	}()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return c.Render("partials/results", fiber.Map{
			"Error": "Failed to read response from backend",
		})
	}

	log.Printf(
		"Response from backend (status %d): %s",
		resp.StatusCode,
		string(respBody),
	)

	// Check for error
	if resp.StatusCode != http.StatusOK {
		log.Printf("Backend returned error status: %d", resp.StatusCode)

		// Try to parse structured error response
		var errorResp map[string]string
		if err := json.Unmarshal(respBody, &errorResp); err != nil {
			// If we can't parse the error JSON, use the raw response
			log.Printf("Could not parse error response: %v", err)
			errorMessage := fmt.Sprintf("Backend error: %s", string(respBody))
			log.Printf("Error message: %s", errorMessage)
			return c.Render("partials/results", fiber.Map{
				"Error": errorMessage,
			})
		}

		// Use error from structured response
		errorMessage := fmt.Sprintf("Backend error: %s", errorResp["error"])
		if details, ok := errorResp["details"]; ok && details != "" {
			errorMessage += fmt.Sprintf(" - %s", details)
		}
		log.Printf("Structured error: %s", errorMessage)
		return c.Render("partials/results", fiber.Map{
			"Error": errorMessage,
		})
	}

	// Parse response
	var response models.EnlistmentResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		log.Printf("Error unmarshaling response: %v", err)
		return c.Render("partials/results", fiber.Map{
			"Error": "Failed to parse response from backend",
		})
	}

	log.Printf(
		"Successfully processed request, returning %d results",
		len(response.Scaled),
	)

	// Return results partial
	return c.Render("partials/results", fiber.Map{
		"Results": response.Scaled,
	})
}
