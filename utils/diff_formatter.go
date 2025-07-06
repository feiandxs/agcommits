package utils

import (
	"fmt"
	"strings"
)

// FormatDiff 格式化Git diff输出，使其更易读
func FormatDiff(diff string) string {
	if diff == "" {
		return "没有发现任何更改"
	}

	// 分割成行
	lines := strings.Split(diff, "\n")
	var formatted strings.Builder
	var currentFile string

	formatted.WriteString("\n=== 代码变更摘要 ===\n")

	for _, line := range lines {
		// 跳过空行
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		// 处理文件头
		if strings.HasPrefix(line, "diff --git") {
			if currentFile != "" {
				formatted.WriteString("\n")
			}
			parts := strings.Split(line, " b/")
			if len(parts) > 1 {
				currentFile = parts[1]
				formatted.WriteString(fmt.Sprintf("\n📄 文件: %s\n", currentFile))
				formatted.WriteString("----------------------------------------\n")
			}
			continue
		}

		// 处理变更信息
		if strings.HasPrefix(line, "index") || strings.HasPrefix(line, "---") || strings.HasPrefix(line, "+++") {
			continue
		}

		// 格式化变更行
		switch {
		case strings.HasPrefix(line, "-"):
			formatted.WriteString(fmt.Sprintf("❌ %s\n", line))
		case strings.HasPrefix(line, "+"):
			formatted.WriteString(fmt.Sprintf("✅ %s\n", line))
		case strings.HasPrefix(line, "@@ "):
			formatted.WriteString(fmt.Sprintf("\n📍 %s\n", line))
		default:
			formatted.WriteString(fmt.Sprintf("   %s\n", line))
		}
	}

	formatted.WriteString("\n=== 变更摘要结束 ===\n")
	return formatted.String()
}
