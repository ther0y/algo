package gitlabCmd

import (
	aw "github.com/deanishe/awgo"
	"github.com/dustin/go-humanize"
	"github.com/xanzy/go-gitlab"
	"strconv"
)

type Projects []*gitlab.Project

type MergeRequests []*gitlab.MergeRequest

func (p Projects) SendAsAlfredItems() {
	wf := aw.New()

	for _, project := range p {
		wf.NewItem(project.Name).
			Subtitle(project.PathWithNamespace).
			Arg().
			Var("project_id", strconv.Itoa(project.ID)).
			Var("project_name", project.Name).
			Var("project_url", project.WebURL).
			Valid(true)
	}
	wf.SendFeedback()
	return
}

func (m MergeRequests) SendAsAlfredItems() {
	wf := aw.New()

	for _, mergeRequest := range m {
		wf.NewItem(mergeRequest.Title).
			Subtitle(humanize.Time(mergeRequest.UpdatedAt.Local())).
			Var("merge_request_id", strconv.Itoa(mergeRequest.ID)).
			Var("merge_request_title", mergeRequest.Title).
			Var("merge_request_url", mergeRequest.WebURL).
			Arg().
			Valid(true)
	}
	wf.SendFeedback()
	return
}
