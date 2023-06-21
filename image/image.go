package image

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rakyll/openai-go"
)

const (
	defaultCreateEndpoint     = "https://api.openai.com/v1/images/generations"
	defaultCreateEditEndpoint = "https://api.openai.com/v1/images/edits"
)

// Client is a client to communicate with Open AI's images API.
type Client struct {
	s *openai.Session

	// CreateEndpoint allows overriding the default
	// for the image generation API endpoint.
	// Set this field before using the client.
	CreateEndpoint string

	// CreateEditEndpoint allows overriding the default
	// for the image edit API endpoint.
	// Set this field before using the client.
	CreateEditEndpoint string
}

func NewClient(session *openai.Session) *Client {
	return &Client{
		s:                  session,
		CreateEndpoint:     defaultCreateEndpoint,
		CreateEditEndpoint: defaultCreateEditEndpoint,
	}
}

type Image struct {
	URL        string `json:"url,omitempty"`
	Base64JSON string `json:"b64_json,omitempty"`
}

func (i *Image) Reader() (io.ReadCloser, error) {
	if i.URL != "" {
		resp, err := http.Get(i.URL)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	}
	if i.Base64JSON != "" {
		decoded, err := base64.StdEncoding.DecodeString(i.Base64JSON)
		if err != nil {
			return nil, err
		}
		return ioutil.NopCloser(bytes.NewBuffer(decoded)), nil
	}
	return nil, errors.New("no image data")
}
