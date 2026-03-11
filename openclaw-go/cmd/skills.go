package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var skillsCmd = &cobra.Command{
	Use:   "skills",
	Short: "Manage skills",
}

var skillsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available skills",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Skills listing not yet implemented.")
		return nil
	},
}

var skillsAddCmd = &cobra.Command{
	Use:   "add <name>",
	Short: "Add a skill",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Adding skill: %s\n", args[0])
		return nil
	},
}

var skillsRemoveCmd = &cobra.Command{
	Use:   "remove <name>",
	Short: "Remove a skill",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Removing skill: %s\n", args[0])
		return nil
	},
}

func init() {
	skillsCmd.AddCommand(skillsListCmd, skillsAddCmd, skillsRemoveCmd)
}
