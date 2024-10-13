#!/bin/bash

# 创建项目根目录
mkdir luck-go
cd luck-go

# 初始化 Go 模块
go mod init github.com/yourusername/luck-go

# 安装 Gin 框架
go get -u github.com/gin-gonic/gin

# 创建目录结构
mkdir -p cmd config database/{migrations,models} daos/{users,products} dtos/{user,product} handlers middleware routes services/{auth,user,product} utils tests vendor

# 创建基本文件
touch cmd/main.go config/config.yaml
touch database/models/{user.go,product.go}
touch daos/users/user_dao.go daos/products/product_dao.go
touch dtos/user/user_dto.go dtos/product/product_dto.go
touch handlers/{user_handler.go,product_handler.go}
touch middleware/auth_middleware.go
touch routes/routes.go
touch services/auth/auth_service.go services/user/user_service.go services/product/product_service.go
touch utils/helpers.go
touch tests/main_test.go
touch README.md

# 创建 main.go 文件内容
cat << EOF > cmd/main.go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
EOF

# 创建 config.yaml 文件内容
cat << EOF > config/config.yaml
server:
  port: 8080
  mode: debug

database:
  driver: mysql
  host: localhost
  port: 3306
  username: root
  password: password
  dbname: myapp

redis:
  host: localhost
  port: 6379
  password: ""
EOF

# 创建 README.md 文件内容
cat << EOF > README.md
# Luck-Go

这是一个使用 Gin 框架的 Go 项目。

## 目录结构

\`\`\`
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
\`\`\`

## 如何运行

1. 确保已安装 Go 1.16 或更高版本。
2. 克隆此仓库。
3. 在项目根目录运行 \`go mod tidy\` 安装依赖。
4. 运行 \`go run cmd/main.go\` 启动服务器。
5. 访问 \`http://localhost:8080/ping\` 测试服务是否正常运行。

## 配置

配置文件位于 \`config/config.yaml\`。请根据您的环境修改相应的配置。

## 贡献

欢迎提交 Pull Requests 来改进这个项目。

## 许可证

[MIT License](https://opensource.org/licenses/MIT)
EOF

# 创建 .gitignore 文件
cat << EOF > .gitignore
# 编译后的二进制文件
*.exe
*.exe~
*.dll
*.so
*.dylib

# 测试二���制文件
*.test

# 输出的性能分析文件
*.prof

# Go工作空间文件
go.work

# 依赖目录
/vendor/

# 输出目录
/bin/
/dist/

# IDE和编辑器文件
.idea/
.vscode/
*.swp
*.swo
*~

# 操作系统生成的文件
.DS_Store
Thumbs.db

# 日志文件
*.log

# 环境变量文件
.env

# Gin框架生成的临时文件
gin-bin

# 其他临时文件
tmp/
temp/

# 缓存文件
.cache/
EOF

# 为各个文件添加适当的包名
cat << EOF > database/models/user.go
package models

// User 结构体定义
type User struct {
    // 添加用户字段
}
EOF

cat << EOF > database/models/product.go
package models

// Product 结构体定义
type Product struct {
    // 添加产品字段
}
EOF

cat << EOF > daos/users/user_dao.go
package users

// UserDAO 接口定义
type UserDAO interface {
    // 添加用户数据访问方法
}
EOF

cat << EOF > daos/products/product_dao.go
package products

// ProductDAO 接口定义
type ProductDAO interface {
    // 添加产品数据访问方法
}
EOF

cat << EOF > dtos/user/user_dto.go
package user

// UserDTO 结构体定义
type UserDTO struct {
    // 添加用户DTO字段
}
EOF

cat << EOF > dtos/product/product_dto.go
package product

// ProductDTO 结构体定义
type ProductDTO struct {
    // 添加产品DTO字段
}
EOF

cat << EOF > handlers/user_handler.go
package handlers

// 用户相关的处理函数
EOF

cat << EOF > handlers/product_handler.go
package handlers

// 产品相关的处理函数
EOF

cat << EOF > middleware/auth_middleware.go
package middleware

// 认证中间件函数
EOF

cat << EOF > routes/routes.go
package routes

// 路由设置函数
EOF

cat << EOF > services/auth/auth_service.go
package auth

// AuthService 接口定义
type AuthService interface {
    // 添加认证服务方法
}
EOF

cat << EOF > services/user/user_service.go
package user

// UserService 接口定义
type UserService interface {
    // 添加用户服务方法
}
EOF

cat << EOF > services/product/product_service.go
package product

// ProductService 接口定义
type ProductService interface {
    // 添加产品服务方法
}
EOF

cat << EOF > utils/helpers.go
package utils

// 辅助函数定义
EOF

cat << EOF > tests/main_test.go
package tests

import "testing"

// 主测试函数
func TestMain(m *testing.M) {
    // 测试设置
}
EOF

echo "项目初始化完成，包括 .gitignore 和包定义！"
