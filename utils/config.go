package utils

import (
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/feiandxs/agcommits/config"
)

func GetConfigValue(cfg *config.Config, fieldName string) string {
	switch fieldName {
	case "openai_key":
		return cfg.OpenAIKey
	case "openai_api_base":
		return cfg.OpenAPIBase
	case "openai_model":
		return cfg.OpenAIModel
	case "commit_locale":
		return cfg.CommitLocale
	case "max_length":
		return strconv.Itoa(cfg.MaxLength)
	case "auto_add":
		return strconv.FormatBool(cfg.AutoAdd)
	case "auto_commit":
		return strconv.FormatBool(cfg.AutoCommit)
	default:
		return ""
	}
}

func AskForOptional(field config.ConfigField) bool {
	if field.Required {
		return true
	}

	var answer bool
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("是否要配置%s？(%s)", field.Help, field.Placeholder),
	}
	survey.AskOne(prompt, &answer)
	return answer
}
func PromptForValue(field config.ConfigField) (string, error) {
	var value string
	prompt := &survey.Input{
		Message: field.Help,
		Default: field.Placeholder,
	}
	if err := survey.AskOne(prompt, &value); err != nil {
		return "", err
	}
	return value, nil
}
