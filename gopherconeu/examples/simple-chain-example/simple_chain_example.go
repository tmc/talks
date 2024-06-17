package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/tools"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	llm, err := ollama.New(ollama.WithModel("mistral"),
		ollama.WithHTTPClient(debugHttpClient),
	)
	if err != nil {
		return err
	}
	promptInput := prompts.PromptTemplate{
		Template: "{{.Name}} is 36 years old and lives in {{.City}}.",
		PartialVariables: map[string]any{
			"Name": "Paul",
			"City": "Arrakeen",
		},
		TemplateFormat: prompts.TemplateFormatGoTemplate,
	}
	var tools []tools.Tool
	agent := agents.NewConversationalAgent(llm, tools, agents.WithPrompt(promptInput))
	agentExecutor := agents.NewExecutor(agent)
	ctx := context.Background()
	res, err := chains.Run(ctx, agentExecutor, "What is the name of the person?",
		chains.WithTemperature(0.8),
	)
	fmt.Println(res)
	return fmt.Errorf("issue in run: %w", err)
}

var debugHttpClient = &http.Client{
	Transport: &logTransport{http.DefaultTransport},
}

type logTransport struct {
	Transport http.RoundTripper
}

// RoundTrip logs the request and response with full contents using httputil.DumpRequest and httputil.DumpResponse
func (t *logTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s\n", dump)
	resp, err := t.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	dump, err = httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s\n", dump)
	return resp, nil
}
