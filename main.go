package main

import (
	"fmt"
	"strings"

	fatihcolor "github.com/fatih/color"
	"github.com/feiandxs/agcommits/config"
	"github.com/feiandxs/agcommits/service/openai_api"
	"github.com/feiandxs/agcommits/utils"
	"github.com/shibukawa/cdiff"
)

func main() {
	// if config file not exists, create it
	exists, err := config.IsConfigFileExists()
	if err != nil {
		fmt.Println("检查配置文件时出错:", err)
		return
	}

	// 如果配置文件不存在，创建默认配置
	if !exists {
		if err := config.CreateConfigFile(); err != nil {
			fatihcolor.Red("创建配置文件失败: %v", err)
			return
		}
		fatihcolor.Green("已创建默认配置文件")
	}

	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		fatihcolor.Red("加载配置文件失败: %v", err)
		return
	}

	// 检查并补充缺失的必填项
	changed := false
	for _, field := range config.ConfigFields {
		value := utils.GetConfigValue(cfg, field.Name)
		if (field.Required && value == "") || (!field.Required && value == "" && utils.AskForOptional(field)) {
			newValue, err := utils.PromptForValue(field)
			if err != nil {
				fatihcolor.Red("获取输入失败: %v", err)
				return
			}
			if newValue != "" {
				if err := config.UpdateConfigField(field.Name, newValue); err != nil {
					fatihcolor.Red("更新配置失败: %v", err)
					return
				}
				changed = true
			}
		}
	}
	if changed {
		fatihcolor.Green("配置已更新")
	}
	// 最终验证
	if err := config.CheckConfig(); err != nil {
		fatihcolor.Red("配置验证失败: %v", err)
		return
	}

	fatihcolor.Green("配置验证通过")

	// 检查是否在 Git 仓库中
	isRepo, err := utils.IsGitRepository()
	if err != nil {
		fatihcolor.Red("检查 Git 仓库时出错: %v", err)
		return
	}

	// 如果不是Git仓库，询问用户是否要初始化一个
	if !isRepo {
		if utils.ConfirmGitInit() {
			fatihcolor.Yellow("正在初始化Git仓库...")
			if err := utils.InitGitRepository(); err != nil {
				fatihcolor.Red("初始化Git仓库失败: %v", err)
				return
			}
			fatihcolor.Green("Git仓库初始化成功")
		} else {
			fatihcolor.Yellow("未初始化Git仓库，程序退出")
			return
		}
	}

	// 检查是否有暂存的更改
	hasStagedChanges, err := utils.HasStagedChanges()
	if err != nil {
		fatihcolor.Red("检查暂存区状态失败: %v", err)
		return
	}

	// 如果没有暂存的更改，询问用户是否要执行 git add .
	if !hasStagedChanges {
		if utils.ConfirmGitAdd() {
			fatihcolor.Yellow("正在执行 git add . 命令...")
			if err := utils.GitAddAll(); err != nil {
				fatihcolor.Red("执行 git add . 命令失败: %v", err)
				return
			}
			fatihcolor.Green("成功将所有更改添加到暂存区")
		} else {
			fatihcolor.Yellow("未执行 git add . 命令，程序退出")
			return
		}
	}

	// 获取 Git 暂存区的 diff 信息
	diff, err := utils.GetGitDiff()
	if err != nil {
		fatihcolor.Red("获取 Git 暂存区 diff 信息失败: %v", err)
		return
	}

	// 打印 diff 信息
	if diff == "" {
		fatihcolor.Yellow("暂存区没有更改，无法生成提交消息")
		return
	} else {
		fatihcolor.Green("已获取暂存区的更改信息")

		// 使用cdiff库美化diff输出
		fmt.Println("\n========== 暂存区更改内容 ==========")
		// 获取当前目录作为文件路径前缀
		diffResult := cdiff.Diff("", diff, cdiff.LineByLine)
		// 使用String方法直接获取格式化后的字符串，然后手动添加颜色
		diffText := diffResult.String()
		printColoredDiff(diffText)
		fmt.Println("===================================\n")
	}

	// 使用 OpenAI API 生成提交消息
	fatihcolor.Yellow("正在使用 AI 生成提交消息...")
	commitMsg, err := openai_api.GenerateCommitMessage(cfg, diff)
	if err != nil {
		fatihcolor.Red("生成提交消息失败: %v", err)
		return
	}

	// 显示提交消息并询问用户是否确认使用
	if utils.ConfirmCommitMessage(commitMsg) {
		// 执行 Git 提交
		fatihcolor.Yellow("正在执行 Git 提交...")
		if err := utils.PerformGitCommit(commitMsg); err != nil {
			fatihcolor.Red("Git 提交失败: %v", err)
			return
		}
		fatihcolor.Green("Git 提交成功")
	} else {
		fatihcolor.Yellow("已取消 Git 提交")
	}
}

// printColoredDiff 打印彩色的 diff 输出
func printColoredDiff(diff string) {
	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			fmt.Println()
			continue
		}

		switch line[0] {
		case '+':
			if len(line) > 1 && line[1] == '+' {
				// 文件头部信息
				fatihcolor.Cyan("%s", line)
			} else {
				// 添加的行 - 绿色
				green := fatihcolor.New(fatihcolor.FgGreen, fatihcolor.Bold)
				green.Print("+ ")
				green.Println(line[1:])
			}
		case '-':
			if len(line) > 1 && line[1] == '-' {
				// 文件头部信息
				fatihcolor.Cyan("%s", line)
			} else {
				// 删除的行 - 红色
				red := fatihcolor.New(fatihcolor.FgRed, fatihcolor.Bold)
				red.Print("- ")
				red.Println(line[1:])
			}
		case '@':
			// 区块信息
			cyan := fatihcolor.New(fatihcolor.FgCyan)
			cyan.Println(line)
		case 'd':
			if strings.HasPrefix(line, "diff --git") {
				// diff 命令行
				yellow := fatihcolor.New(fatihcolor.FgYellow)
				yellow.Println(line)
			} else {
				fmt.Println(line)
			}
		case 'i':
			if strings.HasPrefix(line, "index ") {
				// index 行
				yellow := fatihcolor.New(fatihcolor.FgYellow)
				yellow.Println(line)
			} else {
				fmt.Println(line)
			}
		default:
			fmt.Println(line)
		}
	}
}
