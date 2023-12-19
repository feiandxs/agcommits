package openai_api

import (
	"agcommits/utils"
	"fmt"
)

// 提交类型到格式的映射
var commitTypeFormats = map[string]string{
	"default":      "",
	"conventional": "<type>(<optional scope>): <commit message>",
}

// 提交类型到详细描述的映射
var commitTypeDescriptions = map[string]string{
	"default": "",
	"conventional": `
Choose a type that best describes the git diff:
- docs: Documentation only changes
- style: Changes that do not affect the meaning of the code (white-space, formatting, etc)
- refactor: A code change that neither fixes a bug nor adds a feature
- perf: A code change that improves performance
- test: Adding missing tests or correcting existing tests
- build: Changes that affect the build system or external dependencies
- ci: Changes to our CI configuration files and scripts
- chore: Other changes that don't modify src or test files
- revert: Reverts a previous commit
- feat: A new feature
- fix: A bug fix`,
}

// generatePrompt 生成提示字符串
func generatePrompt(config *utils.Config) string {
	// 确定提交类型
	commitType := "default"
	if config.CommitType == "conventional" {
		commitType = "conventional"
	}

	format := commitTypeFormats[commitType]
	description := commitTypeDescriptions[commitType]

	prompt := fmt.Sprintf(
		"Generate a concise git commit message written in present tense for the following code diff with the given specifications below:\n"+
			"Message language: %s\n"+
			"Commit message must be a maximum of %d characters.\n"+
			"Exclude anything unnecessary such as translation. Your entire response will be passed directly into git commit.\n"+
			"%s\n%s",
		config.CommitLocale, config.MaxLength, description, format,
	)
	return prompt
}
