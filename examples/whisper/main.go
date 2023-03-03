package main

import (
	"context"
	"log"
	"os"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/whisper"
)

func main() {
	sesh := openai.NewSession(os.Getenv("OPENAI_API_KEY"))
	wc := whisper.NewClient(sesh, "")
	filePath := os.Getenv("AUDIO_FILE_PATH")
	if filePath == "" {
		log.Fatal("must provide an AUDIO_FILE_PATH env var")
	}
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("error opening audio file: %v", err)
	}
	defer f.Close()
	resp, err := wc.Transcribe(context.TODO(), &whisper.CreateCompletionParams{
		Language:    "en",
		Audio:       f,
		AudioFormat: "mp3",
	})
	if err != nil {
		log.Fatalf("error transcribing file: %v", err)
	}
	log.Println(resp.Text)
}
