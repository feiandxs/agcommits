package openai_api

import (
	"fmt"

	"github.com/feiandxs/agcommits/constants"
	"github.com/feiandxs/agcommits/utils"
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
- fix: A bug fix
  
  you should generate a commit message like : 
  feat: xxxxx
  or 
  fix: xxxxx
  or 
  refactor: xxxxx
  or 
  perf: xxxxx
  or 
  test: xxxxx
`,
}

// generatePrompt 生成提示字符串
func generatePrompt(config *utils.Config, diff string) string {
	// 确定提交类型
	commitType := "conventional"
	if config.CommitType == "conventional" {
		commitType = "conventional"
	}
	// fmt.Println("commitType is ", commitType)
	format := commitTypeFormats[commitType]
	description := commitTypeDescriptions[commitType]

	// 获取语言的AI提示名称
	languageName := constants.GetLanguagePromptName(constants.LanguageCode(config.CommitLocale))

	// 根据语言添加额外要求
	languageRequirement := ""
	if config.CommitLocale == "en" {
		languageRequirement = "IMPORTANT: Use only lowercase letters in the commit message. No uppercase letters allowed.\n"
	}

	prompt := fmt.Sprintf(
		"You are an experienced programmer who writes great commit messages.\n"+
			"Generate a concise git commit message written in present tense for the following code diff with the given specifications below:\n"+
			"Message language: %s\n"+
			"Commit message must be a maximum of %d characters.\n"+
			"%s"+
			"Exclude anything unnecessary such as translation. Your entire response will be passed directly into git commit.\n"+
			"IMPORTANT: Return ONLY the commit message itself. Do NOT include any markdown formatting, code blocks, or ``` symbols.\n"+
			"%s\n%s\n\n"+
			"Git Diff:\n%s",
		languageName, config.MaxLength, languageRequirement, description, format, diff,
	)
	return prompt
}
