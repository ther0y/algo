package code_review

import (
	"fmt"
	"masood.dev/algo/cmd"

	"github.com/spf13/cobra"
	"masood.dev/algo/agents"
)

var Query string
var AsFilter bool

func init() {
	cmd.RootCmd.AddCommand(mrCmd)

	mrCmd.PersistentFlags().StringVarP(&Query, "query", "q", "", "Query to search")
	mrCmd.PersistentFlags().BoolVarP(&AsFilter, "as-filter", "f", false, "Return results as filter")
}

var mrCmd = &cobra.Command{
	Use:   "mr",
	Short: "mr command",
	Long:  `mr command`,
	Run: func(cmd *cobra.Command, args []string) {
		developerAgent := agents.NewDeveloperAgent()
		resp := developerAgent.Ask(Query)

		if AsFilter {
			AiChoices(resp.Choices).SendAsAlfredFilter()
		} else {
			fmt.Println(resp.Choices[0].Message.Content)
		}
	},
}
