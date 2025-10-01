package scaffold

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

const TemplateRepo = "https://github.com/welliardiansyah/wells-go.git"

func CopyTemplate(targetDir string) error {
	tmpDir, err := os.MkdirTemp("", "wells-template-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	fmt.Println("Mengunduh template dari GitHub...")
	cmd := exec.Command("git", "clone", "--depth=1", TemplateRepo, tmpDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("gagal clone repo: %w", err)
	}

	return filepath.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(tmpDir, path)
		if err != nil {
			return err
		}
		if relPath == ".git" || filepath.HasPrefix(relPath, ".git"+string(filepath.Separator)) {
			return nil
		}

		destPath := filepath.Join(targetDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		return err
	})
}
