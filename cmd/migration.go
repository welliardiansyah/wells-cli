package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrationPostgresCmd = &cobra.Command{
	Use:   "migration-postgres [name]",
	Short: "Generate migration untuk Postgres",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fmt.Println("🔹 Membuat migration Postgres:", name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(migrationPostgresCmd)
}
