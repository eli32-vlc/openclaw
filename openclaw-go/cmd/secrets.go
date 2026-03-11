package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/openclaw/openclaw-go/internal/config"
	"github.com/openclaw/openclaw-go/internal/secrets"
	"github.com/openclaw/openclaw-go/internal/terminal"
)

var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Manage secrets",
}

var secretsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List secret keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := secrets.LoadSecretStore(filepath.Join(config.CredentialsDir(), "secrets.json"))
		if err != nil {
			return err
		}
		keys := store.List()
		if len(keys) == 0 {
			fmt.Println(terminal.Muted("No secrets stored."))
			return nil
		}
		for _, k := range keys {
			fmt.Println(" ", k)
		}
		return nil
	},
}

var secretsSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a secret",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := secrets.LoadSecretStore(filepath.Join(config.CredentialsDir(), "secrets.json"))
		if err != nil {
			return err
		}
		store.Set(args[0], args[1])
		return store.Save()
	},
}

var secretsDeleteCmd = &cobra.Command{
	Use:   "delete <key>",
	Short: "Delete a secret",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := secrets.LoadSecretStore(filepath.Join(config.CredentialsDir(), "secrets.json"))
		if err != nil {
			return err
		}
		store.Delete(args[0])
		return store.Save()
	},
}

var secretsAuditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Audit secrets usage",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Secrets audit not yet implemented.")
		return nil
	},
}

func init() {
	secretsCmd.AddCommand(secretsListCmd, secretsSetCmd, secretsDeleteCmd, secretsAuditCmd)
}
