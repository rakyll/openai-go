package moderation

import (
	"context"

	"github.com/rakyll/openai-go"
)

const defaultCreateModerationEndpoint = "https://api.openai.com/v1/moderations"

// Client is a client to communicate with Open AI's moderation API.
type Client struct {
	s     *openai.Session
	model string

	// CreateModerationEndpoint allows overriding the default API endpoint.
	// Set this field before using the client.
	CreateModerationEndpoint string
}

func NewClient(session *openai.Session, model string) *Client {
	return &Client{
		s:                        session,
		model:                    model,
		CreateModerationEndpoint: defaultCreateModerationEndpoint,
	}
}

type CreateModerationParameters struct {
	Model string   `json:"model,omitempty"`
	Input []string `json:"input,omitempty"`
}

type CreateModerationResponse struct {
	ID      string    `json:"id,omitempty"`
	Results []*Result `json:"results,omitempty"`
}

type Result struct {
	Categories     map[string]bool    `json:"categories,omitempty"`
	CategoryScores map[string]float64 `json:"category_scores,omitempty"`
	Flagged        bool               `json:"flagged,omitempty"`
}

func (c *Client) CreateModeration(ctx context.Context, p *CreateModerationParameters) (*CreateModerationResponse, error) {
	if p.Model == "" {
		p.Model = c.model
	}

	var r CreateModerationResponse
	if err := c.s.MakeRequest(ctx, c.CreateModerationEndpoint, p, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
