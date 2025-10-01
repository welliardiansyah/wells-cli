// cmd/version.go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Tampilkan versi CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wells-cli v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
