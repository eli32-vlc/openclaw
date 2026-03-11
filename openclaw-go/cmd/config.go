package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/openclaw/openclaw-go/internal/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage OpenClaw configuration",
}

var configGetCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a config value",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}
		val, err := config.GetConfigValue(cfg, args[0])
		if err != nil {
			return err
		}
		if val == nil {
			fmt.Printf("%s = (not set)\n", args[0])
			return nil
		}
		fmt.Printf("%s = %v\n", args[0], val)
		return nil
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a config value",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}
		if err := config.SetConfigValue(cfg, args[0], args[1]); err != nil {
			return err
		}
		return config.Save(cfg)
	},
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the full config",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(cfg)
	},
}

var configValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the config",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("config load error: %w", err)
		}
		if err := config.Validate(cfg); err != nil {
			return fmt.Errorf("config validation error: %w", err)
		}
		fmt.Println("Config is valid.")
		return nil
	},
}

func init() {
	configCmd.AddCommand(configGetCmd, configSetCmd, configShowCmd, configValidateCmd)
}
