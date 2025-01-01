package agent

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"go-agents/models"
)

type BaseAgent struct {
	APIURL      string
	APIKey      *string
	Client      *http.Client
	Model       string
	Provider    *string
	Temperature *float64
	MaxTokens   *int
	Messages    []models.Message
}

func NewBaseAgent(apiURL string, apiKey *string, systemContent *string, model *string, provider *string) *BaseAgent {
	defaultMessage := "You are a helpful assistant"
	if systemContent != nil {
		defaultMessage = *systemContent
	}

	return &BaseAgent{
		APIURL:   apiURL,
		APIKey:   apiKey,
		Client:   &http.Client{},
		Model:    defaultString(model, "gpt-3.5-turbo"),
		Provider: provider,
		Messages: []models.Message{{Role: "system", Content: defaultMessage}},
	}
}

func (b *BaseAgent) SendMessage(userMessage string) (string, error) {
	b.Messages = append(b.Messages, models.Message{Role: "user", Content: userMessage})

	requestBody, _ := json.Marshal(models.GPTRequest{
		Model:       b.Model,
		APIKey:      b.APIKey,
		Provider:    b.Provider,
		Messages:    b.Messages,
		Temperature: b.Temperature,
		MaxTokens:   b.MaxTokens,
	})

	req, _ := http.NewRequest("POST", b.APIURL, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	if b.APIKey != nil {
		req.Header.Set("Authorization", "Bearer "+*b.APIKey)
	}

	resp, err := b.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("HTTP error: " + resp.Status)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var gptResponse models.GPTResponse
	if err := json.Unmarshal(body, &gptResponse); err != nil {
		return "", err
	}

	if len(gptResponse.Choices) > 0 {
		reply := gptResponse.Choices[0].Message.Content
		b.Messages = append(b.Messages, models.Message{Role: "assistant", Content: reply})
		return reply, nil
	}

	return "", errors.New("no choices in response")
}

func defaultString(ptr *string, defaultValue string) string {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}
