package scaffold

import (
	"os"
	"path/filepath"
)

func TemplatePath() string {
	dir, _ := os.Getwd()
	return filepath.Join(dir, "wells-go")
}
