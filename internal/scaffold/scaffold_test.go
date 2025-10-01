package scaffold

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyTemplate(t *testing.T) {
	tmp := t.TempDir()
	opts := Options{ModuleName: "github.com/example/demo"}
	if err := CopyTemplate(tmp, opts); err != nil {
		t.Fatalf("CopyTemplate failed: %v", err)
	}

	want := []string{
		"go.mod",
		"main.go",
		"../../templates/wells-go/application/usecase.go",
		"domain/entity.go",
		"interfaces/http/handler.go",
	}
	for _, p := range want {
		if _, err := os.Stat(filepath.Join(tmp, p)); os.IsNotExist(err) {
			t.Fatalf("expected file missing: %s", p)
		}
	}
}
