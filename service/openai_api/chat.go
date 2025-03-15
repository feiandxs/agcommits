package openai_api

import (
	"context"
	"fmt"

	"github.com/feiandxs/agcommits/config"
	"github.com/sashabaranov/go-openai"
)

// GenerateCommitMessage 使用 OpenAI API 生成提交信息
func GenerateCommitMessage(cfg *config.Config, diff string) (string, error) {
	client := openai.NewClient(cfg.OpenAI.APIKey)
	if cfg.OpenAI.BaseURL != "" {
		config := openai.DefaultConfig(cfg.OpenAI.APIKey)
		config.BaseURL = cfg.OpenAI.BaseURL
		client = openai.NewClientWithConfig(config)
	}

	// 构建提示词
	prompt := buildPrompt(diff, cfg.Commit.Language)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: cfg.OpenAI.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: cfg.Commit.MaxLength,
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

// buildPrompt 构建用于生成提交信息的提示词
func buildPrompt(diff string, language string) string {
	// 基础提示词
	basePrompt := "请根据以下 Git diff 内容生成一个简洁的提交信息。\n" +
		"要求：\n" +
		"1. 使用 %s 语言\n" +
		"2. 简洁明了，突出重点\n" +
		"3. 使用祈使句\n" +
		"4. 不要超过 50 个字符\n\n" +
		"Git Diff 内容：\n%s"

	return fmt.Sprintf(basePrompt, language, diff)
}
