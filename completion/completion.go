package completion

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/rakyll/openai-go"
)

const defaultEndpoint = "https://api.openai.com/v1/completions"

type Client struct {
	s     *openai.Session
	model string

	// Endpoint allows overriding the default API endpoint.
	// Set this field before using the client.
	Endpoint string
}

func NewClient(session *openai.Session, model string) *Client {
	return &Client{
		s:        session,
		model:    model,
		Endpoint: defaultEndpoint,
	}
}

type Parameters struct {
	Prompt []string `json:"prompt,omitempty"`
	Stop   []string `json:"stop,omitempty"`
	Suffix string   `json:"suffix,omitempty"`
	User   string   `json:"user,omitempty"`

	MaxTokens   int     `json:"max_tokens,omitempty"`
	N           int     `json:"n,omitempty"`
	TopP        float64 `json:"top_n,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`

	Stream bool `json:"stream,omitempty"`
	Echo   bool `json:"echo,omitempty"`

	LogProbs         int     `json:"logprobs,omitempty"`
	PresencePenalty  float64 `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`
	BestOf           int     `json:"best_of,omitempty"`
}

type Response struct {
	ID        string    `json:"id,omitempty"`
	Object    string    `json:"object,omitempty"`
	CreatedAt int64     `json:"created_at,omitempty"`
	Choices   []*Choice `json:"choices,omitempty"`

	Usage *openai.Usage `json:"usage,omitempty"`
}

type Choice struct {
	Text         string `json:"text,omitempty"`
	Index        int    `json:"index,omitempty"`
	LogProbs     int    `json:"logprobs,omitempty"`
	FinishReason string `json:"finish_reason,omitempty"`
}

func (c *Client) Complete(ctx context.Context, p *Parameters) (*Response, error) {
	// TODO: Make sure we omit zero fields correctly.
	buf, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", c.Endpoint, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	resp, err := c.s.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r Response
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&r); err != nil {
		return nil, err
	}
	return &r, nil
}
