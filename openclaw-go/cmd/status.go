package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/openclaw/openclaw-go/internal/channels"
	"github.com/openclaw/openclaw-go/internal/config"
	"github.com/openclaw/openclaw-go/internal/daemon"
	"github.com/openclaw/openclaw-go/internal/terminal"
	"github.com/openclaw/openclaw-go/internal/version"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show overall system status",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("OpenClaw %s\n\n", version.Version)

		// Gateway status
		fmt.Printf("%-20s", "Gateway:")
		gStatus := daemon.GetStatus()
		if gStatus == daemon.DaemonStatusRunning {
			fmt.Println(terminal.Success("running"))
		} else {
			fmt.Println(terminal.Muted("not running"))
		}

		// Config
		fmt.Printf("%-20s%s\n", "Config:", config.ConfigPath())

		// Channels
		fmt.Printf("\n%s\n", terminal.Header("Channels"))
		cfg, err := config.Load()
		if err != nil {
			fmt.Println(terminal.Error("Failed to load config: " + err.Error()))
		} else {
			for _, id := range channels.ChatChannelOrder {
				meta, ok := channels.GetChannelMeta(id)
				if !ok {
					continue
				}
				enabled := isChannelEnabled(cfg, id)
				status := terminal.Muted("disabled")
				if enabled {
					status = terminal.Success("enabled")
				}
				fmt.Printf("  %-15s %s\n", meta.Label, status)
			}
		}
		return nil
	},
}

func isChannelEnabled(cfg *config.OpenClawConfig, id channels.ChatChannelID) bool {
	if cfg.Channels == nil {
		return false
	}
	switch id {
	case channels.ChannelTelegram:
		return cfg.Channels.Telegram != nil && cfg.Channels.Telegram.Enabled != nil && *cfg.Channels.Telegram.Enabled
	case channels.ChannelWhatsApp:
		return cfg.Channels.WhatsApp != nil && cfg.Channels.WhatsApp.Enabled != nil && *cfg.Channels.WhatsApp.Enabled
	case channels.ChannelDiscord:
		return cfg.Channels.Discord != nil && cfg.Channels.Discord.Enabled != nil && *cfg.Channels.Discord.Enabled
	case channels.ChannelSlack:
		return cfg.Channels.Slack != nil && cfg.Channels.Slack.Enabled != nil && *cfg.Channels.Slack.Enabled
	case channels.ChannelSignal:
		return cfg.Channels.Signal != nil && cfg.Channels.Signal.Enabled != nil && *cfg.Channels.Signal.Enabled
	default:
		return false
	}
}
