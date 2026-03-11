package config

import (
	"os"
	"path/filepath"
)

func ConfigDir() string {
	if d := os.Getenv("OPENCLAW_CONFIG_DIR"); d != "" {
		return d
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ".openclaw"
	}
	return filepath.Join(home, ".openclaw")
}

func ConfigPath() string {
	return filepath.Join(ConfigDir(), "config.json5")
}

func SessionsPath() string {
	return filepath.Join(ConfigDir(), "sessions.json")
}

func CredentialsDir() string {
	return filepath.Join(ConfigDir(), "credentials")
}

func LogsDir() string {
	return filepath.Join(ConfigDir(), "logs")
}

func WorkspaceDir() string {
	return filepath.Join(ConfigDir(), "workspace")
}

func AgentsDir() string {
	return filepath.Join(ConfigDir(), "agents")
}

func PluginsDir() string {
	return filepath.Join(ConfigDir(), "plugins")
}
