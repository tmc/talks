package main

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// START OMIT
func main() {
	ctx := context.Background()
	llm, err := openai.New(openai.WithModel("gpt-4o"))
	if err != nil {
		fmt.Println(err)
	}
	prompt := "Why is go a good language for LLM tools? Be concise."
	fmt.Println("generating..")
	out, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
}

// END OMIT
