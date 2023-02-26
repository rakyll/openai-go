package main

import (
	"context"
	"log"
	"os"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/completion"
)

func main() {
	ctx := context.Background()
	s := openai.NewSession(os.Getenv("OPENAI_API_KEY"), "")

	client := completion.NewClient(s, "text-davinci-003")
	resp, err := client.Create(ctx, &completion.CreateParameters{
		N:         1,
		MaxTokens: 200,
		Prompt:    []string{"say this is a test"},
	})
	if err != nil {
		log.Fatalf("Failed to complete: %v", err)
	}

	for _, choice := range resp.Choices {
		log.Println(choice.Text)
	}
}
