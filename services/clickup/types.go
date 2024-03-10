package clickupservice

import (
	aw "github.com/deanishe/awgo"
	"github.com/dustin/go-humanize"
	"github.com/raksul/go-clickup/clickup"
	"strconv"
	"time"
)

type ByUpdated []clickup.Task

func (a ByUpdated) Len() int      { return len(a) }
func (a ByUpdated) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByUpdated) Less(i, j int) bool {
	iAsInt, _ := strconv.Atoi(a[i].DateUpdated)
	jAsInt, _ := strconv.Atoi(a[j].DateUpdated)

	return iAsInt > jAsInt
}

type Tasks []clickup.Task

func (t Tasks) SendAsAlfredFilter() {
	wf := aw.New()

	for _, task := range t {
		unixUpdatedAt, _ := strconv.Atoi(task.DateUpdated)
		humanizedUpdatedAt := humanize.Time(time.Unix(int64(unixUpdatedAt)/1000, 0))

		// task status to circle colored emoji
		statusEmoji := map[string]string{
			"completed":    "🟢",
			"production":   "🟣",
			"approved":     "☑️",
			"in progress":  "🔵",
			"pending test": "🟡",
			"in review":    "🟠",
			"closed":       "🔴",
			"to do":        "⚪️",
			"backlog":      "🔘",
			"blocked":      "⚫️",
		}

		wf.NewItem(statusEmoji[task.Status.Status]+" "+task.Name).
			Subtitle("Updated: "+humanizedUpdatedAt+", Assignee: "+task.Assignees[0].Username+", List: "+task.Project.Name+" / "+task.List.Name).
			Var("task_id", task.ID).
			Var("task_name", task.Name).
			Var("task_assignee", task.Assignees[0].Username).
			Var("task_status", task.Status.Status).
			Var("task_url", task.URL).
			Arg().
			Valid(true)
	}

	wf.SendFeedback()
}
