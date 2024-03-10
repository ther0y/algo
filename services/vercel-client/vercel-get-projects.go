package vercelService

import "fmt"

func GetProjects() ([]Project, error) {
	var result GetProjectsResponse

	resp, err := client.R().
		SetSuccessResult(&result).
		Get("/v9/projects")

	if err != nil {
		return nil, err // Return the error instead of panicking
	}

	if !resp.IsSuccessState() {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	return result.Projects, nil
}
