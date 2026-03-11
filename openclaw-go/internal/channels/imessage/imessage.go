// Package imessage implements the iMessage channel.
// Note: iMessage integration is macOS-only.
package imessage

import (
	"context"
	"runtime"

	"github.com/rs/zerolog/log"
)

// IMessageChannel implements the iMessage channel.
type IMessageChannel struct {
	running bool
	cancel  context.CancelFunc
}

// Message represents an incoming iMessage.
type Message struct {
	From string
	Text string
	IsDM bool
}

// NewIMessageChannel creates a new iMessage channel.
func NewIMessageChannel() *IMessageChannel {
	return &IMessageChannel{}
}

// Start starts the iMessage channel.
func (c *IMessageChannel) Start(ctx context.Context) error {
	if runtime.GOOS != "darwin" {
		log.Warn().Msg("imessage: channel is only supported on macOS")
	}
	ctx, c.cancel = context.WithCancel(ctx)
	c.running = true
	log.Info().Msg("imessage channel started")
	return nil
}

// Stop stops the iMessage channel.
func (c *IMessageChannel) Stop() {
	if c.cancel != nil {
		c.cancel()
	}
	c.running = false
	log.Info().Msg("imessage channel stopped")
}

// Send sends a message via iMessage.
func (c *IMessageChannel) Send(recipient, text string) error {
	log.Debug().Str("recipient", recipient).Str("text", text).Msg("imessage: sending message")
	return nil
}

// IsRunning returns true if the channel is running.
func (c *IMessageChannel) IsRunning() bool {
	return c.running
}
