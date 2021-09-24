# webapi 替代方案

项目地址: https://github.com/PostgREST/postgrest

## 服务准备
```bash
# 启动 `postgres` 数据库服务
docker-compose up -d

# 下载 webapi 转换程序
# download from https://github.com/begriffs/postgrest/releases/latest
tar xfJ postgrest-<version>-<platform>.tar.xz

# 进入 postgres 客户端
docker exec -it postgresttest_db_1 psql -U postgres
```

## 数据准备
```sql
create schema api;
create table api.todos (
  id serial primary key,
  done boolean not null default false,
  task text not null,
  due timestamptz
);

insert into api.todos (task) values
  ('finish tutorial 0'), ('pat self on back');

create role web_anon nologin;
grant web_anon to postgres;

grant usage on schema api to web_anon;
grant select on api.todos to web_anon;

\q # 退出
```

## 启动 webapi 服务
```bash
./postgrest test.conf
```

## 访问测试
```bash
curl http://localhost:3000/todos

#output:
#[{"id":1,"done":false,"task":"finish tutorial 0","due":null}, {"id":2,"done":false,"task":"pat self on back","due":null}]
```

## 更多资料
https://postgrest-docs-chinese.readthedocs.io/zh/latest/tutorials/tut1.html


## CRUD

添加一个受信的用户
```sql
create role todo_user nologin;
grant todo_user to postgres;

grant usage on schema api to todo_user;
grant all on api.todos to todo_user;
grant usage, select on sequence api.todos_id_seq to todo_user;
\q
```

生成一个密码
```bash
openssl rand -base64 32
#output:pZTvJiRKEzXWa4DYNn2FNIxM4o/ndefyq4GzeunYICQ=

#add this line to test.conf
#jwt-secret = "pZTvJiRKEzXWa4DYNn2FNIxM4o/ndefyq4GzeunYICQ="
```

生成 token
```
访问 https://jwt.io
填写 payload:  {"role": "todo_user"}
填写 secret:   pZTvJiRKEzXWa4DYNn2FNIxM4o/ndefyq4GzeunYICQ=

复制左侧的 token
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidG9kb191c2VyIn0.tLpcxzsFRKtGLG74eVqfuyZVvyE8wWRSHsmeWGwzsw0

数据容易逆向，勿写入隐私信息
```

进行添加新任务请求
```bash
export TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidG9kb191c2VyIn0.tLpcxzsFRKtGLG74eVqfuyZVvyE8wWRSHsmeWGwzsw0"
curl http://localhost:3000/todos -X POST \
     -H "Authorization: Bearer $TOKEN"   \
     -H "Content-Type: application/json" \
     -d '{"task": "learn how to auth"}'
```

```bash
curl http://localhost:3000/todos

curl http://localhost:3000/todos -X PATCH \
     -H "Authorization: Bearer $TOKEN"    \
     -H "Content-Type: application/json"  \
     -d '{"done": true}'
```

会过期 token
```sql
select extract(epoch from now() + '5 minutes'::interval) :: integer;
```

```
访问 https://jwt.io
填写 payload:  {"role": "todo_user","exp":"1632401835"}
填写 secret:   pZTvJiRKEzXWa4DYNn2FNIxM4o/ndefyq4GzeunYICQ=

复制左侧的 token
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidG9kb191c2VyIiwiZXhwIjoiMTYzMjQwMTgzNSJ9.5p8_cPtWFneSNy1wIsfQUYsu9cc348Jp821kzbyV370

数据容易逆向，勿写入隐私信息
```


```bash
export TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidG9kb191c2VyIn0.tLpcxzsFRKtGLG74eVqfuyZVvyE8wWRSHsmeWGwzsw0"
curl http://localhost:3000/todos -X POST \
     -H "Authorization: Bearer $TOKEN"   \
     -H "Content-Type: application/json" \
     -d '{"task": "learn how to auth23456"}'

# 过期的返回
# {"message":"JWT expired"}
```

