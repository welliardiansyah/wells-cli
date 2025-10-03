package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the project (default: API server)",
	RunE: func(cmd *cobra.Command, args []string) error {
		target := "cmd/api/main.go"
		if len(args) > 0 {
			switch args[0] {
			case "worker":
				target = "cmd/worker/main.go"
			case "api":
				target = "cmd/api/main.go"
			default:
				fmt.Println("[WARNING] Unknown argument, defaulting to API server")
			}
		}

		if _, err := os.Stat(target); os.IsNotExist(err) {
			return fmt.Errorf("%s not found. Make sure your project is generated correctly", target)
		}

		fmt.Println("🚀 Running project:", target)
		run := exec.Command("go", "run", target)
		run.Stdout = os.Stdout
		run.Stderr = os.Stderr
		run.Stdin = os.Stdin

		return run.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
