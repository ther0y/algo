package code_review

import (
	aw "github.com/deanishe/awgo"
	"github.com/sashabaranov/go-openai"
)

type AiChoices []openai.ChatCompletionChoice

func (ac AiChoices) SendAsAlfredItems() {
	wf := aw.New()

	for _, choice := range ac {
		wf.NewItem(choice.Message.Content).
			Subtitle("").
			Arg(choice.Message.Content).
			Valid(true)
	}
	wf.SendFeedback()
	return
}
