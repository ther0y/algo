package utils

import (
	"context"
	"github.com/joho/godotenv"
	"os"

	"github.com/sashabaranov/go-openai"
)

var (
	openaiToken string
	client      *openai.Client
)

type OpenAiClient struct {
}

func init() {
	_ = godotenv.Load(".env")
	openaiToken = os.Getenv("OPENAI_API_KEY")

	client = openai.NewClient(openaiToken)
}

func Ask(model string, systemComment string, query string) *openai.ChatCompletionResponse {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       model,
			Temperature: 0.7,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemComment,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)

	if err != nil {
		panic(err)
	}

	return &resp
}
