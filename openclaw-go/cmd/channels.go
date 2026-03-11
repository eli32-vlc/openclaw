package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/openclaw/openclaw-go/internal/channels"
	"github.com/openclaw/openclaw-go/internal/terminal"
)

var (
	channelsProbe bool
	channelsAll   bool
)

var channelsCmd = &cobra.Command{
	Use:   "channels",
	Short: "Manage messaging channels",
}

var channelsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available channels",
	RunE: func(cmd *cobra.Command, args []string) error {
		cols := []terminal.TableColumn{
			{Header: "ID", Width: 12},
			{Header: "Label", Width: 15},
			{Header: "Docs", Width: 25},
		}
		var rows [][]string
		for _, id := range channels.ChatChannelOrder {
			meta, ok := channels.GetChannelMeta(id)
			if !ok {
				continue
			}
			rows = append(rows, []string{
				string(meta.ID),
				meta.Label,
				"https://docs.openclaw.ai" + meta.DocsPath,
			})
		}
		terminal.PrintTable(cols, rows)
		return nil
	},
}

var channelsStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show channel statuses",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Channel statuses:")
		for _, id := range channels.ChatChannelOrder {
			meta, ok := channels.GetChannelMeta(id)
			if !ok {
				continue
			}
			fmt.Printf("  %-12s %s\n", string(id), terminal.Muted(meta.Label))
		}
		return nil
	},
}

var channelsAddCmd = &cobra.Command{
	Use:   "add <type>",
	Short: "Add a channel configuration",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		channelType := args[0]
		meta, ok := channels.GetChannelMeta(channels.ChatChannelID(channelType))
		if !ok {
			return fmt.Errorf("unknown channel type: %s", channelType)
		}
		fmt.Printf("Adding %s channel...\n", meta.Label)
		fmt.Printf("See docs: https://docs.openclaw.ai%s\n", meta.DocsPath)
		return nil
	},
}

func init() {
	channelsStatusCmd.Flags().BoolVar(&channelsProbe, "probe", false, "probe channel connections")
	channelsStatusCmd.Flags().BoolVar(&channelsAll, "all", false, "show all channels including disabled")
	channelsCmd.AddCommand(channelsListCmd, channelsStatusCmd, channelsAddCmd)
}
