package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(notionCmd)
}

var notionCmd = &cobra.Command{
	Use:   "notion",
	Short: "notion command",
	Long:  `notion command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Notion cmd tools")
	},
}
