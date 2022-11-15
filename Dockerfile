FROM golang:latest
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

WORKDIR /app
COPY . /app

RUN go build .

EXPOSE 3001

ENTRYPOINT ["./ginblog"]