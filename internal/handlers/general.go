package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grzadr/refscaler-service/internal/services"
	"github.com/grzadr/refscaler-service/internal/models"
)

func GetVersion(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"version": services.GetServiceVersion(),
	})
}

func GetCurrentStatus(c *fiber.Ctx) error {
	current := services.GetCurrentStatus()

	return c.JSON(models.HealthResponse{
		Status: current.Status,
		Uptime: current.UpTime,
		Timestamp: current.Timestamp.Unix(),
		StartTime: current.StartTime.Unix(),
		Version: current.Version,
	})
}
