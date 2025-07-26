# AGCOMMITS

AI Generated Commits

[中文文档](./README_ZH_CN.md)

## Installation
To install `agcommits` use the `go install` command:

```shell
go install github.com/feiandxs/agcommits@latest
```

Then you can add `agcommits` binary to PATH environment variable in your ~/.bashrc or ~/.bash_profile file:

>If you already have `agcommits` installed, updating `agcommits` is simple:

```
go get -u github.com/feiandxs/agcommits
```

## Initialization
```shell
agcommits
```
Enter your APIKEY and other information as prompted.

## Usage
```shell
cd /path/to/your/project
```

```shell
agcommits
```

Then you can see the generated commit message in the terminal.

That's all.

## Configuration

AGCOMMITS supports both global and project-specific configurations:

- **Global config**: `~/.agcommitsrc.yaml`
- **Project config**: `.agcommits.yaml` in your project root

For detailed configuration options, see [.agcommits.yaml.example](./.agcommits.yaml.example).

### Quick Configuration Example

```yaml
# Required fields
openai_key: "your-api-key"
openai_api_base: "https://api.siliconflow.cn"
openai_model: "Qwen/Qwen2.5-Coder-7B-Instruct"

# Optional fields
commit_locale: "en"      # Language: zh (Chinese) or en (English)
max_length: 150          # Maximum commit message length
auto_add: false          # Auto-execute git add
auto_commit: false       # Auto-execute git commit
```
