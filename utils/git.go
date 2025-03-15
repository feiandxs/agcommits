package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// IsGitRepository 检查当前目录是否为 Git 仓库。
func IsGitRepository() (bool, error) {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()
	if err != nil {
		return false, nil // 不是 Git 仓库，但不视为错误
	}
	return true, nil
}

// HasStagedChanges 检查是否有暂存的更改
func HasStagedChanges() (bool, error) {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")
	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("执行 git diff 命令失败: %v", err)
	}
	return len(strings.TrimSpace(string(output))) > 0, nil
}

// GetGitDiff 获取 Git 暂存区的 diff 信息
func GetGitDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("执行 git diff 命令失败: %v", err)
	}
	return string(output), nil
}

// ConfirmCommitMessage 显示提交消息并询问用户是否确认使用。
func ConfirmCommitMessage(commitMsg string) bool {
	fmt.Println("AI 生成的 Git 提交消息如下：")
	fmt.Println(commitMsg)
	fmt.Println("是否使用此提交消息进行 Git 提交？(y/n，默认: y)")

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入时出错:", err)
		return false
	}
	response = strings.TrimSpace(response)

	return response == "y" || response == ""
}

// PerformGitCommit 执行 Git 提交
func PerformGitCommit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git commit 失败: %v, %s", err, stderr.String())
	}
	return nil
}

// GitAddAll 执行 git add . 命令
func GitAddAll() error {
	cmd := exec.Command("git", "add", ".")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git add . 失败: %v, %s", err, stderr.String())
	}
	return nil
}

// InitGitRepository 在当前目录初始化一个新的 Git 仓库
func InitGitRepository() error {
	cmd := exec.Command("git", "init")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git init 失败: %v, %s", err, stderr.String())
	}
	return nil
}

// ConfirmGitInit 询问用户是否要初始化Git仓库
func ConfirmGitInit() bool {
	fmt.Println("当前目录不是Git仓库，是否要初始化一个新的Git仓库？(y/n，默认: n)")

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入时出错:", err)
		return false
	}
	response = strings.TrimSpace(response)

	return response == "y" || response == "Y"
}

// ConfirmGitAdd 询问用户是否要执行 git add . 命令
func ConfirmGitAdd() bool {
	fmt.Println("暂存区没有更改，是否要执行 git add . 命令？(y/n，默认: n)")

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入时出错:", err)
		return false
	}
	response = strings.TrimSpace(response)

	return response == "y" || response == "Y"
}
