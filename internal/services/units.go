package services

import (
	"github.com/grzadr/refscaler/units"
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
func (s *UnitService) GetAllUnits() units.UnitRegistryJSON {
	return s.Registry
}

// GetUnitGroup retrieves a specific unit group by name
func (s *UnitService) GetUnitGroup(name string) (units.UnitGroupJSON, bool) {
	group, exists := s.Registry[name]
	return group, exists
}
