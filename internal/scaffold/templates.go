package scaffold

import (
	"os"
	"path/filepath"
)

func TemplatePath() string {
	if path := os.Getenv("WELLS_TEMPLATE_PATH"); path != "" {
		return path
	}

	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	return filepath.Join(exeDir, "wells-go")
}
