package main

import (
	"context"
	"log"
	"os"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/embedding"
)

func main() {
	ctx := context.Background()
	s := openai.NewSession(os.Getenv("OPENAI_API_KEY"))

	client := embedding.NewClient(s, "text-embedding-ada-002")
	resp, err := client.Create(ctx, &embedding.CreateParams{
		Input: []string{"The food was delicious and the waiter..."},
	})
	if err != nil {
		log.Fatalf("Failed to complete: %v", err)
	}

	for _, data := range resp.Data {
		log.Printf("index=%d, len(embedding)=%d", data.Index, len(data.Embedding))
	}
}
