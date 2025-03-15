package config

// ConfigItemType 表示配置项的类型
type ConfigItemType string

const (
	TypeString ConfigItemType = "string"
	TypeInt    ConfigItemType = "int"
	TypeBool   ConfigItemType = "bool"
	TypeSelect ConfigItemType = "select" // 用于选择类选项，如语言选择
)

// Config 表示完整的配置结构
type Config struct {
	// 基础配置
	Version string `json:"version"` // 配置文件版本

	// OpenAI 配置
	OpenAI struct {
		APIKey  string `json:"api_key"`
		BaseURL string `json:"base_url"`
		Model   string `json:"model"`
	} `json:"openai"`

	// 提交配置
	Commit struct {
		Language  string `json:"language"`   // 提交信息语言
		MaxLength int    `json:"max_length"` // 最大长度
		Type      string `json:"type"`       // 提交类型
	} `json:"commit"`

	// 用户偏好
	Preferences struct {
		DefaultBranch string `json:"default_branch"` // 默认分支
		AutoAdd       bool   `json:"auto_add"`       // 是否自动 git add
	} `json:"preferences"`
}

// ConfigItem 表示单个配置项的元数据
type ConfigItem struct {
	Key         string                  // 配置项键名
	Type        ConfigItemType          // 配置项类型
	Required    bool                    // 是否必填
	Default     interface{}             // 默认值
	Prompt      string                  // 提示信息
	Validation  func(interface{}) error // 验证函数
	Options     []interface{}           // 可选值（用于 TypeSelect 类型）
	Description string                  // 配置项描述
}

// ConfigResult 表示配置操作的结果
type ConfigResult struct {
	Config *Config
	IsNew  bool  // 是否是新创建的配置
	Error  error // 错误信息
}

// ConfigManager 定义配置管理器的接口
type ConfigManager interface {
	// 加载或初始化配置
	LoadOrInit() ConfigResult

	// 交互式配置初始化
	InitializeInteractive() ConfigResult

	// 验证配置
	Validate() error

	// 保存配置
	Save(*Config) error

	// 获取单个配置项
	GetConfigItem(key string) (interface{}, error)

	// 设置单个配置项
	SetConfigItem(key string, value interface{}) error
}
