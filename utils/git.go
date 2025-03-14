package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// IsGitRepository 检查当前目录是否为 Git 仓库。
func IsGitRepository() (bool, error) {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	output, err := cmd.CombinedOutput()
	outputStr := strings.TrimSpace(string(output))

	// 如果命令执行成功，检查输出是否为 'true'
	if err == nil {
		return outputStr == "true", nil
	}

	// 如果有错误，并且输出不是 'true'，则可能是因为当前目录不是 Git 仓库
	if outputStr != "true" {
		return false, nil
	}

	// 其他错误
	return false, err
}

// HasStagedChanges 检查是否执行过 'git add .'
func HasStagedChanges() (bool, error) {
	cmd := exec.Command("git", "diff", "--cached")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, err
	}
	return strings.TrimSpace(string(output)) != "", nil
}

// GetGitDiff 获取当前 Git 仓库已暂存的 diff 信息。
func GetGitDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
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

// PerformGitCommit 执行 Git 提交操作。
func PerformGitCommit(commitMsg string) error {
	cmd := exec.Command("git", "commit", "-m", commitMsg)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("执行 git commit 时出错: %v", err)
	}

	fmt.Println("Git 提交成功！")
	return nil
}
