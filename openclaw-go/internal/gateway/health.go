package gateway

import "time"

// HealthStatus represents the health status of a component.
type HealthStatus string

const (
	HealthStatusOK      HealthStatus = "ok"
	HealthStatusWarning HealthStatus = "warning"
	HealthStatusError   HealthStatus = "error"
	HealthStatusUnknown HealthStatus = "unknown"
)

// ChannelHealth holds health info for a single channel.
type ChannelHealth struct {
	ChannelID string       `json:"channelId"`
	Status    HealthStatus `json:"status"`
	Message   string       `json:"message,omitempty"`
	LastCheck time.Time    `json:"lastCheck"`
}

// HealthResponse is the response from a health check.
type HealthResponse struct {
	Status    HealthStatus             `json:"status"`
	Version   string                   `json:"version"`
	Uptime    float64                  `json:"uptime"`
	Channels  map[string]ChannelHealth `json:"channels,omitempty"`
	Timestamp time.Time                `json:"timestamp"`
}

// GetHealthStatus returns a HealthResponse for the gateway.
func GetHealthStatus(version string, startTime time.Time, channels map[string]ChannelHealth) *HealthResponse {
	overall := HealthStatusOK
	for _, ch := range channels {
		if ch.Status == HealthStatusError {
			overall = HealthStatusError
			break
		}
		if ch.Status == HealthStatusWarning && overall != HealthStatusError {
			overall = HealthStatusWarning
		}
	}
	return &HealthResponse{
		Status:    overall,
		Version:   version,
		Uptime:    time.Since(startTime).Seconds(),
		Channels:  channels,
		Timestamp: time.Now(),
	}
}
