// Package image contains a client for OpenAI's images API.
package image

import (
	"context"
	"fmt"
	"io"
	"net/url"
)

const (
	imageFormat = "png"
)

// CreateEditParams are create image edit parameters. Refer to OpenAI documentation
// at https://platform.openai.com/docs/api-reference/images/create-edit for reference.
type CreateEditParams struct {
	Prompt         string
	Size           string
	ResponseFormat string
	User           string
	N              int

	Image io.Reader
	Mask  io.Reader
}

type CreateEditResponse struct {
	Created int64    `json:"created,omitempty"`
	Data    []*Image `json:"data,omitempty"`
}

func (c *Client) CreateEdit(ctx context.Context, p *CreateEditParams) (*CreateEditResponse, error) {
	params := url.Values{}
	if p.Prompt != "" {
		params.Set("prompt", p.Prompt)
	}
	if p.Size != "" {
		params.Set("size", p.Size)
	}
	if p.N != 0 {
		params.Set("n", fmt.Sprintf("%d", p.N))
	}
	if p.ResponseFormat != "" {
		params.Set("response_format", p.ResponseFormat)
	}
	if p.User != "" {
		params.Set("user", p.User)
	}
	var r CreateEditResponse
	return &r, c.s.Upload(ctx, c.CreateEditEndpoint, map[string]io.Reader{
		"image": p.Image,
		"mask":  p.Mask,
	}, imageFormat, params, &r)
}
