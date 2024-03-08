package cmd

import (
	"github.com/spf13/cobra"
	vercelclient "masood.dev/algo/services/vercel-client"
)

var (
	AstFilter bool
)

func init() {
	RootCmd.AddCommand(vercelCmd)

	vercelCmd.PersistentFlags().BoolVarP(&AstFilter, "as-filter", "f", false, "Return results as filter")
}

var vercelCmd = &cobra.Command{
	Use:   "vercel",
	Short: "vercel command",
	Long:  `vercel command`,
	Run: func(cmd *cobra.Command, args []string) {
		deps, err := vercelclient.GetDeployments()
		if err != nil {
			panic(err)
		}

		if AstFilter {
			vercelclient.SendDeploymentsAsAlfredItems(deps.Deployments)
		} else {
			for _, dep := range deps.Deployments {
				println(dep.URL)
			}
		}
	},
}
