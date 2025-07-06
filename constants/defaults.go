package constants

// OpenAI相关默认值
const (
	// DefaultOpenAIBaseURL OpenAI API的默认基础URL
	DefaultOpenAIBaseURL = "https://api.openai.com"

	// DefaultOpenAIModel OpenAI的默认模型
	DefaultOpenAIModel = "gpt-3.5-turbo-1106"

	// DefaultMaxLength 提交消息的默认最大长度
	DefaultMaxLength = 150
)

// 提交类型相关默认值
const (
	// DefaultCommitType 默认的提交类型
	DefaultCommitType = ""

	// ConventionalCommitType Conventional Commits规范的提交类型
	ConventionalCommitType = "conventional"
)
