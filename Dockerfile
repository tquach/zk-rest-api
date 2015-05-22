FROM golang:1.4.2
MAINTAINER Tan Quach <tan.quach@gmail.com>
RUN go get github.com/tquach/zk-rest-api
RUN go install github.com/tquach/zk-rest-api

EXPOSE 8080
ENTRYPOINT ["zk-rest-api -zk $ZK_HOSTS"]
