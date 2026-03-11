package whatsapp

import (
	"context"

	"github.com/rs/zerolog/log"
)

// WhatsAppChannel implements the WhatsApp channel.
type WhatsAppChannel struct {
	running bool
	cancel  context.CancelFunc
}

// Message represents an incoming WhatsApp message.
type Message struct {
	From    string
	GroupID string
	Text    string
	IsDM    bool
}

// NewWhatsAppChannel creates a new WhatsApp channel.
func NewWhatsAppChannel() *WhatsAppChannel {
	return &WhatsAppChannel{}
}

// Start starts the WhatsApp channel.
func (c *WhatsAppChannel) Start(ctx context.Context) error {
	ctx, c.cancel = context.WithCancel(ctx)
	c.running = true
	log.Info().Msg("whatsapp channel started")
	return nil
}

// Stop stops the WhatsApp channel.
func (c *WhatsAppChannel) Stop() {
	if c.cancel != nil {
		c.cancel()
	}
	c.running = false
	log.Info().Msg("whatsapp channel stopped")
}

// Send sends a message via WhatsApp.
func (c *WhatsAppChannel) Send(recipient, text string) error {
	log.Debug().Str("recipient", recipient).Str("text", text).Msg("whatsapp: sending message")
	return nil
}

// IsRunning returns true if the channel is running.
func (c *WhatsAppChannel) IsRunning() bool {
	return c.running
}

// GenerateQRCode generates a QR code for pairing.
func (c *WhatsAppChannel) GenerateQRCode() (string, error) {
	return "QR code placeholder - implement with go-whatsmeow or similar", nil
}
