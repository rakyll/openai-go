// Package openai contains Go client libraries for OpenAI libraries.
package openai

import (
	"net/http"
	"time"
)

// Session is a session created to communicate with OpenAI.
type Session struct {
	apiKey string

	// HTTPClient providing a custom HTTP client.
	// This field must be set before session is used.
	HTTPClient *http.Client
}

// NewSession creates a new session.
func NewSession(apiKey string) *Session {
	return &Session{
		apiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// MakeRequest makes HTTP requests and authenticates them with
// session's API key.
func (s *Session) MakeRequest(r *http.Request) (*http.Response, error) {
	if s.apiKey != "" {
		r.Header.Set("Authorization", "Bearer "+s.apiKey)
	}
	r.Header.Set("Content-Type", "application/json")
	// TODO: Handle JSON errors.
	return s.HTTPClient.Do(r)
}

// Usage reports the API usage.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}
