package vercelclient

import (
	"os"

	"github.com/imroc/req/v3"
	"github.com/joho/godotenv"
)

const (
	baseURL     = "https://api.vercel.com"
	appName     = "global-website" // This could be passed as a parameter or made configurable
	deployLimit = "15"             // This could be passed as a parameter or made configurable
)

var (
	client      *req.Client
	vercelToken string // Moved to package-level variable
)

func init() {
	_ = godotenv.Load(".env")

	vercelToken = os.Getenv("VERCEL_TOKEN")
	teamid := os.Getenv("VERCEL_TEAM_ID")

	client = req.C().
		SetBaseURL(baseURL).
		SetCommonHeader("Authorization", "Bearer "+vercelToken). // Move the token to a common header setup
		SetCommonQueryParam("teamId", teamid)
}
