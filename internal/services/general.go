package services

import (
	"sync"
	"time"
)

type ServiceStatus struct {
	mu        sync.RWMutex
	startTime time.Time
	ready     bool
	version   string
}

func NewServiceStatus() *ServiceStatus {
	return &ServiceStatus{
		startTime: time.Now(),
		ready:     false,
		version:   "",
	}
}

func (s *ServiceStatus) SetReady(ready bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ready = ready
}

type CurrentStatus struct {
	Timestamp string
	StartTime string
	UpTime    float64
	Version   string
	Ready     bool
}

// GetStatus returns the current service status
func (s *ServiceStatus) GetCurrentStatus() CurrentStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return CurrentStatus{
		Timestamp: time.Now().Format(time.DateTime),
		StartTime: s.startTime.Format(time.DateTime),
		Version:   s.version,
		UpTime:    s.Uptime(),
		Ready:     s.ready,
	}
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
var (
	defaultStatus *ServiceStatus
	once          sync.Once
)

func GetServiceVersion() string {
	return GetServiceStatus().version
}

func GetCurrentStatus() CurrentStatus {
	return GetServiceStatus().GetCurrentStatus()
}

func GetServiceStatus() *ServiceStatus {
	return defaultStatus
}

func SetupService(version string) {
	status := GetServiceStatus()
	status.SetReady(true)
	status.version = version
}

func init() {
	once.Do(func() {
		defaultStatus = NewServiceStatus()
	})
}
