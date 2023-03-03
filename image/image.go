package image

import (
	"github.com/rakyll/openai-go"
)

const (
	defaultCreateEndpoint = "https://api.openai.com/v1/images/generations"
)

// Client is a client to communicate with Open AI's images API.
type Client struct {
	s *openai.Session

	// CreateEndpoint allows overriding the default
	// for the image generation API endpoint.
	// Set this field before using the client.
	CreateEndpoint string
}

func NewClient(session *openai.Session) *Client {
	return &Client{
		s:              session,
		CreateEndpoint: defaultCreateEndpoint,
	}
}

type Image struct {
	URL        string `json:"url,omitempty"`
	Base64JSON string `json:"b64_json,omitempty"`
}
