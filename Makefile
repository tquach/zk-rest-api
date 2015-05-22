all             : build

MAKE_PATH      ?= .

include $(MAKE_PATH)/zk.make $(MAKE_PATH)/format.make

APP_ENV        ?= dev
DOCKER_IP      ?= docker1
APP_NODE_ID    ?= $(shell hostname)

# Useful when developing outside current dir (libs, other services)
nocachebuild    :
		go build -o $(NAME)

local           : nocachebuild start

re              : clean all

clean_base      : clean_zk

.PHONY          : clean_base re local nocachebuild test
