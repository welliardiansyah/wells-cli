package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/welliardiansyah/wells-cli/internal/intelligence/entropy"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Analyze architecture health (entropy)",
	RunE: func(cmd *cobra.Command, args []string) error {
		results, err := entropy.Analyze(".")
		if err != nil {
			return err
		}

		if len(results) == 0 {
			fmt.Println("✅ Architecture looks healthy")
			return nil
		}

		for file, score := range results {
			if score > 0.6 {
				fmt.Printf("⚠️ High entropy: %s (%.2f)\n", file, score)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
