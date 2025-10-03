package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

const (
	appName    = "Wells Go"
	appVersion = "v1.0.1"
	appStatus  = "Release"
	buildDate  = "2025-10-03"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show application version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("====================================")
		fmt.Printf("%s %s / %s\n", appName, appVersion, appStatus)
		fmt.Printf("Build Date: %s\n", buildDate)
		fmt.Printf("Current Time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("Note: Release version")
		fmt.Println("====================================")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
