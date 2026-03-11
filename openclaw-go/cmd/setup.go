package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Initial setup wizard",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("OpenClaw Setup Wizard")
		fmt.Println()
		fmt.Println("This wizard will help you set up OpenClaw for the first time.")
		fmt.Println()
		fmt.Println("Not yet fully implemented. Please see: https://docs.openclaw.ai")
		return nil
	},
}
