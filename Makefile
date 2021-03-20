test:
	go test ./... -v --count=1

fmt:
	go fmt ./...

tidy:
	go mod tidy
	go mod vendor