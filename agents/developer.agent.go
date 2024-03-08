package agents

var developerComment = "You are a great programmer with great knowledge of javascript, typescript, react, nextjs and anything related to web development. you have a keen eye while developing and doing code reviews and always ready to point to subtle but very important details and change requests on pull requests and merge requests."

// DeveloperAgent struct
type DeveloperAgent struct {
	*Agent
}

// NewDeveloperAgent function
func NewDeveloperAgent() *DeveloperAgent {
	agent := NewAgent(developerComment)

	return &DeveloperAgent{agent}
}

// add ask function to developer agent struct
