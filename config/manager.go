package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

const (
	currentVersion = "1.0.0"
)

type configManagerImpl struct {
	configPath string
	config     *Config
}

// NewConfigManager 创建新的配置管理器
func NewConfigManager() ConfigManager {
	return &configManagerImpl{
		configPath: filepath.Join(os.Getenv("HOME"), ".agcommits", "config.json"),
	}
}

// LoadOrInit 加载或初始化配置
func (cm *configManagerImpl) LoadOrInit() ConfigResult {
	// 1. 尝试加载现有配置
	config, err := cm.loadExisting()
	if err == nil {
		return ConfigResult{Config: config, IsNew: false}
	}

	// 2. 如果配置不存在，创建新配置
	if os.IsNotExist(err) {
		return cm.InitializeInteractive()
	}

	// 3. 其他错误
	return ConfigResult{Error: err}
}

// InitializeInteractive 交互式初始化配置
func (cm *configManagerImpl) InitializeInteractive() ConfigResult {
	config := &Config{Version: currentVersion}

	fmt.Println("\n欢迎使用 AGCommits！")
	fmt.Println("让我们开始进行配置初始化...\n")

	// 按顺序处理每个配置项
	for _, item := range getOrderedConfigItems() {
		value, err := cm.promptForValue(item)
		if err != nil {
			return ConfigResult{Error: fmt.Errorf("配置 %s 失败: %v", item.Key, err)}
		}

		if err := cm.SetConfigItem(item.Key, value); err != nil {
			return ConfigResult{Error: fmt.Errorf("设置 %s 失败: %v", item.Key, err)}
		}
	}

	cm.config = config

	// 创建配置目录
	if err := os.MkdirAll(filepath.Dir(cm.configPath), 0755); err != nil {
		return ConfigResult{Error: fmt.Errorf("创建配置目录失败: %v", err)}
	}

	// 保存配置
	if err := cm.Save(config); err != nil {
		return ConfigResult{Error: fmt.Errorf("保存配置失败: %v", err)}
	}

	fmt.Println("\n配置初始化完成！")
	return ConfigResult{Config: config, IsNew: true}
}

// promptForValue 提示用户输入配置值
func (cm *configManagerImpl) promptForValue(item ConfigItem) (interface{}, error) {
	// 如果是非必填且用户选择跳过
	if !item.Required {
		fmt.Printf("\n%s\n[%s]\n", item.Description, item.Default)
		fmt.Print("是否要修改此配置项？(y/N): ")
		if !promptYesNo(false) {
			return item.Default, nil
		}
	}

	for {
		fmt.Printf("\n%s\n", item.Prompt)

		switch item.Type {
		case TypeSelect:
			return cm.handleSelectPrompt(item)
		case TypeString:
			return cm.handleStringPrompt(item)
		case TypeInt:
			return cm.handleIntPrompt(item)
		case TypeBool:
			return cm.handleBoolPrompt(item)
		default:
			return nil, fmt.Errorf("不支持的配置类型: %s", item.Type)
		}
	}
}

