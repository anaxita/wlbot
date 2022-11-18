export VERSION = $(shell git describe --tags --always --abbrev=0)

win-build:
	GOOS=windows GOARCH=amd64 go build -ldflags="-X 'wlbot/pkg/version.V=${VERSION}'" -o bin/wlbot-${VERSION}-windows-amd64.exe cmd/main.go

run:
	go run -ldflags="-X 'wlbot/pkg/version.V=${VERSION}'" cmd/main.go

lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.50.1 golangci-lint run