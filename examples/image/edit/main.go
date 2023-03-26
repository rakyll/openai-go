package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/image"
)

func main() {
	ctx := context.Background()

	s := openai.NewSession(os.Getenv("OPENAI_API_KEY"))
	client := image.NewClient(s)

	imageFilePath := os.Getenv("IMAGE_FILE_PATH")
	if imageFilePath == "" {
		log.Fatal("must provide an IMAGE_FILE_PATH env var")
	}
	img, err := os.Open(imageFilePath)
	if err != nil {
		log.Fatalf("error opening image file: %v", err)
	}
	defer img.Close()

	maskFilePath := os.Getenv("MASK_FILE_PATH")
	if maskFilePath == "" {
		log.Fatal("must provide an MASK_FILE_PATH env var")
	}
	mask, err := os.Open(maskFilePath)
	if err != nil {
		log.Fatalf("error opening mask file: %v", err)
	}
	defer mask.Close()

	var n *int
	resp, err := client.CreateEdit(ctx, &image.CreateEditParams{
		N:      n,
		Prompt: "A cute baby sea otter wearing a beret",
		Size:   "1024x1024",
		Image:  img,
		Mask:   mask,
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
		_ = data // use data

	}
}
