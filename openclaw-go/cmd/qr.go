package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var qrCmd = &cobra.Command{
	Use:   "qr",
	Short: "Show QR codes for channel pairing",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("QR code display not yet implemented.")
		fmt.Println("Use 'openclaw channels add whatsapp' to pair WhatsApp.")
		return nil
	},
}
