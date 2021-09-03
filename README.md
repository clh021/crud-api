# crud-api

将可连接数据库实例，一次变换为可操作的 WebAPI 的方式。

## 使用

参考 `.crud-api.yaml.example` 定义配置

`./crud-api` 默认启动 `web` 服务

## 规则

### 创建
访问: POST: /{tablename}
数据:
```json
    {
        "title": "Black is the new red",
        "content": "This is the second post.",
        "created": "2018-03-06T21:34:01Z"
    }
```
返回: 新增数据的 主键值

### 读取
访问: GET: /{tablename}/{primaryKeyVal}
返回:
```json
    {
        "id": 1,
        "title": "Black is the new red",
        "content": "This is the second post.",
        "created": "2018-03-06T21:34:01Z"
    }
```

### 更新
访问: PUT: /{tablename}/{primaryKeyVal}
数据:
```json
    {
        "title": "Black is the new red2",
    }
```
返回: 受影响的行数

### 删除
访问: DELETE: /{tablename}/{primaryKeyVal}
返回: 受影响的行数


### 列表
访问: GET: /{tablename}
返回:
```json
    {
        "records":[
            {
                "id": 1,
                "title": "Black is the new red",
                "content": "This is the second post.",
                "created": "2018-03-06T21:34:01Z"
            }
        ]
    }
```

### 过滤器
过滤器使用 `filter` 参数在 `列表`请求  中提供搜索功能。使用`英文逗号`间隔`列名`,`条件类型`,`条件值`，支持的匹配条件如下：

  - "cs": contain string (包含字符串)
  - "sw": start with (字符串以值开头)
  - "ew": end with (字符串以值结尾)
  - "eq": equal (字符串完全匹配)
  - "lt": lower than (数字低于)
  - "le": lower or equal (数字低于或等于)
  - "ge": greater or equal (数字大于或等于)
  - "gt": greater than (数字大于)
  - "bt": between (数字介于两个值之间)
  - "in": in (数字或字符串处于值列表中)
  - "is": is null (字段为 "NULL" 的值)

您可以通过在前面加上一个 `n` 字符来表示否定，如 `eq` 变成 `neq`。

示例:

    GET /records/categories?filter=name,eq,Internet
    GET /records/categories?filter=name,sw,Inter
    GET /records/categories?filter=id,le,1
    GET /records/categories?filter=id,ngt,1
    GET /records/categories?filter=id,bt,0,1
    GET /records/categories?filter=id,in,0,1

返回:

    {
        "records":[
            {
                "id": 1
                "name": "Internet"
            }
        ]
    }

### 多重过滤

#### 并列关系
并列关系的过滤条件可以通过重复 URL 中的 `filter` 参数来应用。 
例如：表示过滤条件 `where id > 1 and id < 3` 的 URL 如下
`GET /records/categories?filter=id,gt,1&filter=id,lt,3`

#### 或者关系
或者关系的过滤条件可以通过为 `filter` 参数添加结尾数字来表示。
 URL 中的 `filter` 参数来应用。 例如：
`GET /records/categories?filter1=id,gt,2&filter2=id,lt,4`

> 还可以通过 重复 `filter1` 来在 `或者语句` 中创建 `并列关系`。
> 注意：所有列只能是当前请求表范围内的字段。

### 排序
通过 `order` 参数指定排序字段。 默认是升序的，通过指定 `desc` 可以倒序：
```
GET /records/categories?order=name,desc
GET /records/categories?order=id,desc&order=name
```

### 限制数量
`size` 参数限制返回的条目数。 可与`order`参数一起用于前 N 个列表（使用降序）。
```
GET /records/categories?order=id,desc&size=1
```
> 如果还想知道记录总数，应使用 `page` 参数。
### 翻页
`page` 参数保存请求的页面。 默认页面大小为 20，但可以调整（例如到 50）。
```
GET /records/categories?order=id&page=1
GET /records/categories?order=id&page=1,50
```
返回
```
    {
        "records":[
            {
                "id": 1
                "name": "Internet"
            },
            {
                "id": 3
                "name": "Web development"
            }
        ],
        "results": 2
    }
```

### 连表查询

TODO