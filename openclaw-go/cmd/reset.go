package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset OpenClaw state",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("WARNING: This will reset your OpenClaw state.")
		fmt.Println("Reset not yet implemented.")
		return nil
	},
}
