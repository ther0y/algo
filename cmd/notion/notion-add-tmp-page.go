package notionCmd

import (
	"github.com/spf13/cobra"
	utils "masood.dev/algo/services"
)

func init() {
	notionCmd.AddCommand(notionAddTmpPage)
}

var notionAddTmpPage = &cobra.Command{
	Use:   "add",
	Short: "notion add tmp page command",
	Long:  `notion add tmp page command`,
	Run: func(cmd *cobra.Command, args []string) {
		notionClient := utils.NewNotionClient()
		notionClient.AddPage()
	},
}
