package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateModel(name string) error {
	entityName := strings.Title(name)
	filePath := filepath.Join("domain", "entities", fmt.Sprintf("%s.go", name))

	content := fmt.Sprintf(`package entities

type %s struct {
	ID   uint   `+"`json:\"id\" gorm:\"primaryKey\"`"+`
	Name string `+"`json:\"name\"`"+`
}
`, entityName)

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return err
	}

	fmt.Println("✅ Model", entityName, "has been created at", filePath)
	return nil
}
