package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/openclaw/openclaw-go/internal/gateway"
	"github.com/openclaw/openclaw-go/internal/infra"
	"github.com/openclaw/openclaw-go/internal/version"
)

var (
	gatewayPort  int
	gatewayBind  string
	gatewayToken string
	gatewayTLS   bool
)

var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Manage the OpenClaw gateway server",
}

var gatewayRunCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the gateway server",
	RunE: func(cmd *cobra.Command, args []string) error {
		if gatewayPort == 0 {
			gatewayPort = infra.DefaultGatewayPort
		}
		if err := infra.EnsurePortAvailable(gatewayPort); err != nil {
			return fmt.Errorf("port %d is not available: %w", gatewayPort, err)
		}
		bindMode := gateway.GatewayBindLoopback
		switch gatewayBind {
		case "lan":
			bindMode = gateway.GatewayBindLAN
		case "loopback":
			bindMode = gateway.GatewayBindLoopback
		case "auto":
			bindMode = gateway.GatewayBindAuto
		case "tailnet":
			bindMode = gateway.GatewayBindTailnet
		}
		opts := gateway.GatewayServerOptions{
			BindMode:   bindMode,
			Token:      gatewayToken,
			TLSEnabled: gatewayTLS,
			Version:    version.Version,
		}
		srv, err := gateway.StartGatewayServer(gatewayPort, opts)
		if err != nil {
			return fmt.Errorf("failed to start gateway: %w", err)
		}
		fmt.Printf("Gateway running on port %d\n", gatewayPort)

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Println("\nShutting down gateway...")
		return srv.Close()
	},
}

var gatewayStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show gateway status",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Gateway status: checking...")
		return nil
	},
}

var gatewayStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the gateway server",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Stopping gateway...")
		return nil
	},
}

func init() {
	gatewayRunCmd.Flags().IntVar(&gatewayPort, "port", 0, fmt.Sprintf("port to listen on (default: %d)", infra.DefaultGatewayPort))
	gatewayRunCmd.Flags().StringVar(&gatewayBind, "bind", "loopback", "bind mode (auto|lan|loopback|tailnet)")
	gatewayRunCmd.Flags().StringVar(&gatewayToken, "token", "", "auth token")
	gatewayRunCmd.Flags().BoolVar(&gatewayTLS, "tls", false, "enable TLS")

	gatewayCmd.AddCommand(gatewayRunCmd, gatewayStatusCmd, gatewayStopCmd)
}
