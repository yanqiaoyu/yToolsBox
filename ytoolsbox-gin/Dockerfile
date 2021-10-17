# 先在golang的alpine中编译好二进制包
FROM golang:alpine AS development
# 设置工作目录
WORKDIR $GOPATH/src/yToolsBox-gin
# 把当前目录的文件全部拷进去
ADD . ./
# 解决国内拉模块的延迟问题
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
# 开始编译
RUN go build -o main .

# 然后把可执行程序放进最小环境里面
FROM alpine:latest AS production
# 设置工作目录
WORKDIR /root/
# 拷贝必须的文件
COPY --from=development /go/src/yToolsBox-gin/main .
COPY --from=development /go/src/yToolsBox-gin/config ./config
COPY --from=development /go/src/yToolsBox-gin/Script ./Script
# 挂载容器目录
VOLUME ["/root/config", "/root/Script"]
# 执行
ENTRYPOINT  ["./main"]
