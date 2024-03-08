package cmd

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
	"masood.dev/algo/agents"
)

func init() {
	RootCmd.AddCommand(mrCmd)

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
		wf := aw.New()

		if AsFilter {
			for _, choice := range resp.Choices {
				wf.NewItem(choice.Message.Content).
					Subtitle("").
					Arg(choice.Message.Content).
					Valid(true)
			}
			wf.SendFeedback()
			return
		} else {
			fmt.Println(resp.Choices[0].Message.Content)
		}
	},
}
