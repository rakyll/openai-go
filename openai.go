// Package openai contains Go client libraries for OpenAI libraries.
package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// Session is a session created to communicate with OpenAI.
type Session struct {
	apiKey string
	orgID  string

	// HTTPClient providing a custom HTTP client.
	// This field must be set before session is used.
	HTTPClient *http.Client
}

// NewSession creates a new session.
func NewSession(apiKey string, orgID string) *Session {
	return &Session{
		apiKey: apiKey,
		orgID:  orgID,
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
	if s.orgID != "" {
		req.Header.Set("OpenAI-Organization", s.orgID)
	}
	req.Header.Set("Content-Type", "application/json")
	// TODO: Handle JSON errors.
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(output)
}

// Usage reports the API usage.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}
