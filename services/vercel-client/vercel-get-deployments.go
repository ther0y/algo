package vercelclient

import (
	"fmt"
	aw "github.com/deanishe/awgo"
)

// GetDeployments retrieves the deployments for the specified app from the Vercel API.
// It returns a pointer to a GetDeploymentsResponse and an error if any.
func GetDeployments() (*GetDeploymentsResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	var result GetDeploymentsResponse
	resp, err := client.R().
		SetSuccessResult(&result). // Assuming success by default; check resp for detailed status
		SetQueryParam("app", appName).
		SetQueryParam("limit", deployLimit).
		Get("/v6/deployments")

	if err != nil {
		return nil, err // Return the error instead of panicking
	}

	if !resp.IsSuccessState() {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	return &result, nil
}

type AlfredItemArg struct {
	URL          string
	InspectorUrl string
}

func SendDeploymentsAsAlfredItems(deployments []Deployment) {
	wf := aw.New()

	alfredItems := make([]*aw.Item, 0, len(deployments))

	statusEmoji := map[string]string{
		"READY":        "🟢",
		"ERROR":        "🔴",
		"BUILDING":     "🔵",
		"QUEUED":       "🟡",
		"INITIALIZING": "🟡",
		"DEPLOYING":    "🟡",
		"UPLOADING":    "🟡",
	}

	for _, deployment := range deployments {
		deploymentStatus := statusEmoji[deployment.ReadyState]
		alfredItems = append(alfredItems, wf.NewItem(deploymentStatus+" "+deployment.Name).
			Subtitle(deployment.URL).
			Arg(deployment.InspectorURL).
			Var("url", deployment.URL).
			Var("inspectorUrl", deployment.InspectorURL).
			Quicklook(deployment.URL).
			Valid(true))
	}

	wf.SendFeedback()
	return
}
