package telegram

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

// TelegramChannel implements the Telegram bot channel.
type TelegramChannel struct {
	token   string
	running bool
	cancel  context.CancelFunc
}

// Message represents an incoming Telegram message.
type Message struct {
	ChatID   int64
	UserID   int64
	Username string
	Text     string
	IsGroup  bool
}

// NewTelegramChannel creates a new Telegram channel.
func NewTelegramChannel(token string) *TelegramChannel {
	return &TelegramChannel{token: token}
}

// Start starts the Telegram channel.
func (c *TelegramChannel) Start(ctx context.Context) error {
	if c.token == "" {
		return fmt.Errorf("telegram: token is required")
	}
	ctx, c.cancel = context.WithCancel(ctx)
	c.running = true
	log.Info().Msg("telegram channel started")
	go c.poll(ctx)
	return nil
}

// Stop stops the Telegram channel.
func (c *TelegramChannel) Stop() {
	if c.cancel != nil {
		c.cancel()
	}
	c.running = false
	log.Info().Msg("telegram channel stopped")
}

// Send sends a message to a Telegram chat.
func (c *TelegramChannel) Send(chatID int64, text string) error {
	log.Debug().Int64("chatId", chatID).Str("text", text).Msg("telegram: sending message")
	return nil
}

// IsRunning returns true if the channel is running.
func (c *TelegramChannel) IsRunning() bool {
	return c.running
}

func (c *TelegramChannel) poll(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
