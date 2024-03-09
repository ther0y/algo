package gitlabCmd

import (
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	gitlabService "masood.dev/algo/services/gitlab-client"
)

var AsFilter bool

func init() {
	gitlabCmd.AddCommand(gitlabProjectsCmd)

	gitlabProjectsCmd.PersistentFlags().BoolVarP(&AsFilter, "as-filter", "f", false, "Return results as filter")
}

var gitlabProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "projects command",
	Long:  `projects command`,
	Run:   runGitlabProjectsCmd,
}

func runGitlabProjectsCmd(cmd *cobra.Command, args []string) {
	projects, _, err := gitlabService.Client.Groups.ListGroupProjects("gwat", &gitlab.ListGroupProjectsOptions{
		Archived: gitlab.Ptr(false),
		OrderBy:  gitlab.Ptr("last_activity_at"),
	})

	if err != nil {
		cmd.Println("Error getting projects:", err)
		return
	}

	if AsFilter {
		Projects(projects).SendAsAlfredItems()
	}

	for _, project := range projects {
		cmd.Println(project.ID, project.Name, project.WebURL)
	}
}
