// Package embedding contains a client for Open AI's Embeddings APIs.
package embedding

import (
	"context"

	"github.com/rakyll/openai-go"
)

const (
	defaultModel          = "text-embedding-ada-002"
	defaultCreateEndpoint = "https://api.openai.com/v1/embeddings"
)

// Client is a client to communicate with Open AI's Embeddings APIs.
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
	if model == "" {
		model = defaultModel
	}
	return &Client{
		s:              session,
		model:          model,
		CreateEndpoint: defaultCreateEndpoint,
	}
}

type CreateParams struct {
	Model string `json:"model,omitempty"`

	Input []string `json:"input,omitempty"`
	User  string   `json:"user,omitempty"`
}

type CreateResponse struct {
	Object string  `json:"object,omitempty"`
	Data   []*Data `json:"data,omitempty"`
	Model  string  `json:"model,omitempty"`

	Usage *openai.Usage `json:"usage,omitempty"`
}

type Data struct {
	Object    string    `json:"object,omitempty"`
	Embedding []float64 `json:"embedding,omitempty"`
	Index     int       `json:"index,omitempty"`
}

func (c *Client) Create(ctx context.Context, p *CreateParams) (*CreateResponse, error) {
	if p.Model == "" {
		p.Model = c.model
	}

	var r CreateResponse
	if err := c.s.MakeRequest(ctx, c.CreateEndpoint, p, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
