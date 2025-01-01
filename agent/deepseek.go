package agent

import (
	"go-agents/models"
	"go-agents/utils"
)

type DeepSeekAgent struct {
	Base *BaseAgent
}

func NewDeepSeekAgent(apiKey string, systemContent string) *DeepSeekAgent {
	apiURL := "https://api.deepseek.com/v1/chat/completions"
	return &DeepSeekAgent{Base: NewBaseAgent(apiURL, &apiKey, &systemContent, nil, nil)}
}

func (d *DeepSeekAgent) SendMessage(userMessage string) (string, error) {
	return d.Base.SendMessage(userMessage)
}
