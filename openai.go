package openai

import (
	"net/http"
	"time"
)

type Session struct {
	apiKey string

	// HTTPClient providing a custom HTTP client.
	// This field must be set before session is used.
	HTTPClient *http.Client
}

func NewSession(apiKey string) *Session {
	return &Session{
		apiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (s *Session) MakeRequest(r *http.Request) (*http.Response, error) {
	if s.apiKey != "" {
		r.Header.Set("Authorization", "Bearer "+s.apiKey)
	}
	return s.HTTPClient.Do(r)
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}
