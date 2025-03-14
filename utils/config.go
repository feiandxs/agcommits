package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/feiandxs/agcommits/constants"
)

func ReadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("无法打开文件: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("关闭文件时出错: %v\n", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	config := &Config{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("配置格式错误: %s", line)
		}

		key, value := parts[0], parts[1]
		switch key {
		case "OPENAI_KEY":
			config.OpenAIKey = value
		case "OPENAI_API_BASE":
			config.OpenAPIBase = value
		case "OPENAI_MODEL":
			config.OpenAIModel = value
		case "COMMIT_LOCALE":
			config.CommitLocale = value
		case "MAX_LENGTH":
			// 将字符串转换为 int
			config.MaxLength, err = strconv.Atoi(value)
			if err != nil {
				// 错误处理
			}
		case "COMMIT_TYPE":
			config.CommitType = value
		default:
			// 处理未知的配置键
			fmt.Printf("未知的配置键: %s\n", key)
		}
	}

	return config, nil
}

// checkConfigKeyExists 检查指定的配置键是否存在于配置文件中。
func checkConfigKeyExists(filePath, key string) (bool, error) {
	config, err := ReadConfig(filePath)
	if err != nil {
		return false, fmt.Errorf("读取配置时出错: %v", err)
	}

	switch key {
	case "OPENAI_KEY":
		return config.OpenAIKey != "", nil
	case "OPENAI_API_BASE":
		return config.OpenAPIBase != "", nil
	case "OPENAI_MODEL":
		return config.OpenAIModel != "", nil
	case "COMMIT_LOCALE":
		return config.CommitLocale != "", nil
	case "MAX_LENGTH":
		return config.MaxLength != 0, nil
	case "COMMIT_TYPE":
		return config.CommitType == "conventional" || config.CommitType == "", nil
	default:
		return false, fmt.Errorf("未知的配置键: %s", key)
	}
}

// CheckRequiredConfigsInFile 检查配置文件中是否存在指定的一组配置键。
func CheckRequiredConfigsInFile(requiredKeys []string) error {
	agCommitsPath, err := constants.GetConfigFilePath()
	if err != nil {
		return fmt.Errorf("获取配置文件路径时出错: %v", err)
	}

	for _, key := range requiredKeys {
		exists, err := checkConfigKeyExists(agCommitsPath, key)
		if err != nil {
			return fmt.Errorf("检查时出错: %v", err)
		}
		if !exists {
			return fmt.Errorf("缺少必要的配置: %s", key)
		}
	}

	fmt.Println("配置正确.")
	return nil
}

// CreateAndSetupConfigFile 创建配置文件，并设置必要的配置项。
func CreateAndSetupConfigFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer func() {
		cerr := file.Close()
		if cerr != nil {
			fmt.Printf("关闭文件时出现错误: %v\n", cerr)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	// 获取 OPENAI_KEY
	fmt.Print("请输入 OPENAI_KEY: ")
	scanner.Scan()
	openAIKey := scanner.Text()

	// 获取 OPENAI_API_BASE，带默认值
	fmt.Print("请输入 OPENAI_API_BASE (默认: https://api.openai.com): ")
	scanner.Scan()
	openAPIBase := scanner.Text()
	if openAPIBase == "" {
		openAPIBase = "https://api.openai.com"
	}

	// 获取 OPENAI_MODEL，带默认值
	fmt.Print("请输入 OPENAI_MODEL (默认: gpt-3.5-turbo-1106): ")
	scanner.Scan()
	openAIModel := scanner.Text()
	if openAIModel == "" {
		openAIModel = "gpt-3.5-turbo-1106"
	}

	// 获取 COMMIT_LOCALE
	fmt.Print("请输入 COMMIT_LOCALE（提交消息的语言环境，如 'en' 或 'zh'）: ")
	scanner.Scan()
	commitLocale := scanner.Text()

	// 获取 MAX_LENGTH，带默认值
	fmt.Print("请输入 MAX_LENGTH（提交消息的最大长度，默认: 150）: ")
	scanner.Scan()
	maxLengthStr := scanner.Text()
	var maxLength int
	if maxLengthStr == "" {
		maxLength = 150
	} else {
		maxLength, err = strconv.Atoi(maxLengthStr)
		if err != nil {
			return fmt.Errorf("无效的 MAX_LENGTH: %v", err)
		}
	}

	// 获取 COMMIT_TYPE
	fmt.Print("请输入 COMMIT_TYPE（默认: '', 可选: 'conventional'）: ")
	scanner.Scan()
	commitType := scanner.Text()
	if commitType != "conventional" {
		commitType = "" // 设置为默认值
	}

	// 写入配置文件
	_, err = file.WriteString(fmt.Sprintf("OPENAI_KEY=%s\nOPENAI_API_BASE=%s\nOPENAI_MODEL=%s\nCOMMIT_LOCALE=%s\nMAX_LENGTH=%d\nCOMMIT_TYPE=%s\n", openAIKey, openAPIBase, openAIModel, commitLocale, maxLength, commitType))
	if err != nil {
		return fmt.Errorf("写入配置失败: %v", err)
	}

	fmt.Println("配置文件写入成功.")
	return nil
}
