package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/openclaw/openclaw-go/internal/config"
	"github.com/openclaw/openclaw-go/internal/sessions"
	"github.com/openclaw/openclaw-go/internal/terminal"
)

var sessionsCmd = &cobra.Command{
	Use:   "sessions",
	Short: "Manage sessions",
}

var sessionsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all sessions",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := sessions.LoadSessionStore(config.SessionsPath())
		if err != nil {
			return err
		}
		list := store.List()
		if len(list) == 0 {
			fmt.Println(terminal.Muted("No sessions found."))
			return nil
		}
		cols := []terminal.TableColumn{
			{Header: "Key", Width: 40},
			{Header: "Agent", Width: 15},
			{Header: "Updated", Width: 25},
		}
		var rows [][]string
		for _, s := range list {
			rows = append(rows, []string{
				s.Key, s.AgentID, s.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}
		terminal.PrintTable(cols, rows)
		return nil
	},
}

var sessionsShowCmd = &cobra.Command{
	Use:   "show <key>",
	Short: "Show a session",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := sessions.LoadSessionStore(config.SessionsPath())
		if err != nil {
			return err
		}
		sess, ok := store.Get(args[0])
		if !ok {
			return fmt.Errorf("session %q not found", args[0])
		}
		fmt.Printf("Key:     %s\n", sess.Key)
		fmt.Printf("Agent:   %s\n", sess.AgentID)
		fmt.Printf("Created: %s\n", sess.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("Updated: %s\n", sess.UpdatedAt.Format("2006-01-02 15:04:05"))
		return nil
	},
}

var sessionsCleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Clean up old sessions",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Session cleanup not yet implemented.")
		return nil
	},
}

func init() {
	sessionsCmd.AddCommand(sessionsListCmd, sessionsShowCmd, sessionsCleanupCmd)
}
