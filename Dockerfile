#源镜像
FROM golang:alpine3.13

#安装supervisor
RUN apk add --no-cache supervisor

#将二进制文件拷贝进容器的GOPATH目录中
RUN mkdir -p /app/logs/
RUN mkdir -p /usr/local/service/ordersystem/
ADD ordersystem /usr/local/service/ordersystem/
COPY script/supervisord.ini /etc/supervisord.d/
COPY script/kick_start.sh /etc/kickStart.d/

#为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

#暴露端口
EXPOSE 8000

#最终运行docker的命令
CMD ["/etc/kickStart.d/kick_start.sh"]
