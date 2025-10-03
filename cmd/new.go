package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/welliardiansyah/wells-cli/internal/scaffold"
)

var newCmd = &cobra.Command{
	Use:   "new [project name]",
	Short: "Generate a new project",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		fmt.Printf("Creating a new project: %s\n", projectName)
		return scaffold.CreateProject(projectName)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
