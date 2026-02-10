package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/welliardiansyah/wells-cli/internal/intelligence/validator"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate architecture rules",
	Run: func(cmd *cobra.Command, args []string) {
		rules := []validator.Rule{
			{
				Layer:        "domain",
				CannotImport: []string{"infrastructure"},
			},
		}

		violations := validator.Validate(".", rules)
		if len(violations) == 0 {
			fmt.Println("✅ Architecture valid")
			return
		}

		for _, v := range violations {
			fmt.Println("❌", v)
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
