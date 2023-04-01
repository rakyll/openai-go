package main

import (
	"context"
	"log"
	"os"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/moderation"
)

func main() {
	ctx := context.Background()
	s := openai.NewSession(os.Getenv("OPENAI_API_KEY"))

	client := moderation.NewClient(s, "text-moderation-latest")
	resp, err := client.Create(ctx, &moderation.CreateParams{
		Input: []string{"I will kill you"},
	})
	if err != nil {
		log.Fatalf("Failed to complete: %v", err)
	}

	for _, result := range resp.Results {
		log.Println("Content moderation is flagged as", result.Flagged)
		if result.Flagged {
			for key, value := range result.Categories {
				if value {
					log.Println("Content category is", key)
				}
			}
		}
	}
}
