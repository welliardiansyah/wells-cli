package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run the HTTP server (alias for 'make http')",
	RunE: func(cmd *cobra.Command, args []string) error {
		run := exec.Command("go", "run", "main.go")
		run.Stdout = os.Stdout
		run.Stderr = os.Stderr
		run.Stdin = os.Stdin

		fmt.Println("🚀 Running the HTTP application...")
		return run.Run()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
