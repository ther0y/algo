package agents

import (
	"github.com/sashabaranov/go-openai"
	utils "masood.dev/algo/services"
)

type Agent struct {
	model         string
	systemComment string
}

// NewAgent function
func NewAgent(systemComment string) *Agent {
	return &Agent{
		model:         openai.GPT3Dot5Turbo1106,
		systemComment: systemComment,
	}
}

func (a *Agent) SetModel(model string) {
	a.model = model
}

func (a *Agent) Ask(query string) *openai.ChatCompletionResponse {
	return utils.Ask(a.model, a.systemComment, query)
}
