package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	messageTo      string
	messageFrom    string
	messageSession string
	messageDeliver string
	messageChannel string
)

var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "Send and read messages",
}

var messageSendCmd = &cobra.Command{
	Use:   "send <text>",
	Short: "Send a message",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		text := args[0]
		fmt.Printf("Sending message: %q\n", text)
		if messageTo != "" {
			fmt.Printf("  To: %s\n", messageTo)
		}
		if messageFrom != "" {
			fmt.Printf("  From: %s\n", messageFrom)
		}
		if messageSession != "" {
			fmt.Printf("  Session: %s\n", messageSession)
		}
		if messageChannel != "" {
			fmt.Printf("  Channel: %s\n", messageChannel)
		}
		return nil
	},
}

var messageReadCmd = &cobra.Command{
	Use:   "read",
	Short: "Read messages",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Reading messages...")
		return nil
	},
}

func init() {
	messageSendCmd.Flags().StringVar(&messageTo, "to", "", "recipient address")
	messageSendCmd.Flags().StringVar(&messageFrom, "from", "", "sender address")
	messageSendCmd.Flags().StringVar(&messageSession, "session", "", "session key")
	messageSendCmd.Flags().StringVar(&messageDeliver, "deliver", "", "deliver to channel")
	messageSendCmd.Flags().StringVar(&messageChannel, "channel", "", "channel ID")
	messageCmd.AddCommand(messageSendCmd, messageReadCmd)
}
