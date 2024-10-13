   #!/bin/bash
   set -e

   # 更新代码
   #git pull

   # 编译项目 机器上有go环境 机器是linux 不需要编译成exe
   #go build -o tmp/main ./cmd/main.go

   # 复制服务文件
   sudo cp -n job.service /etc/systemd/system/ || echo "服务文件已存在，跳过复制"

   # 重新加载 systemd
   sudo systemctl daemon-reload

   # 重启服务
   sudo systemctl restart carline-go

   # 确保服务已启用
   sudo systemctl enable carline-go

   # 给deploy.sh执行权限
   # chmod +x deploy.sh

   echo "部署完成！"
