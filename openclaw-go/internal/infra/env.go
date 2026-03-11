package infra

import (
	"os"
	"strings"
)

// NormalizeEnv normalizes environment variable aliases.
func NormalizeEnv() {
	if strings.TrimSpace(os.Getenv("ZAI_API_KEY")) == "" {
		if v := strings.TrimSpace(os.Getenv("Z_AI_API_KEY")); v != "" {
			os.Setenv("ZAI_API_KEY", v)
		}
	}
}

// IsTruthyEnvValue returns true if the env value is truthy.
func IsTruthyEnvValue(value string) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}

// IsDebugMode returns true if debug mode is enabled.
func IsDebugMode() bool {
	return IsTruthyEnvValue(os.Getenv("OPENCLAW_DEBUG")) ||
		IsTruthyEnvValue(os.Getenv("DEBUG"))
}
