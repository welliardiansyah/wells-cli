package main

import (
	"fmt"
	"os"

	"github.com/welliardiansyah/wells-cli/internal/scaffold"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: generate-test <name>")
		os.Exit(1)
	}

	// Nama entity / module yang ingin di-generate test-nya
	name := os.Args[1]

	// Module project, bisa disesuaikan sesuai go.mod
	moduleName := "github.com/welliardiansyah/wells-go"

	// Generate unit test untuk Usecase
	scaffold.GenerateUnitTest(
		name,                   // nama file / entity
		"application/usecases", // folder path
		fmt.Sprintf("%sUsecase", capitalize(name)), // structName
		moduleName,     // module name
		"usecase_test", // test type
	)

	// Generate unit test untuk Repository
	scaffold.GenerateUnitTest(
		name,
		"infrastructure/persistence",
		fmt.Sprintf("%sRepositoryGorm", capitalize(name)),
		moduleName,
		"repository_test",
	)

	// Generate unit test untuk Handler
	scaffold.GenerateUnitTest(
		name,
		"interfaces/http/"+name+"s",
		fmt.Sprintf("%sHandler", capitalize(name)),
		moduleName,
		"handler_test",
	)

	fmt.Println("✅ Unit test files generated successfully")
}

// capitalize huruf pertama jadi uppercase
func capitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	return string(str[0]-32) + str[1:]
}
