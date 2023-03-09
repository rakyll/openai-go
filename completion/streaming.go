package completion

import (
	"context"

	"github.com/rakyll/openai-go"
)

// StreamingClient is a client to communicate with Open AI's completions API.
type StreamingClient struct {
	s     *openai.Session
	model string

	// CreateEndpoint allows overriding the default API endpoint.
	// Set this field before using the client.
	CreateEndpoint string
}

// NewStreamingClient creates a new default streaming client that uses the given session
// and defaults to the given model.
func NewStreamingClient(session *openai.Session, model string) *StreamingClient {
	return &StreamingClient{
		s:              session,
		model:          model,
		CreateEndpoint: defaultCreateEndpoint,
	}
}

// Create creates a completion for the provided parameters.
func (c *StreamingClient) Create(ctx context.Context, p *CreateParams, fn func(r *CreateResponse)) error {
	if p.Model == "" {
		p.Model = c.model
	}
	p.Stream = true

	var r CreateResponse
	return c.s.MakeStreamingRequest(ctx, c.CreateEndpoint, p, &r, func(r any) {
		fn(r.(*CreateResponse))
	})
}
