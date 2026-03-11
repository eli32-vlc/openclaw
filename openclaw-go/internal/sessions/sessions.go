package sessions

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Session holds state for a single session.
type Session struct {
	Key       string                 `json:"key"`
	AgentID   string                 `json:"agentId,omitempty"`
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// SessionStore manages all sessions.
type SessionStore struct {
	mu       sync.RWMutex
	path     string
	sessions map[string]*Session
}

// LoadSessionStore loads sessions from the given path.
func LoadSessionStore(path string) (*SessionStore, error) {
	store := &SessionStore{
		path:     path,
		sessions: make(map[string]*Session),
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return store, nil
		}
		return nil, err
	}
	var sessions []*Session
	if err := json.Unmarshal(data, &sessions); err != nil {
		return nil, err
	}
	for _, s := range sessions {
		store.sessions[s.Key] = s
	}
	return store, nil
}

// Save persists the session store to disk.
func (s *SessionStore) Save() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sessions := make([]*Session, 0, len(s.sessions))
	for _, sess := range s.sessions {
		sessions = append(sessions, sess)
	}
	data, err := json.MarshalIndent(sessions, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(s.path), 0o700); err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0o600)
}

// Get retrieves a session by key.
func (s *SessionStore) Get(key string) (*Session, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sess, ok := s.sessions[key]
	return sess, ok
}

// Set stores a session.
func (s *SessionStore) Set(sess *Session) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sess.UpdatedAt = time.Now()
	s.sessions[sess.Key] = sess
}

// Delete removes a session.
func (s *SessionStore) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, key)
}

// List returns all sessions.
func (s *SessionStore) List() []*Session {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sessions := make([]*Session, 0, len(s.sessions))
	for _, sess := range s.sessions {
		sessions = append(sessions, sess)
	}
	return sessions
}

// DeriveSessionKey creates a session key from components.
func DeriveSessionKey(channel, chatType, peerID string) string {
	return channel + ":" + chatType + ":" + peerID
}

// ResolveSessionKey resolves the effective session key.
func ResolveSessionKey(agentID, requestKey string) string {
	if requestKey == "" {
		requestKey = "main"
	}
	return "agent:" + agentID + ":" + requestKey
}
