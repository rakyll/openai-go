// Package edit contains a client for OpenAI's edits API.
package edit

import (
	"context"

	"github.com/rakyll/openai-go"
)

const defaultCreateEndpoint = "https://api.openai.com/v1/edits"

// Client is a client to communicate with Open AI's edits API.
type Client struct {
	s     *openai.Session
	model string

	// CreateEndpoint allows overriding the default API endpoint.
	// Set this field before using the client.
	CreateEndpoint string
}

func NewClient(session *openai.Session, model string) *Client {
	return &Client{
		s:              session,
		model:          model,
		CreateEndpoint: defaultCreateEndpoint,
	}
}

// CreateParameters are completion parameters. Refer to OpenAI documentation
// at https://platform.openai.com/docs/api-reference/edits/create
// for reference.
type CreateParameters struct {
	Model       string `json:"model,omitempty"`
	Input       string `json:"input,omitempty"`
	Instruction string `json:"instruction,omitempty"`

	N           int     `json:"n,omitempty"`
	TopP        float64 `json:"top_n,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
}

// CreateResponse is a response to a completion. Refer to OpenAI documentation
// at https://platform.openai.com/docs/api-reference/edits/create
// for reference.
type CreateResponse struct {
	Object    string    `json:"object,omitempty"`
	CreatedAt int64     `json:"created_at,omitempty"`
	Choices   []*Choice `json:"choices,omitempty"`

	Usage *openai.Usage `json:"usage,omitempty"`
}

type Choice struct {
	Text  string `json:"text,omitempty"`
	Index int    `json:"index,omitempty"`
}

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
