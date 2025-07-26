package config

// Config 应用程序配置结构体
type Config struct {
	// OpenAI API 密钥，用于调用 AI 服务生成提交消息
	OpenAIKey string `yaml:"openai_key"`

	// OpenAI API 基础 URL，支持自定义 API 端点（如 SiliconFlow）
	OpenAPIBase string `yaml:"openai_api_base"`

	// OpenAI 模型名称，指定使用的 AI 模型
	OpenAIModel string `yaml:"openai_model"`

	// 提交消息的语言设置：zh（中文）或 en（英文）
	CommitLocale string `yaml:"commit_locale"`

	// 提交消息的最大字符长度限制
	MaxLength int `yaml:"max_length"`

	// 提交消息格式类型：conventional（约定式提交）或 default（默认格式）
	CommitType string `yaml:"commit_type"`

	// 是否自动执行 git add 命令，跳过用户确认（true：自动执行，false：需要确认）
	AutoAdd bool `yaml:"auto_add"`

	// 是否自动执行 git commit 命令，跳过用户确认（true：自动提交，false：需要确认）
	AutoCommit bool `yaml:"auto_commit"`
}

// NewDefaultConfig returns default configuration
func NewDefaultConfig() *Config {
	return &Config{
		OpenAIKey:    "",
		OpenAPIBase:  "",
		OpenAIModel:  "",
		CommitLocale: "zh",
		MaxLength:    150,
		CommitType:   "conventional",
		AutoAdd:      false, // 默认需要确认
		AutoCommit:   false, // 默认需要确认
	}
}
