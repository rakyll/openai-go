package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/image"
)

func main() {
	ctx := context.Background()
	s := openai.NewSession(os.Getenv("OPENAI_API_KEY"))

	client := image.NewClient(s)
	resp, err := client.Create(ctx, &image.CreateParams{
		N:      3,
		Prompt: "a cute baby",
		Size:   "1024x1024",
		Format: "b64_json",
	})
	if err != nil {
		log.Fatalf("Failed to generate image: %v", err)
	}

	for _, image := range resp.Data {
		reader, err := image.Reader()
		if err != nil {
			log.Fatalf("Failed to read image data: %v", err)
		}
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Fatalf("ReadAll error: %v", err)
		}
		_ = data // use data
	}
}
