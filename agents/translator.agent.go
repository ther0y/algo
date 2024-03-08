package agents

var translatorComment = "You are a great translator with great knowledge of languages, you can translate from any language to any language. you have a keen eye while translating and always ready to point to subtle but very important details and change requests on translations."

// TranslatorAgent struct
type TranslatorAgent struct {
	*Agent
}

// NewTranslatorAgent function
func NewTranslatorAgent() *TranslatorAgent {
	agent := NewAgent(translatorComment)

	return &TranslatorAgent{agent}
}
