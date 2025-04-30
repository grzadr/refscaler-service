package models

type HealthResponse struct {
	Ready     bool    `json:"ready"`
	StartTime string  `json:"starttime"`
	Uptime    float64 `json:"uptime"`
	Timestamp string  `json:"timestamp"`
	Version   string  `json:"version"`
}
