package auth

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

// ProfileMode identifies the credential mode for an auth profile.
type ProfileMode string

const (
	ProfileModeAPIKey ProfileMode = "api_key"
	ProfileModeOAuth  ProfileMode = "oauth"
	ProfileModeToken  ProfileMode = "token"
)

// Credential holds stored credentials.
type Credential struct {
	ProfileID    string      `json:"profileId"`
	Provider     string      `json:"provider"`
	Mode         ProfileMode `json:"mode"`
	APIKey       string      `json:"apiKey,omitempty"`
	AccessToken  string      `json:"accessToken,omitempty"`
	RefreshToken string      `json:"refreshToken,omitempty"`
	ExpiresAt    int64       `json:"expiresAt,omitempty"`
}

// AuthStore manages stored auth credentials.
type AuthStore struct {
	mu          sync.RWMutex
	dir         string
	credentials map[string]*Credential
}

// NewAuthStore creates a new auth store.
func NewAuthStore(dir string) *AuthStore {
	return &AuthStore{
		dir:         dir,
		credentials: make(map[string]*Credential),
	}
}

// Load loads credentials from disk.
func (s *AuthStore) Load() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	path := filepath.Join(s.dir, "auth-profiles.json")
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &s.credentials)
}

// Save persists credentials to disk.
func (s *AuthStore) Save() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := os.MkdirAll(s.dir, 0o700); err != nil {
		return err
	}
	path := filepath.Join(s.dir, "auth-profiles.json")
	data, err := json.MarshalIndent(s.credentials, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o600)
}

// Get retrieves a credential by profile ID.
func (s *AuthStore) Get(profileID string) (*Credential, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	cred, ok := s.credentials[profileID]
	return cred, ok
}

// Set stores a credential.
func (s *AuthStore) Set(cred *Credential) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.credentials[cred.ProfileID] = cred
}

// Delete removes a credential.
func (s *AuthStore) Delete(profileID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.credentials, profileID)
}

// List returns all credential profile IDs.
func (s *AuthStore) List() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ids := make([]string, 0, len(s.credentials))
	for id := range s.credentials {
		ids = append(ids, id)
	}
	return ids
}
