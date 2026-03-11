package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: "Manage hooks",
}

var hooksListCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured hooks",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hooks listing not yet implemented.")
		return nil
	},
}

var hooksTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Test hooks",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hook testing not yet implemented.")
		return nil
	},
}

func init() {
	hooksCmd.AddCommand(hooksListCmd, hooksTestCmd)
}
