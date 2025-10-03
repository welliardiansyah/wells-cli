package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the project (alias for 'go run main.go')",
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat("main.go"); os.IsNotExist(err) {
			return fmt.Errorf("main.go not found in the current directory")
		}
		run := exec.Command("go", "run", "main.go")
		run.Stdout = os.Stdout
		run.Stderr = os.Stderr
		run.Stdin = os.Stdin

		fmt.Println("🚀 Running the project with 'go run main.go' ...")
		return run.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
