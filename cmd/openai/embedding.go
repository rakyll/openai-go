package main

import (
	"context"
	"log"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/embedding"
	"github.com/spf13/cobra"
)

type embeddingCmd struct {
	*cobra.Command
	Prompt string
	Model  string
}

func (c embeddingCmd) Init() error {
	c.Command = &cobra.Command{
		Use:   "embedding",
		Short: "Interact with OpenAI's embedding API",
		Long:  `Interact with OpenAI's embedding API`,
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()
			s := openai.NewSession(apiKey)

			client := embedding.NewClient(s, c.Model)
			resp, err := client.Create(ctx, &embedding.CreateParams{
				Input: []string{c.Prompt},
			})
			if err != nil {
				log.Fatalf("Failed to complete: %v", err)
			}

			for _, data := range resp.Data {
				log.Printf("index=%d, len(embedding)=%d", data.Index, len(data.Embedding))
			}

		},
	}
	flags := c.Flags()

	c.Prompt = *flags.String("prompt", "", "prompt to send to OpenAI")
	_ = c.MarkFlagRequired("prompt")

	c.Model = *flags.String("model", "text-embedding-ada-002", "model of the interaction")

	rootCmd.AddCommand(c.Command)
	return nil
}

var _ = embeddingCmd{}.Init()
