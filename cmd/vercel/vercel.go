package vercelCmd

import (
	"github.com/spf13/cobra"
	"masood.dev/algo/cmd"
)

var (
	AstFilter bool
)

func init() {
	cmd.RootCmd.AddCommand(vercelCmd)

	vercelCmd.PersistentFlags().BoolVarP(&AstFilter, "as-filter", "f", false, "Return results as filter")
}

var vercelCmd = &cobra.Command{
	Use:   "vercel",
	Short: "vercel command",
	Long:  `vercel command`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
