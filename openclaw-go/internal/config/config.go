package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// Load reads the config from the default config path.
func Load() (*OpenClawConfig, error) {
	return LoadFrom(ConfigPath())
}

// LoadFrom reads the config from the given path.
func LoadFrom(path string) (*OpenClawConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &OpenClawConfig{}, nil
		}
		return nil, err
	}
	var cfg OpenClawConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Save writes the config to the default config path.
func Save(cfg *OpenClawConfig) error {
	return SaveTo(ConfigPath(), cfg)
}

// SaveTo writes the config to the given path.
func SaveTo(path string, cfg *OpenClawConfig) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o600)
}

// Validate checks the config for errors.
func Validate(cfg *OpenClawConfig) error {
	if cfg == nil {
		return errors.New("config is nil")
	}
	return nil
}

// GetGatewayPort returns the configured gateway port or the default.
func GetGatewayPort(cfg *OpenClawConfig) int {
	if cfg.Gateway != nil && cfg.Gateway.Port != nil {
		return *cfg.Gateway.Port
	}
	return 18789
}

// GetGatewayMode returns the configured gateway bind mode or the default.
func GetGatewayMode(cfg *OpenClawConfig) GatewayBindMode {
	if cfg.Gateway != nil && cfg.Gateway.Mode != "" {
		return cfg.Gateway.Mode
	}
	return GatewayBindModeAuto
}

// GetOrCreateGateway returns the gateway config, creating it if nil.
func GetOrCreateGateway(cfg *OpenClawConfig) *GatewayConfig {
	if cfg.Gateway == nil {
		cfg.Gateway = &GatewayConfig{}
	}
	return cfg.Gateway
}

// GetOrCreateChannels returns the channels config, creating it if nil.
func GetOrCreateChannels(cfg *OpenClawConfig) *ChannelsConfig {
	if cfg.Channels == nil {
		cfg.Channels = &ChannelsConfig{}
	}
	return cfg.Channels
}

// GetOrCreateAgents returns the agents config, creating it if nil.
func GetOrCreateAgents(cfg *OpenClawConfig) *AgentsConfig {
	if cfg.Agents == nil {
		cfg.Agents = &AgentsConfig{}
	}
	return cfg.Agents
}

// SetConfigValue sets a config value by dot-separated key path.
func SetConfigValue(cfg *OpenClawConfig, key, value string) error {
	// For now, support a simple flat key set via JSON re-encoding
	raw, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(raw, &m); err != nil {
		return err
	}
	m[key] = value
	updated, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(updated, cfg)
}

// GetConfigValue retrieves a config value by key.
func GetConfigValue(cfg *OpenClawConfig, key string) (interface{}, error) {
	raw, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, err
	}
	val, ok := m[key]
	if !ok {
		return nil, nil
	}
	return val, nil
}
