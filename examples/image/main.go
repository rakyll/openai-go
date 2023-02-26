package main

import (
	"context"
	"log"
	"os"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/image"
)

func main() {
	ctx := context.Background()
	s := openai.NewSession(os.Getenv("OPENAI_API_KEY"), "")

	client := image.NewClient(s)
	resp, err := client.Create(ctx, &image.CreateParameters{
		N:      3,
		Prompt: "a cute baby",
	})
	if err != nil {
		log.Fatalf("Failed to generate image: %v", err)
	}

	for _, image := range resp.Data {
		log.Println(image.URL)
	}
}
