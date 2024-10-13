   #!/bin/bash
   set -e

   # 更新代码
   #git pull

   # 编译项目 机器上有go环境 机器是linux 不需要编译成exe
   #go build -o tmp/main ./cmd/main.go

   # 复制服务文件
   sudo cp  job.service /etc/systemd/system/ 

   # 重新加载 systemd
   sudo systemctl daemon-reload
   # 停止服务
    #  sudo systemctl stop job.service

   # 重启服务
   sudo systemctl restart job.service

   # 确保服务已启用
   sudo systemctl enable job.service

   # 给deploy.sh执行权限
   # chmod +x deploy.sh

   echo "部署完成！"
