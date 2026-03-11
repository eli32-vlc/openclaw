package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/openclaw/openclaw-go/internal/config"
	"github.com/openclaw/openclaw-go/internal/daemon"
	"github.com/openclaw/openclaw-go/internal/terminal"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Run health checks",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(terminal.Header("OpenClaw Doctor"))
		fmt.Println()

		// Check config
		fmt.Print("Checking config... ")
		cfg, err := config.Load()
		if err != nil {
			fmt.Println(terminal.Error("FAIL: " + err.Error()))
		} else if err := config.Validate(cfg); err != nil {
			fmt.Println(terminal.Warning("WARN: " + err.Error()))
		} else {
			fmt.Println(terminal.Success("OK"))
		}

		// Check config dir
		fmt.Print("Checking config dir... ")
		configDir := config.ConfigDir()
		fmt.Println(terminal.Success("OK: " + configDir))

		// Check gateway
		fmt.Print("Checking gateway... ")
		status := daemon.GetStatus()
		if status == daemon.DaemonStatusRunning {
			fmt.Println(terminal.Success("running"))
		} else {
			fmt.Println(terminal.Muted("not running"))
		}

		fmt.Println()
		fmt.Println(terminal.Muted("Run 'openclaw gateway run' to start the gateway."))
		return nil
	},
}
