package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pairingCmd = &cobra.Command{
	Use:   "pairing",
	Short: "Manage channel pairings",
}

var pairingAddCmd = &cobra.Command{
	Use:   "add <channel>",
	Short: "Add a pairing for a channel",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Adding pairing for channel: %s\n", args[0])
		fmt.Println("Pairing wizard not yet implemented.")
		return nil
	},
}

var pairingListCmd = &cobra.Command{
	Use:   "list",
	Short: "List active pairings",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Pairings: not yet implemented")
		return nil
	},
}

var pairingRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a pairing",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Pairing removal not yet implemented.")
		return nil
	},
}

func init() {
	pairingCmd.AddCommand(pairingAddCmd, pairingListCmd, pairingRemoveCmd)
}
