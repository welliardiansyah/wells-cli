package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wells",
	Short: "Wells Go Framework CLI",
	Long:  "CLI untuk generate project Go boilerplate dengan Wells Go Framework",
}

func Execute() error {
	return rootCmd.Execute()
}
