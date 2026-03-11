package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var onboardCmd = &cobra.Command{
	Use:   "onboard",
	Short: "Interactive onboarding wizard",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Welcome to OpenClaw!")
		fmt.Println()
		fmt.Println("This wizard will help you set up OpenClaw.")
		fmt.Println("Interactive onboarding not yet fully implemented.")
		fmt.Println()
		fmt.Println("To get started:")
		fmt.Println("  1. Configure a provider: openclaw models auth")
		fmt.Println("  2. Start the gateway:    openclaw gateway run")
		fmt.Println("  3. Launch the TUI:       openclaw tui")
		return nil
	},
}
