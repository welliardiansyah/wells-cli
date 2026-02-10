package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var timelineCmd = &cobra.Command{
	Use:   "timeline [entity]",
	Short: "Show evolution timeline of an entity",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		entity := args[0]
		fmt.Println("📈 Timeline for entity:", entity)

		fmt.Println("No snapshots yet (feature ready)")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(timelineCmd)
}
