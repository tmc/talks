package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	ctx := context.Background()
	// START OMIT
	llm, err := ollama.New(ollama.WithModel("mistral"))
	if err != nil {
		log.Fatal(err)
	}
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem,
			"You are a company branding design wizard."),
		llms.TextParts(llms.ChatMessageTypeHuman,
			"What would be a good company name a company that makes colorful socks?"),
	}
	completion, err := llm.GenerateContent(ctx, messages, llms.WithStreamingFunc(
		func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	// END OMIT
	if err != nil {
		log.Fatal(err)
	}
	_ = completion
}
