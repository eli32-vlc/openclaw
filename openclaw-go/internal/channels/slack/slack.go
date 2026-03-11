package slack

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

// SlackChannel implements the Slack bot channel.
type SlackChannel struct {
	botToken string
	appToken string
	running  bool
	cancel   context.CancelFunc
}

// Message represents an incoming Slack message.
type Message struct {
	TeamID    string
	ChannelID string
	UserID    string
	Username  string
	Text      string
	IsDM      bool
}

// NewSlackChannel creates a new Slack channel.
func NewSlackChannel(botToken, appToken string) *SlackChannel {
	return &SlackChannel{botToken: botToken, appToken: appToken}
}

// Start starts the Slack channel.
func (c *SlackChannel) Start(ctx context.Context) error {
	if c.botToken == "" {
		return fmt.Errorf("slack: botToken is required")
	}
	ctx, c.cancel = context.WithCancel(ctx)
	c.running = true
	log.Info().Msg("slack channel started")
	return nil
}

// Stop stops the Slack channel.
func (c *SlackChannel) Stop() {
	if c.cancel != nil {
		c.cancel()
	}
	c.running = false
	log.Info().Msg("slack channel stopped")
}

// Send sends a message to a Slack channel.
func (c *SlackChannel) Send(channelID, text string) error {
	log.Debug().Str("channelId", channelID).Str("text", text).Msg("slack: sending message")
	return nil
}

// IsRunning returns true if the channel is running.
func (c *SlackChannel) IsRunning() bool {
	return c.running
}
