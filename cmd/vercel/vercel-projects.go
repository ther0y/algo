package vercelCmd

import (
	"github.com/spf13/cobra"
	. "masood.dev/algo/services/vercel-client"
)

var asFilter bool

func init() {
	vercelCmd.AddCommand(vercelProjectsCmd)

	vercelProjectsCmd.PersistentFlags().BoolVarP(&asFilter, "as-filter", "f", false, "Return results as filter")
}

var vercelProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "projects command",
	Long:  `projects command`,
	Run:   runVercelProjectsCmd,
}

func runVercelProjectsCmd(cmd *cobra.Command, args []string) {
	projects, err := GetProjects()

	if err != nil {
		cmd.Println("Error getting projects:", err)
		return
	}

	if asFilter {
		Projects(projects).SendAsAlfredFilter()
		return
	}

	printPrettyJSON(projects)
}
