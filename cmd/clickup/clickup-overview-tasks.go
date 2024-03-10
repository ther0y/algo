package clickupcmd

import (
	"fmt"
	"github.com/spf13/cobra"
	clickupservice "masood.dev/algo/services/clickup"
)

var asFilter bool

func init() {
	clickupCmd.AddCommand(clickupOverviewTasksCmd)
	clickupOverviewTasksCmd.PersistentFlags().BoolVarP(&asFilter, "as-filter", "f", false, "Return results as filter")
}

var clickupOverviewTasksCmd = &cobra.Command{
	Use:   "overview-tasks",
	Short: "Get all tasks from overview list",
	Long:  `Get all tasks from overview list from ClickUp`,
	Run:   runClickUpOverviewTasksCmd,
}

func runClickUpOverviewTasksCmd(cmd *cobra.Command, args []string) {
	tasks, err := clickupservice.GetOverviewTasks()

	if err != nil {
		cmd.Println("Error getting tasks:", err)
		return
	}

	if asFilter {
		clickupservice.Tasks(tasks).SendAsAlfredFilter()
		return
	}

	for _, task := range tasks {
		fmt.Printf("%+v %+v %+v %+v\n", task.Name, task.Assignees[0].Username, task.Status.Status, task.Project.Name)
	}
}
