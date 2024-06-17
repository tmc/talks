package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/llms"
	// START OMIT
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {
	//llm, _ := openai.New(openai.WithModel("gpt-4o"))
	ctx := context.Background()
	llm, _ := googleai.New(ctx, googleai.WithAPIKey(os.Getenv("GOOGLE_AI_KEY")))
	prompt := "Say hello from the stage at GopherCon.eu!"
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, prompt),
	}
	fmt.Println("generating..")
	_, err := llm.GenerateContent(ctx, messages,
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	fmt.Println(err)
}

// END OMIT
