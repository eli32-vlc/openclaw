package discord

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

// DiscordChannel implements the Discord bot channel.
type DiscordChannel struct {
	token   string
	running bool
	cancel  context.CancelFunc
}

// Message represents an incoming Discord message.
type Message struct {
	GuildID   string
	ChannelID string
	UserID    string
	Username  string
	Text      string
	IsDM      bool
}

// NewDiscordChannel creates a new Discord channel.
func NewDiscordChannel(token string) *DiscordChannel {
	return &DiscordChannel{token: token}
}

// Start starts the Discord channel.
func (c *DiscordChannel) Start(ctx context.Context) error {
	if c.token == "" {
		return fmt.Errorf("discord: token is required")
	}
	ctx, c.cancel = context.WithCancel(ctx)
	c.running = true
	log.Info().Msg("discord channel started")
	return nil
}

// Stop stops the Discord channel.
func (c *DiscordChannel) Stop() {
	if c.cancel != nil {
		c.cancel()
	}
	c.running = false
	log.Info().Msg("discord channel stopped")
}

// Send sends a message to a Discord channel.
func (c *DiscordChannel) Send(channelID, text string) error {
	log.Debug().Str("channelId", channelID).Str("text", text).Msg("discord: sending message")
	return nil
}

// IsRunning returns true if the channel is running.
func (c *DiscordChannel) IsRunning() bool {
	return c.running
}
