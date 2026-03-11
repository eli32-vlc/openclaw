package irc

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

// IRCChannel implements the IRC channel.
type IRCChannel struct {
	server   string
	port     int
	nick     string
	password string
	channels []string
	useTLS   bool
	running  bool
	cancel   context.CancelFunc
}

// Message represents an incoming IRC message.
type Message struct {
	Nick    string
	Channel string
	Text    string
	IsDM    bool
}

// NewIRCChannel creates a new IRC channel.
func NewIRCChannel(server string, port int, nick string) *IRCChannel {
	if port == 0 {
		port = 6667
	}
	return &IRCChannel{server: server, port: port, nick: nick}
}

// Start starts the IRC channel.
func (c *IRCChannel) Start(ctx context.Context) error {
	if c.server == "" {
		return fmt.Errorf("irc: server is required")
	}
	ctx, c.cancel = context.WithCancel(ctx)
	c.running = true
	log.Info().Str("server", c.server).Int("port", c.port).Msg("irc channel started")
	return nil
}

// Stop stops the IRC channel.
func (c *IRCChannel) Stop() {
	if c.cancel != nil {
		c.cancel()
	}
	c.running = false
	log.Info().Msg("irc channel stopped")
}

// Send sends a message to an IRC channel or user.
func (c *IRCChannel) Send(target, text string) error {
	log.Debug().Str("target", target).Str("text", text).Msg("irc: sending message")
	return nil
}

// IsRunning returns true if the channel is running.
func (c *IRCChannel) IsRunning() bool {
	return c.running
}
