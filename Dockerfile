FROM golang:1.4.2-onbuild
MAINTAINER Tan Quach <tan.quach@gmail.com>

RUN go get github.com/tquach/zk-rest-api
RUN go install github.com/tquach/zk-rest-api

ADD scripts/run.sh /usr/local/bin/start-zk.sh

EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/start-zk.sh"]
