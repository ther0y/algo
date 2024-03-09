package gitlabService

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	gitlab "github.com/xanzy/go-gitlab"
)

var (
	Client *gitlab.Client
	err    error
)

func init() {
	Client, err = newClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}

// newClient creates and initializes a GitLab Client, ensuring it's created only once.
func newClient() (*gitlab.Client, error) {
	_ = godotenv.Load(".env")

	c, err := gitlab.NewClient(os.Getenv("GITLAB_TOKEN"))

	if err != nil {
		return nil, err
	}

	return c, nil
}
