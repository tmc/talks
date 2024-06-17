package main

// START OMIX
import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// END OMIX

// START OMIT
func main() {
	llm, _ := openai.New(openai.WithModel("gpt-4o"))
	ctx := context.Background()
	prompt := "Say hello from the stage at GopherCon.eu!"
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, prompt),
	}
	fmt.Println("generating..")
	llm.GenerateContent(ctx, messages,
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
}

// END OMIT
