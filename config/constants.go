package config

import "errors"

const (
	ConfigFileName = ".agcommitsrc.yaml"
)

var (
	ErrConfigNotFound     = errors.New("配置文件不存在")
	ErrConfigInvalid      = errors.New("配置文件格式不正确")
	ErrRequiredFieldEmpty = errors.New("必填字段为空")
)
