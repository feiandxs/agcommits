package constants

import (
	"os/user"
	"path/filepath"
)

// ConfigFileName 配置文件名
const ConfigFileName = ".agcommitsrc"

// GetConfigFilePath 获取配置文件的完整路径
func GetConfigFilePath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, ConfigFileName), nil
}
