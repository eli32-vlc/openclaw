package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/openclaw/openclaw-go/internal/daemon"
	"github.com/openclaw/openclaw-go/internal/terminal"
)

var daemonPort int

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Manage the OpenClaw daemon",
}

var daemonStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the daemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		if daemonPort == 0 {
			daemonPort = 18789
		}
		if err := daemon.StartDaemon(daemonPort, nil); err != nil {
			return err
		}
		fmt.Printf("Daemon started on port %d\n", daemonPort)
		return nil
	},
}

var daemonStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the daemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := daemon.StopDaemon(); err != nil {
			return err
		}
		fmt.Println("Daemon stopped.")
		return nil
	},
}

var daemonStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show daemon status",
	RunE: func(cmd *cobra.Command, args []string) error {
		status := daemon.GetStatus()
		switch status {
		case daemon.DaemonStatusRunning:
			fmt.Println(terminal.Success("Daemon is running."))
		case daemon.DaemonStatusStopped:
			fmt.Println(terminal.Muted("Daemon is not running."))
		default:
			fmt.Println(terminal.Warning("Daemon status unknown."))
		}
		return nil
	},
}

var daemonInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install as system service",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("System service installation not yet implemented.")
		return nil
	},
}

func init() {
	daemonStartCmd.Flags().IntVar(&daemonPort, "port", 0, "port to run on (default: 18789)")
	daemonCmd.AddCommand(daemonStartCmd, daemonStopCmd, daemonStatusCmd, daemonInstallCmd)
}
