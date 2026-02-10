package entropy

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func Analyze(root string) (map[string]float64, error) {
	results := map[string]float64{}

	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || !strings.HasSuffix(path, ".go") {
			return nil
		}

		fs := token.NewFileSet()
		file, err := parser.ParseFile(fs, path, nil, parser.ImportsOnly)
		if err != nil {
			return nil
		}

		m := Metrics{
			FanOut: len(file.Imports),
		}

		for _, imp := range file.Imports {
			if strings.Contains(imp.Path.Value, "infrastructure") &&
				strings.Contains(path, "domain") {
				m.LayerViolations++
			}
		}

		results[path] = m.Score()
		return nil
	})

	return results, nil
}
