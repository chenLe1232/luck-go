```bash
# 清除代理
go env -u GOPROXY
```

```bash
# 安装依赖
go mod tidy
```

```bash
# 清除 GOPROXY
go env -u GOPROXY

# 清除 GOSUMDB（如果设置了的话）
go env -u GOSUMDB

# 清除 GONOSUMDB（如果设置了的话）
go env -u GONOSUMDB

# 清除 GOPRIVATE（如果设置了的话）
go env -u GOPRIVATE
```

```bash
# gvm 设置默认版本
gvm use go1.17.4 --default
```

[m 芯片](https://soulteary.com/2022/05/12/better-golang-usage-on-m1-mac.html)
