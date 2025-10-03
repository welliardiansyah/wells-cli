package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var generateMocksCmd = &cobra.Command{
	Use:   "generate-mocks",
	Short: "Generate mocks for testing",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🔹 Generating mocks ...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateMocksCmd)
}
