package gitlabClient

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/xanzy/go-gitlab"
)

var (
	client *Client
	err    error
)

func init() {
	client, err = newClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}

// Client encapsulates the GitLab client and its operations.
type Client struct {
	*gitlab.Client
}

// newClient creates and initializes a GitLab client, ensuring it's created only once.
func newClient() (*Client, error) {
	_ = godotenv.Load(".env")

	c, err := gitlab.NewClient(os.Getenv("GITLAB_TOKEN"))

	if err != nil {
		return nil, err
	}

	return &Client{c}, nil
}

// GetProject retrieves a project by its ID.
func GetProject(projectId int) (*gitlab.Project, error) {
	if err != nil {
		return nil, err // Return the error from client creation
	}

	project, _, err := client.Client.Projects.GetProject(projectId, nil)
	if err != nil {
		return nil, errors.New("failed to get project")
	}
	return project, nil
}
