package main

import (
	"context"
	"log"
	"os"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/audio"
	"github.com/spf13/cobra"
)

type audioCmd struct {
	*cobra.Command
	AudioFile string
	Language  string
}

func (c audioCmd) Init() error {

	c.Command = &cobra.Command{
		Use:   "audio",
		Short: "Interact with OpenAI's Whisper Engine",
		Long:  `Interact with OpenAI's Whisper Engine`,
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()

			s := openai.NewSession(apiKey)
			client := audio.NewClient(s, "")
			f, err := os.Open(c.AudioFile)
			if err != nil {
				log.Fatalf("error opening audio file: %v", err)
			}
			defer f.Close()
			resp, err := client.CreateTranscription(ctx, &audio.CreateTranscriptionParams{
				Language:    c.Language,
				Audio:       f,
				AudioFormat: "mp3",
			})
			if err != nil {
				log.Fatalf("error transcribing file: %v", err)
			}
			log.Println(resp.Text)

		},
	}

	flags := c.Flags()

	c.AudioFile = *flags.String("audio_file", "", "Path to the audio file to be transcribed")
	_ = c.MarkFlagRequired("audio_file")

	c.AudioFile = *flags.String("language", "en", "Language of the audio file")

	rootCmd.AddCommand(c.Command)
	return nil
}

var _ = audioCmd{}.Init()
