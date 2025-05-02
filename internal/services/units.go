package services

import (
	"maps"
	"slices"
	"github.com/grzadr/refscaler/units"
	"github.com/grzadr/refscaler-service/internal/models"
)

// UnitService handles unit registry data access
type UnitService struct {
	Registry units.UnitRegistryJSON
}

// NewUnitService initializes the unit service with data
func NewUnitService() *UnitService {
	return &UnitService{
		Registry: units.EmbeddedUnitRegistry.Serialize(),
	}
}

// GetAllUnits returns the complete unit registry
func (s *UnitService) GetAllUnits() models.AllUnitsResponse {
	return slices.Collect(maps.Keys(s.Registry))
}

// GetUnitGroup retrieves a specific unit group by name
func (s *UnitService) GetUnitGroup(name string) (models.UnitGroupResponse, bool) {
	group, exists := s.Registry[name]
	if !exists {
		return nil, false
	}

	response := make(models.UnitGroupResponse, 0, len(group))

	for _, unit := range group {
		response = append(response, models.UnitJSON{
			Name: unit.Name,
			Value: unit.Value,
			Aliases: unit.Aliases,
		})
	}
	return response, true
}
