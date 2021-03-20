help:
	@ cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## Runs test within project
	go test ./... -v --count=1

fmt: ## formats code
	go fmt ./...

tidy: ## updates dependencies and save them in vendor folder
	go mod tidy
	go mod vendor

build: tidy ## build local binary
	go build

test-example: build ## Run sample example with rectangle={leftBottom: {x: 0, y: 0}, rightTop: {x: 10, y: 10}}, rectangle={leftBottom: {x: 10, y: 0}, rightTop: {x: 7, y: 7}}
	./rectangle-assessment execute 0 0 10 10 10 0 7 7

.PHONY: help test fmt tidy build
.DEFAULT_GOAL := help
