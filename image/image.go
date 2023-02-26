package image

import (
	"github.com/rakyll/openai-go"
)

const (
	defaultCreateImageEndpoint = "https://api.openai.com/v1/images/generations"
)

// Client is a client to communicate with Open AI's images API.
type Client struct {
	s *openai.Session

	// CreateImageEndpoint allows overriding the default
	// for the image generation API endpoint.
	// Set this field before using the client.
	CreateImageEndpoint string
}

func NewClient(session *openai.Session) *Client {
	return &Client{
		s:                   session,
		CreateImageEndpoint: defaultCreateImageEndpoint,
	}
}

type Image struct {
	URL string `json:"url,omitempty"`
}
