package plugins

import "fmt"

// Plugin defines the interface for OpenClaw plugins.
type Plugin interface {
	ID() string
	Name() string
	Version() string
	Init(config map[string]interface{}) error
	Close() error
}

// PluginRegistry manages loaded plugins.
type PluginRegistry struct {
	plugins map[string]Plugin
}

// NewPluginRegistry creates a new plugin registry.
func NewPluginRegistry() *PluginRegistry {
	return &PluginRegistry{plugins: make(map[string]Plugin)}
}

// Register registers a plugin.
func (r *PluginRegistry) Register(p Plugin) error {
	if _, exists := r.plugins[p.ID()]; exists {
		return fmt.Errorf("plugin %q already registered", p.ID())
	}
	r.plugins[p.ID()] = p
	return nil
}

// Get retrieves a plugin by ID.
func (r *PluginRegistry) Get(id string) (Plugin, bool) {
	p, ok := r.plugins[id]
	return p, ok
}

// List returns all registered plugins.
func (r *PluginRegistry) List() []Plugin {
	plugins := make([]Plugin, 0, len(r.plugins))
	for _, p := range r.plugins {
		plugins = append(plugins, p)
	}
	return plugins
}

// CloseAll closes all plugins.
func (r *PluginRegistry) CloseAll() {
	for _, p := range r.plugins {
		p.Close()
	}
}
