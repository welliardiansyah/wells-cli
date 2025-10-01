package scaffold

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyTemplate(t *testing.T) {
	tmp := t.TempDir()

	err := CopyTemplate(tmp)
	if err != nil {
		t.Fatalf("CopyTemplate gagal: %v", err)
	}

	if _, err := os.Stat(filepath.Join(tmp, "main.go")); os.IsNotExist(err) {
		t.Fatal("File main.go tidak ditemukan")
	}

	if _, err := os.Stat(filepath.Join(tmp, "README.md")); os.IsNotExist(err) {
		t.Fatal("File README.md tidak ditemukan")
	}
}
