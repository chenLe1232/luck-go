# Luck-Go

这是一个使用 Gin 框架的 Go 项目。

## 目录结构

```
luck-go
├── cmd                # 主程序入口
│   └── main.go        # 启动文件
├── config             # 配置文件
│   └── config.yaml    # 配置示例
├── database           # 数据库相关代码
│   ├── migrations     # 数据库迁移文件
│   └── models         # 数据模型
├── daos               # 数据访问对象
│   ├── users          # 用户相关数据访问对象
│   └── products       # 产品相关数据访问对象
├── dtos               # 数据传输对象
│   ├── user           # 用户 DTO
│   └── product        # 产品 DTO
├── handlers           # 控制器函数，处理 HTTP 请求
├── middleware         # 中间件函数，用于请求和响应的预处理
├── routes             # 路由规则，映射 URL 到控制器函数
├── services           # 业务逻辑和应用服务
│   ├── auth           # 认证和授权服务
│   ├── user           # 用户相关服务（如注册、个人资料管理）
│   └── product        # 产品相关服务（如产品管理、下单）
├── utils              # 通用工具函数和辅助代码
├── tests              # 测试代码
└── vendor             # 第三方库代码
```

## 如何运行

1. 确保已安装 Go 1.16 或更高版本。
2. 克隆此仓库。
3. 在项目根目录运行 `go mod tidy` 安装依赖。
4. 运行以下命令启动服务器：
   ```
   go run cmd/main.go
   ```
5. 服务器将在 `http://localhost:8080` 上运行。

## 本地请求测试

您可以使用以下方法测试本地服务：

### 使用浏览器（仅适用于 GET 请求）

直接在浏览器地址栏输入：
- http://localhost:8080/ping
- http://localhost:8080/hello

### 使用 curl 命令

1. GET 请求:
   ```
   curl http://localhost:8080/ping
   curl http://localhost:8080/hello
   ```

2. POST 请求:
   ```
   curl -X POST -H "Content-Type: application/json" -d '{"message":"测试消息"}' http://localhost:8080/echo
   ```

### 使用 Postman 或其他 API 测试工具

1. 打开 Postman
2. 创建新请求
3. 输入 URL（例如：http://localhost:8080/ping）
4. 选择适当的 HTTP 方法（GET 或 POST）
5. 对于 POST 请求，在 Body 选项卡中选择 raw 和 JSON，然后输入数据
6. 点击 Send 发送请求

## 配置

配置文件位于 `config/config.yaml`。请根据您的环境修改相应的配置。

## 贡献

欢迎提交 Pull Requests 来改进这个项目。

## 许可证

[MIT License](https://opensource.org/licenses/MIT)
