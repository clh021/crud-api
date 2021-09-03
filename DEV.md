# DEV History
----
## 创建项目
```bash
# 安装 cobra
go get -u github.com/spf13/cobra
# 生成项目初始文件
cobra init --pkg-name github.com/clh021/crud-api
# 初始化 项目
go mod init github.com/clh021/crud-api
go mod tidy
# 初次编译
go build
# 初次执行
./crud-api
```


## 创建命令
### 概念
Cobra 结构由三部分组成：命令 (commands)、参数 (arguments)、标志 (flags)。最好的应用程序在使用时读起来像句子，要遵循的模式是APPNAME VERB NOUN --ADJECTIVE。

```bash
# 使用 cobra 创建 `web`,`conf` 三个命令
cobra add web # 用于启动 web 服务
cobra add conf # 用于创建和测试 配置文件
```