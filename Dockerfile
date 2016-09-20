FROM golang:1.6.2

EXPOSE 18080

ENV TIME_ZONE=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

COPY . /go/src/github.com/zonesan/go-hello

WORKDIR /go/src/github.com/zonesan/go-hello

RUN go build

CMD ["sh", "-c", "./go-hello"]
