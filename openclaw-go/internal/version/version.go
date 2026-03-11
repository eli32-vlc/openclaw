package version

import (
	"os"
	"strings"
)

const CurrentVersion = "2026.3.9"

const RuntimeServiceVersionFallback = "unknown"

func ResolveBinaryVersion(injected, bundled, fallback string) string {
	if v := firstNonEmpty(injected); v != "" {
		return v
	}
	if v := firstNonEmpty(bundled); v != "" {
		return v
	}
	if fallback != "" {
		return fallback
	}
	return "0.0.0"
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		trimmed := trimSpace(v)
		if trimmed != "" {
			return trimmed
		}
	}
	return ""
}

func trimSpace(s string) string {
	return strings.TrimSpace(s)
}

func ResolveRuntimeServiceVersion(fallback string) string {
	if fallback == "" {
		fallback = RuntimeServiceVersionFallback
	}
	if v := os.Getenv("OPENCLAW_VERSION"); v != "" {
		return v
	}
	if v := CurrentVersion; v != "" && v != "0.0.0" {
		return v
	}
	if v := os.Getenv("OPENCLAW_SERVICE_VERSION"); v != "" {
		return v
	}
	if v := os.Getenv("npm_package_version"); v != "" {
		return v
	}
	return fallback
}

// Version is the resolved binary version.
var Version = func() string {
	if v := os.Getenv("OPENCLAW_BUNDLED_VERSION"); v != "" {
		return ResolveBinaryVersion("", v, CurrentVersion)
	}
	return CurrentVersion
}()
