package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/welliardiansyah/wells-cli/internal/scaffold"
)

var generateTestCmd = &cobra.Command{
	Use:   "generate-test [name]",
	Short: "Generate unit test skeleton for entity",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		moduleName := "github.com/welliardiansyah/wells-go"

		scaffold.GenerateUnitTest(
			name,
			"application/usecases",
			fmt.Sprintf("%sUsecase", capitalize(name)),
			moduleName,
			"usecase_test",
		)

		scaffold.GenerateUnitTest(
			name,
			"infrastructure/persistence",
			fmt.Sprintf("%sRepositoryGorm", capitalize(name)),
			moduleName,
			"repository_test",
		)

		scaffold.GenerateUnitTest(
			name,
			fmt.Sprintf("interfaces/http/%ss", strings.ToLower(name)),
			fmt.Sprintf("%sHandler", capitalize(name)),
			moduleName,
			"handler_test",
		)

		fmt.Println("✅ Unit test files generated successfully")
	},
}

func init() {
	rootCmd.AddCommand(generateTestCmd)
}

func capitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.ToUpper(str[:1]) + str[1:]
}
