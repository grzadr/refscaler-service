package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grzadr/refscaler-service/internal/models"
	"github.com/grzadr/refscaler-service/internal/services"
)

// GetVersion returns the current service version
// @Summary Get service version
// @Description Returns the current version of the service
// @Tags general
// @Produce json
// @Success 200 {object} models.VersionResponse
// @Router /version [get]
func GetVersion(c *fiber.Ctx) error {
	return c.JSON(models.VersionResponse{
		Version: services.GetServiceVersion(),
	})
}

// GetCurrentStatus returns the health status of the service
// @Summary Get service health
// @Description Returns the current health status of the service
// @Tags general
// @Produce json
// @Success 200 {object} models.HealthResponse
// @Router /health [get]
func GetCurrentStatus(c *fiber.Ctx) error {
	current := services.GetCurrentStatus()

	return c.JSON(models.HealthResponse{
		Ready:     current.Ready,
		Uptime:    current.UpTime,
		Timestamp: current.Timestamp,
		StartTime: current.StartTime,
		Version:   current.Version,
	})
}