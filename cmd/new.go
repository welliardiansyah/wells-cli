package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/welliardiansyah/wells-cli/internal/scaffold"
)

var (
	flagModule string
	flagGit    bool
	flagTidy   bool
)

var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create new project scaffold",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		dest := filepath.Join(cwd, name)
		moduleName := flagModule
		if moduleName == "" {
			moduleName = name
		}

		opts := scaffold.Options{
			ModuleName: moduleName,
			GitInit:    flagGit,
			RunTidy:    flagTidy,
		}

		if err := scaffold.CopyTemplate(dest, opts); err != nil {
			return err
		}

		if flagGit {
			if err := scaffold.GitInit(dest); err != nil {
				fmt.Println("⚠️ Git init failed:", err)
			} else {
				fmt.Println("✅ Git initialized")
			}
		}

		if flagTidy {
			if err := scaffold.GoModTidy(dest); err != nil {
				fmt.Println("⚠️ go mod tidy failed:", err)
			} else {
				fmt.Println("✅ go mod tidy executed")
			}
		} else {
			fmt.Println("ℹ️ Run 'go mod tidy' inside the new project to download deps.")
		}

		fmt.Println("✅ Project created at", dest)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringVar(&flagModule, "module", "", "module name to set in go.mod (default: project name)")
	newCmd.Flags().BoolVar(&flagGit, "git", false, "run 'git init' in target folder")
	newCmd.Flags().BoolVar(&flagTidy, "tidy", false, "run 'go mod tidy' in target folder after scaffold")
}
