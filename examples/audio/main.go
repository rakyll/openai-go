package main

import (
	"context"
	"log"
	"os"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/audio"
)

func main() {
	ctx := context.Background()

	s := openai.NewSession(os.Getenv("OPENAI_API_KEY"))
	client := audio.NewClient(s, "")
	filePath := os.Getenv("AUDIO_FILE_PATH")
	if filePath == "" {
		log.Fatal("must provide an AUDIO_FILE_PATH env var")
	}
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("error opening audio file: %v", err)
	}
	defer f.Close()
	resp, err := client.CreateTranscription(ctx, &audio.CreateTranscriptionParams{
		Language:    "en",
		Audio:       f,
		AudioFormat: "mp3",
	})
	if err != nil {
		log.Fatalf("error transcribing file: %v", err)
	}
	log.Println(resp.Text)
}
