package nobitex

import (
	"fmt"
	"os"

	"github.com/imroc/req/v3"
	"github.com/joho/godotenv"
)

var (
	client       *req.Client
	nobitexToken string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	nobitexToken = os.Getenv("NOBITEX_TOKEN")

	client = req.C().
		SetCommonHeader("Authorization", "Token "+nobitexToken).
		SetBaseURL("https://api.nobitex.ir").
		SetCommonHeader("Content-Type", "application/json").
		SetCommonHeader("Accept", "application/json")
}
