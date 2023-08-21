// Package chat contains a client for Open AI's ChatGPT APIs.
package chat

import (
	"context"
	"errors"

	"github.com/rakyll/openai-go"
)

const defaultModel = "gpt-3.5-turbo"

const defaultCreateCompletionsEndpoint = "https://api.openai.com/v1/chat/completions"

// Client is a client to communicate with Open AI's ChatGPT APIs.
type Client struct {
	s     *openai.Session
	model string

	// CreateCompletionsEndpoint allows overriding the default API endpoint.
	// Set this field before using the client.
	CreateCompletionEndpoint string
}

// NewClient creates a new default client that uses the given session
// and defaults to the given model.
func NewClient(session *openai.Session, model string) *Client {
	if model == "" {
		model = defaultModel
	}
	return &Client{
		s:                        session,
		model:                    model,
		CreateCompletionEndpoint: defaultCreateCompletionsEndpoint,
	}
}

type CreateCompletionParams struct {
	Model string `json:"model,omitempty"`

	Messages []*Message `json:"messages,omitempty"`
	Stop     []string   `json:"stop,omitempty"`
	Stream   bool       `json:"stream,omitempty"`

	N           int     `json:"n,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	MaxTokens   int     `json:"max_tokens,omitempty"`

	PresencePenalty  float64 `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`

	User string `json:"user,omitempty"`
}

type CreateCompletionResponse struct {
	ID        string    `json:"id,omitempty"`
	Object    string    `json:"object,omitempty"`
	CreatedAt int64     `json:"created_at,omitempty"`
	Choices   []*Choice `json:"choices,omitempty"`

	Usage *openai.Usage `json:"usage,omitempty"`
}

type Choice struct {
	Message      *Message `json:"message,omitempty"`
	Index        int      `json:"index,omitempty"`
	LogProbs     int      `json:"logprobs,omitempty"`
	FinishReason string   `json:"finish_reason,omitempty"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name,omitempty"`
}

func (c *Client) CreateCompletion(ctx context.Context, p *CreateCompletionParams) (*CreateCompletionResponse, error) {
	if p.Model == "" {
		p.Model = c.model
	}
	if p.Stream {
		return nil, errors.New("use StreamingClient instead")
	}

	var r CreateCompletionResponse
	if err := c.s.MakeRequest(ctx, c.CreateCompletionEndpoint, p, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
