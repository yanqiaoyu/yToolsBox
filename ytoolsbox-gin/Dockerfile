FROM golang:alpine AS development

WORKDIR $GOPATH/src/yToolsBox-gin
ADD . ./
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
RUN go build -o main .

FROM alpine:latest AS production
WORKDIR /root/
COPY --from=development /go/src/yToolsBox-gin .
EXPOSE 8080
ENTRYPOINT  ["./main"]
