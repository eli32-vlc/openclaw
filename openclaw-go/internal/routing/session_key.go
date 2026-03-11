package routing

import (
	"fmt"
	"regexp"
	"strings"
)

const DefaultAgentID = "main"
const DefaultMainKey = "main"

var validIDRe = regexp.MustCompile(`^[a-z0-9][a-z0-9_-]{0,63}$`)
var invalidCharsRe = regexp.MustCompile(`[^a-z0-9_-]+`)
var leadingDashRe = regexp.MustCompile(`^-+`)
var trailingDashRe = regexp.MustCompile(`-+$`)

// ParsedAgentSessionKey holds the parsed components of an agent session key.
type ParsedAgentSessionKey struct {
	AgentID string
	Rest    string
	Raw     string
}

// BuildAgentMainSessionKey builds the main session key for an agent.
func BuildAgentMainSessionKey(agentID, mainKey string) string {
	if mainKey == "" {
		mainKey = DefaultMainKey
	}
	return fmt.Sprintf("agent:%s:%s", agentID, mainKey)
}

// ParseAgentSessionKey parses an agent session key.
func ParseAgentSessionKey(key string) *ParsedAgentSessionKey {
	if !strings.HasPrefix(key, "agent:") {
		return nil
	}
	parts := strings.SplitN(key[6:], ":", 2)
	if len(parts) != 2 {
		return nil
	}
	return &ParsedAgentSessionKey{
		AgentID: parts[0],
		Rest:    parts[1],
		Raw:     key,
	}
}

// NormalizeAgentID normalizes an agent ID to a safe format.
func NormalizeAgentID(value string) string {
	normalized := strings.TrimSpace(strings.ToLower(value))
	normalized = invalidCharsRe.ReplaceAllString(normalized, "-")
	normalized = leadingDashRe.ReplaceAllString(normalized, "")
	normalized = trailingDashRe.ReplaceAllString(normalized, "")
	if normalized == "" {
		return DefaultAgentID
	}
	if !validIDRe.MatchString(normalized) {
		return DefaultAgentID
	}
	return normalized
}

// ToAgentStoreSessionKey converts a request session key to a store session key.
func ToAgentStoreSessionKey(agentID, requestKey, mainKey string) string {
	raw := strings.TrimSpace(requestKey)
	if raw == "" || strings.ToLower(raw) == DefaultMainKey {
		return BuildAgentMainSessionKey(agentID, mainKey)
	}
	parsed := ParseAgentSessionKey(raw)
	if parsed != nil {
		return raw
	}
	return fmt.Sprintf("agent:%s:%s", agentID, raw)
}
