package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var messageCmd = &cobra.Command{
	Use:   "message [subscriber]",
	Short: "Run message consumer",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		subscriber := args[0]
		fmt.Println("📩 Menjalankan message consumer:", subscriber)

		run := exec.Command("go", "run", "main.go", "message", subscriber)
		run.Stdout = os.Stdout
		run.Stderr = os.Stderr
		run.Stdin = os.Stdin

		return run.Run()
	},
}

func init() {
	rootCmd.AddCommand(messageCmd)
}
