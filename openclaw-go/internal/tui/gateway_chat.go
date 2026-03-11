package tui

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

// GatewayConnectionOptions holds options for connecting to the gateway.
type GatewayConnectionOptions struct {
	URL      string
	Token    string
	Password string
}

// ChatSendOptions holds options for sending a chat message.
type ChatSendOptions struct {
	Message    string
	SessionKey string
	AgentID    string
	Thinking   string
	Deliver    string
	TimeoutMs  int
}

// GatewayEventType identifies the type of a gateway event.
type GatewayEventType string

const (
	GatewayEventTypeChat        GatewayEventType = "chat"
	GatewayEventTypeStream      GatewayEventType = "stream"
	GatewayEventTypeToolStart   GatewayEventType = "tool_start"
	GatewayEventTypeToolEnd     GatewayEventType = "tool_end"
	GatewayEventTypeError       GatewayEventType = "error"
	GatewayEventTypeDone        GatewayEventType = "done"
	GatewayEventTypeConnected   GatewayEventType = "connected"
	GatewayEventTypeSessionList GatewayEventType = "session_list"
	GatewayEventTypeAgentList   GatewayEventType = "agent_list"
	GatewayEventTypeModelList   GatewayEventType = "model_list"
)

// GatewayEvent is an event from the gateway.
type GatewayEvent struct {
	Type    GatewayEventType       `json:"type"`
	Payload map[string]interface{} `json:"payload,omitempty"`
	Text    string                 `json:"text,omitempty"`
	Error   string                 `json:"error,omitempty"`
}

// SessionInfo holds info about a single session.
type SessionInfo struct {
	Key       string    `json:"key"`
	AgentID   string    `json:"agentId"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// AgentInfo holds info about an agent.
type AgentInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Default bool   `json:"default"`
}

// ModelChoice holds info about a model choice.
type ModelChoice struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
	Name     string `json:"name"`
}

// GatewayChatClient manages the WebSocket connection to the gateway.
type GatewayChatClient struct {
	opts           GatewayConnectionOptions
	conn           *websocket.Conn
	mu             sync.Mutex
	connected      bool
	OnEvent        func(event GatewayEvent)
	OnConnected    func()
	OnDisconnected func(err error)
}

// NewGatewayChatClient creates a new gateway chat client.
func NewGatewayChatClient(opts GatewayConnectionOptions) *GatewayChatClient {
	return &GatewayChatClient{opts: opts}
}

// Connect establishes a WebSocket connection to the gateway.
func (c *GatewayChatClient) Connect() error {
	dialer := websocket.DefaultDialer
	headers := http.Header{}
	if c.opts.Token != "" {
		headers.Set("Authorization", "Bearer "+c.opts.Token)
	}
	wsURL := c.opts.URL
	if wsURL == "" {
		wsURL = fmt.Sprintf("ws://127.0.0.1:%d/ws", 18789)
	}
	conn, _, err := dialer.Dial(wsURL, headers)
	if err != nil {
		return fmt.Errorf("failed to connect to gateway: %w", err)
	}
	c.mu.Lock()
	c.conn = conn
	c.connected = true
	c.mu.Unlock()

	if c.OnConnected != nil {
		c.OnConnected()
	}
	go c.readLoop()
	return nil
}

// Disconnect closes the WebSocket connection.
func (c *GatewayChatClient) Disconnect() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
	c.connected = false
}

// SendChat sends a chat message to the gateway.
func (c *GatewayChatClient) SendChat(opts ChatSendOptions) error {
	msg := map[string]interface{}{
		"type": "chat",
		"payload": map[string]interface{}{
			"message":    opts.Message,
			"sessionKey": opts.SessionKey,
			"agentId":    opts.AgentID,
			"thinking":   opts.Thinking,
		},
	}
	return c.sendJSON(msg)
}

// AbortChat sends an abort signal to the gateway.
func (c *GatewayChatClient) AbortChat(sessionKey string) error {
	return c.sendJSON(map[string]interface{}{
		"type":    "abort",
		"payload": map[string]interface{}{"sessionKey": sessionKey},
	})
}

// ListSessions requests the list of sessions.
func (c *GatewayChatClient) ListSessions() error {
	return c.sendJSON(map[string]interface{}{"type": "list_sessions"})
}

// ListAgents requests the list of agents.
func (c *GatewayChatClient) ListAgents() error {
	return c.sendJSON(map[string]interface{}{"type": "list_agents"})
}

// ListModels requests the list of models.
func (c *GatewayChatClient) ListModels() error {
	return c.sendJSON(map[string]interface{}{"type": "list_models"})
}

// GetStatus requests gateway status.
func (c *GatewayChatClient) GetStatus() error {
	return c.sendJSON(map[string]interface{}{"type": "status"})
}

// IsConnected returns true if the client is connected.
func (c *GatewayChatClient) IsConnected() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.connected
}

func (c *GatewayChatClient) sendJSON(v interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn == nil {
		return fmt.Errorf("not connected")
	}
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return c.conn.WriteMessage(websocket.TextMessage, data)
}

func (c *GatewayChatClient) readLoop() {
	for {
		c.mu.Lock()
		conn := c.conn
		c.mu.Unlock()
		if conn == nil {
			return
		}
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Debug().Err(err).Msg("gateway ws read error")
			c.mu.Lock()
			c.connected = false
			c.mu.Unlock()
			if c.OnDisconnected != nil {
				c.OnDisconnected(err)
			}
			return
		}
		var event GatewayEvent
		if err := json.Unmarshal(msg, &event); err != nil {
			log.Debug().Err(err).Msg("failed to parse gateway event")
			continue
		}
		if c.OnEvent != nil {
			c.OnEvent(event)
		}
	}
}

// ResolveGatewayConnection resolves the gateway connection URL from config/env.
func ResolveGatewayConnection(urlOverride, token, password string, port int) GatewayConnectionOptions {
	if urlOverride != "" {
		return GatewayConnectionOptions{URL: urlOverride, Token: token, Password: password}
	}
	wsURL := fmt.Sprintf("ws://127.0.0.1:%d/ws", port)
	return GatewayConnectionOptions{URL: wsURL, Token: token, Password: password}
}
