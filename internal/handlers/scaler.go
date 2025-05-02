package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/grzadr/refscaler-service/internal/models"
	"github.com/grzadr/refscaler-service/internal/services"
)

// PostScaledEnlistment scales an enlistment based on the provided scale
// @Summary Scale an enlistment
// @Description Scales an enlistment based on the provided scale factor
// @Tags scaler
// @Accept json
// @Produce json
// @Param request body models.EnlistmentRequest true "Enlistment scale request"
// @Success 200 {object} models.EnlistmentResponse
// @Failure 400 {object} map[string]string
// @Router /scale [post]
func PostScaledEnlistment(c *fiber.Ctx) error {
	req := new(models.EnlistmentRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Cannot parse JSON",
			"details": err.Error(),
		})
	}

	scaled, err := services.GetScaled(req.Enlistment, req.Scale)

	switch {
	case errors.Is(err, services.ErrEnlistmentCreate):
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Cannot parse enlistment",
			"details": err.Error(),
		})
	case errors.Is(err, services.ErrScaleConvert):
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Cannot parse scale",
			"details": err.Error(),
		})
	default:
		return c.JSON(models.EnlistmentResponse{Scaled: scaled})
	}
}