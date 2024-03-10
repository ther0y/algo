package vercelService

import (
	"fmt"
)

type AlfredItemArg struct {
	URL          string
	InspectorUrl string
}

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

func GetDeploymentsByProjectID(projectID string) (Deployments, error) {
	if client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	var result GetDeploymentsResponse
	resp, err := client.R().
		SetSuccessResult(&result). // Assuming success by default; check resp for detailed status
		SetQueryParam("projectId", projectID).
		Get("/v6/deployments")

	if err != nil {
		return nil, err // Return the error instead of panicking
	}

	if !resp.IsSuccessState() {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	return result.Deployments, nil
}
