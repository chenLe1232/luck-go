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

## 可用路由

- http://localhost:7001/ping
  ```
  curl http://localhost:7001/ping
  ``` 
- http://localhost:7001/hello
  ```
  curl http://localhost:7001/hello 
  ```
- http://localhost:7001/api/wallpaper (POST)
  ```
  curl -X POST "http://localhost:7001/api/wallpaper?category=machine&pageSize=10&pageStart=1"
  ```

## 如何运行

1. 确保已安装 Go 1.16 或更高版本。
2. 克隆此仓库。
3. 在项目根目录运行 `go mod tidy` 安装依赖。
4. 运行以下命令启动服务器：
   ```
   go run cmd/main.go
   ```
5. 服务器将在 `http://localhost:7001` 上运行。

## 测试路由

您可以使用以下方法测试新添加的路由：

1. 使用浏览器访问：
   - http://localhost:7001/ping
   - http://localhost:7001/hello

2. 使用 curl 命令：
   ```
   curl http://localhost:7001/ping
   curl http://localhost:7001/hello
   ```

3. 对于 `/api/wallpaper` 路由，您可以使用 POST 请求：
   ```
   curl -X POST "http://localhost:7001/api/wallpaper?category=nature&pageSize=10&pageStart=1"
   ```
4. 安装 Air 用于热重载（可选）：
   ```
   go install github.com/air-verse/air@latest
   ```
5. 运行以下命令启动服务器：
   - 不使用 Air：
     ```
     go run cmd/main.go
     ```
   - 使用 Air（如果已安装）：
     ```
     air
     ```
> [!WARNING]
> ### 环境差异
> - **macOS**: 在 `.air.toml` 中使用 `bin` 配置：
>   ```toml
>   bin = "./tmp/main.exe"
>   ```
> - **Windows 11**: 使用 `full_bin` 配置可能会报错，详情请查看 [issue #659](https://github.com/air-verse/air/issues/659)：
>   ```toml
>   full_bin = "D:\\code-work\\carline-go\\luck-go\\tmp\\main.exe"
>   ```
6. 服务器将在 `http://localhost:7001` 上运行。
## 配置

配置文件位于 `config/config.yaml`。请根据您的环境修改相应的配置。

## 贡献

欢迎提交 Pull Requests 来改进这个项目。

## 许可证

[MIT License](https://opensource.org/licenses/MIT)
