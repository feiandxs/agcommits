package main

import (
	"fmt"
	"strings"

	"github.com/feiandxs/agcommits/config"
	"github.com/feiandxs/agcommits/service/openai_api"
	"github.com/feiandxs/agcommits/utils"
)

func main() {
	// 创建配置管理器
	configManager := config.NewConfigManager()

	// 加载或初始化配置
	result := configManager.LoadOrInit()
	if result.Error != nil {
		fmt.Printf("配置初始化失败: %v\n", result.Error)
		return
	}

	// 使用配置继续后续流程
	cfg := result.Config

	// 检查是否在 Git 仓库中
	isInGitRepo, err := utils.IsGitRepository()
	if err != nil {
		fmt.Println("检查 Git 仓库时出错:", err)
		return
	}

	if !isInGitRepo {
		fmt.Println("当前目录不是 Git 仓库")
		// TODO: 根据配置决定是否自动初始化 Git 仓库
		return
	}

	// 检查暂存区
	hasStagedChanges, err := utils.HasStagedChanges()
	if err != nil {
		fmt.Println("检查暂存区时出错:", err)
		return
	}

	if !hasStagedChanges {
		if cfg.Preferences.AutoAdd {
			// 如果配置了自动 add，则执行 git add .
			if err := utils.GitAddAll(); err != nil {
				fmt.Println("执行 git add . 时出错:", err)
				return
			}
			fmt.Println("已自动执行 git add .")
		} else {
			fmt.Println("暂存区为空，请执行 'git add .'")
			return
		}
	}

	// 获取 Git Diff 信息
	diff, err := utils.GetGitDiff()
	if err != nil {
		fmt.Println("获取 Git Diff 信息时出错:", err)
		return
	}

	// 检查是否有实际的更改
	if strings.TrimSpace(diff) == "" {
		fmt.Println("没有发现任何文件更改")
		return
	}

	// 显示格式化的diff信息
	fmt.Println(utils.FormatDiff(diff))
	fmt.Println("正在生成提交信息...")

	// 调用 OpenAI API 生成提交信息
	res, err := openai_api.GenerateCommitMessage(cfg, diff)
	if err != nil {
		fmt.Println("调用 AI 生成时出错:", err)
		return
	}

	// 显示提交消息并获取用户确认
	if utils.ConfirmCommitMessage(res) {
		// 用户确认，执行 Git 提交
		if err := utils.PerformGitCommit(res); err != nil {
			fmt.Println("Git 提交时出错:", err)
			return
		}
		fmt.Println("Git 提交成功！")
	} else {
		fmt.Println("用户取消了提交操作。")
	}
}
