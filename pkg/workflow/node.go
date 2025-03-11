package workflow

import (
	"context"
	"github.com/aiagt/aiagt/pkg/schema"
	"github.com/cloudwego/eino-ext/components/model/openai"
	openaigo "github.com/cloudwego/eino-ext/libs/acl/openai"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	einoschema "github.com/cloudwego/eino/schema"
)

type Node struct {
	Params *NodeParams `json:"params"`
	Runner NodeRunner  `json:"-"`
}

func NewNode(params *NodeParams, runner NodeRunner) *Node {
	return &Node{Params: params, Runner: runner}
}

type NodeParams struct {
	Name         string             `gorm:"column:name;NOT NULL" json:"name"`
	InputMapper  ObjectMapper       `gorm:"column:input_mapper;serializer:json;type:json" json:"input_mapper"`
	OutputSchema *schema.Definition `gorm:"column:output_schema;serializer:json;type:json" json:"output_schema"`
	BatchField   *ObjectField       `gorm:"column:batch_field;serializer:json;type:json" json:"batch_field"`
	Start        bool               `gorm:"column:start" json:"start"`
	End          bool               `gorm:"column:end" json:"end"`
}

func (node *Node) Lambda() *compose.Lambda {
	run := func(ctx context.Context, input Object) (Object, error) {
		return node.Runner.Run(ctx, node.Params, input)
	}

	if node.Params.BatchField != nil {
		return NodeLambdaBatch(node.Params.Name, node.Params.InputMapper, ArraySplitter(node.Params.BatchField), run)
	}

	if node.Params.Start {
		return NodeLambdaStart(run)
	}

	if node.Params.End {
		return NodeLambdaEnd(node.Params.InputMapper, run)
	}

	return NodeLambda(node.Params.Name, node.Params.InputMapper, run)
}

func NewStartNode() *Node {
	return &Node{
		Params: &NodeParams{
			Name:  NodeNameStart,
			Start: true,
		},
		Runner: NewDirectNodeRunner(),
	}
}

func NewEndNode(inputMapper ObjectMapper) *Node {
	return &Node{
		Params: &NodeParams{
			Name:        NodeNameEnd,
			InputMapper: inputMapper,
			End:         true,
		},
		Runner: NewDirectNodeRunner(),
	}
}

type NodeRunner interface {
	Run(ctx context.Context, params *NodeParams, input Object) (Object, error)
}

type FunctionNodeRunner struct {
	runner func(ctx context.Context, input Object) (Object, error)
}

func NewFunctionNodeRunner(runner func(ctx context.Context, input Object) (Object, error)) *FunctionNodeRunner {
	return &FunctionNodeRunner{runner: runner}
}

func (r *FunctionNodeRunner) Run(ctx context.Context, _ *NodeParams, input Object) (Object, error) {
	return r.runner(ctx, input)
}

type LLMNodeRunner struct {
	baseURL      string
	apiKey       string
	model        string
	systemPrompt string
	userPrompt   string
}

func NewLLMNodeRunner(baseURL, apiKey, model, systemPrompt, userPrompt string) *LLMNodeRunner {
	return &LLMNodeRunner{
		baseURL:      baseURL,
		apiKey:       apiKey,
		model:        model,
		systemPrompt: systemPrompt,
		userPrompt:   userPrompt,
	}
}

func (r *LLMNodeRunner) Run(ctx context.Context, params *NodeParams, input Object) (Object, error) {
	template := prompt.FromMessages(einoschema.FString,
		&einoschema.Message{
			Role:    einoschema.System,
			Content: r.systemPrompt,
		},
		&einoschema.Message{
			Role:    einoschema.User,
			Content: r.userPrompt,
		},
	)

	messages, err := template.Format(ctx, input)
	if err != nil {
		return nil, err
	}

	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL: r.baseURL,
		APIKey:  r.apiKey,
		Model:   r.model,
		ResponseFormat: &openaigo.ChatCompletionResponseFormat{
			Type: openaigo.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openaigo.ChatCompletionResponseFormatJSONSchema{
				Name:   "output",
				Schema: params.OutputSchema.SchemaV3(),
				Strict: true,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	result, err := chatModel.Generate(ctx, messages)
	if err != nil {
		return nil, err
	}

	output, err := NewJSONObject([]byte(result.Content))
	if err != nil {
		return nil, err
	}

	return output, nil
}

type DirectNodeRunner struct{}

func NewDirectNodeRunner() *DirectNodeRunner {
	return &DirectNodeRunner{}
}

func (r *DirectNodeRunner) Run(ctx context.Context, _ *NodeParams, input Object) (Object, error) {
	return input, nil
}
