package handlers

import (
	
	"github.com/gofiber/fiber/v2"
	"github.com/grzadr/refscaler-service/internal/services"
)

// UnitHandler manages HTTP requests for unit data
type UnitHandler struct {
	unitService *services.UnitService
}

// NewUnitHandler creates a new unit handler with the required service
func NewUnitHandler(unitService *services.UnitService) *UnitHandler {
	return &UnitHandler{
		unitService: unitService,
	}
}

// GetAllUnits returns all units as JSON
// @Summary Get all units
// @Description Returns all available units in the registry
// @Tags units
// @Produce json
// @Success 200 {object} models.AllUnitsResponse
// @Router /units [get]
func (h *UnitHandler) GetAllUnits(c *fiber.Ctx) error {
	return c.JSON(h.unitService.GetAllUnits())
}

// GetUnitByName returns a specific unit group by name
// @Summary Get unit by name
// @Description Returns a specific unit group by its name
// @Tags units
// @Produce json
// @Param name path string true "Unit group name"
// @Success 200 {object} models.UnitGroupResponse
// @Failure 404 {object} map[string]string
// @Router /units/{name} [get]
func (h *UnitHandler) GetUnitByName(c *fiber.Ctx) error {
	name := c.Params("name")

	group, exists := h.unitService.GetUnitGroup(name)
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Unit not found",
			"name":  name,
		})
	}

	return c.JSON(group)
}