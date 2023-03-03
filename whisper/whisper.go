// Package whisper implements a client for OpenAI's Whisper
// audio transcriber.
package whisper

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"github.com/rakyll/openai-go"
)

const defaultCreateCompletionsEndpoint = "https://api.openai.com/v1/audio/transcriptions"

// Client is a client to communicate with Open AI's ChatGPT APIs.
type Client struct {
	s     *openai.Session
	model string

	// CreateCompletionsEndpoint allows overriding the default API endpoint.
	// Set this field before using the client.
	CreateCompletionEndpoint string
}

// NewClient creates a new default client that uses the given session
// and defaults to the given model.
func NewClient(session *openai.Session, model string) *Client {
	if model == "" {
		model = "whisper-1"
	}
	return &Client{
		s:                        session,
		model:                    model,
		CreateCompletionEndpoint: defaultCreateCompletionsEndpoint,
	}
}

type CreateCompletionParams struct {
	Model       string
	Language    string
	Audio       io.Reader
	AudioFormat string // such as "mp3" or "wav", etc.
}

type CreateCompletionResponse struct {
	Text string `json:"text,omitempty"`
}

func (c *Client) Transcribe(ctx context.Context, p *CreateCompletionParams) (*CreateCompletionResponse, error) {
	if p.AudioFormat == "" {
		return nil, fmt.Errorf("audio format is required")
	}
	if p.Model == "" {
		p.Model = c.model
	}
	params := url.Values{}
	params.Set("model", p.Model)
	if p.Language != "" {
		params.Set("language", p.Language)
	}
	var r CreateCompletionResponse
	return &r, c.s.Upload(ctx, c.CreateCompletionEndpoint, p.Audio, p.AudioFormat, params, &r)
}
