package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var outboundDatabasePostgresCmd = &cobra.Command{
	Use:   "outbound-database-postgres [name]",
	Short: "Generate outbound database Postgres integration",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🔹 Creating outbound database Postgres:", args[0])
		return nil
	},
}

var outboundHttpCmd = &cobra.Command{
	Use:   "outbound-http [name]",
	Short: "Generate outbound HTTP integration",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🔹 Creating outbound HTTP:", args[0])
		return nil
	},
}

var outboundMessageRabbitCmd = &cobra.Command{
	Use:   "outbound-message-rabbitmq [name]",
	Short: "Generate outbound message for RabbitMQ",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🔹 Creating outbound message RabbitMQ:", args[0])
		return nil
	},
}

var outboundCacheRedisCmd = &cobra.Command{
	Use:   "outbound-cache-redis [name]",
	Short: "Generate outbound cache Redis integration",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🔹 Creating outbound cache Redis:", args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(outboundDatabasePostgresCmd)
	rootCmd.AddCommand(outboundHttpCmd)
	rootCmd.AddCommand(outboundMessageRabbitCmd)
	rootCmd.AddCommand(outboundCacheRedisCmd)
}
