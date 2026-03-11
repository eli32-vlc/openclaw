package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/openclaw/openclaw-go/internal/config"
)

var pluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "Manage plugins",
}

var pluginsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed plugins",
	RunE: func(cmd *cobra.Command, args []string) error {
		pluginsDir := config.PluginsDir()
		fmt.Printf("Plugins directory: %s\n", pluginsDir)
		fmt.Println("Plugin listing not yet implemented.")
		return nil
	},
}

var pluginsInstallCmd = &cobra.Command{
	Use:   "install <name>",
	Short: "Install a plugin",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Installing plugin: %s\n", args[0])
		fmt.Println("Plugin installation not yet implemented.")
		return nil
	},
}

var pluginsRemoveCmd = &cobra.Command{
	Use:   "remove <name>",
	Short: "Remove a plugin",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Removing plugin: %s\n", args[0])
		fmt.Println("Plugin removal not yet implemented.")
		return nil
	},
}

func init() {
	pluginsCmd.AddCommand(pluginsListCmd, pluginsInstallCmd, pluginsRemoveCmd)
}
