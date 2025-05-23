package openai

import (
	"github.com/bit8bytes/gogantic/llm"
)

type Model struct {
	Model       string    `json:"model"`
	APIKey      string    `json:"api_key"`
	Temperature *float32  `json:"temperature,omitempty"`
	Stop        *[]string `json:"stop,omitempty"`
	Stream      *bool     `json:"stream,omitempty"`
}

type Chat struct {
	Model       string        `json:"model"`
	Messages    []llm.Message `json:"messages"`
	Temperature float32       `json:"temperature,omitempty"`
	Stop        []string      `json:"stop,omitempty"`
	Stream      bool          `json:"stream,omitempty"`
}

type Response struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Usage   UsageResponse
	Choices []ChoiceResponse
}

type ChoiceResponse struct {
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
	Index        int         `json:"index"`
}

type UsageResponse struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type EmbeddingResponseData struct {
	Object    string    `json:"object"`
	Embedding []float32 `json:"embedding"`
	Index     int       `json:"index"`
}

type EmbeddingResponseUsage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

type EmbeddingResponse struct {
	Object string                  `json:"object"`
	Data   []EmbeddingResponseData `json:"data"`
	Model  string                  `json:"model"`
	Usage  EmbeddingResponseUsage  `json:"usage"`
}

type Embedding struct {
	Input          string `json:"input"`
	Model          string `json:"model"`
	EncodingFormat string `json:"encoding_format,omitempty"`
}
