package gollm

import (
	"context"
	"fmt"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/bedrock"
	"github.com/anthropics/anthropic-sdk-go/option"
	"k8s.io/klog/v2"
	"os"
)

func init() {
	if err := RegisterProvider("anthropic-bedrock", anthropicFactory); err != nil {
		klog.Fatalf("Failed to register awsbedrock provider: %v", err)
	}
}

// anthropicFactory is the provider factory function for Anthropic
// ClientOptions is not used, similar to Gemini
func anthropicFactory(ctx context.Context, opts ClientOptions) (Client, error) {
	opt := AnthropicAPIOptions{}
	return NewAnthropicAPIClient(ctx, opt)
}

// AnthropicAPIOptions contains options for the AnthropicBedrock API client
type AnthropicAPIOptions struct {
	// API Key for AnthropicBedrock
	APIKey string
	// Anthropic Client option for bedrock service
	AWSBedrock bool // Doc: https://docs.anthropic.com/en/api/claude-on-amazon-bedrock
}

// AnthropicAPIClient is a client for Anthropic Bedrock API
type AnthropicAPIClient struct {
	client anthropic.Client
}

type AnthropicAIChat struct {
	client  anthropic.Client
	model   string
	history []anthropic.TextBlock
	system  string
}

type AnthropicContent struct {
	Content string
	Role    string
}

func (c *AnthropicAIChat) Send(ctx context.Context, contents ...any) (ChatResponse, error) {
	msgNewParam := anthropic.MessageNewParams{
		MaxTokens:     2048,
		Messages:      nil,
		Model:         c.model,
		Temperature:   anthropic.Float(1.0),
		Metadata:      anthropic.MetadataParam{},
		StopSequences: nil,
		System: []anthropic.TextBlockParam{
			anthropic.TextBlockParam{
				Text:         c.system,
				Citations:    nil,
				CacheControl: anthropic.CacheControlEphemeralParam{},
				Type:         "text",
			},
		},
		Thinking:   anthropic.ThinkingConfigParamUnion{},
		ToolChoice: anthropic.ToolChoiceUnionParam{},
		Tools:      nil,
	}
	message, err := c.client.Messages.New(ctx, msgNewParam)
	if err != nil {
		return nil, err
	}
	for _, msgContent := range message.Content {
		c.history = append(c.history, msgContent.AsResponseTextBlock())
	}

}

func (c *AnthropicAIChat) partsToClaude(content ...any) {

}

func (c *AnthropicAIChat) SendStreaming(ctx context.Context, contents ...any) (ChatResponseIterator, error) {
	//TODO implement me
	panic("implement me")
}

func (c *AnthropicAIChat) SetFunctionDefinitions(functionDefinitions []*FunctionDefinition) error {
	//TODO implement me
	panic("implement me")
}

func (c *AnthropicAIChat) IsRetryableError(err error) bool {
	//TODO implement me
	panic("implement me")
}

func NewAnthropicAPIClient(ctx context.Context, opts AnthropicAPIOptions) (Client, error) {
	if opts.AWSBedrock {
		client := anthropic.NewClient(
			// bedrock supports aws.Config as well with `bedrock.WithConfig`
			bedrock.WithLoadDefaultConfig(ctx),
		)
		return &AnthropicAPIClient{client: client}, nil
	}

	apiKey := opts.APIKey
	if apiKey == "" {
		apiKey = os.Getenv("ANTHROPIC_API_KEY")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("ANTHROPIC_API_KEY environment varaiable not set")
	}
	client := anthropic.NewClient(
		option.WithAPIKey(apiKey),
	)
	return &AnthropicAPIClient{client: client}, nil
}

var _ Client = &AnthropicAPIClient{}

// Close free resource
func (c *AnthropicAPIClient) Close() error {
	// There are no actual way to close client in `anthropic-go-sdk`
	return nil
}

func (c *AnthropicAPIClient) StartChat(systemPrompt string, model string) Chat {
	return AnthropicAIChat{
		client:  c.client,
		model:   model,
		history: []string{},
		system:  systemPrompt,
	}
}

func (c *AnthropicAPIClient) GenerateCompletion(ctx context.Context, req *CompletionRequest) (CompletionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *AnthropicAPIClient) SetResponseSchema(responseSchema *Schema) error {
	return nil
}

func (c *AnthropicAPIClient) ListModels(ctx context.Context) ([]string, error) {
	// AutoPaging can also be used here for easier iterating, but I want to handel the error respond here
	// AutoPaging Doc: https://pkg.go.dev/github.com/anthropics/anthropic-sdk-go@v0.2.0-beta.3#readme-pagination
	modelInfoPage, err := c.client.Models.List(ctx, anthropic.ModelListParams{
		// Default values for limit is 20 max range is 1 - 1000
		// https://github.com/anthropics/anthropic-sdk-go/blob/v0.2.0-beta.3/model.go#L116
		Limit: anthropic.Int(1000),
	})
	if err != nil {
		return nil, err
	}
	var models []string
	for modelInfoPage != nil {
		for _, m := range modelInfoPage.Data {
			models = append(models, m.DisplayName)
		}
		// GetNextPage returns nil if there are no more pages left, will not error
		modelInfoPage, err = modelInfoPage.GetNextPage()
		if err != nil {
			return nil, err
		}
	}
	return models, nil
}
