package cmd

import (
	"github.com/spf13/cobra"
	"github.com/welliardiansyah/wells-cli/internal/scaffold"
)

var modelCmd = &cobra.Command{
	Use:   "model [name]",
	Short: "Generate a model/entity",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		scaffold.LogInfo("Creating model %s ...", name)
		if err := scaffold.CreateModel(name); err != nil {
			scaffold.LogError("Failed to create model", err)
			return err
		}
		scaffold.LogSuccess("Model %s successfully created", name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
}
