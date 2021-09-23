# DEV History

数据结构在一个项目的前后端经常需要写几次，项目越大越冗余，当数据结构需要修改的，同样的也需要修改前后端对应的表结构模型。
市面上已经有了一些针对设计好的数据表结构生成表结构模型的工具，比如：[grom](https://github.com/sliveryou/grom),[gormt](https://github.com/xxjwxc/gormt),[gormat](https://github.com/airplayx/gormat)。
这个项目从另一个思路来解决问题：数据库一旦设计好，前端立刻可以访问，所有数据操作逻辑都由前端负责，接口只负责表字段的读写权限验证。

前端显示表字段的对应逻辑: 表字段的注释(仅`Mysql`有字段注释概念)，`注释表`中对应表字段的记录，驼峰间隔英文
接口权限控制字段的逻辑：  配置文件，`注释表`中对应表字段的记录，默认逻辑，允许读写

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
cobra add main # 用于启动 主 服务（包含 api, web ui）
cobra add conf # 用于创建和测试 配置文件
```

## 使用目标

```
/db/mysql1/conn  链接 mysql1
/table/:tablename 查询数据         header 传递 使用哪个链接哪个数据库
```

## 鸣谢
[go-gin-api](https://github.com/xinliangnote/go-gin-api)
[go-crud-api](https://github.com/mevdschee/go-crud-api)

## 不同思路的工具
