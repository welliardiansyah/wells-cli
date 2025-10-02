package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/welliardiansyah/wells-cli/internal/scaffold"
)

var modelCmd = &cobra.Command{
	Use:   "model [name]",
	Short: "Generate model/entity",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fmt.Printf("🔨 Membuat model %s ...\n", name)
		return scaffold.CreateModel(name)
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
}
