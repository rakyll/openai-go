// Package image contains a client for OpenAI's images API.
package image

import "context"

type CreateParams struct {
	Prompt string `json:"prompt,omitempty"`
	N      int    `json:"n,omitempty"`
	Size   string `json:"size,omitempty"`
	Format string `json:"response_format,omitempty"`
	User   string `json:"user,omitempty"`
	Model  string `json:"model,omitempty"`
}

type CreateResponse struct {
	CreatedAt int64    `json:"created_at,omitempty"`
	Data      []*Image `json:"data,omitempty"`
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
