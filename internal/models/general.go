package models

type HealthResponse struct {
	Status    string  `json:"status"`
	StartTime int64   `json:"starttime"`
	Uptime    float64 `json:"uptime"`
	Timestamp int64   `json:"timestamp"`
	Version   string  `json:"version"`
}
