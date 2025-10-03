package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var inboundMessageRabbitCmd = &cobra.Command{
	Use:   "inbound-message-rabbitmq [name]",
	Short: "Generate inbound message consumer for RabbitMQ",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🔹 Creating inbound message consumer:", args[0])
		return nil
	},
}

var inboundCommandCmd = &cobra.Command{
	Use:   "inbound-command [name]",
	Short: "Generate inbound command handler",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🔹 Creating inbound command handler:", args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(inboundMessageRabbitCmd)
	rootCmd.AddCommand(inboundCommandCmd)
}
