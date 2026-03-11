package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/openclaw/openclaw-go/internal/agents"
	"github.com/openclaw/openclaw-go/internal/terminal"
)

var modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "Manage AI models",
}

var modelsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available models",
	RunE: func(cmd *cobra.Command, args []string) error {
		cols := []terminal.TableColumn{
			{Header: "Provider", Width: 15},
			{Header: "Model", Width: 35},
			{Header: "Default", Width: 8},
		}
		var rows [][]string
		for _, p := range agents.KnownProviders {
			for i, m := range p.Models() {
				isDefault := ""
				if i == 0 {
					isDefault = "✓"
				}
				rows = append(rows, []string{p.Name(), m, isDefault})
			}
		}
		terminal.PrintTable(cols, rows)
		return nil
	},
}

var modelsSetCmd = &cobra.Command{
	Use:   "set <provider> <model>",
	Short: "Set the default model for a provider",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Setting default model for %s to %s\n", args[0], args[1])
		return nil
	},
}

var modelsAuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with AI providers",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Provider authentication wizard not yet implemented.")
		return nil
	},
}

func init() {
	modelsCmd.AddCommand(modelsListCmd, modelsSetCmd, modelsAuthCmd)
}
