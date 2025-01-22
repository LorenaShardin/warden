package models

type GPTRequest struct {
	APIKey      *string   `json:"api_key,omitempty"`
	Messages    []Message `json:"messages"`
	Model       string    `json:"model"`
	Provider    string    `json:"provider"`
	Temperature *float64  `json:"temperature,omitempty"`
	MaxTokens   *int      `json:"max_tokens,omitempty"`
}

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type GPTResponse struct {
	ID        string    `json:"id"`
	Object    string    `json:"object"`
	Created   int64     `json:"created"`
	Model     *string   `json:"model,omitempty"`
	Provider  *string   `json:"provider,omitempty"`
	Choices   []Choice  `json:"choices"`
	Usage     *Usage    `json:"usage,omitempty"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
