all             : build

APP_NAME   		?= zk-rest-api
DOCKER_IP      	?= $(shell docker-machine ip)

build    		: 
	godep go build -o $(APP_NAME)

local           : nocachebuild start

re              : clean all

clean_base      : clean_zk

.PHONY          : clean_base re local build test
