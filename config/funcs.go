package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// GetConfigFilePath 获取配置文件在操作系统的完整路径
func GetConfigFilePath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, ConfigFileName), nil
}

// ValidateConfig 验证配置是否完整
func (c *Config) ValidateConfig() error {
	if c.OpenAIKey == "" {
		return fmt.Errorf("%w: OpenAIKey", ErrRequiredFieldEmpty)
	}
	// 可以添加其他必填字段的验证
	return nil
}

// 检查配置文件是否存在
func IsConfigFileExists() (bool, error) {
	configPath, err := GetConfigFilePath()
	if err != nil {
		return false, err
	}
	// 检查文件是否已存在
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		return true, nil
	}
	return false, nil
}

// CreateConfigFile 创建新的配置文件
func CreateConfigFile() error {
	configPath, err := GetConfigFilePath()
	if err != nil {
		return err
	}
	// 检查文件是否已存在
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		return fmt.Errorf("配置文件已存在: %s", configPath)
	}
	// 创建默认配置
	fmt.Println("创建默认配置")
	config := NewDefaultConfig()
	return SaveConfig(config)
}

// LoadConfig 加载配置文件
func LoadConfig() (*Config, error) {
	configPath, err := GetConfigFilePath()
	if err != nil {
		return nil, err
	}
	// 检查文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, ErrConfigNotFound
	}
	// 读取文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	// 解析YAML
	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, ErrConfigInvalid
	}
	return config, nil
}

// SaveConfig 保存配置到文件
func SaveConfig(config *Config) error {
	configPath, err := GetConfigFilePath()
	if err != nil {
		return err
	}
	// 转换为YAML
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	// 添加注释
	content := []byte(`# AGCommit Configuration File
# Generated automatically - DO NOT EDIT MANUALLY unless you know what you're doing
` + string(data))
	// 写入文件
	return os.WriteFile(configPath, content, 0644)
}

// UpdateConfigField 更新单个配置字段
func UpdateConfigField(fieldName, value string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}
	// 根据字段名更新相应的值
	switch fieldName {
	case "openai_key":
		config.OpenAIKey = value
	case "openai_api_base":
		config.OpenAPIBase = value
	case "openai_model":
		config.OpenAIModel = value
	case "commit_locale":
		config.CommitLocale = value
	case "commit_type":
		config.CommitType = value
	case "max_length":
		// 需要转换为int
		var length int
		if _, err := fmt.Sscanf(value, "%d", &length); err != nil {
			return fmt.Errorf("invalid max_length value: %s", value)
		}
		config.MaxLength = length
	case "auto_add":
		// 需要转换为bool
		if value == "true" {
			config.AutoAdd = true
		} else if value == "false" {
			config.AutoAdd = false
		} else {
			return fmt.Errorf("invalid auto_add value: %s (should be true or false)", value)
		}
	case "auto_commit":
		// 需要转换为bool
		if value == "true" {
			config.AutoCommit = true
		} else if value == "false" {
			config.AutoCommit = false
		} else {
			return fmt.Errorf("invalid auto_commit value: %s (should be true or false)", value)
		}
	default:
		return fmt.Errorf("unknown field: %s", fieldName)
	}
	return SaveConfig(config)
}

// ListConfigFields 列出所有配置字段及其值
func ListConfigFields() (map[string]interface{}, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"openai_key":      config.OpenAIKey,
		"openai_api_base": config.OpenAPIBase,
		"openai_model":    config.OpenAIModel,
		"commit_locale":   config.CommitLocale,
		"max_length":      config.MaxLength,
		"commit_type":     config.CommitType,
		"auto_add":        config.AutoAdd,
		"auto_commit":     config.AutoCommit,
	}, nil
}

// RemoveConfig 删除配置文件
func RemoveConfig() error {
	configPath, err := GetConfigFilePath()
	if err != nil {
		return err
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return ErrConfigNotFound
	}
	return os.Remove(configPath)
}

// CheckConfig 检查配置文件是否存在且格式正确
func CheckConfig() error {
	config, err := LoadConfig()
	if err != nil {
		fmt.Println("检查配置文件时出错:", err)
		return err
	}
	return config.ValidateConfig()
}
