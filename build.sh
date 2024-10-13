   #!/bin/bash
   set -e

   # 编译项目 机器上有go环境 机器是linux 不需要编译成exe
   #go build -o build/main ./cmd/main.go
   # linux 产物
   GOOS=linux GOARCH=amd64 go build -o build/main ./cmd/main.go