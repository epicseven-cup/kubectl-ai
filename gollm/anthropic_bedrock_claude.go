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

func NewAnthropicAPIClient(ctx context.Context, opts AnthropicAPIOptions) (Client, error) {
	if opts.AWSBedrock {
		client := anthropic.NewClient(
			bedrock.WithLoadDefaultConfig(ctx), // bedrock supports aws.Config as well with `bedrock.WithConfig`
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

func (a AnthropicAPIClient) Close() error {
	//TODO implement me
	panic("implement me")
}

func (a AnthropicAPIClient) StartChat(systemPrompt, model string) Chat {
	//TODO implement me
	panic("implement me")
}

func (a AnthropicAPIClient) GenerateCompletion(ctx context.Context, req *CompletionRequest) (CompletionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AnthropicAPIClient) SetResponseSchema(schema *Schema) error {
	//TODO implement me
	panic("implement me")
}

func (a AnthropicAPIClient) ListModels(ctx context.Context) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

var _ Client = &AnthropicAPIClient{}
