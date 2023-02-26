package image

import "context"

type GenerateParameters struct {
	Prompt string `json:"prompt,omitempty"`
	N      int    `json:"n,omitempty"`
	Size   string `json:"size,omitempty"`
	Format string `json:"response_format,omitempty"`
	User   string `json:"user,omitempty"`
}

type GenerateResponse struct {
	CreatedAt int64    `json:"created_at,omitempty"`
	Data      []*Image `json:"data,omitempty"`
}

func (c *Client) Generate(ctx context.Context, p *GenerateParameters) (*GenerateResponse, error) {
	var r GenerateResponse
	if err := c.s.MakeRequest(ctx, c.GenerationEndpoint, p, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
