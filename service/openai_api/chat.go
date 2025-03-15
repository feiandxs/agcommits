package openai_api

import (
	"context"
	"fmt"

	"github.com/feiandxs/agcommits/config"
	"github.com/feiandxs/agcommits/utils"
	"github.com/sashabaranov/go-openai"
)

// 将config.Config转换为utils.Config
func convertConfig(cfg *config.Config) *utils.Config {
	return &utils.Config{
		OpenAIKey:    cfg.OpenAIKey,
		OpenAPIBase:  cfg.OpenAPIBase,
		OpenAIModel:  cfg.OpenAIModel,
		CommitLocale: cfg.CommitLocale,
		MaxLength:    cfg.MaxLength,
		CommitType:   cfg.CommitType,
	}
}

// GenerateCommitMessage 使用 OpenAI API 生成提交信息
func GenerateCommitMessage(cfg *config.Config, diff string) (string, error) {
	client := openai.NewClient(cfg.OpenAIKey)
	if cfg.OpenAPIBase != "" {
		config := openai.DefaultConfig(cfg.OpenAIKey)
		config.BaseURL = cfg.OpenAPIBase
		client = openai.NewClientWithConfig(config)
	}

	// 将config.Config转换为utils.Config
	utilsConfig := convertConfig(cfg)
	// 构建提示词
	prompt := generatePrompt(utilsConfig, diff)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: cfg.OpenAIModel,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: cfg.MaxLength,
		},
	)

	if err != nil {
		return "", fmt.Errorf("OpenAI API 调用失败: %v", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("OpenAI API 返回结果为空")
	}

	return resp.Choices[0].Message.Content, nil
}
