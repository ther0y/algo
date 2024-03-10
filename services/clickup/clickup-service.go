package clickupservice

import (
	"github.com/joho/godotenv"
	"github.com/raksul/go-clickup/clickup"
	"os"
)

var client *clickup.Client

func getClient() *clickup.Client {
	_ = godotenv.Load(".env")

	if client == nil {
		_ = godotenv.Load(".env")
		client = clickup.NewClient(nil, os.Getenv("CLICKUP_TOKEN"))
	}

	return client
}
