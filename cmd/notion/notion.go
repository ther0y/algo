package notionCmd

import (
	"fmt"
	"masood.dev/algo/cmd"

	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(notionCmd)
}

var notionCmd = &cobra.Command{
	Use:   "notion",
	Short: "notion command",
	Long:  `notion command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Notion cmd tools")
	},
}
