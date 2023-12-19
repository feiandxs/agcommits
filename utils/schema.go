package utils

// Config 表示配置文件中的配置项
type Config struct {
	OpenAIKey    string
	OpenAPIBase  string
	OpenAIModel  string
	CommitLocale string
	MaxLength    int
	CommitType   string
}
