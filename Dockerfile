FROM golang:latest
MAINTAINER "lk"
WORKDIR /go/src/code.meikeland.com/wanghejun/cruddemo
# 将代码从代码库复制到打包环境的WORKDIR
COPY . .
# 将main文件，从cmd中复制到WORKDIR的根目录
COPY cmd/cruddemo/main.go ./main.go
RUN go env -w GO111MODULE="on"
RUN go env -w GOPROXY="https://goproxy.cn,direct"
RUN go build .
EXPOSE 8081
ENTRYPOINT ["./cruddemo"]