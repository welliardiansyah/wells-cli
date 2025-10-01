// cmd/root.go
package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wells",
	Short: "Wells CLI Framework",
	Long:  "Wells CLI adalah framework untuk generate project Go boilerplate",
}

func Execute() error {
	return rootCmd.Execute()
}
