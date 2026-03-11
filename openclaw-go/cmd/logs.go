package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/openclaw/openclaw-go/internal/config"
)

var (
	logsFollow  bool
	logsLevel   string
	logsLines   int
	logsChannel string
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Show logs",
	RunE: func(cmd *cobra.Command, args []string) error {
		logsDir := config.LogsDir()
		fmt.Printf("Logs directory: %s\n", logsDir)
		if logsFollow {
			fmt.Println("Following logs (Ctrl+C to stop)...")
		} else {
			fmt.Printf("Showing last %d lines\n", logsLines)
		}
		fmt.Println("Log streaming not yet implemented.")
		return nil
	},
}

func init() {
	logsCmd.Flags().BoolVarP(&logsFollow, "follow", "f", false, "follow log output")
	logsCmd.Flags().StringVar(&logsLevel, "level", "info", "minimum log level")
	logsCmd.Flags().IntVarP(&logsLines, "lines", "n", 100, "number of lines to show")
	logsCmd.Flags().StringVar(&logsChannel, "channel", "", "filter by channel")
}
