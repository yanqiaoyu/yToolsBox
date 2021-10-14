# 先在golang的alpine中编译好
FROM golang:alpine AS development

WORKDIR $GOPATH/src/yToolsBox-gin
ADD . ./
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
RUN go build -o main .

FROM alpine:latest AS production
WORKDIR /root/
COPY --from=development /go/src/yToolsBox-gin/main .
COPY --from=development /go/src/yToolsBox-gin/config ./config
COPY --from=development /go/src/yToolsBox-gin/Script ./Script
ENTRYPOINT  ["./main"]
