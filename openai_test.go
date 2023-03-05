package openai

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitSession(t *testing.T) {
	s := NewSession("test")
	require.NotNil(t, s, "you must set an API key")
}

func TestMakeRequest(t *testing.T) {
	tests := []struct {
		name        string
		statusCode  int
		apiKey      string
		input       any
		output      any
		expectError bool
	}{
		{
			name:        "success",
			statusCode:  http.StatusOK,
			apiKey:      "test-key",
			input:       map[string]string{"hello": "world"},
			output:      &struct{ Foo string }{},
			expectError: false,
		},
		{
			name:        "error:empty_api_key",
			statusCode:  http.StatusUnauthorized,
			apiKey:      "",
			input:       map[string]string{"hello": "world"},
			output:      &struct{ Foo string }{},
			expectError: true,
		},
		{
			name:        "error:nil_input",
			statusCode:  http.StatusOK,
			apiKey:      "test-key",
			input:       nil,
			output:      &struct{ Foo string }{},
			expectError: true,
		},
		{
			name:        "error:invalid_output",
			statusCode:  http.StatusInternalServerError,
			apiKey:      "test-key",
			input:       map[string]string{"hello": "world"},
			output:      &struct{ Foo int }{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			mockServer := httpMockServer(tt.statusCode)
			defer mockServer.Close()

			// Create a new session with a test API key and the mock server's HTTP client.
			session := &Session{
				apiKey:     tt.apiKey,
				HTTPClient: mockServer.Client(),
			}

			// Make the request.
			err := session.MakeRequest(context.Background(), mockServer.URL, tt.input, tt.output)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func httpMockServer(statusCode int) *httptest.Server {
	// Set up a mock server that returns a response with the given status code.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(`{"foo": "bar"}`))
	})

	return httptest.NewServer(handler)
}
