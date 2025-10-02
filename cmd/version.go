package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Tampilkan versi",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wells go v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
