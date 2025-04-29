package services

import (
	"time"
	"sync"
)

type Status int

// Status enum values
const (
	StatusStarting Status = iota
	StatusUp
	StatusDegraded
	StatusMaintenance
	StatusDown
	StatusShuttingDown
)

// String implements the fmt.Stringer interface for automatic string conversion
func (s Status) String() string {
	return [...]string{
		"STARTING",
		"UP",
		"DEGRADED",
		"MAINTENANCE",
		"DOWN",
		"SHUTTING_DOWN",
	}[s]
}

// MarshalJSON enables direct JSON serialization
func (s Status) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

type ServiceStatus struct {
	mu        sync.RWMutex
	startTime time.Time
	ready     bool
	version   string
	status    Status
}

func NewServiceStatus(version string) *ServiceStatus {
	return &ServiceStatus{
		startTime: time.Now(),
		ready:     false,
		version:   version,
		status:    StatusStarting,
	}
}

func (s *ServiceStatus) SetReady(ready bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ready = ready
	if ready && s.status != StatusDegraded && s.status != StatusMaintenance {
		s.status = StatusUp
	} else if !ready && s.status == StatusUp {
		s.status = StatusDown
	}
}

func (s *ServiceStatus) SetStatus(status Status) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status = status

	switch status {
	case StatusUp:
		s.ready = true
	case StatusStarting, StatusDown, StatusDegraded, StatusMaintenance, StatusShuttingDown:
		s.ready = false
	}
}

type CurrentStatus struct {
	Timestamp time.Time
	StartTime time.Time
	UpTime float64
	Version   string
	Status    string
}

// GetStatus returns the current service status
func (s *ServiceStatus) GetCurrentStatus() CurrentStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return CurrentStatus{
		Timestamp: time.Now(),
		StartTime: s.startTime,
		Version: s.version,
		Status: s.StatusString(),
		UpTime: s.Uptime(),
	}
}

// StatusString returns the string representation of current status
func (s *ServiceStatus) StatusString() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.status.String()
}

// IsReady returns whether the service is ready for traffic
func (s *ServiceStatus) IsReady() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.ready
}

// Uptime returns the service uptime in seconds
func (s *ServiceStatus) Uptime() float64 {
	return time.Since(s.startTime).Seconds()
}

// Global instance - unexported
var defaultStatus *ServiceStatus
var once sync.Once
var Version string


func GetServiceVersion() string {
	return Version
}

// GetServiceStatus returns the singleton service status
// This is the factory method to access the global instance
func GetServiceStatus() *ServiceStatus {
    // Initialize on first access using sync.Once (thread-safe)
    once.Do(func() {
        defaultStatus = NewServiceStatus(GetServiceVersion())
    })
    return defaultStatus
}

func GetCurrentStatus() CurrentStatus {
	return GetServiceStatus().GetCurrentStatus()
}
