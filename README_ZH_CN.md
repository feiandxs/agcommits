# AGCOMMITS

使用 AI 自动生成 git commit messages

[英文文档](./README.md)

## 安装
要安装 `agcommits`，使用`go install` 命令:

```shell
go install github.com/feiandxs/agcommits@latest
```

然后，您可以将 `agcommits` 可执行文件添加到您的 ~/.bashrc 或 ~/.bash_profile 文件中的 PATH 环境变量:

>如果您已经安装了  `agcommits` ，更新 `agcommits` 很简单:

```
go get -u github.com/feiandxs/agcommits
```

## 初始化
```shell
agcommits
```
根据提示输入您的 APIKEY 等信息。

## 使用方法
```shell
cd /path/to/your/project
```

```shell
agcommits
```

然后您可以在终端中看到生成的提交信息。

就是这些。

## 配置说明

AGCOMMITS 支持全局配置和项目级配置：

- **全局配置**：`~/.agcommitsrc.yaml`
- **项目配置**：项目根目录下的 `.agcommits.yaml`

详细配置选项请参考 [.agcommits.yaml.example](./.agcommits.yaml.example)。

### 快速配置示例

```yaml
# 必填字段
openai_key: "your-api-key"
openai_api_base: "https://api.siliconflow.cn"
openai_model: "Qwen/Qwen2.5-Coder-7B-Instruct"

# 可选字段
commit_locale: "zh"      # 语言：zh（中文）或 en（英文）
max_length: 150          # 提交信息最大长度
auto_add: false          # 自动执行 git add
auto_commit: false       # 自动执行 git commit
```
