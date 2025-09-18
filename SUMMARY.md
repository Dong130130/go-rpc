#  项目实现步骤

## 1. 配置管理

使用 config.json 文件配置服务端监听地址、端口以及认证信息，例如：

```
{
  "host": "127.0.0.1",
  "port": 8080,
  "auth": {
    "user1": "12345"
  }
}
```

配置文件统一管理服务端信息和用户认证，实现灵活配置和安全认证。


## 2. 服务端实现

使用 Gin 框架 搭建 HTTP 服务。

实现三个主要接口：

### 2.1 Add 接口

- 定义请求参数结构体 Args， 包含 a、b 和用户认证信息 (username / password)。

- 实现 addHandler 函数：

  - 使用 c.ShouldBindJSON(&a) 解析请求 JSON，校验 JSON 参数合法性。

  - 调用 checkAuth(a.Username, a.Password) 进行用户认证

  - 返回两数之和。

- 错误处理：

  - JSON 解析失败 → 返回 400。

  - 认证失败 → 返回 401。


### 2.2 Division 接口

- 复用 Args 结构体

- 增加除数为 0 的校验，避免运行时错误。

- 请求处理流程与 Add 接口类似：

  - 参数解析 → 用户认证 → 商计算。

- 错误处理：

  - 除数为 0 → 返回 400。

  - JSON 解析失败 → 返回 400。

  - 认证失败 → 返回 401。


### 2.3 System 接口     

- 定义请求参数结构体 SystemCmd ，包含命令字符串及认证信息。

- 实现 systemHandler 函数：

  - 参数解析 → c.ShouldBindJSON(&cmd)

  - 用户认证 → checkAuth(cmd.Username, cmd.Password)

- 执行系统命令：

  - Linux/macOS → exec.Command("bash", "-c", cmd.Cmd)

  - Windows → exec.Command("cmd", "/C", cmd.Cmd)

- 获取标准输出和错误输出 → out, err := ec.CombinedOutput()

- 返回命令输出。

- 错误处理：

  - 参数解析失败 → 400。

  - 认证失败 → 401。

  - 命令执行异常 → 500，并返回错误信息和输出。


## 3. 认证机制

- 所有 RPC 方法调用前进行认证：

  - 校验 username 和 password 是否匹配配置文件中的记录

  - 认证失败 → 返回 HTTP 401

- 保证只有通过认证的客户端才能调用接口


## 4. HTTP 服务初始化

- 使用 Gin 框架 创建 HTTP 引擎实例。

- 注册 RPC 接口路由：

  - /add → addHandler：计算两个数的和

  - /division → divisionHandler：计算两个数的商

  - /system → systemHandler：执行系统命令



## 5. 接口文档

- 接口注解：项目中每个 RPC 接口均使用 Swaggo 注解（注释形式）标注，包含接口简介、请求参数、响应格式和错误码说明。

- 文档生成：通过 Swaggo 自动生成 Swagger UI 文档，无需手动编写接口文档。

- 功能：支持交互式调用和测试，方便开发和验证接口正确性。

访问方式：

```
http://127.0.0.1:8080/swagger/index.html
```

## 6. 总结

本项目通过 Go + Gin + Swaggo 实现了一个带认证机制的 RPC 服务端示例：

接口清晰、功能完整（Add / Division / System）。

支持安全认证和错误处理。

配置灵活，可通过 config.json 修改端口、认证信息。

提供 Swagger UI 文档，方便测试和开发验证。
