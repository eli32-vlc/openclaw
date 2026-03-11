package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/openclaw/openclaw-go/internal/version"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update OpenClaw",
}

var updateCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for updates",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Current version: %s\n", version.Version)
		fmt.Println("Update check not yet implemented.")
		return nil
	},
}

var updateRunCmd = &cobra.Command{
	Use:   "run",
	Short: "Update to the latest version",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Updating OpenClaw...")
		fmt.Println("Auto-update not yet implemented. Please use: go install github.com/openclaw/openclaw-go@latest")
		return nil
	},
}

func init() {
	updateCmd.AddCommand(updateCheckCmd, updateRunCmd)
	// Default action is check
	updateCmd.RunE = updateCheckCmd.RunE
}
