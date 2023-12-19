# AGCOMMITS

使用 AI 自动生成 git commit messages

[英文文档](./README.md)

## 安装
要安装 `agcommits`，使用`go insall` 命令:

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
