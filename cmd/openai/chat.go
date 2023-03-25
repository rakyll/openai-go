package main

import (
	"context"
	"log"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/chat"
	"github.com/spf13/cobra"
)

type chatCmd struct {
	*cobra.Command
	Prompt string
	Role   string
	Engine string
}

func (c chatCmd) Init() error {
	c.Command = &cobra.Command{
		Use:   "chat",
		Short: "chat with OpenAI (ChatGPT)",
		Long:  `Chat with OpenAI (ChatGPT)`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			s := openai.NewSession(apiKey)

			client := chat.NewClient(s, c.Engine)
			resp, err := client.CreateCompletion(ctx, &chat.CreateCompletionParams{
				Messages: []*chat.Message{
					{Role: c.Role, Content: c.Prompt},
				},
			})
			if err != nil {
				log.Fatalf("Failed to complete: %v", err)
			}

			for _, choice := range resp.Choices {
				msg := choice.Message
				log.Printf("role=%q, content=%q", msg.Role, msg.Content)
			}
		},
	}

	flags := c.Flags()

	c.Prompt = *flags.String("prompt", "", "prompt to send to OpenAI")
	_ = c.MarkFlagRequired("prompt")

	c.Role = *flags.String("role", "user", "Role of the message")

	c.Engine = *flags.String("engine", "gpt-3.5-turbo", "Engine to use for the chat")

	rootCmd.AddCommand(c.Command)
	return nil
}

var _ = chatCmd{}.Init()
