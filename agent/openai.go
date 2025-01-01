package agent

// OpenAIAgent communicates with OpenAI's API
type OpenAIAgent struct {
	Base *BaseAgent
}

// NewOpenAIAgent initializes OpenAIAgent
func NewOpenAIAgent(systemContent, apiKey string) *OpenAIAgent {
	apiURL := "https://api.openai.com/v1/chat/completions"
	return &OpenAIAgent{Base: NewBaseAgent(apiURL, &apiKey, &systemContent, nil, nil)}
}

// SendMessage forwards the message request
func (o *OpenAIAgent) SendMessage(userMessage string) (string, error) {
	return o.Base.SendMessage(userMessage)
}
