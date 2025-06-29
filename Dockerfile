FROM golang:1.24.4

ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /workspace

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

CMD ["go","run","main.go"]

