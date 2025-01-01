package agent

// GPT4FreeAgent is a wrapper around BaseAgent
type GPT4FreeAgent struct {
	Base *BaseAgent
}

// NewGPT4FreeAgent initializes GPT4FreeAgent
func NewGPT4FreeAgent(systemContent string) *GPT4FreeAgent {
	apiURL := "http://localhost:1337/v1/chat/completions"
	return &GPT4FreeAgent{Base: NewBaseAgent(apiURL, nil, &systemContent, nil, nil)}
}

// SendMessage forwards the message request
func (g *GPT4FreeAgent) SendMessage(userMessage string) (string, error) {
	return g.Base.SendMessage(userMessage)
}
