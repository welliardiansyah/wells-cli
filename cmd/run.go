package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Jalankan project (alias go run main.go)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat("main.go"); os.IsNotExist(err) {
			return fmt.Errorf("main.go tidak ditemukan di direktori saat ini")
		}

		run := exec.Command("go", "run", "main.go")
		run.Stdout = os.Stdout
		run.Stderr = os.Stderr
		run.Stdin = os.Stdin

		fmt.Println("🚀 Menjalankan project dengan go run main.go ...")
		return run.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
