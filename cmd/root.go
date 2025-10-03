package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wells",
	Short: "Wells CLI - generate project and files",
	Long:  `Wells CLI is a developer tool to generate Go project, modules, and unit tests.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()
}
