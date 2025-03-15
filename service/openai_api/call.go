// service/openai_api/openai.go

package openai_api

import (
	"context"
	"fmt"

	"github.com/feiandxs/agcommits/utils" // 导入 utils 包以使用 Config 结构体
	"github.com/sashabaranov/go-openai"
)

var client *openai.Client
var model string

func generateClient(config *utils.Config) {
	clientConfig := openai.DefaultConfig(config.OpenAIKey)
	clientConfig.BaseURL = config.OpenAPIBase
	client = openai.NewClientWithConfig(clientConfig)
	model = config.OpenAIModel
}

// ChatCompletionBlocking 接收一个 Config 类型的参数，并打印其值。
func ChatCompletionBlocking(config *utils.Config, diff string) (string, error) {
	generateClient(config)

	// 构建提示词
	prompt := generatePrompt(config, diff)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: config.MaxLength,
		})
	if err != nil {
		fmt.Printf("completion error: %v\n", err)
		return "", err
	}

	// 检查 resp.Choices 是否有元素并且第一个元素的 Message.Content 是否有效
	if len(resp.Choices) > 0 && resp.Choices[0].Message.Content != "" {
		return resp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("未接收到有效的 AI 响应")
}
