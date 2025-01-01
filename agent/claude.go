package agent

import (
	"go-agents/models"
	"go-agents/utils"
)

type ClaudeAgent struct {
	Base *BaseAgent
}

func NewClaudeAgent(apiKey string, systemContent string) *ClaudeAgent {
	apiURL := "https://api.anthropic.com/v1/complete"
	return &ClaudeAgent{Base: NewBaseAgent(apiURL, &apiKey, &systemContent, nil, nil)}
}

func (c *ClaudeAgent) SendMessage(userMessage string) (string, error) {
	return c.Base.SendMessage(userMessage)
}
