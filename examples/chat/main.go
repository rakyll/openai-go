package main

import (
	"context"
	"log"
	"os"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/chat"
)

func main() {
	ctx := context.Background()
	s := openai.NewSession(os.Getenv("OPENAI_API_KEY"))

	client := chat.NewClient(s, "gpt-3.5-turbo")
	resp, err := client.CreateCompletion(ctx, &chat.CreateCompletionParams{
		Messages: []*chat.Message{
			{Role: "user", Content: "hello"},
		},
	})
	if err != nil {
		log.Fatalf("Failed to complete: %v", err)
	}

	for _, choice := range resp.Choices {
		msg := choice.Message
		log.Printf("role=%q, content=%q", msg.Role, msg.Content)
	}
}
