package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrationPostgresCmd = &cobra.Command{
	Use:   "migration-postgres [name]",
	Short: "Generate a migration for Postgres",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fmt.Println("🔹 Creating Postgres migration:", name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(migrationPostgresCmd)
}
