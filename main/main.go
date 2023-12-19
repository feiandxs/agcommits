package main

import (
	"agcommits/service/openai_api"
	"agcommits/utils"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("获取当前用户时出错:", err)
		return
	}

	agCommitsPath := filepath.Join(usr.HomeDir, ".agcommits")

	// 检查配置文件是否存在
	if _, err := os.Stat(agCommitsPath); os.IsNotExist(err) {
		fmt.Println("配置文件不存在，将进行创建和设置。")
		err = utils.CreateAndSetupConfigFile(agCommitsPath)
		if err != nil {
			fmt.Println("创建和设置配置文件时出错:", err)
			return
		}
	}

	// 读取配置文件
	config, err := utils.ReadConfig(agCommitsPath)
	if err != nil {
		fmt.Println("读取配置文件时出错:", err)
		return
	}

	isInGitRepo, err := utils.IsGitRepository()
	if err != nil {
		fmt.Println("检查 Git 仓库时出错:", err)
		return
	}

	if !isInGitRepo {
		fmt.Println("当前目录不是 Git 仓库")
		return
	}

	hasStagedChanges, err := utils.HasStagedChanges()
	if err != nil {
		fmt.Println("检查暂存区时出错:", err)
		return
	}
	if !hasStagedChanges {
		fmt.Println("暂存区为空，请执行 'git add .'")
		return
	}

	// 获取 Git Diff 信息
	diff, err := utils.GetGitDiff()
	if err != nil {
		fmt.Println("获取 Git Diff 信息时出错:", err)
		return
	}

	// 调用 ChatCompletionBlocking 函数
	res, err := openai_api.ChatCompletionBlocking(config, diff)

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
	} else {
		fmt.Println("用户取消了提交操作。")
	}
}
