## 项目简介

这是一个基于 Go 语言实现的 RPC（Remote Procedure Call）简单应用，
支持 add、division、system 三个远程调用方法，并提供认证机制和接口文档（Swagger UI）。

##  功能列表

- add：计算两个数的和

- division：计算两个数的商（支持除数校验）

- system：执行系统命令并返回结果

- 认证机制（用户名/密码）

- Swagger UI 接口文档

## 项目目录结构

```
├── docs/ # Swagger 自动生成的 API 文档
├── README.md # 项目说明文档
├── SUMMARY.md # 项目总结文档
├── config.json # 项目配置文件
├── go.mod # Go 模块定义
├── go.sum # Go 依赖版本锁定文件
├── handlers.go # 业务处理逻辑
└── main.go # 项目入口
```


## 配置文件

config.json

```
{
  "host": "127.0.0.1",
  "port": 8080,
  "auth": {
    "user1": "12345",
    "admin": "admin123"
  }
}
```

这是**默认配置**，可以通过修改 `config.json` 来改变端口或认证用户。 

## 启动方式

```
go run main.go handlers.go
```

启动成功后访问接口文档：http://127.0.0.1:8080/swagger/index.html

## 示例请求

```
POST /add
{
  "a": 10,
  "b": 20,
  "username": "user1",
  "password": "12345"
}
```

响应：

```
{
  "result": 30
}
```
[项目总结文档请见][SUMMARY.md](./SUMMARY.md)
