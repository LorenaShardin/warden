package agent

import (
	"go-agents/models"
	"go-agents/utils"
)

type GrokAgent struct {
	Base *BaseAgent
}

func NewGrokAgent(apiKey string, systemContent string) *GrokAgent {
	apiURL := "https://api.grok.com/v1/chat/completions"
	return &GrokAgent{Base: NewBaseAgent(apiURL, &apiKey, &systemContent, nil, nil)}
}

func (g *GrokAgent) SendMessage(userMessage string) (string, error) {
	return g.Base.SendMessage(userMessage)
}
