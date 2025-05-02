package models

// @Description A list of unit groups
type AllUnitsResponse []string

// UnitJSON represents a single Unit from UnitsGroup
// @Description A unit with value and available aliases
type UnitJSON struct {
	Name    string   `json:"name"`
	Value   float64  `json:"value"`
	Aliases []string `json:"aliases"`
}

// @Description A list of units representing a group
type UnitGroupResponse []UnitJSON
