# Zookeeper REST API
[![Build Status](https://travis-ci.org/tquach/zk-rest-api.svg?branch=master)](https://travis-ci.org/tquach/zk-rest-api)

For managing Zookeeper configuration data, start up this thin REST API server and connect.

## Installing
```
go get github.com/tquach/zk-rest-api
```

Or use Docker. Specify a comma-separated list of Zookeeper hosts as `ZK_HOSTS`:
```
docker run --name zkapi -e ZK_HOSTS=docker1 -p 8009:8080 tquach/zk-rest-api
```

## Usage
```
Usage of ./zk-rest-api:
  -addr=":8080": address and port to listen on
  -zk=[]:        comma-separated list of hosts to zookeeper.
```

## Example
```
[~/projects/zk-rest-api] ⚡  zk-rest-api -addr localhost:8001 -zk zookeeper-1,zookeeper-2:2081
```

## License
MIT