package scaffold

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateProject(t *testing.T) {
	tmpDir := t.TempDir()
	projectName := filepath.Join(tmpDir, "testproject")

	err := CreateProject(projectName)
	if err != nil {
		t.Fatalf("CreateProject gagal: %v", err)
	}

	expectedFolders := []string{
		"application/dtos",
		"application/mappers",
		"application/usecases",
		"domain/entities",
		"domain/migrate",
		"domain/repository",
		"interfaces/http",
		"infrastructure/config",
		"infrastructure/database",
		"infrastructure/middleware",
		"infrastructure/persistence",
		"infrastructure/redis",
		"response",
		"util",
	}

	for _, folder := range expectedFolders {
		path := filepath.Join(projectName, folder)
		if fi, err := os.Stat(path); err != nil || !fi.IsDir() {
			t.Fatalf("Folder %s tidak ditemukan atau bukan directory", folder)
		}
	}

	expectedFiles := []string{
		"main.go",
		"go.mod",
		"interfaces/http/server.go",
	}

	for _, file := range expectedFiles {
		path := filepath.Join(projectName, file)
		if fi, err := os.Stat(path); os.IsNotExist(err) || fi.IsDir() {
			t.Fatalf("File %s tidak ditemukan", file)
		}
	}
}
