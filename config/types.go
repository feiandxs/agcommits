package config

type ConfigField struct {
	Name        string
	Required    bool
	Placeholder string
	Help        string
}

var ConfigFields = []ConfigField{
	{
		Name:        "openai_key",
		Required:    true,
		Placeholder: "sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		Help:        "OpenAI API密钥",
	},
	{
		Name:        "openai_api_base",
		Required:    true,
		Placeholder: "https://api.siliconflow.cn",
		Help:        "OpenAI API基础URL",
	},
	{
		Name:        "openai_model",
		Required:    true,
		Placeholder: "Qwen/Qwen2.5-Coder-7B-Instruct",
		Help:        "OpenAI模型名称",
	},
	{
		Name:        "commit_locale",
		Required:    false,
		Placeholder: "zh",
		Help:        "提交信息语言(zh/en)",
	},
	{
		Name:        "max_length",
		Required:    false,
		Placeholder: "150",
		Help:        "提交信息最大长度",
	},
}
