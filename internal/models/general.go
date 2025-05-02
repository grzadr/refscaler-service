package models

// VersionResponse represents the service version
// @Description Service version information
type VersionResponse struct {
	Version string `json:"version"`
}

// HealthResponse represents the service health status
// @Description Service health information
type HealthResponse struct {
	Ready     bool    `json:"ready"`
	StartTime string  `json:"starttime"`
	Uptime    float64 `json:"uptime"`
	Timestamp string  `json:"timestamp"`
	Version   string  `json:"version"`
}
