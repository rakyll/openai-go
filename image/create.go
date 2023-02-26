package image

import "context"

type CreateParameters struct {
	Prompt string `json:"prompt,omitempty"`
	N      int    `json:"n,omitempty"`
	Size   string `json:"size,omitempty"`
	Format string `json:"response_format,omitempty"`
	User   string `json:"user,omitempty"`
}

type CreateResponse struct {
	CreatedAt int64    `json:"created_at,omitempty"`
	Data      []*Image `json:"data,omitempty"`
}

func (c *Client) Create(ctx context.Context, p *CreateParameters) (*CreateResponse, error) {
	var r CreateResponse
	if err := c.s.MakeRequest(ctx, c.CreateImageEndpoint, p, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
