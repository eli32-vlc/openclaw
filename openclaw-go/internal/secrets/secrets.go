package secrets

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

// SecretStore manages secrets.
type SecretStore struct {
	mu      sync.RWMutex
	path    string
	secrets map[string]string
}

// LoadSecretStore loads secrets from the given path.
func LoadSecretStore(path string) (*SecretStore, error) {
	store := &SecretStore{
		path:    path,
		secrets: make(map[string]string),
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return store, nil
		}
		return nil, err
	}
	if err := json.Unmarshal(data, &store.secrets); err != nil {
		return nil, err
	}
	return store, nil
}

// Save persists the secrets to disk.
func (s *SecretStore) Save() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	data, err := json.MarshalIndent(s.secrets, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(s.path), 0o700); err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0o600)
}

// ResolveSecret resolves a secret value, checking env vars and store.
func (s *SecretStore) ResolveSecret(input interface{}) string {
	switch v := input.(type) {
	case string:
		// Check if it's an env var reference
		if len(v) > 4 && v[:4] == "env:" {
			return os.Getenv(v[4:])
		}
		return v
	case map[string]interface{}:
		if secret, ok := v["secret"].(string); ok {
			return s.Get(secret)
		}
	}
	return ""
}

// Get retrieves a secret by key.
func (s *SecretStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.secrets[key]
}

// Set stores a secret.
func (s *SecretStore) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.secrets[key] = value
}

// Delete removes a secret.
func (s *SecretStore) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.secrets, key)
}

// List returns all secret keys (not values).
func (s *SecretStore) List() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	keys := make([]string, 0, len(s.secrets))
	for k := range s.secrets {
		keys = append(keys, k)
	}
	return keys
}
