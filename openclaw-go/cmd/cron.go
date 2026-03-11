package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "Manage cron jobs",
}

var cronListCmd = &cobra.Command{
	Use:   "list",
	Short: "List cron jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Cron jobs listing not yet implemented.")
		return nil
	},
}

var cronAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a cron job",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Cron job addition not yet implemented.")
		return nil
	},
}

var cronDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a cron job",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Deleting cron job: %s\n", args[0])
		return nil
	},
}

func init() {
	cronCmd.AddCommand(cronListCmd, cronAddCmd, cronDeleteCmd)
}
