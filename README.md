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
4. 运行 `go run cmd/main.go` 启动服务器。
5. 访问 `http://localhost:8080/ping` 测试服务是否正常运行。

## 配置

配置文件位于 `config/config.yaml`。请根据您的环境修改相应的配置。

## 贡献

欢迎提交 Pull Requests 来改进这个项目。

## 许可证

[MIT License](https://opensource.org/licenses/MIT)
