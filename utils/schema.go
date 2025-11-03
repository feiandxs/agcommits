package utils

// Config 表示配置文件中的配置项
type Config struct {
	OpenAIKey         string
	OpenAPIBase       string
	OpenAIModel       string
	CommitLocale      string
	ResponseMaxLength int
	MaxLength         int
	CommitType        string
}

// EffectiveResponseMaxLength 返回用于限制 AI 响应长度的有效值
func (c *Config) EffectiveResponseMaxLength() int {
	if c.ResponseMaxLength > 0 {
		return c.ResponseMaxLength
	}
	return c.MaxLength
}
