package config

import (
	"fmt"
	"strings"

	"github.com/feiandxs/agcommits/constants"
)

// 配置项定义
var ConfigItems = map[string]ConfigItem{
	"openai.api_key": {
		Key:         "openai.api_key",
		Type:        TypeString,
		Required:    true,
		Prompt:      "请输入您的 OpenAI API Key",
		Description: "用于访问 OpenAI API 的密钥",
		Validation: func(v interface{}) error {
			key, ok := v.(string)
			if !ok {
				return fmt.Errorf("API Key 必须是字符串")
			}
			if len(key) < 20 {
				return fmt.Errorf("API Key 长度不正确")
			}
			return nil
		},
	},
	"openai.base_url": {
		Key:         "openai.base_url",
		Type:        TypeString,
		Required:    false,
		Default:     constants.DefaultOpenAIBaseURL,
		Prompt:      "请输入 OpenAI API 基础 URL（可选，回车使用默认值）",
		Description: "OpenAI API 的基础 URL",
		Validation: func(v interface{}) error {
			url, ok := v.(string)
			if !ok {
				return fmt.Errorf("URL 必须是字符串")
			}
			if url != "" && !strings.HasPrefix(url, "http") {
				return fmt.Errorf("URL 必须以 http 或 https 开头")
			}
			return nil
		},
	},
	"openai.model": {
		Key:         "openai.model",
		Type:        TypeSelect,
		Required:    false,
		Default:     constants.DefaultOpenAIModel,
		Prompt:      "请选择 OpenAI 模型",
		Description: "用于生成提交信息的 OpenAI 模型",
		Options: []interface{}{
			"gpt-4",
			"gpt-4-turbo-preview",
			"gpt-3.5-turbo",
		},
	},
	"commit.language": {
		Key:         "commit.language",
		Type:        TypeSelect,
		Required:    false,
		Default:     string(constants.DefaultLanguage),
		Prompt:      "请选择提交信息的语言",
		Description: "生成提交信息时使用的语言",
		Options:     getSupportedLanguages(),
	},
	"commit.max_length": {
		Key:         "commit.max_length",
		Type:        TypeInt,
		Required:    false,
		Default:     constants.DefaultMaxLength,
		Prompt:      "请输入提交信息最大长度（可选，回车使用默认值）",
		Description: "生成的提交信息最大长度",
		Validation: func(v interface{}) error {
			length, ok := v.(int)
			if !ok {
				return fmt.Errorf("长度必须是整数")
			}
			if length < 20 || length > 500 {
				return fmt.Errorf("长度必须在 20-500 之间")
			}
			return nil
		},
	},
	"commit.type": {
		Key:         "commit.type",
		Type:        TypeString,
		Required:    false,
		Default:     constants.DefaultCommitType,
		Prompt:      "请输入默认的提交类型（可选，回车使用默认值）",
		Description: "默认的提交类型",
	},
	"preferences.default_branch": {
		Key:         "preferences.default_branch",
		Type:        TypeString,
		Required:    false,
		Default:     "main",
		Prompt:      "请输入默认分支名（可选，回车使用默认值）",
		Description: "默认的 Git 分支名",
	},
	"preferences.auto_add": {
		Key:         "preferences.auto_add",
		Type:        TypeBool,
		Required:    false,
		Default:     false,
		Prompt:      "是否自动执行 git add（y/n，可选，回车使用默认值）",
		Description: "是否自动执行 git add 命令",
	},
}

// 获取有序的配置项列表
func getOrderedConfigItems() []ConfigItem {
	// 定义配置项的顺序
	order := []string{
		"openai.api_key",
		"openai.base_url",
		"openai.model",
		"commit.language",
		"commit.max_length",
		"commit.type",
		"preferences.default_branch",
		"preferences.auto_add",
	}

	items := make([]ConfigItem, 0, len(order))
	for _, key := range order {
		if item, ok := ConfigItems[key]; ok {
			items = append(items, item)
		}
	}
	return items
}

// 获取支持的语言列表
func getSupportedLanguages() []interface{} {
	languages := constants.GetSupportedLanguages()
	result := make([]interface{}, len(languages))
	for i, lang := range languages {
		result[i] = string(lang)
	}
	return result
}
