FROM golang:1.7-onbuild
MAINTAINER Tan Quach <tan.quach@gmail.com>

RUN go install github.com/tquach/zk-rest-api

ADD scripts/run.sh /usr/local/bin/start-zk.sh

EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/start-zk.sh"]
