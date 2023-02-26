// Package completion contains a client for OpenAI's completion API.
package completion

import (
	"context"

	"github.com/rakyll/openai-go"
)

const defaultCreateEndpoint = "https://api.openai.com/v1/completions"

// Client is a client to communicate with Open AI's completions API.
type Client struct {
	s     *openai.Session
	model string

	// CreateEndpoint allows overriding the default API endpoint.
	// Set this field before using the client.
	CreateEndpoint string
}

// NewClient creates a new default client that uses the given session
// and defaults to the given model.
func NewClient(session *openai.Session, model string) *Client {
	return &Client{
		s:              session,
		model:          model,
		CreateEndpoint: defaultCreateEndpoint,
	}
}

// CreateParameters are completion parameters. Refer to OpenAI documentation
// at https://platform.openai.com/docs/api-reference/completions/create
// for reference.
type CreateParameters struct {
	Model  string   `json:"model,omitempty"`
	Prompt []string `json:"prompt,omitempty"`
	Stop   []string `json:"stop,omitempty"`
	Suffix string   `json:"suffix,omitempty"`
	User   string   `json:"user,omitempty"`

	MaxTokens   int     `json:"max_tokens,omitempty"`
	N           int     `json:"n,omitempty"`
	TopP        float64 `json:"top_n,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`

	Stream bool `json:"stream,omitempty"`
	Echo   bool `json:"echo,omitempty"`

	LogProbs         int     `json:"logprobs,omitempty"`
	PresencePenalty  float64 `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`
	BestOf           int     `json:"best_of,omitempty"`
}

// CreateResponse is a response to a completion. Refer to OpenAI documentation
// at https://platform.openai.com/docs/api-reference/completions/create
// for reference.
type CreateResponse struct {
	ID        string    `json:"id,omitempty"`
	Object    string    `json:"object,omitempty"`
	CreatedAt int64     `json:"created_at,omitempty"`
	Choices   []*Choice `json:"choices,omitempty"`

	Usage *openai.Usage `json:"usage,omitempty"`
}

type Choice struct {
	Text         string `json:"text,omitempty"`
	Index        int    `json:"index,omitempty"`
	LogProbs     int    `json:"logprobs,omitempty"`
	FinishReason string `json:"finish_reason,omitempty"`
}

// Create creates a completion for the provided parameters.
func (c *Client) Create(ctx context.Context, p *CreateParameters) (*CreateResponse, error) {
	if p.Model == "" {
		p.Model = c.model
	}

	var r CreateResponse
	if err := c.s.MakeRequest(ctx, c.CreateEndpoint, p, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
