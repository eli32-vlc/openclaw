package agents

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Agent represents an OpenClaw agent.
type Agent struct {
	ID        string            `json:"id"`
	Name      string            `json:"name,omitempty"`
	Workspace string            `json:"workspace,omitempty"`
	AgentDir  string            `json:"agentDir,omitempty"`
	Model     *AgentModel       `json:"model,omitempty"`
	Skills    []string          `json:"skills,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

// AgentModel represents model configuration for an agent.
type AgentModel struct {
	Provider  string   `json:"provider,omitempty"`
	Model     string   `json:"model,omitempty"`
	Fallbacks []string `json:"fallbacks,omitempty"`
}

// AgentManager manages agents.
type AgentManager struct {
	agentsDir string
	agents    map[string]*Agent
}

// NewAgentManager creates a new agent manager.
func NewAgentManager(agentsDir string) *AgentManager {
	return &AgentManager{
		agentsDir: agentsDir,
		agents:    make(map[string]*Agent),
	}
}

// Load loads agents from the agents directory.
func (m *AgentManager) Load() error {
	entries, err := os.ReadDir(m.agentsDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		agentFile := filepath.Join(m.agentsDir, entry.Name(), "agent.json")
		data, err := os.ReadFile(agentFile)
		if err != nil {
			continue
		}
		var agent Agent
		if err := json.Unmarshal(data, &agent); err != nil {
			continue
		}
		m.agents[agent.ID] = &agent
	}
	return nil
}

// ListAgents returns all agents.
func (m *AgentManager) ListAgents() []*Agent {
	agents := make([]*Agent, 0, len(m.agents))
	for _, a := range m.agents {
		agents = append(agents, a)
	}
	return agents
}

// GetAgent retrieves an agent by ID.
func (m *AgentManager) GetAgent(id string) (*Agent, bool) {
	a, ok := m.agents[id]
	return a, ok
}

// CreateAgent creates a new agent.
func (m *AgentManager) CreateAgent(agent *Agent) error {
	if agent.ID == "" {
		return errors.New("agent ID is required")
	}
	if _, exists := m.agents[agent.ID]; exists {
		return fmt.Errorf("agent %q already exists", agent.ID)
	}
	agentDir := filepath.Join(m.agentsDir, agent.ID)
	if err := os.MkdirAll(agentDir, 0o700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(agent, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(agentDir, "agent.json"), data, 0o600); err != nil {
		return err
	}
	m.agents[agent.ID] = agent
	return nil
}

// DeleteAgent deletes an agent.
func (m *AgentManager) DeleteAgent(id string) error {
	if _, exists := m.agents[id]; !exists {
		return fmt.Errorf("agent %q not found", id)
	}
	agentDir := filepath.Join(m.agentsDir, id)
	if err := os.RemoveAll(agentDir); err != nil {
		return err
	}
	delete(m.agents, id)
	return nil
}
