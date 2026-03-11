package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/openclaw/openclaw-go/internal/infra"
	"github.com/openclaw/openclaw-go/internal/version"
)

var (
	cfgFile  string
	logLevel string
	noColor  bool
)

// rootCmd is the base command.
var rootCmd = &cobra.Command{
	Use:   "openclaw",
	Short: "OpenClaw – multi-channel AI gateway",
	Long: `OpenClaw is a multi-channel AI gateway that routes messages from chat
platforms (Telegram, WhatsApp, Discord, Slack, Signal, iMessage, IRC, Google Chat,
MS Teams) to AI agents.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		infra.NormalizeEnv()
		infra.InitLogging(infra.LogLevel(logLevel))
	},
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ~/.openclaw/config.json5)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "log level (silent|fatal|error|warn|info|debug|trace)")
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable color output")

	rootCmd.Version = version.Version
	rootCmd.AddCommand(
		gatewayCmd,
		tuiCmd,
		channelsCmd,
		messageCmd,
		agentsCmd,
		configCmd,
		modelsCmd,
		sessionsCmd,
		doctorCmd,
		secretsCmd,
		onboardCmd,
		daemonCmd,
		statusCmd,
		pairingCmd,
		logsCmd,
		updateCmd,
		pluginsCmd,
		skillsCmd,
		hooksCmd,
		cronCmd,
		memoryCmd,
		setupCmd,
		qrCmd,
		resetCmd,
	)
}
