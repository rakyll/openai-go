package edit

import (
	"context"

	"github.com/rakyll/openai-go"
)

const defaultEndpoint = "https://api.openai.com/v1/edits"

// Client is a client to communicate with Open AI's edits API.
type Client struct {
	s     *openai.Session
	model string

	// Endpoint allows overriding the default API endpoint.
	// Set this field before using the client.
	Endpoint string
}

// Parameters are completion parameters. Refer to OpenAI documentation
// at https://platform.openai.com/docs/api-reference/edits/create
// for reference.
type Parameters struct {
	Model       string `json:"model,omitempty"`
	Input       string `json:"input,omitempty"`
	Instruction string `json:"instruction,omitempty"`

	N           int     `json:"n,omitempty"`
	TopP        float64 `json:"top_n,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
}

// Response is a response to a completion. Refer to OpenAI documentation
// at https://platform.openai.com/docs/api-reference/edits/create
// for reference.
type Response struct {
	Object    string    `json:"object,omitempty"`
	CreatedAt int64     `json:"created_at,omitempty"`
	Choices   []*Choice `json:"choices,omitempty"`

	Usage *openai.Usage `json:"usage,omitempty"`
}

type Choice struct {
	Text  string `json:"text,omitempty"`
	Index int    `json:"index,omitempty"`
}

func (c *Client) Edit(ctx context.Context, p *Parameters) (*Response, error) {
	if p.Model == "" {
		p.Model = c.model
	}

	var r Response
	if err := c.s.MakeRequest(ctx, c.Endpoint, p, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
