package cmd

import (
	"github.com/spf13/cobra"

	"github.com/openclaw/openclaw-go/internal/tui"
)

var (
	tuiURL          string
	tuiToken        string
	tuiPassword     string
	tuiSession      string
	tuiDeliver      string
	tuiThinking     string
	tuiMessage      string
	tuiTimeoutMs    int
	tuiHistoryLimit int
	tuiPort         int
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch the interactive TUI chat",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := tui.TUIOptions{
			URL:          tuiURL,
			Token:        tuiToken,
			Password:     tuiPassword,
			SessionKey:   tuiSession,
			Deliver:      tuiDeliver,
			Thinking:     tuiThinking,
			Message:      tuiMessage,
			TimeoutMs:    tuiTimeoutMs,
			HistoryLimit: tuiHistoryLimit,
			GatewayPort:  tuiPort,
		}
		return tui.RunTUI(opts)
	},
}

func init() {
	tuiCmd.Flags().StringVar(&tuiURL, "url", "", "gateway WebSocket URL")
	tuiCmd.Flags().StringVar(&tuiToken, "token", "", "auth token")
	tuiCmd.Flags().StringVar(&tuiPassword, "password", "", "auth password")
	tuiCmd.Flags().StringVar(&tuiSession, "session", "", "session key")
	tuiCmd.Flags().StringVar(&tuiDeliver, "deliver", "", "deliver assistant replies to channel")
	tuiCmd.Flags().StringVar(&tuiThinking, "thinking", "", "thinking level (low|medium|high)")
	tuiCmd.Flags().StringVar(&tuiMessage, "message", "", "initial message to send")
	tuiCmd.Flags().IntVar(&tuiTimeoutMs, "timeout-ms", 0, "agent timeout in milliseconds")
	tuiCmd.Flags().IntVar(&tuiHistoryLimit, "history-limit", 200, "max history entries")
	tuiCmd.Flags().IntVar(&tuiPort, "port", 18789, "gateway port")
}
