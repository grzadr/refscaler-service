package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grzadr/refscaler-service/internal/models"
	"github.com/grzadr/refscaler-service/internal/services"
)

func GetVersion(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"version": services.GetServiceVersion(),
	})
}

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
