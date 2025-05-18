package gollm

import (
	"context"
	"k8s.io/klog/v2"
)

func init() {
	if err := RegisterProvider("anthropic-bedrock", anthropicBedrockFactory); err != nil {
		klog.Fatalf("Failed to register awsbedrock provider: %v", err)
	}
}

// anthropicBedrockFactory is the provider factory function for Anthropic
func anthropicBedrockFactory(ctx context.Context, opts ClientOptions) (Client, error) {
	return NewAnthropicBedrockClient(ctx, opts)
}

// AnthropicBedrockOptions contains options for the AnthropicBedrock API client
type AnthropicBedrockOptions struct{}

// AnthropicBedrockClient is a client for Anthropic Bedrock API
type AnthropicBedrockClient struct{}

func (a AnthropicBedrockClient) Close() error {
	//TODO implement me
	panic("implement me")
}

func (a AnthropicBedrockClient) StartChat(systemPrompt, model string) Chat {
	//TODO implement me
	panic("implement me")
}

func (a AnthropicBedrockClient) GenerateCompletion(ctx context.Context, req *CompletionRequest) (CompletionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AnthropicBedrockClient) SetResponseSchema(schema *Schema) error {
	//TODO implement me
	panic("implement me")
}

func (a AnthropicBedrockClient) ListModels(ctx context.Context) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func NewAnthropicBedrockClient(ctx context.Context, opts ClientOptions) (Client, error) {
	return AnthropicBedrockClient{}, nil
}
