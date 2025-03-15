package config

type Config struct {
	OpenAIKey    string `yaml:"openai_key"`
	OpenAPIBase  string `yaml:"openai_api_base"`
	OpenAIModel  string `yaml:"openai_model"`
	CommitLocale string `yaml:"commit_locale"`
	MaxLength    int    `yaml:"max_length"`
	CommitType   string `yaml:"commit_type"`
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
	}
}
