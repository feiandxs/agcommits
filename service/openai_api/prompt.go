package openai_api

import (
	"fmt"

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
- fix: A bug fix`,
}

// generatePrompt 生成提示字符串
func generatePrompt(config *utils.Config) string {
	commitType := "default"
	if config.CommitType == "conventional" {
		commitType = "conventional"
	}

	format := commitTypeFormats[commitType]
	description := commitTypeDescriptions[commitType]

	prompt := fmt.Sprintf(
		`Analyze the following code diff carefully and generate a meaningful git commit message:

1. First, examine the changes thoroughly:
   - What files were modified?
   - What is the core purpose of these changes?
   - How do these changes improve the codebase?

2. Then, generate a concise git commit message that:
   - Is written in present tense
   - Clearly explains the main purpose of the changes
   - Highlights any important technical decisions
   - Follows these specifications:
     - Language: %s
     - Maximum length: %d characters
     - Format: %s

3. Requirements:
   - Be specific and meaningful
   - Focus on the "what" and "why", not just the "how"
   - Exclude unnecessary information or translations
   %s

Your response should be a single, well-thought-out commit message that can be used directly in git commit.`,
		config.CommitLocale, config.MaxLength, format, description,
	)
	return prompt
}
