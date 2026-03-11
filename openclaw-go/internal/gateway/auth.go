package gateway

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"
)

// GatewayAuthMode controls gateway authentication mode.
type GatewayAuthMode string

const (
	GatewayAuthModeToken    GatewayAuthMode = "token"
	GatewayAuthModePassword GatewayAuthMode = "password"
	GatewayAuthModeDisabled GatewayAuthMode = "disabled"
)

// AuthRateLimiter tracks failed auth attempts.
type AuthRateLimiter struct {
	mu           sync.Mutex
	attempts     map[string][]time.Time
	maxPerWindow int
	window       time.Duration
}

// NewAuthRateLimiter creates a new rate limiter.
func NewAuthRateLimiter(maxPerWindow int, window time.Duration) *AuthRateLimiter {
	return &AuthRateLimiter{
		attempts:     make(map[string][]time.Time),
		maxPerWindow: maxPerWindow,
		window:       window,
	}
}

// Allow checks if the given key is allowed to attempt auth.
func (r *AuthRateLimiter) Allow(key string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	now := time.Now()
	cutoff := now.Add(-r.window)
	attempts := r.attempts[key]
	var recent []time.Time
	for _, t := range attempts {
		if t.After(cutoff) {
			recent = append(recent, t)
		}
	}
	if len(recent) >= r.maxPerWindow {
		r.attempts[key] = recent
		return false
	}
	r.attempts[key] = append(recent, now)
	return true
}

// GenerateToken generates a random auth token.
func GenerateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// ValidateAuth validates an auth token or password.
func ValidateAuth(provided, expected string, mode GatewayAuthMode) bool {
	if mode == GatewayAuthModeDisabled {
		return true
	}
	if expected == "" {
		return true
	}
	return provided == expected
}
