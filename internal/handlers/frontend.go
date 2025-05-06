// cmd/frontend/handlers/handler.go
package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	if enlistment == "" || scale == "" {
		return c.Status(fiber.StatusBadRequest).
			Render("partials/results", fiber.Map{
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
		return c.Status(fiber.StatusInternalServerError).
			Render("partials/results", fiber.Map{
				"Error": "Failed to marshal request",
			})
	}

	// Send request to backend
	resp, err := http.Post(
		fmt.Sprintf("%s/scale", h.backendURL),
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			Render("partials/results", fiber.Map{
				"Error": fmt.Sprintf(
					"Failed to communicate with backend: %v",
					err,
				),
			})
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			Render("partials/results", fiber.Map{
				"Error": "Failed to read response from backend",
			})
	}

	// Check for error
	if resp.StatusCode != http.StatusOK {
		var errorResp map[string]string
		if err := json.Unmarshal(respBody, &errorResp); err != nil {
			return c.Status(fiber.StatusInternalServerError).
				Render("partials/results", fiber.Map{
					"Error": fmt.Sprintf("Backend error: %s", respBody),
				})
		}
		return c.Status(resp.StatusCode).Render("partials/results", fiber.Map{
			"Error": fmt.Sprintf("Backend error: %s", errorResp["error"]),
		})
	}

	// Parse response
	var response models.EnlistmentResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			Render("partials/results", fiber.Map{
				"Error": "Failed to parse response from backend",
			})
	}

	// Return results partial
	return c.Render("partials/results", fiber.Map{
		"Results": response.Scaled,
	})
}
