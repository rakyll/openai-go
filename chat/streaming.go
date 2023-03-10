package chat

import (
	"context"

	"github.com/rakyll/openai-go"
)

// StreamingClient is a client to communicate with Open AI's ChatGPT APIs.
type StreamingClient struct {
	s     *openai.Session
	model string

	// CreateCompletionsEndpoint allows overriding the default API endpoint.
	// Set this field before using the client.
	CreateCompletionEndpoint string
}

// NewStreamingClient creates a new default streaming client that uses the given session
// and defaults to the given model.
func NewStreamingClient(session *openai.Session, model string) *StreamingClient {
	if model == "" {
		model = defaultModel
	}
	return &StreamingClient{
		s:                        session,
		model:                    model,
		CreateCompletionEndpoint: defaultCreateCompletionsEndpoint,
	}
}

type CreateCompletionStreamingResponse struct {
	ID        string             `json:"id,omitempty"`
	Object    string             `json:"object,omitempty"`
	CreatedAt int64              `json:"created_at,omitempty"`
	Choices   []*StreamingChoice `json:"choices,omitempty"`
}

type StreamingChoice struct {
	Delta        *Message `json:"delta,omitempty"`
	Index        int      `json:"index,omitempty"`
	LogProbs     int      `json:"logprobs,omitempty"`
	FinishReason string   `json:"finish_reason,omitempty"`
}

func (c *StreamingClient) CreateCompletion(ctx context.Context, p *CreateCompletionParams, fn func(r *CreateCompletionStreamingResponse)) error {
	if p.Model == "" {
		p.Model = c.model
	}
	p.Stream = true

	var r CreateCompletionStreamingResponse
	return c.s.MakeStreamingRequest(ctx, c.CreateCompletionEndpoint, p, &r, func(r any) {
		fn(r.(*CreateCompletionStreamingResponse))
	})
}
