package gitlabCmd

import (
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	gitlabService "masood.dev/algo/services/gitlab-client"
)

var ProjectId int

func init() {
	gitlabCmd.AddCommand(gitlabMergeRequestsCmd)

	gitlabMergeRequestsCmd.PersistentFlags().IntVarP(&ProjectId, "project-id", "p", 0, "Project ID")
	gitlabMergeRequestsCmd.PersistentFlags().BoolVarP(&AsFilter, "as-filter", "f", false, "Return results as filter")
}

var gitlabMergeRequestsCmd = &cobra.Command{
	Use:   "merge-requests",
	Short: "merge-requests command",
	Long:  `merge-requests command`,
	Run:   runGitlabMergeRequestsCmd,
}

func runGitlabMergeRequestsCmd(cmd *cobra.Command, args []string) {
	if ProjectId == 0 {
		cmd.Println("Project ID is required")
		return
	}

	mergeRequests, _, err := gitlabService.Client.MergeRequests.ListProjectMergeRequests(ProjectId, &gitlab.ListProjectMergeRequestsOptions{
		State:   gitlab.Ptr("opened"),
		OrderBy: gitlab.Ptr("created_at"),
	})

	if err != nil {
		cmd.Println("Error getting merge requests:", err)
		return
	}

	if AsFilter {
		MergeRequests(mergeRequests).SendAsAlfredFilter()
		return
	}

	for _, mergeRequest := range mergeRequests {
		cmd.Println(mergeRequest.ID, mergeRequest.Title, mergeRequest.WebURL)
	}
}