// handleSelectPrompt 处理选择类型的配置项
func (cm *configManagerImpl) handleSelectPrompt(item ConfigItem) (interface{}, error) {
	fmt.Println("可选项：")
	for i, opt := range item.Options {
		fmt.Printf("%d. %v\n", i+1, opt)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入选项编号: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	input = strings.TrimSpace(input)
	if input == "" && !item.Required {
		return item.Default, nil
	}

	selection, err := strconv.Atoi(input)
	if err != nil || selection < 1 || selection > len(item.Options) {
		return nil, fmt.Errorf("无效的选择")
	}

	return item.Options[selection-1], nil
}

// handleStringPrompt 处理字符串类型的配置项
func (cm *configManagerImpl) handleStringPrompt(item ConfigItem) (interface{}, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	input = strings.TrimSpace(input)
	if input == "" && !item.Required {
		return item.Default, nil
	}

	if item.Validation != nil {
		if err := item.Validation(input); err != nil {
			return nil, err
		}
	}

	return input, nil
}

// handleIntPrompt 处理整数类型的配置项
func (cm *configManagerImpl) handleIntPrompt(item ConfigItem) (interface{}, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	input = strings.TrimSpace(input)
	if input == "" && !item.Required {
		return item.Default, nil
	}

	value, err := strconv.Atoi(input)
	if err != nil {
		return nil, fmt.Errorf("请输入有效的整数")
	}

	if item.Validation != nil {
		if err := item.Validation(value); err != nil {
			return nil, err
		}
	}

	return value, nil
}

// handleBoolPrompt 处理布尔类型的配置项
func (cm *configManagerImpl) handleBoolPrompt(item ConfigItem) (interface{}, error) {
	return promptYesNo(false), nil
}

// promptYesNo 提示用户输入是/否
func promptYesNo(defaultValue bool) bool {
	reader := bufio.NewReader(os.Stdin)
	defaultStr := "N"
	if defaultValue {
		defaultStr = "Y"
	}

	fmt.Printf("[y/n] (%s) ", defaultStr)
	input, err := reader.ReadString('\n')
	if err != nil {
		return defaultValue
	}

	input = strings.ToLower(strings.TrimSpace(input))
	if input == "" {
		return defaultValue
	}

	return input == "y" || input == "yes"
}

// loadExisting 加载现有配置
func (cm *configManagerImpl) loadExisting() (*Config, error) {
	data, err := os.ReadFile(cm.configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return &config, nil
}

// Save 保存配置
func (cm *configManagerImpl) Save(config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化配置失败: %v", err)
	}

	if err := os.MkdirAll(filepath.Dir(cm.configPath), 0755); err != nil {
		return fmt.Errorf("创建配置目录失败: %v", err)
	}

	if err := os.WriteFile(cm.configPath, data, 0644); err != nil {
		return fmt.Errorf("写入配置文件失败: %v", err)
	}

	return nil
}

// Validate 验证配置
func (cm *configManagerImpl) Validate() error {
	if cm.config == nil {
		return fmt.Errorf("配置未加载")
	}

	// 验证每个必填项
	for _, item := range ConfigItems {
		if !item.Required {
			continue
		}

		value, err := cm.GetConfigItem(item.Key)
		if err != nil {
			return fmt.Errorf("获取配置项 %s 失败: %v", item.Key, err)
		}

		if value == nil || (reflect.ValueOf(value).Kind() == reflect.String && value.(string) == "") {
			return fmt.Errorf("必填项 %s 未设置", item.Key)
		}

		if item.Validation != nil {
			if err := item.Validation(value); err != nil {
				return fmt.Errorf("配置项 %s 验证失败: %v", item.Key, err)
			}
		}
	}

	return nil
}

// GetConfigItem 获取配置项值
func (cm *configManagerImpl) GetConfigItem(key string) (interface{}, error) {
	if cm.config == nil {
		return nil, fmt.Errorf("配置未加载")
	}

	parts := strings.Split(key, ".")
	value := reflect.ValueOf(cm.config)

	for _, part := range parts {
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		if value.Kind() != reflect.Struct {
			return nil, fmt.Errorf("无效的配置路径: %s", key)
		}

		value = value.FieldByName(strings.Title(part))
		if !value.IsValid() {
			return nil, fmt.Errorf("配置项不存在: %s", key)
		}
	}

	return value.Interface(), nil
}

// SetConfigItem 设置配置项值
func (cm *configManagerImpl) SetConfigItem(key string, value interface{}) error {
	if cm.config == nil {
		cm.config = &Config{Version: currentVersion}
	}

	parts := strings.Split(key, ".")
	target := reflect.ValueOf(cm.config)

	if target.Kind() == reflect.Ptr {
		target = target.Elem()
	}

	for i, part := range parts {
		if i == len(parts)-1 {
			field := target.FieldByName(strings.Title(part))
			if !field.IsValid() {
				return fmt.Errorf("配置项不存在: %s", key)
			}

			val := reflect.ValueOf(value)
			if field.Type() != val.Type() {
				return fmt.Errorf("配置项类型不匹配: %s", key)
			}

			field.Set(val)
			break
		}

		target = target.FieldByName(strings.Title(part))
		if !target.IsValid() {
			return fmt.Errorf("配置路径无效: %s", key)
		}
	}

	return nil
}
