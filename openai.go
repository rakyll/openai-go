// Package openai contains Go client libraries for OpenAI libraries.
package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Session is a session created to communicate with OpenAI.
type Session struct {
	// OrganizationID is the ID optionally to be included as
	// a header to requests made from this session.
	// This field must be set before session is used.
	OrganizationID string

	// HTTPClient providing a custom HTTP client.
	// This field must be set before session is used.
	HTTPClient *http.Client

	apiKey string
}

// NewSession creates a new session. Organization IDs are optional,
// use an empty string when you don't want to set one.
func NewSession(apiKey string) *Session {
	return &Session{
		apiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// MakeRequest make HTTP requests and authenticates them with
// session's API key. MakeRequest marshals input as the request body,
// and unmarshals the response as output.
func (s *Session) MakeRequest(ctx context.Context, endpoint string, input, output interface{}) error {
	buf, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(buf))
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	if s.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+s.apiKey)
	}
	if s.OrganizationID != "" {
		req.Header.Set("OpenAI-Organization", s.OrganizationID)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		respBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}
		return &APIError{
			StatusCode: resp.StatusCode,
			Payload:    respBody,
		}
	}
	return json.NewDecoder(resp.Body).Decode(output)
}

// APIError is returned from API requests if the API
// responds with an error.
type APIError struct {
	StatusCode int
	Payload    []byte
}

func (e *APIError) Error() string {
	return fmt.Sprintf("status_code=%d, payload=%s", e.StatusCode, e.Payload)
}

// Usage reports the API usage.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}
