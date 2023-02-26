package image

import (
	"github.com/rakyll/openai-go"
)

const (
	defaultGenerationEndpoint = "https://api.openai.com/v1/images/generations"
)

// Client is a client to communicate with Open AI's images API.
type Client struct {
	s *openai.Session

	// GenerationEndpoint allows overriding the default
	// for the image generation API endpoint.
	// Set this field before using the client.
	GenerationEndpoint string
}

func NewClient(session *openai.Session) *Client {
	return &Client{
		s:                  session,
		GenerationEndpoint: defaultGenerationEndpoint,
	}
}

type Image struct {
	URL string `json:"url,omitempty"`
}
