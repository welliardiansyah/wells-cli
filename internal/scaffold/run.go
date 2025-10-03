package scaffold

import (
	"fmt"
	"os"
	"os/exec"
)

func RunMain(args ...string) error {
	fmt.Println("🚀 Running the project with 'go run main.go' ...")
	cmd := exec.Command("go", append([]string{"run", "main.go"}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Println("[ERROR] Failed to run main.go:", err)
		return err
	}

	fmt.Println("[SUCCESS] Project successfully ran")
	return nil
}
