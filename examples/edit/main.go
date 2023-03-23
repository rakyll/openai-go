package main

import (
	"context"
	"log"
	"os"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/edit"
)

func main() {
	ctx := context.Background()
	s := openai.NewSession(os.Getenv("OPENAI_API_KEY"))

	client := edit.NewClient(s, "text-davinci-edit-001")
	resp, err := client.Create(ctx, &edit.CreateParams{
		N:           1,
		Input:       "What day of the wek is it?",
		Instruction: "Fix the spelling mistakes",
	})
	if err != nil {
		log.Fatalf("Failed to create an edit: %v", err)
	}

	for _, choice := range resp.Choices {
		log.Println(choice.Text)
	}
}
