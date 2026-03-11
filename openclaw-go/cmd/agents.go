package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/openclaw/openclaw-go/internal/agents"
	"github.com/openclaw/openclaw-go/internal/config"
	"github.com/openclaw/openclaw-go/internal/terminal"
)

var agentsCmd = &cobra.Command{
	Use:   "agents",
	Short: "Manage AI agents",
}

var agentsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all agents",
	RunE: func(cmd *cobra.Command, args []string) error {
		mgr := agents.NewAgentManager(config.AgentsDir())
		if err := mgr.Load(); err != nil {
			return err
		}
		list := mgr.ListAgents()
		if len(list) == 0 {
			fmt.Println(terminal.Muted("No agents configured."))
			return nil
		}
		cols := []terminal.TableColumn{
			{Header: "ID", Width: 15},
			{Header: "Name", Width: 20},
			{Header: "Workspace", Width: 30},
		}
		var rows [][]string
		for _, a := range list {
			rows = append(rows, []string{a.ID, a.Name, a.Workspace})
		}
		terminal.PrintTable(cols, rows)
		return nil
	},
}

var agentsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new agent",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Adding agent (interactive wizard not yet implemented)")
		return nil
	},
}

var agentsDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete an agent",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		mgr := agents.NewAgentManager(config.AgentsDir())
		if err := mgr.Load(); err != nil {
			return err
		}
		return mgr.DeleteAgent(args[0])
	},
}

var agentsStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show agent status",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Agent status: not implemented")
		return nil
	},
}

func init() {
	agentsCmd.AddCommand(agentsListCmd, agentsAddCmd, agentsDeleteCmd, agentsStatusCmd)
}
