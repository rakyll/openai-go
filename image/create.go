// Package image contains a client for OpenAI's images API.
package image

import "context"

// CreateParams are create image parameters. Refer to OpenAI documentation
// at https://platform.openai.com/docs/api-reference/images/create
// for reference.
type CreateParams struct {
	Prompt         string `json:"prompt,omitempty"`
	N              int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}

type CreateResponse struct {
	CreatedAt int64    `json:"created,omitempty"`
	Data      []*Image `json:"data,omitempty"`
}

func (c *Client) Create(ctx context.Context, p *CreateParams) (*CreateResponse, error) {
	var r CreateResponse
	if err := c.s.MakeRequest(ctx, c.CreateEndpoint, p, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
