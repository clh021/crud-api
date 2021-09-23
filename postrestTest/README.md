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