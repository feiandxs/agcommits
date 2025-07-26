package config

// ConfigField 配置字段定义结构体
type ConfigField struct {
	Name        string // 字段名称
	Required    bool   // 是否必填
	Placeholder string // 占位符/示例值
	Help        string // 帮助信息
}

// ConfigFields 所有配置字段的定义列表
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
	{
		Name:        "auto_add",
		Required:    false,
		Placeholder: "false",
		Help:        "是否自动执行git add命令，跳过确认步骤(true/false)",
	},
	{
		Name:        "auto_commit",
		Required:    false,
		Placeholder: "false",
		Help:        "是否自动执行git commit命令，跳过确认步骤(true/false)",
	},
}
