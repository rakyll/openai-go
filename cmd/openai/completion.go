package main

import (
	"context"
	"log"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/completion"
	"github.com/spf13/cobra"
)

type completionCmd struct {
	*cobra.Command
	Model     string
	Prompt    string
	MaxTokens uint
}

func (c completionCmd) Init() error {
	c.Command = &cobra.Command{
		Use:   "completion",
		Short: "Interact with OpenAI's completion API",
		Long:  `Interact with OpenAI's completion API`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			s := openai.NewSession(apiKey)

			client := completion.NewClient(s, c.Model)
			resp, err := client.Create(ctx, &completion.CreateParams{
				N:         1,
				MaxTokens: int(c.MaxTokens),
				Prompt:    []string{c.Prompt},
			})
			if err != nil {
				log.Fatalf("Failed to complete: %v", err)
			}

			for _, choice := range resp.Choices {
				log.Println(choice.Text)
			}
		},
	}
	flags := c.Flags()

	c.Prompt = *flags.String("prompt", "", "prompt to send to OpenAI")
	_ = c.MarkFlagRequired("prompt")

	c.Model = *flags.String("model", "text-davinci-003", "model of the interaction")
	c.MaxTokens = *flags.Uint("max-tokens", 200, "maximum number of tokens to generate")

	rootCmd.AddCommand(c.Command)
	return nil
}

var _ = completionCmd{}.Init()
