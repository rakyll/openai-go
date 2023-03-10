// Package openai contains Go client libraries for OpenAI libraries.
package openai

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const userAgent = "openai-go/1"

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
func (s *Session) MakeRequest(ctx context.Context, endpoint string, input, output any) error {
	buf, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(buf))
	if err != nil {
		return err
	}

	body, err := s.sendRequest(req, "application/json")
	if err != nil {
		return err
	}
	defer body.Close()

	return json.NewDecoder(body).Decode(output)
}

func (s *Session) MakeStreamingRequest(ctx context.Context, endpoint string, input any, output any, fn func(any)) error {
	const (
		streamPrefix = "data: "
		streamEnd    = "[DONE]"
	)

	buf, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(buf))
	if err != nil {
		return err
	}

	body, err := s.sendRequest(req, "application/json")
	if err != nil {
		return err
	}
	defer body.Close()

	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		line := strings.Replace(scanner.Text(), streamPrefix, "", 1)
		if line == "" {
			continue
		}
		if line == streamEnd {
			return nil
		}
		if err := json.Unmarshal([]byte(line), output); err != nil {
			return fmt.Errorf("failed to unmarshal streaming response: %w", err)
		}
		fn(output)
	}
	return scanner.Err()
}

// Upload makes a multi-part form data upload them with
// session's API key. Upload combines the file with the given params
// and unmarshals the response as output.
func (s *Session) Upload(ctx context.Context, endpoint string, file io.Reader, fileExt string, params url.Values, output any) error {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	go func() {
		err := upload(mw, file, fileExt, params)
		pw.CloseWithError(err)
	}()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, pr)
	if err != nil {
		return err
	}
	body, err := s.sendRequest(req, mw.FormDataContentType())
	if err != nil {
		return err
	}
	defer body.Close()

	return json.NewDecoder(body).Decode(output)
}

func (s *Session) sendRequest(req *http.Request, contentType string) (io.ReadCloser, error) {
	if s.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+s.apiKey)
	}
	if s.OrganizationID != "" {
		req.Header.Set("OpenAI-Organization", s.OrganizationID)
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Add("User-Agent", userAgent)

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Payload:    respBody,
		}
	}
	return resp.Body, nil
}

func upload(mw *multipart.Writer, file io.Reader, fileExt string, params url.Values) error {
	for key := range params {
		w, err := mw.CreateFormField(key)
		if err != nil {
			return fmt.Errorf("error creating %q field: %w", key, err)
		}
		_, err = fmt.Fprint(w, params.Get(key))
		if err != nil {
			return fmt.Errorf("error writing %q field: %w", key, err)
		}
	}
	w, err := mw.CreateFormFile("file", "audio."+fileExt)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	_, err = io.Copy(w, file)
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}
	err = mw.Close()
	if err != nil {
		return fmt.Errorf("error closing multipart writer: %w", err)
	}
	return nil
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
