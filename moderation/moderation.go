package moderation

import (
	"context"

	"github.com/rakyll/openai-go"
)

const defaultEndpoint = "https://api.openai.com/v1/moderations"

// Client is a client to communicate with Open AI's moderation API.
type Client struct {
	s     *openai.Session
	model string

	// Endpoint allows overriding the default API endpoint.
	// Set this field before using the client.
	Endpoint string
}

type Parameters struct {
	Model string   `json:"model,omitempty"`
	Input []string `json:"input,omitempty"`
}

type Response struct {
	ID      string    `json:"id,omitempty"`
	Results []*Result `json:"results,omitempty"`
}

type Result struct {
	Categories     map[string]bool    `json:"categories,omitempty"`
	CategoryScores map[string]float64 `json:"category_scores,omitempty"`
	Flagged        bool               `json:"flagged,omitempty"`
}

func (c *Client) Moderate(ctx context.Context, p *Parameters) (*Response, error) {
	if p.Model == "" {
		p.Model = c.model
	}

	var r Response
	if err := c.s.MakeRequest(ctx, c.Endpoint, p, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
