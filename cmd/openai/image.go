package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/image"
	"github.com/spf13/cobra"
)

type imageCmd struct {
	*cobra.Command
	Prompt  string
	Size    string
	Format  string
	OutFile string
}

func (c imageCmd) Init() error {
	c.Command = &cobra.Command{
		Use:   "image",
		Short: "Interact with OpenAI's image API",
		Long:  `Interact with OpenAI's image API`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			s := openai.NewSession(apiKey)

			client := image.NewClient(s)
			resp, err := client.Create(ctx, &image.CreateParams{
				N:      3,
				Prompt: c.Prompt,
				Size:   c.Size,
				Format: c.Format,
			})
			if err != nil {
				log.Fatalf("Failed to generate image: %v", err)
			}

			for _, image := range resp.Data {
				reader, err := image.Reader()
				if err != nil {
					log.Fatalf("Failed to read image data: %v", err)
				}
				data, err := io.ReadAll(reader)
				if err != nil {
					log.Fatalf("ReadAll error: %v", err)
				}
				//write data to file
				if c.OutFile != "" {
					err = ioutil.WriteFile(c.OutFile, data, 0644)
					if err != nil {
						log.Fatalf("WriteFile error: %v", err)
					}
				}
			}
		},
	}
	flags := c.Flags()

	c.Prompt = *flags.String("prompt", "", "prompt to send to OpenAI")
	_ = c.MarkFlagRequired("prompt")

	c.Size = *flags.String("size", "1024x1024", "image size")
	c.Format = *flags.String("format", "b64_json", "image format")

	c.OutFile = *flags.String("out", "", "output file")

	rootCmd.AddCommand(c.Command)
	return nil
}

var _ = imageCmd{}.Init()
