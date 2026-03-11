package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var memoryCmd = &cobra.Command{
	Use:   "memory",
	Short: "Manage agent memory",
}

var memorySearchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search memory",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Searching memory for: %q\n", args[0])
		fmt.Println("Memory search not yet implemented.")
		return nil
	},
}

var memoryListCmd = &cobra.Command{
	Use:   "list",
	Short: "List memories",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Memory listing not yet implemented.")
		return nil
	},
}

func init() {
	memoryCmd.AddCommand(memorySearchCmd, memoryListCmd)
}
