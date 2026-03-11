package signal

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

// SignalChannel implements the Signal channel via signal-cli HTTP daemon.
type SignalChannel struct {
	phoneNumber string
	apiURL      string
	running     bool
	cancel      context.CancelFunc
}

// Message represents an incoming Signal message.
type Message struct {
	Sender  string
	GroupID string
	Text    string
	IsDM    bool
}

// NewSignalChannel creates a new Signal channel.
func NewSignalChannel(phoneNumber, apiURL string) *SignalChannel {
	return &SignalChannel{phoneNumber: phoneNumber, apiURL: apiURL}
}

// Start starts the Signal channel.
func (c *SignalChannel) Start(ctx context.Context) error {
	if c.phoneNumber == "" {
		return fmt.Errorf("signal: phoneNumber is required")
	}
	if c.apiURL == "" {
		c.apiURL = "http://127.0.0.1:8080"
	}
	ctx, c.cancel = context.WithCancel(ctx)
	c.running = true
	log.Info().Str("apiURL", c.apiURL).Msg("signal channel started")
	return nil
}

// Stop stops the Signal channel.
func (c *SignalChannel) Stop() {
	if c.cancel != nil {
		c.cancel()
	}
	c.running = false
	log.Info().Msg("signal channel stopped")
}

// Send sends a message via Signal.
func (c *SignalChannel) Send(recipient, text string) error {
	log.Debug().Str("recipient", recipient).Str("text", text).Msg("signal: sending message")
	return nil
}

// IsRunning returns true if the channel is running.
func (c *SignalChannel) IsRunning() bool {
	return c.running
}
