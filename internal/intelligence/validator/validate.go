package validator

import (
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"
)

func Validate(root string, rules []Rule) []string {
	var violations []string

	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		if strings.Contains(path, "/vendor/") {
			return nil
		}

		fsToken := token.NewFileSet()
		file, err := parser.ParseFile(fsToken, path, nil, parser.ImportsOnly)
		if err != nil {
			return nil
		}

		for _, rule := range rules {
			if strings.Contains(path, rule.Layer) {
				for _, imp := range file.Imports {
					for _, forbidden := range rule.CannotImport {
						if strings.Contains(imp.Path.Value, forbidden) {
							violations = append(
								violations,
								path+" must not import "+forbidden,
							)
						}
					}
				}
			}
		}
		return nil
	})

	return violations
}
