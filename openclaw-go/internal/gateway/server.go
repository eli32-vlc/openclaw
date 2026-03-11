package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

// GatewayBindMode controls how the gateway binds to network interfaces.
type GatewayBindMode string

const (
	GatewayBindAuto     GatewayBindMode = "auto"
	GatewayBindLAN      GatewayBindMode = "lan"
	GatewayBindLoopback GatewayBindMode = "loopback"
	GatewayBindCustom   GatewayBindMode = "custom"
	GatewayBindTailnet  GatewayBindMode = "tailnet"
)

// GatewayServerOptions configures the gateway server.
type GatewayServerOptions struct {
	BindMode   GatewayBindMode
	Host       string
	Token      string
	Password   string
	TLSEnabled bool
	CertPath   string
	KeyPath    string
	ControlUI  bool
	Version    string
}

// GatewayServer is the running gateway server.
type GatewayServer struct {
	opts        GatewayServerOptions
	httpServer  *http.Server
	upgrader    websocket.Upgrader
	clients     map[*websocket.Conn]bool
	mu          sync.RWMutex
	startTime   time.Time
	rateLimiter *AuthRateLimiter
	channels    map[string]ChannelHealth
}

// WSMessage is a message sent over WebSocket.
type WSMessage struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

// StartGatewayServer starts the gateway server on the given port.
func StartGatewayServer(port int, opts GatewayServerOptions) (*GatewayServer, error) {
	host := opts.Host
	if host == "" {
		switch opts.BindMode {
		case GatewayBindLoopback:
			host = "127.0.0.1"
		case GatewayBindLAN:
			host = ""
		default:
			host = "127.0.0.1"
		}
	}
	addr := fmt.Sprintf("%s:%d", host, port)

	srv := &GatewayServer{
		opts:        opts,
		clients:     make(map[*websocket.Conn]bool),
		startTime:   time.Now(),
		rateLimiter: NewAuthRateLimiter(10, time.Minute),
		channels:    make(map[string]ChannelHealth),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", srv.handleHealth)
	mux.HandleFunc("/ws", srv.handleWS)
	mux.HandleFunc("/", srv.handleRoot)

	srv.httpServer = &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	go func() {
		log.Info().Str("addr", addr).Msg("gateway server listening")
		if err := srv.httpServer.Serve(ln); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("gateway server error")
		}
	}()

	return srv, nil
}

// Close shuts down the gateway server.
func (s *GatewayServer) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.mu.Lock()
	for conn := range s.clients {
		conn.Close()
	}
	s.mu.Unlock()
	return s.httpServer.Shutdown(ctx)
}

func (s *GatewayServer) handleHealth(w http.ResponseWriter, r *http.Request) {
	health := GetHealthStatus(s.opts.Version, s.startTime, s.channels)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

func (s *GatewayServer) handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"name":    "openclaw-gateway",
		"version": s.opts.Version,
	})
}

func (s *GatewayServer) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error().Err(err).Msg("websocket upgrade failed")
		return
	}
	s.mu.Lock()
	s.clients[conn] = true
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.clients, conn)
		s.mu.Unlock()
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		var wsMsg WSMessage
		if err := json.Unmarshal(msg, &wsMsg); err != nil {
			continue
		}
		s.handleWSMessage(conn, wsMsg)
	}
}

func (s *GatewayServer) handleWSMessage(conn *websocket.Conn, msg WSMessage) {
	switch msg.Type {
	case "ping":
		resp, _ := json.Marshal(WSMessage{Type: "pong"})
		conn.WriteMessage(websocket.TextMessage, resp)
	default:
		log.Debug().Str("type", msg.Type).Msg("unhandled ws message type")
	}
}

// Broadcast sends a message to all connected WebSocket clients.
func (s *GatewayServer) Broadcast(msg WSMessage) {
	data, err := json.Marshal(msg)
	if err != nil {
		return
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	for conn := range s.clients {
		conn.WriteMessage(websocket.TextMessage, data)
	}
}
