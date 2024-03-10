package clickupservice

import (
	"context"
	"github.com/raksul/go-clickup/clickup"
	"sort"
)

func GetOverviewTasks() ([]clickup.Task, error) {
	client := getClient()

	tasks, _, _, err := client.Views.GetViewTasks(context.Background(), "6jyze-22552", 0)

	if err != nil {
		return nil, err
	}

	sort.Sort(ByUpdated(tasks))

	return tasks, nil
}
