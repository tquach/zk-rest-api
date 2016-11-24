REPO		?= tquach
APP_NAME	?= zk-rest-api
TAG			?= latest

deps:
	@go get -u github.com/tools/godep
	@go get -u github.com/alecthomas/gometalinter
	@gometalinter --install

all: $(APP_NAME)

$(APP_NAME): test
	@echo "Running tests..."
	@godep go build .

build: test
	@echo "Building ${APP_NAME}/${APP_NAME}:${TAG} ..."
	docker build -t $(REPO)/$(APP_NAME):$(TAG) .

test: deps lint
	@godep go test ./... 

lint: 
	@gometalinter --vendor --disable gotype --fast --errors ./...

clean:
	@rm -f $(APP_NAME)

deploy: build
	docker push $(REPO)/$(APP_NAME):$(TAG) 

.PHONY:	start clean test build deploy lint deps